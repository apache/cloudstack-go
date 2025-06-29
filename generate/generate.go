//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

const pkg = "cloudstack"

// detailsRequireKeyValue is a prefilled map with a list of details
// that need to be encoded using an explicit key and a value entry.
var detailsRequireKeyValue = map[string]bool{
	"addGuestOs":                  true,
	"addImageStore":               true,
	"addResourceDetail":           true,
	"createSecondaryStagingStore": true,
	"updateCloudToUseObjectStore": true,
	"updateGuestOs":               true,
	"updateZone":                  true,
	"addObjectStoragePool":        true,
}

// detailsRequireZeroIndex is a prefilled map with a list of details
// that need to be encoded using zero indexing
var detailsRequireZeroIndex = map[string]bool{
	"registerTemplate": true,
	"updateTemplate":   true,
	"createAccount":    true,
	"updateAccount":    true,
}

// requiresPost is a prefilled set of API names that require POST
// for security or size purposes
var requiresPostMethod = map[string]bool{
	"login":                            true,
	"deployVirtualMachine":             true,
	"updateVirtualMachine":             true,
	"createUser":                       true,
	"updateUser":                       true,
	"addVpnUser":                       true,
	"registerUserData":                 true,
	"setupUserTwoFactorAuthentication": true,
	"validateUserTwoFactorAuthenticationCode": true,
	"quotaTariffCreate":                       true,
}

var mapRequireList = map[string]map[string]bool{
	"deployVirtualMachine": map[string]bool{
		"dhcpoptionsnetworklist": true,
		"iptonetworklist":        true,
		"nicnetworklist":         true,
	},
	"updateVirtualMachine": map[string]bool{
		"dhcpoptionsnetworklist": true,
	},
	"migrateVirtualMachineWithVolume": map[string]bool{
		"migrateto": true,
	},
	"importVm": map[string]bool{
		"networklist":          true,
		"nicipaddresslist":     true,
		"datadiskofferinglist": true,
	},
	"registerOauthProvider": map[string]bool{
		"details": true,
	},
}

// nestedResponse is a prefilled map with the list of endpoints
// that responses fields are nested in a parent object. The map value
// gives the object field name.
var nestedResponse = map[string]string{
	"getUploadParamsForTemplate":       "getuploadparams",
	"getUploadParamsForVolume":         "getuploadparams",
	"createRole":                       "role",
	"createRolePermission":             "rolepermission",
	"getCloudIdentifier":               "cloudidentifier",
	"getKubernetesClusterConfig":       "clusterconfig",
	"getPathForVolume":                 "apipathforvolume",
	"createConsoleEndpoint":            "consoleendpoint",
	"addVmwareDc":                      "vmwaredc",
	"updateVmwareDc":                   "vmwaredc",
	"createProjectRole":                "projectrole",
	"updateProjectRole":                "projectrole",
	"registerUserData":                 "userdata",
	"updateSecurityGroup":              "securitygroup",
	"updateOauthProvider":              "oauthprovider",
	"readyForShutdown":                 "readyforshutdown",
	"updateObjectStoragePool":          "objectstore",
	"addObjectStoragePool":             "objectstore",
	"updateImageStore":                 "imagestore",
	"linkUserDataToTemplate":           "template",
	"assignVolume":                     "volume",
	"createVMSchedule":                 "vmschedule",
	"updateVMSchedule":                 "vmschedule",
	"setupUserTwoFactorAuthentication": "setup2fa",
	"updateSecondaryStorageSelector":   "heuristics",
	"createSecondaryStorageSelector":   "heuristics",
}

// longToStringConvertedParams is a prefilled map with the list of
// response fields that migrated from long to string within
// the current major baseline. This fields will be parsed from
// json as string and then fallback on long.
var longToStringConvertedParams = map[string]bool{
	"managementserverid": true,
}

// customResponseStructTypes maps the API call to a custom struct name
// This is to change the struct type name to something other than the API name
var customResponseStructTypes = map[string]string{
	"findHostsForMigration": "HostForMigration",
}

// We prefill this one value to make sure it is not
// created twice, as this is also a top level type.
var typeNames = map[string]bool{"Nic": true}

type apiInfo map[string][]string

type allServices struct {
	services services
}

type apiInfoNotFoundError struct {
	api string
}

func (e *apiInfoNotFoundError) Error() string {
	return fmt.Sprintf("Could not find API details for: %s", e.api)
}

type generateError struct {
	service *service
	error   error
}

func (e *generateError) Error() string {
	return fmt.Sprintf("API %s failed to generate code: %v", e.service.name, e.error)
}

type goimportError struct {
	output string
	error  error
}

func (e *goimportError) Error() string {
	return fmt.Sprintf("GoImport failed to format:\n%v\n%v", e.output, e.error)
}

type service struct {
	name string
	apis []*API

	p  func(format string, args ...interface{}) // print raw
	pn func(format string, args ...interface{}) // print with indent and newline
}

type services []*service

// Add functions for the Sort interface
func (s services) Len() int {
	return len(s)
}

func (s services) Less(i, j int) bool {
	return s[i].name < s[j].name
}

func (s services) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// APIParams represents a list of API params
type APIParams []*APIParam

// Len implements the Sort interface
func (s APIParams) Len() int {
	return len(s)
}

// Less implements the Sort interface
func (s APIParams) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// Swap implements the Sort interface
func (s APIParams) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// API represents an API endpoint we can call
type API struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Isasync     bool         `json:"isasync"`
	Params      APIParams    `json:"params"`
	Response    APIResponses `json:"response"`
}

// APIParam represents a single API parameter
type APIParam struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Required    bool   `json:"required"`
}

// APIResponse represents a API response
type APIResponse struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Type        string       `json:"type"`
	Response    APIResponses `json:"response"`
}

// APIResponses represents a list of API responses
type APIResponses []*APIResponse

// Len implements the Sort interface
func (s APIResponses) Len() int {
	return len(s)
}

// Less implements the Sort interface
func (s APIResponses) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// Swap implements the Sort interface
func (s APIResponses) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	listApis := flag.String("api", "listApis.json", "path to the saved JSON output of listApis")
	flag.Parse()

	as, errors, err := getAllServices(*listApis)
	if err != nil {
		log.Fatal(err)
	}

	if err = as.WriteGeneralCode(); err != nil {
		log.Fatal(err)
	}

	for _, s := range as.services {
		if err = s.WriteGeneratedCode(); err != nil {
			errors = append(errors, &generateError{s, err})
		}
	}

	outdir, err := sourceDir()
	if err != nil {
		log.Fatal(err)
	}
	out, err := exec.Command("goimports", "-w", outdir).CombinedOutput()
	if err != nil {
		errors = append(errors, &goimportError{string(out), err})
	}

	testdir, err := testDir()
	if err != nil {
		log.Fatal(err)
	}
	out, err = exec.Command("goimports", "-w", testdir).CombinedOutput()
	if err != nil {
		errors = append(errors, &goimportError{string(out), err})
	}

	if len(errors) > 0 {
		log.Printf("%d API(s) failed to generate:", len(errors))
		for _, ce := range errors {
			log.Print(ce.Error())
		}
		os.Exit(1)
	}
}

func (as *allServices) WriteGeneralCode() error {
	outdir, err := sourceDir()
	if err != nil {
		log.Fatalf("Failed to get source dir: %s", err)
	}

	code, err := as.GeneralCode()
	if err != nil {
		return err
	}

	file := path.Join(outdir, "cloudstack.go")
	return ioutil.WriteFile(file, code, 0644)
}

func (as *allServices) GeneralCode() ([]byte, error) {
	// Buffer the output in memory, for gofmt'ing later in the defer.
	var buf bytes.Buffer
	p := func(format string, args ...interface{}) {
		_, err := fmt.Fprintf(&buf, format, args...)
		if err != nil {
			panic(err)
		}
	}
	pn := func(format string, args ...interface{}) {
		p(format+"\n", args...)
	}
	pn("//")
	pn("// Licensed to the Apache Software Foundation (ASF) under one")
	pn("// or more contributor license agreements.  See the NOTICE file")
	pn("// distributed with this work for additional information")
	pn("// regarding copyright ownership.  The ASF licenses this file")
	pn("// to you under the Apache License, Version 2.0 (the")
	pn("// \"License\"); you may not use this file except in compliance")
	pn("// with the License.  You may obtain a copy of the License at")
	pn("//")
	pn("//   http://www.apache.org/licenses/LICENSE-2.0")
	pn("//")
	pn("// Unless required by applicable law or agreed to in writing,")
	pn("// software distributed under the License is distributed on an")
	pn("// \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY")
	pn("// KIND, either express or implied.  See the License for the")
	pn("// specific language governing permissions and limitations")
	pn("// under the License.")
	pn("//")
	pn("")
	pn("package %s", pkg)
	pn("")
	pn("// UnlimitedResourceID is a special ID to define an unlimited resource")
	pn("const UnlimitedResourceID = \"-1\"")
	pn("")
	pn("var idRegex = regexp.MustCompile(`^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|-1)$`)")
	pn("")
	pn("// IsID return true if the passed ID is either a UUID or a UnlimitedResourceID")
	pn("func IsID(id string) bool {")
	pn("	return idRegex.MatchString(id)")
	pn("}")
	pn("")
	pn("// ClientOption can be passed to new client functions to set custom options")
	pn("type ClientOption func(*CloudStackClient)")
	pn("")
	pn("// OptionFunc can be passed to the courtesy helper functions to set additional parameters")
	pn("type OptionFunc func(*CloudStackClient, interface{}) error")
	pn("")
	pn("type CSError struct {")
	pn("	ErrorCode   int    `json:\"errorcode\"`")
	pn("	CSErrorCode int    `json:\"cserrorcode\"`")
	pn("	ErrorText   string `json:\"errortext\"`")
	pn("}")
	pn("")
	pn("func (e *CSError) Error() error {")
	pn("	return fmt.Errorf(\"CloudStack API error %%d (CSExceptionErrorCode: %%d): %%s\", e.ErrorCode, e.CSErrorCode, e.ErrorText)")
	pn("}")
	pn("")

	pn("type UUID string")
	pn("")
	pn("func (c UUID) MarshalJSON() ([]byte, error) {")
	pn("	return json.Marshal(string(c))")
	pn("}")
	pn("")
	pn("func (c *UUID) UnmarshalJSON(data []byte) error {")
	pn("	value := strings.Trim(string(data), \"\\\"\")")
	pn("	if strings.HasPrefix(string(data), \"\\\"\") {")
	pn("		*c = UUID(value)")
	pn("	  return nil")
	pn("	}")
	pn("	_, err := strconv.ParseInt(value, 10, 64)")
	pn("	if err != nil {")
	pn("		return err")
	pn("	}")
	pn("	*c = UUID(value)")
	pn("	return nil")
	pn("}")

	pn("type CloudStackClient struct {")
	pn("	HTTPGETOnly bool // If `true` only use HTTP GET calls")
	pn("")
	pn("	client  *http.Client // The http client for communicating")
	pn("	baseURL string       // The base URL of the API")
	pn("	apiKey  string       // Api key")
	pn("	secret  string       // Secret key")
	pn("	async   bool         // Wait for async calls to finish")
	pn("	options []OptionFunc // A list of option functions to apply to all API calls")
	pn("	timeout int64        // Max waiting timeout in seconds for async jobs to finish; defaults to 300 seconds")
	pn("")
	for _, s := range as.services {
		pn("  %s %sIface", strings.TrimSuffix(s.name, "Service"), s.name)
	}
	pn("}")
	pn("")
	pn("// Creates a new client for communicating with CloudStack")
	pn("func newClient(apiurl string, apikey string, secret string, async bool, verifyssl bool, options ...ClientOption) *CloudStackClient {")
	pn("	jar, _ := cookiejar.New(nil)")
	pn("	cs := &CloudStackClient{")
	pn("		client: &http.Client{")
	pn("			Jar: jar,")
	pn("			Transport: &http.Transport{")
	pn("				Proxy: http.ProxyFromEnvironment,")
	pn("				DialContext: (&net.Dialer{")
	pn("					Timeout:   30 * time.Second,")
	pn("					KeepAlive: 30 * time.Second,")
	pn("					DualStack: true,")
	pn("				}).DialContext,")
	pn("				MaxIdleConns:          100,")
	pn("				IdleConnTimeout:       90 * time.Second,")
	pn("				TLSClientConfig:       &tls.Config{InsecureSkipVerify: !verifyssl},")
	pn("				TLSHandshakeTimeout:   10 * time.Second,")
	pn("				ExpectContinueTimeout: 1 * time.Second,")
	pn("			},")
	pn("		Timeout: time.Duration(60 * time.Second),")
	pn("		},")
	pn("		baseURL: apiurl,")
	pn("		apiKey:  apikey,")
	pn("		secret:  secret,")
	pn("		async:   async,")
	pn("		options: []OptionFunc{},")
	pn("		timeout: 300,")
	pn("	}")
	pn("")
	pn("	for _, fn := range options {")
	pn("		fn(cs)")
	pn("	}")
	pn("")
	for _, s := range as.services {
		pn("	cs.%s = New%s(cs)", strings.TrimSuffix(s.name, "Service"), s.name)
	}
	pn("")
	pn("	return cs")
	pn("}")
	pn("")
	pn("// Creates a new mock client for communicating with CloudStack")
	pn("func newMockClient(ctrl *gomock.Controller) *CloudStackClient {")
	pn("	cs := &CloudStackClient{}")
	pn("")
	for _, s := range as.services {
		pn("	cs.%s = NewMock%sIface(ctrl)", strings.TrimSuffix(s.name, "Service"), s.name)
	}
	pn("")
	pn("	return cs")
	pn("}")
	pn("")

	pn("// Default non-async client. So for async calls you need to implement and check the async job result yourself. When using")
	pn("// HTTPS with a self-signed certificate to connect to your CloudStack API, you would probably want to set 'verifyssl' to")
	pn("// false so the call ignores the SSL errors/warnings.")
	pn("func NewClient(apiurl string, apikey string, secret string, verifyssl bool, options ...ClientOption) *CloudStackClient {")
	pn("	cs := newClient(apiurl, apikey, secret, false, verifyssl, options...)")
	pn("	return cs")
	pn("}")
	pn("")
	pn("// For sync API calls this client behaves exactly the same as a standard client call, but for async API calls")
	pn("// this client will wait until the async job is finished or until the configured AsyncTimeout is reached. When the async")
	pn("// job finishes successfully it will return actual object received from the API and nil, but when the timeout is")
	pn("// reached it will return the initial object containing the async job ID for the running job and a warning.")
	pn("func NewAsyncClient(apiurl string, apikey string, secret string, verifyssl bool, options ...ClientOption) *CloudStackClient {")
	pn("	cs := newClient(apiurl, apikey, secret, true, verifyssl, options...)")
	pn("	return cs")
	pn("}")
	pn("")
	pn("// Creates a new mock client for communicating with CloudStack")
	pn("func NewMockClient(ctrl *gomock.Controller) *CloudStackClient {")
	pn("	cs := newMockClient(ctrl)")
	pn("	return cs")
	pn("}")
	pn("")
	pn("// When using the async client an api call will wait for the async call to finish before returning. The default is to poll for 300 seconds")
	pn("// seconds, to check if the async job is finished.")
	pn("func (cs *CloudStackClient) AsyncTimeout(timeoutInSeconds int64) {")
	pn("	cs.timeout = timeoutInSeconds")
	pn("}")
	pn("")
	pn("// Sets timeout when using sync api calls. Default is 60 seconds")
	pn("func (cs *CloudStackClient) Timeout(timeout time.Duration) {")
	pn("	cs.client.Timeout = timeout")
	pn("}")
	pn("")
	pn("// Set any default options that would be added to all API calls that support it.")
	pn("func (cs *CloudStackClient) DefaultOptions(options ...OptionFunc) {")
	pn("	if options != nil {")
	pn("		cs.options = options")
	pn("	} else {")
	pn("		cs.options = []OptionFunc{}")
	pn("	}")
	pn("}")
	pn("")
	pn("var AsyncTimeoutErr = errors.New(\"Timeout while waiting for async job to finish\")")
	pn("")
	pn("// A helper function that you can use to get the result of a running async job. If the job is not finished within the configured")
	pn("// timeout, the async job returns a AsyncTimeoutErr.")
	pn("func (cs *CloudStackClient) GetAsyncJobResult(jobid string, timeout int64) (json.RawMessage, error) {")
	pn("	var timer time.Duration")
	pn("	currentTime := time.Now().Unix()")
	pn("")
	pn("		for {")
	pn("		p := cs.Asyncjob.NewQueryAsyncJobResultParams(jobid)")
	pn("		r, err := cs.Asyncjob.QueryAsyncJobResult(p)")
	pn("		if err != nil {")
	pn("			return nil, err")
	pn("		}")
	pn("")
	pn("		// Status 1 means the job is finished successfully")
	pn("		if r.Jobstatus == 1 {")
	pn("			return r.Jobresult, nil")
	pn("		}")
	pn("")
	pn("		// When the status is 2, the job has failed")
	pn("		if r.Jobstatus == 2 {")
	pn("			if r.Jobresulttype == \"text\" {")
	pn("				return nil, fmt.Errorf(string(r.Jobresult))")
	pn("			} else {")
	pn("				return nil, fmt.Errorf(\"Undefined error: %%s\", string(r.Jobresult))")
	pn("			}")
	pn("		}")
	pn("")
	pn("		if time.Now().Unix()-currentTime > timeout {")
	pn("			return nil, AsyncTimeoutErr")
	pn("		}")
	pn("")
	pn("		// Add an (extremely simple) exponential backoff like feature to prevent")
	pn("		// flooding the CloudStack API")
	pn("		if timer < 15 {")
	pn("			timer++")
	pn("		}")
	pn("")
	pn("		time.Sleep(timer * time.Second)")
	pn("	}")
	pn("}")
	pn("")
	pn("// Execute the request against a CS API. Will return the raw JSON data returned by the API and nil if")
	pn("// no error occurred. If the API returns an error the result will be nil and the HTTP error code and CS")
	pn("// error details. If a processing (code) error occurs the result will be nil and the generated error")
	pn("func (cs *CloudStackClient) newRequest(api string, params url.Values) (json.RawMessage, error) {")
	pn("		return cs.newRawRequest(api, false, params)")
	pn("}")
	pn("")
	pn("// Execute the request against a CS API using POST. Will return the raw JSON data returned by the API and")
	pn("// nil if no error occurred. If the API returns an error the result will be nil and the HTTP error code")
	pn("// and CS error details. If a processing (code) error occurs the result will be nil and the generated error")
	pn("func (cs *CloudStackClient) newPostRequest(api string, params url.Values) (json.RawMessage, error) {")
	pn("		return cs.newRawRequest(api, true, params)")
	pn("}")
	pn("")
	pn("// Execute a raw request against a CS API. Will return the raw JSON data returned by the API and nil if")
	pn("// no error occurred. If the API returns an error the result will be nil and the HTTP error code and CS")
	pn("// error details. If a processing (code) error occurs the result will be nil and the generated error")
	pn("func (cs *CloudStackClient) newRawRequest(api string, post bool, params url.Values) (json.RawMessage, error) {")
	pn("	params.Set(\"apiKey\", cs.apiKey)")
	pn("	params.Set(\"command\", api)")
	pn("	params.Set(\"response\", \"json\")")
	pn("	params.Set(\"signatureversion\", \"3\")")
	pn("	params.Set(\"expires\", time.Now().UTC().Add(15*time.Minute).Format(time.RFC3339))")
	pn("")
	pn("	// Generate signature for API call")
	pn("	// * Serialize parameters, URL encoding only values and sort them by key, done by EncodeValues")
	pn("	// * Convert the entire argument string to lowercase")
	pn("	// * Replace all instances of '+' to '%%20'")
	pn("	// * Calculate HMAC SHA1 of argument string with CloudStack secret")
	pn("	// * URL encode the string and convert to base64")
	pn("	s := EncodeValues(params)")
	pn("	s2 := strings.ToLower(s)")
	pn("	mac := hmac.New(sha1.New, []byte(cs.secret))")
	pn("	mac.Write([]byte(s2))")
	pn("	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))")
	pn("")
	pn("	var err error")
	pn("	var resp *http.Response")
	pn("	if !cs.HTTPGETOnly && post {")
	pn("		// The deployVirtualMachine API should be called using a POST call")
	pn("  	// so we don't have to worry about the userdata size")
	pn("")
	pn("		// Add the unescaped signature to the POST params")
	pn("		params.Set(\"signature\", signature)")
	pn("")
	pn("		// Make a POST call")
	pn("		resp, err = cs.client.PostForm(cs.baseURL, params)")
	pn("	} else {")
	pn("		// Create the final URL before we issue the request")
	pn("		url := cs.baseURL + \"?\" + s + \"&signature=\" + url.QueryEscape(signature)")
	pn("")
	pn("		// Make a GET call")
	pn("		resp, err = cs.client.Get(url)")
	pn("	}")
	pn("	if err != nil {")
	pn("		return nil, err")
	pn("	}")
	pn("	defer resp.Body.Close()")
	pn("")
	pn("	b, err := ioutil.ReadAll(resp.Body)")
	pn("	if err != nil {")
	pn("		return nil, err")
	pn("	}")
	pn("")
	pn("	// Need to get the raw value to make the result play nice")
	pn("	b, err = getRawValue(b)")
	pn("	if err != nil {")
	pn("		return nil, err")
	pn("	}")
	pn("")
	pn("	if resp.StatusCode != 200 {")
	pn("		var e CSError")
	pn("		if err := json.Unmarshal(b, &e); err != nil {")
	pn("			return nil, err")
	pn("		}")
	pn("		return nil, e.Error()")
	pn("	}")
	pn("	return b, nil")
	pn("}")
	pn("")
	pn("// Custom version of net/url Encode that only URL escapes values")
	pn("// Unmodified portions here remain under BSD license of The Go Authors: https://go.googlesource.com/go/+/master/LICENSE")
	pn("func EncodeValues(v url.Values) string {")
	pn("	if v == nil {")
	pn("		return \"\"")
	pn("	}")
	pn("	var buf bytes.Buffer")
	pn("	keys := make([]string, 0, len(v))")
	pn("	for k := range v {")
	pn("		keys = append(keys, k)")
	pn("	}")
	pn("	sort.Strings(keys)")
	pn("	for _, k := range keys {")
	pn("		vs := v[k]")
	pn("		prefix := k + \"=\"")
	pn("		for _, v := range vs {")
	pn("			if buf.Len() > 0 {")
	pn("				buf.WriteByte('&')")
	pn("			}")
	pn("			buf.WriteString(prefix)")
	pn("			escaped := url.QueryEscape(v)")
	pn("			// we need to ensure + (representing a space) is encoded as %%20")
	pn("			escaped = strings.Replace(escaped, \"+\", \"%%20\", -1)")
	pn("			// we need to ensure * is not escaped")
	pn("			escaped = strings.Replace(escaped, \"%%2A\", \"*\", -1)")
	pn("			buf.WriteString(escaped)")
	pn("		}")
	pn("	}")
	pn("	return buf.String()")
	pn("}")
	pn("")
	pn("// Generic function to get the first non-count raw value from a response as json.RawMessage")
	pn("func getRawValue(b json.RawMessage) (json.RawMessage, error) {")
	pn("	var m map[string]json.RawMessage")
	pn("	if err := json.Unmarshal(b, &m); err != nil {")
	pn("		return nil, err")
	pn("	}")
	pn("	getArrayResponse := false")
	pn("	for k := range m {")
	pn("		if k == \"count\" {")
	pn("			getArrayResponse = true")
	pn("		}")
	pn("	}")
	pn("	if getArrayResponse {")
	pn("		var resp []json.RawMessage")
	pn("		for k, v := range m {")
	pn("			if k != \"count\" {")
	pn("				if err := json.Unmarshal(v, &resp); err != nil {")
	pn("					return nil, err")
	pn("				}")
	pn("				return resp[0], nil")
	pn("			}")
	pn("		}")
	pn("")
	pn("	} else {")
	pn("		for _, v := range m {")
	pn("			return v, nil")
	pn("		}")
	pn("	}")
	pn("	return nil, fmt.Errorf(\"Unable to extract the raw value from:\\n\\n%%s\\n\\n\", string(b))")
	pn("}")
	pn("")
	pn("// getSortedKeysFromMap returns the keys from m in increasing order.")
	pn("func getSortedKeysFromMap(m map[string]string) (keys []string) {")
	pn("	for k := range m {")
	pn("		keys = append(keys, k)")
	pn("	}")
	pn("	sort.Strings(keys)")
	pn("	return keys")
	pn("}")
	pn("")
	pn("// WithAsyncTimeout takes a custom timeout to be used by the CloudStackClient")
	pn("func WithAsyncTimeout(timeout int64) ClientOption {")
	pn("	return func(cs *CloudStackClient) {")
	pn("		if timeout != 0 {")
	pn("			cs.timeout = timeout")
	pn("		}")
	pn("	}")
	pn("}")
	pn("")
	pn("// DomainIDSetter is an interface that every type that can set a domain ID must implement")
	pn("type DomainIDSetter interface {")
	pn("	SetDomainid(string)")
	pn("}")
	pn("")
	pn("// WithDomain takes either a domain name or ID and sets the `domainid` parameter")
	pn("func WithDomain(domain string) OptionFunc {")
	pn("	return func(cs *CloudStackClient, p interface{}) error {")
	pn("		ps, ok := p.(DomainIDSetter)")
	pn("")
	pn(" 		if !ok || domain == \"\" {")
	pn("			return nil")
	pn("		}")
	pn("")
	pn(" 		if !IsID(domain) {")
	pn("			id, _, err := cs.Domain.GetDomainID(domain)")
	pn("			if err != nil {")
	pn("				return err")
	pn("			}")
	pn("			domain = id")
	pn("		}")
	pn("")
	pn(" 		ps.SetDomainid(domain)")
	pn("")
	pn(" 		return nil")
	pn("	}")
	pn("}")
	pn("")
	pn("// WithHTTPClient takes a custom HTTP client to be used by the CloudStackClient")
	pn("func WithHTTPClient(client *http.Client) ClientOption {")
	pn("	return func(cs *CloudStackClient) {")
	pn("		if client != nil {")
	pn("			if client.Jar == nil {")
	pn("				client.Jar = cs.client.Jar")
	pn("			}")
	pn("			cs.client = client")
	pn("		}")
	pn("	}")
	pn("}")
	pn("")
	pn("// ListallSetter is an interface that every type that can set listall must implement")
	pn("type ListallSetter interface {")
	pn("	SetListall(bool)")
	pn("}")
	pn("")
	pn("// WithListall takes either a project name or ID and sets the `listall` parameter")
	pn("func WithListall(listall bool) OptionFunc {")
	pn("	return func(cs *CloudStackClient, p interface{}) error {")
	pn("		ps, ok := p.(ListallSetter)")
	pn("")
	pn("		if !ok {")
	pn("			return nil")
	pn("		}")
	pn("")
	pn("		ps.SetListall(listall)")
	pn("")
	pn("		return nil")
	pn("	}")
	pn("}")
	pn("// ProjectIDSetter is an interface that every type that can set a project ID must implement")
	pn("type ProjectIDSetter interface {")
	pn("	SetProjectid(string)")
	pn("}")
	pn("")
	pn("// WithProject takes either a project name or ID and sets the `projectid` parameter")
	pn("func WithProject(project string) OptionFunc {")
	pn("	return func(cs *CloudStackClient, p interface{}) error {")
	pn("		ps, ok := p.(ProjectIDSetter)")
	pn("")
	pn("		if !ok || project == \"\" {")
	pn("			return nil")
	pn("		}")
	pn("")
	pn("		if !IsID(project) {")
	pn("			id, _, err := cs.Project.GetProjectID(project)")
	pn("			if err != nil {")
	pn("				return err")
	pn("			}")
	pn("			project = id")
	pn("		}")
	pn("")
	pn("		ps.SetProjectid(project)")
	pn("")
	pn("		return nil")
	pn("	}")
	pn("}")
	pn("")
	pn("// VPCIDSetter is an interface that every type that can set a vpc ID must implement")
	pn("type VPCIDSetter interface {")
	pn("	SetVpcid(string)")
	pn("}")
	pn("")
	pn("// WithVPCID takes a vpc ID and sets the `vpcid` parameter")
	pn("func WithVPCID(id string) OptionFunc {")
	pn("	return func(cs *CloudStackClient, p interface{}) error {")
	pn("		vs, ok := p.(VPCIDSetter)")
	pn("")
	pn("		if !ok || id == \"\" {")
	pn("			return nil")
	pn("		}")
	pn("")
	pn("		vs.SetVpcid(id)")
	pn("")
	pn("		return nil")
	pn("	}")
	pn("}")
	pn("")
	pn("// ZoneIDSetter is an interface that every type that can set a zone ID must implement")
	pn("type ZoneIDSetter interface {")
	pn("	SetZoneid(string)")
	pn("}")
	pn("")
	pn("// WithZone takes either a zone name or ID and sets the `zoneid` parameter")
	pn("func WithZone(zone string) OptionFunc {")
	pn("	return func(cs *CloudStackClient, p interface{}) error {")
	pn("		zs, ok := p.(ZoneIDSetter)")
	pn("")
	pn("		if !ok || zone == \"\" {")
	pn("			return nil")
	pn("		}")
	pn("")
	pn("		if !IsID(zone) {")
	pn("			id, _, err := cs.Zone.GetZoneID(zone)")
	pn("			if err != nil {")
	pn("				return err")
	pn("			}")
	pn("			zone = id")
	pn("		}")
	pn("")
	pn("		zs.SetZoneid(zone)")
	pn("")
	pn("		return nil")
	pn("	}")
	pn("}")
	pn("")
	for _, s := range as.services {
		pn("type %s struct {", s.name)
		pn("  cs *CloudStackClient")
		pn("}")
		pn("")
		pn("func New%s(cs *CloudStackClient) %sIface {", s.name, s.name)
		pn("	return &%s{cs: cs}", s.name)
		pn("}")
		pn("")
	}

	clean, err := format.Source(buf.Bytes())
	if err != nil {
		return buf.Bytes(), err
	}
	return clean, err
}

func (s *service) WriteGeneratedCode() error {
	outdir, err := sourceDir()
	if err != nil {
		log.Fatalf("Failed to get source dir: %s", err)
	}

	code, err := s.GenerateCode()
	if err != nil {
		return err
	}

	if s.name != "CustomService" {
		tests, err := s.GenerateTestCode()
		if err != nil {
			return err
		}
		testdir, err := testDir()
		file := path.Join(testdir, s.name+"_test.go")
		ioutil.WriteFile(file, tests, 0644)
	}

	file := path.Join(outdir, s.name+".go")
	return ioutil.WriteFile(file, code, 0644)
}

func (s *service) GenerateCode() ([]byte, error) {
	// Buffer the output in memory, for gofmt'ing later in the defer.
	var buf bytes.Buffer
	s.p = func(format string, args ...interface{}) {
		_, err := fmt.Fprintf(&buf, format, args...)
		if err != nil {
			panic(err)
		}
	}
	s.pn = func(format string, args ...interface{}) {
		s.p(format+"\n", args...)
	}
	pn := s.pn

	pn("//")
	pn("// Licensed to the Apache Software Foundation (ASF) under one")
	pn("// or more contributor license agreements.  See the NOTICE file")
	pn("// distributed with this work for additional information")
	pn("// regarding copyright ownership.  The ASF licenses this file")
	pn("// to you under the Apache License, Version 2.0 (the")
	pn("// \"License\"); you may not use this file except in compliance")
	pn("// with the License.  You may obtain a copy of the License at")
	pn("//")
	pn("//   http://www.apache.org/licenses/LICENSE-2.0")
	pn("//")
	pn("// Unless required by applicable law or agreed to in writing,")
	pn("// software distributed under the License is distributed on an")
	pn("// \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY")
	pn("// KIND, either express or implied.  See the License for the")
	pn("// specific language governing permissions and limitations")
	pn("// under the License.")
	pn("//")
	pn("")
	pn("package %s", pkg)
	pn("")
	if s.name == "FirewallService" {
		pn("// Helper function for maintaining backwards compatibility")
		pn("func convertFirewallServiceResponse(b []byte) ([]byte, error) {")
		pn("	var raw map[string]interface{}")
		pn("	if err := json.Unmarshal(b, &raw); err != nil {")
		pn("		return nil, err")
		pn("	}")
		pn("")
		pn("	if _, ok := raw[\"firewallrule\"]; ok {")
		pn("		return convertFirewallServiceListResponse(b)")
		pn("	}")
		pn("")
		pn("	for _, k := range []string{\"endport\", \"startport\"} {")
		pn("		if sVal, ok := raw[k].(string); ok {")
		pn("			iVal, err := strconv.Atoi(sVal)")
		pn("			if err != nil {")
		pn("				return nil, err")
		pn("			}")
		pn("			raw[k] = iVal")
		pn("		}")
		pn("	}")
		pn("")
		pn("	return json.Marshal(raw)")
		pn("}")
		pn("")
		pn("// Helper function for maintaining backwards compatibility")
		pn("func convertFirewallServiceListResponse(b []byte) ([]byte, error) {")
		pn("	var rawList struct {")
		pn("		Count         int                      `json:\"count\"`")
		pn("		FirewallRules []map[string]interface{} `json:\"firewallrule\"`")
		pn("	}")
		pn("")
		pn("	if err := json.Unmarshal(b, &rawList); err != nil {")
		pn("		return nil, err")
		pn("	}")
		pn("")
		pn("	for _, r := range rawList.FirewallRules {")
		pn("		for _, k := range []string{\"endport\", \"startport\"} {")
		pn("			if sVal, ok := r[k].(string); ok {")
		pn("				iVal, err := strconv.Atoi(sVal)")
		pn("				if err != nil {")
		pn("					return nil, err")
		pn("				}")
		pn("				r[k] = iVal")
		pn("			}")
		pn("		}")
		pn("	}")
		pn("")
		pn("	return json.Marshal(rawList)")
		pn("}")
		pn("")
	}
	if s.name == "SecurityGroupService" {
		pn("// Helper function for maintaining backwards compatibility")
		pn("func convertAuthorizeSecurityGroupIngressResponse(b []byte) ([]byte, error) {")
		pn("	var raw struct {")
		pn("		Ingressrule []interface{} `json:\"ingressrule\"`")
		pn("	}")
		pn("	if err := json.Unmarshal(b, &raw); err != nil {")
		pn("		return nil, err")
		pn("	}")
		pn("")
		pn("	if len(raw.Ingressrule) != 1 {")
		pn("		return b, nil")
		pn("	}")
		pn("")
		pn("	return json.Marshal(raw.Ingressrule[0])")
		pn("}")
		pn("")
		pn("// Helper function for maintaining backwards compatibility")
		pn("func convertAuthorizeSecurityGroupEgressResponse(b []byte) ([]byte, error) {")
		pn("	var raw struct {")
		pn("		Egressrule []interface{} `json:\"egressrule\"`")
		pn("	}")
		pn("	if err := json.Unmarshal(b, &raw); err != nil {")
		pn("		return nil, err")
		pn("	}")
		pn("")
		pn("	if len(raw.Egressrule) != 1 {")
		pn("		return b, nil")
		pn("	}")
		pn("")
		pn("	return json.Marshal(raw.Egressrule[0])")
		pn("}")
		pn("")
	}
	if s.name == "CustomService" {
		pn("type CustomServiceParams struct {")
		pn("	p map[string]interface{}")
		pn("}")
		pn("")
		pn("func (p *CustomServiceParams) toURLValues() url.Values {")
		pn("	u := url.Values{}")
		pn("	if p.p == nil {")
		pn("		return u")
		pn("	}")
		pn("")
		pn("	for k, v := range p.p {")
		pn("		switch t := v.(type) {")
		pn("		case bool:")
		pn("			u.Set(k, strconv.FormatBool(t))")
		pn("		case int:")
		pn("			u.Set(k, strconv.Itoa(t))")
		pn("		case int64:")
		pn("			vv := strconv.FormatInt(t, 10)")
		pn("			u.Set(k, vv)")
		pn("		case string:")
		pn("			u.Set(k, t)")
		pn("		case []string:")
		pn("			u.Set(k, strings.Join(t, \", \"))")
		pn("		case map[string]string:")
		pn("			i := 0")
		pn("			for kk, vv := range t {")
		pn("				u.Set(fmt.Sprintf(\"%%s[%%d].%%s\", k, i, kk), vv)")
		pn("				i++")
		pn("			}")
		pn("		}")
		pn("	}")
		pn("")
		pn("	return u")
		pn("}")
		pn("")
		pn("func (p *CustomServiceParams) SetParam(param string, v interface{}) {")
		pn("	if p.p == nil {")
		pn("		p.p = make(map[string]interface{})")
		pn("	}")
		pn("	p.p[param] = v")
		pn("}")
		pn("func (p *CustomServiceParams) GetParam(param string) (interface{}, bool) {")
		pn("	if p.p == nil {")
		pn("		p.p = make(map[string]interface{})")
		pn("	}")
		pn("	value, ok := p.p[param].(interface{})")
		pn("	return value, ok")
		pn("}")
		pn("")
		pn("func (s *CustomService) CustomRequest(api string, p *CustomServiceParams, result interface{}) error {")
		pn("	resp, err := s.cs.newPostRequest(api, p.toURLValues())")
		pn("	if err != nil {")
		pn("		return err")
		pn("	}")
		pn("")
		pn("	return json.Unmarshal(resp, result)")
		pn("}")
		pn("func (s *CustomService) CustomPostRequest(api string, p *CustomServiceParams, result interface{}) error {")
		pn("	resp, err := s.cs.newPostRequest(api, p.toURLValues())")
		pn("	if err != nil {")
		pn("		return err")
		pn("	}")
		pn("")
		pn("	return json.Unmarshal(resp, result)")
		pn("}")
	}

	s.generateInterfaceType()

	for _, a := range s.apis {
		s.generateParamType(a)
		s.generateToURLValuesFunc(a)
		s.generateParamGettersAndSettersFunc(a)
		s.generateNewParamTypeFunc(a)
		s.generateHelperFuncs(a)
		s.generateNewAPICallFunc(a)
		s.generateResponseType(a)
	}

	clean, err := format.Source(buf.Bytes())
	if err != nil {
		buf.WriteTo(os.Stdout)
		return buf.Bytes(), err
	}
	return clean, nil
}

func (s *service) GenerateTestCode() ([]byte, error) {
	var buf bytes.Buffer
	s.p = func(format string, args ...interface{}) {
		_, err := fmt.Fprintf(&buf, format, args...)
		if err != nil {
			panic(err)
		}
	}
	s.pn = func(format string, args ...interface{}) {
		s.p(format+"\n", args...)
	}
	pn := s.pn

	pn("//")
	pn("// Licensed to the Apache Software Foundation (ASF) under one")
	pn("// or more contributor license agreements.  See the NOTICE file")
	pn("// distributed with this work for additional information")
	pn("// regarding copyright ownership.  The ASF licenses this file")
	pn("// to you under the Apache License, Version 2.0 (the")
	pn("// \"License\"); you may not use this file except in compliance")
	pn("// with the License.  You may obtain a copy of the License at")
	pn("//")
	pn("//   http://www.apache.org/licenses/LICENSE-2.0")
	pn("//")
	pn("// Unless required by applicable law or agreed to in writing,")
	pn("// software distributed under the License is distributed on an")
	pn("// \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY")
	pn("// KIND, either express or implied.  See the License for the")
	pn("// specific language governing permissions and limitations")
	pn("// under the License.")
	pn("//")
	pn("")
	pn("package test")
	pn("")

	pn("func Test%s(t *testing.T) {", s.name)
	pn("	service := \"%s\"", s.name)
	pn("	response, err := readData(service)")
	pn("	if err != nil {")
	pn("		t.Skipf(\"Skipping test as %%v\", err)")
	pn("	}")
	pn("	server := CreateTestServer(t, response)")
	pn("	client := cloudstack.NewClient(server.URL, \"APIKEY\", \"SECRETKEY\", true)")
	pn("	defer server.Close()")
	pn("")

	for _, a := range s.apis {
		s.generateAPITest(a)
	}
	pn("}")

	clean, err := format.Source(buf.Bytes())
	if err != nil {
		buf.WriteTo(os.Stdout)
		return buf.Bytes(), err
	}
	return clean, nil
}

func (s *service) generateAPITest(a *API) {
	p, pn := s.p, s.pn
	tn := capitalize(a.Name + "Params")
	rp := APIParams{}
	pn("	test%s := func(t *testing.T) {", a.Name)
	pn("		if _, ok := response[\"%s\"]; !ok {", a.Name)
	pn("			t.Skipf(\"Skipping as no json response is provided in testdata\")")
	pn("		}")
	p("		p := client.%s.New%s(", strings.TrimSuffix(s.name, "Service"), tn)
	for _, ap := range a.Params {
		if ap.Required || isRequiredParam(a, ap) {
			rp = append(rp, ap)
			p("%s, ", getDefaultValueForType(a.Name, ap.Name, ap.Type))
		}
	}
	pn(")")
	idPresent := false
	if !(strings.HasPrefix(a.Name, "list") || a.Name == "registerTemplate" || a.Name == "findHostsForMigration" || a.Name == "quotaTariffList") {
		for _, ap := range a.Response {
			if ap.Name == "id" && ap.Type == "string" {
				pn("		r, err := client.%s.%s(p)", strings.TrimSuffix(s.name, "Service"), capitalize(a.Name))
				idPresent = true
				break
			}
		}
	}
	if !idPresent {
		pn("		_, err := client.%s.%s(p)", strings.TrimSuffix(s.name, "Service"), capitalize(a.Name))
	}
	pn("		if err != nil {")
	pn("			t.Errorf(err.Error())")
	pn("		}")
	if idPresent {
		pn("		if r.Id == \"\" {")
		pn("			t.Errorf(\"Failed to parse response. ID not found\")")
		pn("		}")
	}
	pn("	}")
	pn("    t.Run(\"%s\", test%s)", capitalize(a.Name), a.Name)
	pn("")
}

func getDefaultValueForType(aName string, pName string, pType string) string {
	switch pType {
	case "boolean":
		return "true"
	case "short", "int", "integer", "long", "float", "double":
		return "0"
	case "list":
		return "[]string{}"
	case "map":
		return "map[string]string{}"
	default:
		return fmt.Sprintf("\"%s\"", pName)
	}
}

func (s *service) generateParamType(a *API) {
	pn := s.pn

	pn("type %s struct {", capitalize(a.Name+"Params"))
	pn("	p map[string]interface{}")
	pn("}\n")
}

func (s *service) generateInterfaceType() {
	p, pn := s.p, s.pn

	pn("type %sIface interface {", capitalize(s.name))
	for _, api := range s.apis {
		n := capitalize(api.Name)
		tn := capitalize(api.Name + "Params")
		// API Calls
		pn("	%s(p *%s) (*%s, error)", n, n+"Params", strings.TrimPrefix(n, "Configure")+"Response")

		// NewParam funcs
		p("New%s(", tn)
		for _, ap := range api.Params {
			if ap.Required || isRequiredParam(api, ap) {
				p("%s %s, ", s.parseParamName(ap.Name), mapType(api.Name, ap.Name, ap.Type))
			}
		}
		pn(") *%s", tn)

		// Helper funcs
		if strings.HasPrefix(api.Name, "list") {
			v, found := hasNameOrKeywordParamField(api.Name, api.Params)
			if found && hasIDAndNameResponseField(api.Name, api.Response) {
				ln := strings.TrimPrefix(api.Name, "list")

				// Check if ID is a required parameters and bail if so
				for _, ap := range api.Params {
					if ap.Required && ap.Name == "id" {
						return
					}
				}

				// Generate the function signature
				p("Get%sID(%s string, ", parseSingular(ln), v)
				for _, ap := range api.Params {
					if ap.Required || isRequiredParam(api, ap) {
						p("%s %s, ", s.parseParamName(ap.Name), mapType(api.Name, ap.Name, ap.Type))
					}
				}
				if parseSingular(ln) == "Iso" {
					p("isofilter string, ")
				}
				if parseSingular(ln) == "Template" || parseSingular(ln) == "Iso" {
					p("zoneid string, ")
				}
				pn("opts ...OptionFunc) (string, int, error)")

				if hasIDParamField(api.Name, api.Params) {
					p("Get%sByName(name string, ", parseSingular(ln))
					for _, ap := range api.Params {
						if ap.Required || isRequiredParam(api, ap) {
							p("%s %s, ", s.parseParamName(ap.Name), mapType(api.Name, ap.Name, ap.Type))
						}
					}
					if parseSingular(ln) == "Iso" {
						p("isofilter string, ")
					}
					if parseSingular(ln) == "Template" || parseSingular(ln) == "Iso" {
						p("zoneid string, ")
					}
					pn("opts ...OptionFunc) (*%s, int, error)", parseSingular(ln))
				}
			}

			if hasIDParamField(api.Name, api.Params) {
				ln := strings.TrimPrefix(api.Name, "list")

				// Generate the function signature
				p("Get%sByID(id string, ", parseSingular(ln))
				for _, ap := range api.Params {
					if ap.Required && s.parseParamName(ap.Name) != "id" {
						p("%s %s, ", ap.Name, mapType(api.Name, ap.Name, ap.Type))
					}
				}
				if ln == "LoadBalancerRuleInstances" {
					pn("opts ...OptionFunc) (*VirtualMachine, int, error)")
				} else {
					pn("opts ...OptionFunc) (*%s, int, error)", parseSingular(ln))
				}
			}
		}
	}
	pn("}\n")
}

func (s *service) generateToURLValuesFunc(a *API) {
	pn := s.pn

	pn("func (p *%s) toURLValues() url.Values {", capitalize(a.Name+"Params"))
	pn("	u := url.Values{}")
	pn("	if p.p == nil {")
	pn("		return u")
	pn("	}")
	for _, ap := range a.Params {
		pn("	if v, found := p.p[\"%s\"]; found {", ap.Name)
		s.generateConvertCode(a.Name, ap.Name, mapType(a.Name, ap.Name, ap.Type))
		pn("	}")
	}
	pn("	return u")
	pn("}")
	pn("")
}

func (s *service) generateConvertCode(cmd, name, typ string) {
	pn := s.pn

	switch typ {
	case "string", "UUID":
		pn("u.Set(\"%s\", v.(string))", name)
	case "int":
		pn("vv := strconv.Itoa(v.(int))")
		pn("u.Set(\"%s\", vv)", name)
	case "int64":
		pn("vv := strconv.FormatInt(v.(int64), 10)")
		pn("u.Set(\"%s\", vv)", name)
	case "float64":
		pn("vv := strconv.FormatFloat(v.(float64), 'f', -1, 64)")
		pn("u.Set(\"%s\", vv)", name)
	case "bool":
		pn("vv := strconv.FormatBool(v.(bool))")
		pn("u.Set(\"%s\", vv)", name)
	case "[]string":
		pn("vv := strings.Join(v.([]string), \",\")")
		pn("u.Set(\"%s\", vv)", name)
	case "[]map[string]string":
		pn("l := v.([]map[string]string)")
		pn("for i, m := range l {")
		pn("  for key, val := range m {")
		pn("	  u.Set(fmt.Sprintf(\"%s[%%d].%%s\", i, key), val)", name)
		pn("  }")
		pn("}")
	case "map[string]string":
		pn("m := v.(map[string]string)")
		zeroIndex := detailsRequireZeroIndex[cmd]
		if zeroIndex {
			pn("for _, k := range getSortedKeysFromMap(m) {")
		} else {
			pn("for i, k := range getSortedKeysFromMap(m) {")
		}
		switch name {
		case "details":
			if detailsRequireKeyValue[cmd] {
				pn("	u.Set(fmt.Sprintf(\"%s[%%d].key\", i), k)", name)
				pn("	u.Set(fmt.Sprintf(\"%s[%%d].value\", i), m[k])", name)
			} else {
				if zeroIndex {
					pn("	u.Set(fmt.Sprintf(\"%s[0].%%s\", k), m[k])", name)
				} else {
					pn("	u.Set(fmt.Sprintf(\"%s[%%d].%%s\", i, k), m[k])", name)
				}
			}
		case "serviceproviderlist":
			pn("	u.Set(fmt.Sprintf(\"%s[%%d].service\", i), k)", name)
			pn("	u.Set(fmt.Sprintf(\"%s[%%d].provider\", i), m[k])", name)
		case "usersecuritygrouplist":
			pn("	u.Set(fmt.Sprintf(\"%s[%%d].account\", i), k)", name)
			pn("	u.Set(fmt.Sprintf(\"%s[%%d].group\", i), m[k])", name)
		case "tags":
			pn("	u.Set(fmt.Sprintf(\"%s[%%d].key\", i), k)", name)
			if cmd == "deleteTags" {
				pn("	if m[k] != \"\" {")
				pn("		u.Set(fmt.Sprintf(\"%s[%%d].value\", i), m[k])", name)
				pn("	}")
			} else {
				pn("	u.Set(fmt.Sprintf(\"%s[%%d].value\", i), m[k])", name)
			}
		case "nicnetworklist":
			pn("	u.Set(fmt.Sprintf(\"%s[%%d].nic\", i), k)", name)
			pn("	u.Set(fmt.Sprintf(\"%s[%%d].network\", i), m[k])", name)
		case "nicipaddresslist":
			pn("	u.Set(fmt.Sprintf(\"%s[%%d].nic\", i), k)", name)
			pn("	u.Set(fmt.Sprintf(\"%s[%%d].ip4Address\", i), m[k])", name)
		case "datadiskofferinglist":
			pn("	u.Set(fmt.Sprintf(\"%s[%%d].disk\", i), k)", name)
			pn("	u.Set(fmt.Sprintf(\"%s[%%d].diskOffering\", i), m[k])", name)
		default:
			if zeroIndex && !detailsRequireKeyValue[cmd] {
				pn("	u.Set(fmt.Sprintf(\"%s[0].%%s\", k), m[k])", name)
			} else {
				pn("	u.Set(fmt.Sprintf(\"%s[%%d].key\", i), k)", name)
				pn("	u.Set(fmt.Sprintf(\"%s[%%d].value\", i), m[k])", name)
			}
		}
		pn("}")
	}
}

func (s *service) parseParamName(name string) string {
	if name != "type" {
		return name
	}
	return uncapitalize(strings.TrimSuffix(s.name, "Service")) + "Type"
}

func (s *service) generateParamGettersAndSettersFunc(a *API) {
	pn := s.pn
	found := make(map[string]bool)

	for _, ap := range a.Params {
		if !found[ap.Name] {
			pn("func (p *%s) Set%s(v %s) {", capitalize(a.Name+"Params"), capitalize(ap.Name), mapType(a.Name, ap.Name, ap.Type))
			pn("	if p.p == nil {")
			pn("		p.p = make(map[string]interface{})")
			pn("	}")
			pn("	p.p[\"%s\"] = v", ap.Name)
			pn("}")
			pn("")

			pn("func (p *%s) Reset%s() {", capitalize(a.Name+"Params"), capitalize(ap.Name))
			pn("	if p.p != nil && p.p[\"%s\"] != nil {", ap.Name)
			pn("		delete(p.p, \"%s\")", ap.Name)
			pn("	}")
			pn("}")
			pn("")

			pn("func (p *%s) Get%s() (%s, bool) {", capitalize(a.Name+"Params"), capitalize(ap.Name), mapType(a.Name, ap.Name, ap.Type))
			pn("	if p.p == nil {")
			pn("		p.p = make(map[string]interface{})")
			pn("	}")
			pn("	value, ok := p.p[\"%s\"].(%s)", ap.Name, mapType(a.Name, ap.Name, ap.Type))
			pn("	return value, ok")
			pn("}")
			pn("")

			if mapRequireList[a.Name] != nil && mapRequireList[a.Name][ap.Name] {
				pn("func (p *%s) Add%s(item map[string]string) {", capitalize(a.Name+"Params"), capitalize(ap.Name))
				pn("	if p.p == nil {")
				pn("		p.p = make(map[string]interface{})")
				pn("	}")
				pn("	val, found := p.p[\"%s\"]", ap.Name)
				pn("	if !found {")
				pn("	  p.p[\"%s\"] = []map[string]string{}", ap.Name)
				pn("	  val = p.p[\"%s\"]", ap.Name)
				pn("	}")
				pn("	l := val.([]map[string]string)")
				pn("	l = append(l, item)")
				pn("	p.p[\"%s\"] = l", ap.Name)
				pn("}")
				pn("")
			}
			found[ap.Name] = true
		}
	}
}

func (s *service) generateNewParamTypeFunc(a *API) {
	p, pn := s.p, s.pn
	tn := capitalize(a.Name + "Params")
	rp := APIParams{}

	// Generate the function signature
	pn("// You should always use this function to get a new %s instance,", tn)
	pn("// as then you are sure you have configured all required params")
	p("func (s *%s) New%s(", s.name, tn)
	for _, ap := range a.Params {
		if ap.Required || isRequiredParam(a, ap) {
			rp = append(rp, ap)
			p("%s %s, ", s.parseParamName(ap.Name), mapType(a.Name, ap.Name, ap.Type))
		}
	}
	pn(") *%s {", tn)

	// Generate the function body
	pn("	p := &%s{}", tn)
	pn("	p.p = make(map[string]interface{})")
	sort.Sort(rp)
	for _, ap := range rp {
		pn("	p.p[\"%s\"] = %s", ap.Name, s.parseParamName(ap.Name))
	}
	pn("	return p")
	pn("}")
	pn("")
}

func (s *service) generateHelperFuncs(a *API) {
	p, pn := s.p, s.pn

	if strings.HasPrefix(a.Name, "list") {
		v, found := hasNameOrKeywordParamField(a.Name, a.Params)
		if found && hasIDAndNameResponseField(a.Name, a.Response) {
			ln := strings.TrimPrefix(a.Name, "list")

			// Check if ID is a required parameters and bail if so
			for _, ap := range a.Params {
				if ap.Required && ap.Name == "id" {
					return
				}
			}

			// Generate the function signature
			pn("// This is a courtesy helper function, which in some cases may not work as expected!")
			p("func (s *%s) Get%sID(%s string, ", s.name, parseSingular(ln), v)
			for _, ap := range a.Params {
				if ap.Required || isRequiredParam(a, ap) {
					p("%s %s, ", s.parseParamName(ap.Name), mapType(a.Name, ap.Name, ap.Type))
				}
			}
			if parseSingular(ln) == "Iso" {
				p("isofilter string, ")
			}
			if parseSingular(ln) == "Template" || parseSingular(ln) == "Iso" {
				p("zoneid string, ")
			}
			pn("opts ...OptionFunc) (string, int, error) {")

			// Generate the function body
			pn("	p := &List%sParams{}", ln)
			pn("	p.p = make(map[string]interface{})")
			pn("")
			pn("	p.p[\"%s\"] = %s", v, v)
			for _, ap := range a.Params {
				if ap.Required || isRequiredParam(a, ap) {
					pn("	p.p[\"%s\"] = %s", ap.Name, s.parseParamName(ap.Name))
				}
			}
			if parseSingular(ln) == "Iso" {
				pn("	p.p[\"isofilter\"] = isofilter")
			}
			if parseSingular(ln) == "Template" || parseSingular(ln) == "Iso" {
				pn("	p.p[\"zoneid\"] = zoneid")
			}
			pn("")
			pn("	for _, fn := range append(s.cs.options, opts...) {")
			pn("		if err := fn(s.cs, p); err != nil {")
			pn("			return \"\", -1, err")
			pn("		}")
			pn("	}")
			pn("")
			pn("	l, err := s.List%s(p)", ln)
			pn("	if err != nil {")
			pn("		return \"\", -1, err")
			pn("	}")
			pn("")
			if ln == "AffinityGroups" {
				pn("	// This is needed because of a bug with the listAffinityGroup call. It reports the")
				pn("	// number of VirtualMachines in the groups as being the number of groups found.")
				pn("	l.Count = len(l.%s)", ln)
				pn("")
			}
			pn("	if l.Count == 0 {")
			pn("	  return \"\", l.Count, fmt.Errorf(\"No match found for %%s: %%+v\", %s, l)", v)
			pn("	}")
			pn("")
			pn("	if l.Count == 1 {")
			pn("	  return l.%s[0].Id, l.Count, nil", ln)
			pn("	}")
			pn("")
			pn(" 	if l.Count > 1 {")
			pn("    for _, v := range l.%s {", ln)
			pn("      if v.Name == %s {", v)
			pn("        return v.Id, l.Count, nil")
			pn("      }")
			pn("    }")
			pn("	}")
			pn("  return \"\", l.Count, fmt.Errorf(\"Could not find an exact match for %%s: %%+v\", %s, l)", v)
			pn("}\n")
			pn("")

			if hasIDParamField(a.Name, a.Params) {
				// Generate the function signature
				pn("// This is a courtesy helper function, which in some cases may not work as expected!")
				p("func (s *%s) Get%sByName(name string, ", s.name, parseSingular(ln))
				for _, ap := range a.Params {
					if ap.Required || isRequiredParam(a, ap) {
						p("%s %s, ", s.parseParamName(ap.Name), mapType(a.Name, ap.Name, ap.Type))
					}
				}
				if parseSingular(ln) == "Iso" {
					p("isofilter string, ")
				}
				if parseSingular(ln) == "Template" || parseSingular(ln) == "Iso" {
					p("zoneid string, ")
				}
				pn("opts ...OptionFunc) (*%s, int, error) {", parseSingular(ln))

				// Generate the function body
				p("  id, count, err := s.Get%sID(name, ", parseSingular(ln))
				for _, ap := range a.Params {
					if ap.Required || isRequiredParam(a, ap) {
						p("%s, ", s.parseParamName(ap.Name))
					}
				}
				if parseSingular(ln) == "Iso" {
					p("isofilter, ")
				}
				if parseSingular(ln) == "Template" || parseSingular(ln) == "Iso" {
					p("zoneid, ")
				}
				pn("opts...)")
				pn("  if err != nil {")
				pn("    return nil, count, err")
				pn("  }")
				pn("")
				p("  r, count, err := s.Get%sByID(id, ", parseSingular(ln))
				for _, ap := range a.Params {
					if ap.Required || isRequiredParam(a, ap) {
						p("%s, ", s.parseParamName(ap.Name))
					}
				}
				pn("opts...)")
				pn("  if err != nil {")
				pn("    return nil, count, err")
				pn("  }")
				pn("	return r, count, nil")
				pn("}")
				pn("")
			}
		}

		if hasIDParamField(a.Name, a.Params) {
			ln := strings.TrimPrefix(a.Name, "list")

			// Generate the function signature
			pn("// This is a courtesy helper function, which in some cases may not work as expected!")
			p("func (s *%s) Get%sByID(id string, ", s.name, parseSingular(ln))
			for _, ap := range a.Params {
				if ap.Required && s.parseParamName(ap.Name) != "id" {
					p("%s %s, ", ap.Name, mapType(a.Name, ap.Name, ap.Type))
				}
			}
			if ln == "LoadBalancerRuleInstances" {
				pn("opts ...OptionFunc) (*VirtualMachine, int, error) {")
			} else {
				pn("opts ...OptionFunc) (*%s, int, error) {", parseSingular(ln))
			}

			// Generate the function body
			pn("	p := &List%sParams{}", ln)
			pn("	p.p = make(map[string]interface{})")
			pn("")
			pn("	p.p[\"id\"] = id")
			for _, ap := range a.Params {
				if ap.Required && s.parseParamName(ap.Name) != "id" {
					pn("	p.p[\"%s\"] = %s", ap.Name, s.parseParamName(ap.Name))
				}
			}
			pn("")
			pn("	for _, fn := range append(s.cs.options, opts...) {")
			pn("		if err := fn(s.cs, p); err != nil {")
			pn("			return nil, -1, err")
			pn("		}")
			pn("	}")
			pn("")
			pn("	l, err := s.List%s(p)", ln)
			pn("	if err != nil {")
			pn("		if strings.Contains(err.Error(), fmt.Sprintf(")
			pn("			\"Invalid parameter id value=%%s due to incorrect long value format, \"+")
			pn("				\"or entity does not exist\", id)) {")
			pn("			return nil, 0, fmt.Errorf(\"No match found for %%s: %%+v\", id, l)")
			pn("		}")
			pn("		return nil, -1, err")
			pn("	}")
			pn("")
			if ln == "AffinityGroups" {
				pn("	// This is needed because of a bug with the listAffinityGroup call. It reports the")
				pn("	// number of VirtualMachines in the groups as being the number of groups found.")
				pn("	l.Count = len(l.%s)", ln)
				pn("")
			}
			pn("	if l.Count == 0 {")
			pn("	  return nil, l.Count, fmt.Errorf(\"No match found for %%s: %%+v\", id, l)")
			pn("	}")
			pn("")
			pn("	if l.Count == 1 {")
			pn("	  return l.%s[0], l.Count, nil", ln)
			pn("	}")
			pn("  return nil, l.Count, fmt.Errorf(\"There is more then one result for %s UUID: %%s!\", id)", parseSingular(ln))
			pn("}\n")
			pn("")
		}
	}
}

func hasNameOrKeywordParamField(aName string, params APIParams) (v string, found bool) {
	for _, p := range params {
		if p.Name == "keyword" && mapType(aName, p.Name, p.Type) == "string" {
			v = "keyword"
			found = true
		}
		if p.Name == "name" && mapType(aName, p.Name, p.Type) == "string" {
			return "name", true
		}

	}
	return v, found
}

func hasIDParamField(aName string, params APIParams) bool {
	for _, p := range params {
		if p.Name == "id" && mapType(aName, p.Name, p.Type) == "string" {
			return true
		}
	}
	return false
}

func hasIDAndNameResponseField(aName string, resp APIResponses) bool {
	id := false
	name := false

	for _, r := range resp {
		if r.Name == "id" && mapType(aName, r.Name, r.Type) == "string" {
			id = true
		}
		if r.Name == "name" && mapType(aName, r.Name, r.Type) == "string" {
			name = true
		}
	}
	return id && name
}

func (s *service) generateNewAPICallFunc(a *API) {
	pn := s.pn
	n := capitalize(a.Name)

	// Generate the function signature
	pn("// %s", a.Description)
	pn("func (s *%s) %s(p *%s) (*%s, error) {", s.name, n, n+"Params", strings.TrimPrefix(n, "Configure")+"Response")

	// Generate the function body
	if n == "QueryAsyncJobResult" {
		pn("	var resp json.RawMessage")
		pn("	var err error")
		pn("")
		pn("	// We should be able to retry on failure as this call is idempotent")
		pn("	for i := 0; i < 3; i++ {")
		pn("		resp, err = s.cs.newRequest(\"%s\", p.toURLValues())", a.Name)
		pn("		if err == nil {")
		pn("			break")
		pn("		}")
		pn("		time.Sleep(500 * time.Millisecond)")
		pn("	}")
	} else {
		isGetRequest, _ := regexp.MatchString("^(get|list|query|find)(\\w+)+$", strings.ToLower(a.Name))
		getRequestList := map[string]struct{}{"isaccountallowedtocreateofferingswithtags": {}, "readyforshutdown": {}, "cloudianisenabled": {}, "quotabalance": {},
			"quotasummary": {}, "quotatarifflist": {}, "quotaisenabled": {}, "quotastatement": {}, "verifyoauthcodeandgetuser": {}}
		_, isInGetRequestList := getRequestList[strings.ToLower(a.Name)]

		if requiresPostMethod[a.Name] || !(isGetRequest || isInGetRequestList) {
			pn("	resp, err := s.cs.newPostRequest(\"%s\", p.toURLValues())", a.Name)
		} else {
			pn("	resp, err := s.cs.newRequest(\"%s\", p.toURLValues())", a.Name)
		}
	}
	pn("	if err != nil {")
	pn("		return nil, err")
	pn("	}")
	pn("")
	switch n {
	case
		"AddCluster",
		"AddImageStore",
		"CreateAccount",
		"CreateDomain",
		"UpdateDomain",
		"CreateNetwork",
		"CreateStoragePool",
		"CreateNetworkOffering",
		"CreateVlanIpRange",
		"UpdateNetworkOffering",
		"UpdateServiceOffering",
		"UpdateConfiguration",
		"UpdateCluster",
		"UpdateVlanIpRange",
		"CreatePod",
		"CreateSSHKeyPair",
		"CreateSecurityGroup",
		"CreateServiceOffering",
		"CreateUser",
		"CreateZone",
		"DedicateGuestVlanRange",
		"EnableUser",
		"GetVirtualMachineUserData",
		"LockUser",
		"RegisterSSHKeyPair",
		"RegisterUserKeys",
		"GetUserKeys",
		"AddAnnotation",
		"RemoveAnnotation",
		"AddKubernetesSupportedVersion",
		"CreateDiskOffering",
		"AddHost",
		"RegisterIso":
		pn("	if resp, err = getRawValue(resp); err != nil {")
		pn("		return nil, err")
		pn("	}")
		pn("")
	}
	if !a.Isasync && s.name == "FirewallService" {
		pn("		resp, err = convertFirewallServiceResponse(resp)")
		pn("		if err != nil {")
		pn("			return nil, err")
		pn("		}")
		pn("")
	}

	if field, isNested := nestedResponse[a.Name]; isNested {
		pn("	var nested struct {")
		pn("			Response %sResponse `json:\"%s\"`", strings.TrimPrefix(n, "Configure"), field)
		pn("	}")
		pn("	if err := json.Unmarshal(resp, &nested); err != nil {")
		pn("		return nil, err")
		pn("	}")
		pn("	r := nested.Response")
	} else {
		pn("	var r %sResponse", strings.TrimPrefix(n, "Configure"))
		pn("	if err := json.Unmarshal(resp, &r); err != nil {")
		pn("		return nil, err")
		pn("	}")
	}
	pn("")
	if a.Isasync {
		pn("	// If we have a async client, we need to wait for the async result")
		pn("	if s.cs.async {")
		pn("		b, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)")
		pn("		if err != nil {")
		pn("			if err == AsyncTimeoutErr {")
		pn("				return &r, err")
		pn("			}")
		pn("			return nil, err")
		pn("		}")
		pn("")
		if !isSuccessOnlyResponse(a.Response) {
			pn("		b, err = getRawValue(b)")
			pn("		if err != nil {")
			pn("		  return nil, err")
			pn("		}")
			pn("")
		}
		if s.name == "FirewallService" {
			pn("		b, err = convertFirewallServiceResponse(b)")
			pn("		if err != nil {")
			pn("			return nil, err")
			pn("		}")
			pn("")
		}
		if n == "AuthorizeSecurityGroupIngress" {
			pn("		b, err = convertAuthorizeSecurityGroupIngressResponse(b)")
			pn("		if err != nil {")
			pn("			return nil, err")
			pn("		}")
			pn("")
		}
		if n == "AuthorizeSecurityGroupEgress" {
			pn("		b, err = convertAuthorizeSecurityGroupEgressResponse(b)")
			pn("		if err != nil {")
			pn("			return nil, err")
			pn("		}")
			pn("")
		}
		pn("		if err := json.Unmarshal(b, &r); err != nil {")
		pn("			return nil, err")
		pn("		}")
		pn("	}")
		pn("")
	}
	pn("	return &r, nil")
	pn("}")
	pn("")
}

func isSuccessOnlyResponse(resp APIResponses) bool {
	success := false
	displaytext := false

	for _, r := range resp {
		if r.Name == "displaytext" {
			displaytext = true
		}
		if r.Name == "success" {
			success = true
		}
	}
	return displaytext && success
}

func (s *service) generateResponseType(a *API) {
	pn := s.pn
	tn := capitalize(strings.TrimPrefix(a.Name, "configure") + "Response")

	// add custom response types for some specific API calls
	if a.Name == "quotaBalance" {
		pn("type QuotaBalanceResponse struct {")
		pn("    Statement QuotaBalanceResponseType `json:\"balance\"`")
		pn("}")
		pn("")
		pn("type QuotaBalanceResponseType struct {")
		pn("    StartQuota float64  `json:\"startquota\"`")
		pn("    Credits    []string `json:\"credits\"`")
		pn("    StartDate  string   `json:\"startdate\"`")
		pn("    Currency   string   `json:\"currency\"`")
		pn("}")
		pn("")
		return
	}
	if a.Name == "quotaStatement" {
		pn("type QuotaStatementResponse struct {")
		pn("    Statement QuotaStatementResponseType `json:\"statement\"`")
		pn("}")
		pn("")
		pn("type QuotaStatementResponseType struct {")
		pn("    QuotaUsage []QuotaUsage `json:\"quotausage\"`")
		pn("    TotalQuota float64      `json:\"totalquota\"`")
		pn("    StartDate  string       `json:\"startdate\"`")
		pn("    EndDate    string       `json:\"enddate\"`")
		pn("    Currency   string       `json:\"currency\"`")
		pn("}")
		pn("")
		pn("type QuotaUsage struct {")
		pn("    Type      int     `json:\"type\"`")
		pn("    Accountid int     `json:\"accountid\"`")
		pn("    Domain    int     `json:\"domain\"`")
		pn("    Name      string  `json:\"name\"`")
		pn("    Unit      string  `json:\"unit\"`")
		pn("    Quota     float64 `json:\"quota\"`")
		pn("}")
		pn("")
		return
	}

	ln := capitalize(strings.TrimPrefix(a.Name, "list"))

	// If this is a 'list' response, we need an separate list struct. There seem to be other
	// types of responses that also need a separate list struct, so checking on exact matches
	// for those once.
	if strings.HasPrefix(a.Name, "list") || a.Name == "registerTemplate" || a.Name == "findHostsForMigration" || a.Name == "registerUserData" ||
		a.Name == "quotaBalance" || a.Name == "quotaSummary" || a.Name == "quotaTariffList" {
		pn("type %s struct {", tn)

		// This nasty check is for some specific response that do not behave consistent
		switch a.Name {
		case "listAsyncJobs":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "asyncjobs")
		case "listCapabilities":
			pn("    %s *%s `json:\"%s\"`", ln, parseSingular(ln), "capability")
		case "listEgressFirewallRules":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "firewallrule")
		case "listLoadBalancerRuleInstances":
			pn("	Count int `json:\"count\"`")
			pn("	LBRuleVMIDIPs []*%s `json:\"%s\"`", parseSingular(ln), "lbrulevmidip")
			pn("	LoadBalancerRuleInstances []*VirtualMachine `json:\"%s\"`", strings.ToLower(parseSingular(ln)))
		case "listVirtualMachinesMetrics":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "virtualmachine")
		case "listManagementServersMetrics":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "managementserver")
		case "listDbMetrics":
			pn("	%s %s `json:\"%s\"`", ln, parseSingular(ln), "dbMetrics")
		case "registerTemplate":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "template")
		case "listDomainChildren":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "domain")
		case "findHostsForMigration":
			pn(" Count int `json:\"count\"`")
			pn(" Host []*%s `json:\"%s\"`", customResponseStructTypes[a.Name], "host")
		case "listVmwareDcVms":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "unmanagedinstance")
		case "registerUserData":
			pn("    Account string `json:\"account\"`")
			pn("    Accountid string `json:\"accountid\"`")
			pn("    Domain string `json:\"domain\"`")
			pn("    Domainid string `json:\"domainid\"`")
			pn("    Hasannotations bool `json:\"hasannotations\"`")
			pn("    Id string `json:\"id\"`")
			pn("    JobID string `json:\"jobid\"`")
			pn("    Jobstatus int `json:\"jobstatus\"`")
			pn("    Name string `json:\"name\"`")
			pn("    Params string `json:\"params\"`")
			pn("    Userdata string `json:\"userdata\"`")
		case "listObjectStoragePools":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "objectstore")
		case "listStoragePoolObjects":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "datastoreobject")
		case "listImageStoreObjects":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "datastoreobject")
		case "listVolumesUsageHistory":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "volume")
		case "listHostHAProviders":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "haprovider")
		case "listSecondaryStorageSelectors":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "heuristics")
		case "listHostHAResources":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "hostha")
		case "listInfrastructure":
			pn("	Count int `json:\"count\"`")
			pn("	%s *%s `json:\"%s\"`", ln, parseSingular(ln), "infrastructure")
		case "listStoragePoolsMetrics":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "storagepool")
		case "quotaTariffList":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "quotatariff")
		case "quotaBalance":
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "balance")
		case "quotaSummary":
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), "summary")
		default:
			pn("	Count int `json:\"count\"`")
			pn("	%s []*%s `json:\"%s\"`", ln, parseSingular(ln), strings.ToLower(parseSingular(ln)))
		}
		pn("}")
		pn("")
		tn = parseSingular(ln)
	}

	sort.Sort(a.Response)
	customMarshal := s.recusiveGenerateResponseType(a.Name, tn, a.Response, a.Isasync)

	if customMarshal {
		pn("func (r *%s) UnmarshalJSON(b []byte) error {", tn)
		pn("	var m map[string]interface{}")
		pn("	err := json.Unmarshal(b, &m)")
		pn("	if err != nil {")
		pn("		return err")
		pn("	}")
		pn("")
		pn("	if success, ok := m[\"success\"].(string); ok {")
		pn("		m[\"success\"] = success == \"true\"")
		pn("		b, err = json.Marshal(m)")
		pn("		if err != nil {")
		pn("			return err")
		pn("		}")
		pn("	}")
		pn("")
		pn("	if ostypeid, ok := m[\"ostypeid\"].(float64); ok {")
		pn("		m[\"ostypeid\"] = strconv.Itoa(int(ostypeid))")
		pn("		b, err = json.Marshal(m)")
		pn("		if err != nil {")
		pn("			return err")
		pn("		}")
		pn("	}")
		pn("")
		pn("	type alias %s", tn)
		pn("	return json.Unmarshal(b, (*alias)(r))")
		pn("}")
		pn("")
	}
}

func parseSingular(n string) string {
	if strings.HasSuffix(n, "ies") {
		return strings.TrimSuffix(n, "ies") + "y"
	}
	if strings.HasSuffix(n, "sses") {
		return strings.TrimSuffix(n, "es")
	}
	return strings.TrimSuffix(n, "s")
}

func (s *service) recusiveGenerateResponseType(aName string, tn string, resp APIResponses, async bool) bool {
	pn := s.pn
	customMarshal := false
	found := make(map[string]bool)

	// Only apply custom response type name if the current tn is not already a custom type
	if val, ok := customResponseStructTypes[aName]; ok {
		// Don't override if tn is already a custom-generated nested type (contains the custom name)
		isNestedType := false
		for _, customType := range customResponseStructTypes {
			if strings.Contains(tn, customType) && tn != customType {
				isNestedType = true
				break
			}
		}
		if !isNestedType {
			tn = val
		}
	}

	pn("type %s struct {", tn)

	for _, r := range resp {
		if r.Name == "" {
			continue
		}
		if r.Name == "secondaryip" {
			pn("%s []struct {", capitalize(r.Name))
			pn("	Id string `json:\"id\"`")
			pn("	Ipaddress string `json:\"ipaddress\"`")
			pn("} `json:\"%s\"`", r.Name)
			continue
		}
		if r.Response != nil {
			sort.Sort(r.Response)
			typeName, create := getUniqueTypeName(tn, r.Name)
			pn("%s []%s `json:\"%s\"`", capitalize(r.Name), typeName, r.Name)
			if create {
				defer s.recusiveGenerateResponseType(aName, typeName, r.Response, false)
			}
		} else {
			if !found[r.Name] {
				switch r.Name {
				case "success":
					// This case is because the response field is different for sync and async calls :(
					pn("%s bool `json:\"%s\"`", capitalize(r.Name), r.Name)
					if !async {
						customMarshal = true
					}
				case "ostypeid":
					// This case is needed for backwards compatibility.
					pn("%s string `json:\"%s\"`", capitalize(r.Name), r.Name)
					customMarshal = true
				default:
					pn("%s %s `json:\"%s\"`", capitalize(r.Name), mapType(aName, r.Name, r.Type), r.Name)
				}
				found[r.Name] = true
			}
		}
	}

	pn("}")
	pn("")

	return customMarshal
}

func getUniqueTypeName(prefix, name string) (string, bool) {
	// We have special cases for [in|e]gressrules, nics and tags as the exact
	// sames types are used used in multiple different locations.
	switch {
	case strings.HasSuffix(name, "gressrule"):
		name = "rule"
	case strings.HasSuffix(name, "nic"):
		prefix = ""
		name = "nic"
	case strings.HasSuffix(name, "tags"):
		prefix = ""
		name = "tags"
	}

	tn := prefix + capitalize(name)
	if !typeNames[tn] {
		typeNames[tn] = true
		return tn, true
	}

	// Return here as this means the type already exists.
	if name == "rule" || name == "nic" || name == "tags" {
		return tn, false
	}

	return getUniqueTypeName(prefix, name+"Internal")
}

func logMissingApis(ai map[string]*API, as *allServices) {
	asMap := make(map[string]*API)
	for _, svc := range as.services {
		for _, api := range svc.apis {
			asMap[api.Name] = api
		}
	}

	for apiName, _ := range ai {
		_, found := asMap[apiName]
		if !found {
			log.Printf("Api missing in layout: %s", apiName)
		}
	}
}

func getAllServices(listApis string) (*allServices, []error, error) {
	// Get a map with all API info
	ai, err := getAPIInfo(listApis)
	if err != nil {
		return nil, nil, err
	}

	// Generate a complete set of services with their methods (APIs)
	as := &allServices{}
	errors := []error{}
	for sn, apis := range layout {
		typeNames[sn] = true
		s := &service{name: sn}
		for _, api := range apis {
			a, found := ai[api]
			if !found {
				errors = append(errors, &apiInfoNotFoundError{api})
				continue
			}
			s.apis = append(s.apis, a)
		}
		for _, apis := range s.apis {
			sort.Sort(apis.Params)
		}
		as.services = append(as.services, s)
	}

	// Add an extra field to enable adding a custom service
	as.services = append(as.services, &service{name: "CustomService"})
	sort.Sort(as.services)

	logMissingApis(ai, as)

	return as, errors, nil
}

func getAPIInfo(listApis string) (map[string]*API, error) {
	apis, err := ioutil.ReadFile(listApis)
	if err != nil {
		return nil, err
	}

	var ar struct {
		Count int    `json:"count"`
		APIs  []*API `json:"api"`
	}
	if err := json.Unmarshal(apis, &ar); err != nil {
		return nil, err
	}

	// Make a map of all retrieved APIs
	ai := make(map[string]*API)
	for _, api := range ar.APIs {
		ai[api.Name] = api
	}
	return ai, nil
}

func sourceDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	outdir := path.Join(wd, pkg)

	if err := os.MkdirAll(outdir, 0755); err != nil {
		return "", fmt.Errorf("Failed to Mkdir %s: %v", outdir, err)
	}
	return outdir, nil
}

func testDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	testdir := path.Join(wd, "test")

	if err := os.MkdirAll(testdir, 0755); err != nil {
		return "", fmt.Errorf("Failed to Mkdir %s: %v", testdir, err)
	}
	return testdir, nil

}

func isRequiredParam(api *API, apiParam *APIParam) bool {
	if params, ok := requiredParams[api.Name]; ok {
		for _, param := range params {
			if param == apiParam.Name {
				return true
			}
		}
	}
	return false
}

func mapType(aName string, pName string, pType string) string {
	if _, ok := longToStringConvertedParams[pName]; ok {
		pType = "UUID"
	}

	switch pType {
	case "UUID":
		return "UUID"
	case "boolean":
		return "bool"
	case "short", "int", "integer":
		return "int"
	case "long":
		return "int64"
	case "float", "double", "bigdecimal":
		return "float64"
	case "list":
		if pName == "downloaddetails" || pName == "owner" {
			return "[]map[string]string"
		} else if pName == "network" {
			return "[]*Network"
		}
		if pName == "virtualmachines" {
			return "[]*VirtualMachine"
		}
		return "[]string"
	case "map":
		if mapRequireList[aName] != nil && mapRequireList[aName][pName] {
			return "[]map[string]string"
		}
		return "map[string]string"
	case "double[]":
		return "[]float64"
	case "long[]":
		return "[]int64"
	case "string[]":
		return "[]string"
	case "set":
		return "[]interface{}"
	case "resourceiconresponse":
		return "interface{}"
	case "responseobject":
		return "json.RawMessage"
	case "uservmresponse":
		// This is a really specific anomaly of the API
		return "*VirtualMachine"
	case "outofbandmanagementresponse":
		return "OutOfBandManagementResponse"
	case "hostharesponse":
		return "HAForHostResponse"
	case "consoleendpointwebsocketresponse":
		return "map[string]interface{}"
	default:
		return "string"
	}
}

func capitalize(s string) string {
	if s == "jobid" {
		return "JobID"
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func uncapitalize(s string) string {
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

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

package cloudstack

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	gomock "github.com/golang/mock/gomock"
)

// UnlimitedResourceID is a special ID to define an unlimited resource
const UnlimitedResourceID = "-1"

var idRegex = regexp.MustCompile(`^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|-1)$`)

// IsID return true if the passed ID is either a UUID or a UnlimitedResourceID
func IsID(id string) bool {
	return idRegex.MatchString(id)
}

// ClientOption can be passed to new client functions to set custom options
type ClientOption func(*CloudStackClient)

// OptionFunc can be passed to the courtesy helper functions to set additional parameters
type OptionFunc func(*CloudStackClient, interface{}) error

type CSError struct {
	ErrorCode   int    `json:"errorcode"`
	CSErrorCode int    `json:"cserrorcode"`
	ErrorText   string `json:"errortext"`
}

func (e *CSError) Error() error {
	return fmt.Errorf("CloudStack API error %d (CSExceptionErrorCode: %d): %s", e.ErrorCode, e.CSErrorCode, e.ErrorText)
}

type UUID string

func (c UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(c))
}

func (c *UUID) UnmarshalJSON(data []byte) error {
	value := strings.Trim(string(data), "\"")
	if strings.HasPrefix(string(data), "\"") {
		*c = UUID(value)
		return nil
	}
	_, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	*c = UUID(value)
	return nil
}

type CloudStackClient struct {
	HTTPGETOnly bool // If `true` only use HTTP GET calls

	client  *http.Client // The http client for communicating
	baseURL string       // The base URL of the API
	apiKey  string       // Api key
	secret  string       // Secret key
	async   bool         // Wait for async calls to finish
	options []OptionFunc // A list of option functions to apply to all API calls
	timeout int64        // Max waiting timeout in seconds for async jobs to finish; defaults to 300 seconds

	APIDiscovery        APIDiscoveryServiceIface
	Account             AccountServiceIface
	Address             AddressServiceIface
	AffinityGroup       AffinityGroupServiceIface
	Alert               AlertServiceIface
	Annotation          AnnotationServiceIface
	Asyncjob            AsyncjobServiceIface
	Authentication      AuthenticationServiceIface
	AutoScale           AutoScaleServiceIface
	Baremetal           BaremetalServiceIface
	BigSwitchBCF        BigSwitchBCFServiceIface
	BrocadeVCS          BrocadeVCSServiceIface
	Certificate         CertificateServiceIface
	CloudIdentifier     CloudIdentifierServiceIface
	Cluster             ClusterServiceIface
	Configuration       ConfigurationServiceIface
	ConsoleEndpoint     ConsoleEndpointServiceIface
	Custom              CustomServiceIface
	Diagnostics         DiagnosticsServiceIface
	DiskOffering        DiskOfferingServiceIface
	Domain              DomainServiceIface
	Event               EventServiceIface
	Firewall            FirewallServiceIface
	GuestOS             GuestOSServiceIface
	Host                HostServiceIface
	Hypervisor          HypervisorServiceIface
	ISO                 ISOServiceIface
	ImageStore          ImageStoreServiceIface
	InfrastructureUsage InfrastructureUsageServiceIface
	InternalLB          InternalLBServiceIface
	Kubernetes          KubernetesServiceIface
	LDAP                LDAPServiceIface
	Limit               LimitServiceIface
	LoadBalancer        LoadBalancerServiceIface
	Management          ManagementServiceIface
	Metrics             MetricsServiceIface
	NAT                 NATServiceIface
	NetworkACL          NetworkACLServiceIface
	NetworkDevice       NetworkDeviceServiceIface
	NetworkOffering     NetworkOfferingServiceIface
	Network             NetworkServiceIface
	Nic                 NicServiceIface
	NiciraNVP           NiciraNVPServiceIface
	Oauth               OauthServiceIface
	ObjectStore         ObjectStoreServiceIface
	OutofbandManagement OutofbandManagementServiceIface
	OvsElement          OvsElementServiceIface
	Pod                 PodServiceIface
	Pool                PoolServiceIface
	PortableIP          PortableIPServiceIface
	Project             ProjectServiceIface
	Quota               QuotaServiceIface
	Region              RegionServiceIface
	Registration        RegistrationServiceIface
	Resourcemetadata    ResourcemetadataServiceIface
	Resourcetags        ResourcetagsServiceIface
	Role                RoleServiceIface
	Router              RouterServiceIface
	SSH                 SSHServiceIface
	SecurityGroup       SecurityGroupServiceIface
	ServiceOffering     ServiceOfferingServiceIface
	Shutdown            ShutdownServiceIface
	Snapshot            SnapshotServiceIface
	StoragePool         StoragePoolServiceIface
	StratosphereSSP     StratosphereSSPServiceIface
	Swift               SwiftServiceIface
	SystemCapacity      SystemCapacityServiceIface
	SystemVM            SystemVMServiceIface
	Template            TemplateServiceIface
	UCS                 UCSServiceIface
	Usage               UsageServiceIface
	User                UserServiceIface
	VLAN                VLANServiceIface
	VMGroup             VMGroupServiceIface
	VPC                 VPCServiceIface
	VPN                 VPNServiceIface
	VirtualMachine      VirtualMachineServiceIface
	Volume              VolumeServiceIface
	Zone                ZoneServiceIface
}

// Creates a new client for communicating with CloudStack
func newClient(apiurl string, apikey string, secret string, async bool, verifyssl bool, options ...ClientOption) *CloudStackClient {
	jar, _ := cookiejar.New(nil)
	cs := &CloudStackClient{
		client: &http.Client{
			Jar: jar,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
					DualStack: true,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: !verifyssl},
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
			Timeout: time.Duration(60 * time.Second),
		},
		baseURL: apiurl,
		apiKey:  apikey,
		secret:  secret,
		async:   async,
		options: []OptionFunc{},
		timeout: 300,
	}

	for _, fn := range options {
		fn(cs)
	}

	cs.APIDiscovery = NewAPIDiscoveryService(cs)
	cs.Account = NewAccountService(cs)
	cs.Address = NewAddressService(cs)
	cs.AffinityGroup = NewAffinityGroupService(cs)
	cs.Alert = NewAlertService(cs)
	cs.Annotation = NewAnnotationService(cs)
	cs.Asyncjob = NewAsyncjobService(cs)
	cs.Authentication = NewAuthenticationService(cs)
	cs.AutoScale = NewAutoScaleService(cs)
	cs.Baremetal = NewBaremetalService(cs)
	cs.BigSwitchBCF = NewBigSwitchBCFService(cs)
	cs.BrocadeVCS = NewBrocadeVCSService(cs)
	cs.Certificate = NewCertificateService(cs)
	cs.CloudIdentifier = NewCloudIdentifierService(cs)
	cs.Cluster = NewClusterService(cs)
	cs.Configuration = NewConfigurationService(cs)
	cs.ConsoleEndpoint = NewConsoleEndpointService(cs)
	cs.Custom = NewCustomService(cs)
	cs.Diagnostics = NewDiagnosticsService(cs)
	cs.DiskOffering = NewDiskOfferingService(cs)
	cs.Domain = NewDomainService(cs)
	cs.Event = NewEventService(cs)
	cs.Firewall = NewFirewallService(cs)
	cs.GuestOS = NewGuestOSService(cs)
	cs.Host = NewHostService(cs)
	cs.Hypervisor = NewHypervisorService(cs)
	cs.ISO = NewISOService(cs)
	cs.ImageStore = NewImageStoreService(cs)
	cs.InfrastructureUsage = NewInfrastructureUsageService(cs)
	cs.InternalLB = NewInternalLBService(cs)
	cs.Kubernetes = NewKubernetesService(cs)
	cs.LDAP = NewLDAPService(cs)
	cs.Limit = NewLimitService(cs)
	cs.LoadBalancer = NewLoadBalancerService(cs)
	cs.Management = NewManagementService(cs)
	cs.Metrics = NewMetricsService(cs)
	cs.NAT = NewNATService(cs)
	cs.NetworkACL = NewNetworkACLService(cs)
	cs.NetworkDevice = NewNetworkDeviceService(cs)
	cs.NetworkOffering = NewNetworkOfferingService(cs)
	cs.Network = NewNetworkService(cs)
	cs.Nic = NewNicService(cs)
	cs.NiciraNVP = NewNiciraNVPService(cs)
	cs.Oauth = NewOauthService(cs)
	cs.ObjectStore = NewObjectStoreService(cs)
	cs.OutofbandManagement = NewOutofbandManagementService(cs)
	cs.OvsElement = NewOvsElementService(cs)
	cs.Pod = NewPodService(cs)
	cs.Pool = NewPoolService(cs)
	cs.PortableIP = NewPortableIPService(cs)
	cs.Project = NewProjectService(cs)
	cs.Quota = NewQuotaService(cs)
	cs.Region = NewRegionService(cs)
	cs.Registration = NewRegistrationService(cs)
	cs.Resourcemetadata = NewResourcemetadataService(cs)
	cs.Resourcetags = NewResourcetagsService(cs)
	cs.Role = NewRoleService(cs)
	cs.Router = NewRouterService(cs)
	cs.SSH = NewSSHService(cs)
	cs.SecurityGroup = NewSecurityGroupService(cs)
	cs.ServiceOffering = NewServiceOfferingService(cs)
	cs.Shutdown = NewShutdownService(cs)
	cs.Snapshot = NewSnapshotService(cs)
	cs.StoragePool = NewStoragePoolService(cs)
	cs.StratosphereSSP = NewStratosphereSSPService(cs)
	cs.Swift = NewSwiftService(cs)
	cs.SystemCapacity = NewSystemCapacityService(cs)
	cs.SystemVM = NewSystemVMService(cs)
	cs.Template = NewTemplateService(cs)
	cs.UCS = NewUCSService(cs)
	cs.Usage = NewUsageService(cs)
	cs.User = NewUserService(cs)
	cs.VLAN = NewVLANService(cs)
	cs.VMGroup = NewVMGroupService(cs)
	cs.VPC = NewVPCService(cs)
	cs.VPN = NewVPNService(cs)
	cs.VirtualMachine = NewVirtualMachineService(cs)
	cs.Volume = NewVolumeService(cs)
	cs.Zone = NewZoneService(cs)

	return cs
}

// Creates a new mock client for communicating with CloudStack
func newMockClient(ctrl *gomock.Controller) *CloudStackClient {
	cs := &CloudStackClient{}

	cs.APIDiscovery = NewMockAPIDiscoveryServiceIface(ctrl)
	cs.Account = NewMockAccountServiceIface(ctrl)
	cs.Address = NewMockAddressServiceIface(ctrl)
	cs.AffinityGroup = NewMockAffinityGroupServiceIface(ctrl)
	cs.Alert = NewMockAlertServiceIface(ctrl)
	cs.Annotation = NewMockAnnotationServiceIface(ctrl)
	cs.Asyncjob = NewMockAsyncjobServiceIface(ctrl)
	cs.Authentication = NewMockAuthenticationServiceIface(ctrl)
	cs.AutoScale = NewMockAutoScaleServiceIface(ctrl)
	cs.Baremetal = NewMockBaremetalServiceIface(ctrl)
	cs.BigSwitchBCF = NewMockBigSwitchBCFServiceIface(ctrl)
	cs.BrocadeVCS = NewMockBrocadeVCSServiceIface(ctrl)
	cs.Certificate = NewMockCertificateServiceIface(ctrl)
	cs.CloudIdentifier = NewMockCloudIdentifierServiceIface(ctrl)
	cs.Cluster = NewMockClusterServiceIface(ctrl)
	cs.Configuration = NewMockConfigurationServiceIface(ctrl)
	cs.ConsoleEndpoint = NewMockConsoleEndpointServiceIface(ctrl)
	cs.Custom = NewMockCustomServiceIface(ctrl)
	cs.Diagnostics = NewMockDiagnosticsServiceIface(ctrl)
	cs.DiskOffering = NewMockDiskOfferingServiceIface(ctrl)
	cs.Domain = NewMockDomainServiceIface(ctrl)
	cs.Event = NewMockEventServiceIface(ctrl)
	cs.Firewall = NewMockFirewallServiceIface(ctrl)
	cs.GuestOS = NewMockGuestOSServiceIface(ctrl)
	cs.Host = NewMockHostServiceIface(ctrl)
	cs.Hypervisor = NewMockHypervisorServiceIface(ctrl)
	cs.ISO = NewMockISOServiceIface(ctrl)
	cs.ImageStore = NewMockImageStoreServiceIface(ctrl)
	cs.InfrastructureUsage = NewMockInfrastructureUsageServiceIface(ctrl)
	cs.InternalLB = NewMockInternalLBServiceIface(ctrl)
	cs.Kubernetes = NewMockKubernetesServiceIface(ctrl)
	cs.LDAP = NewMockLDAPServiceIface(ctrl)
	cs.Limit = NewMockLimitServiceIface(ctrl)
	cs.LoadBalancer = NewMockLoadBalancerServiceIface(ctrl)
	cs.Management = NewMockManagementServiceIface(ctrl)
	cs.Metrics = NewMockMetricsServiceIface(ctrl)
	cs.NAT = NewMockNATServiceIface(ctrl)
	cs.NetworkACL = NewMockNetworkACLServiceIface(ctrl)
	cs.NetworkDevice = NewMockNetworkDeviceServiceIface(ctrl)
	cs.NetworkOffering = NewMockNetworkOfferingServiceIface(ctrl)
	cs.Network = NewMockNetworkServiceIface(ctrl)
	cs.Nic = NewMockNicServiceIface(ctrl)
	cs.NiciraNVP = NewMockNiciraNVPServiceIface(ctrl)
	cs.Oauth = NewMockOauthServiceIface(ctrl)
	cs.ObjectStore = NewMockObjectStoreServiceIface(ctrl)
	cs.OutofbandManagement = NewMockOutofbandManagementServiceIface(ctrl)
	cs.OvsElement = NewMockOvsElementServiceIface(ctrl)
	cs.Pod = NewMockPodServiceIface(ctrl)
	cs.Pool = NewMockPoolServiceIface(ctrl)
	cs.PortableIP = NewMockPortableIPServiceIface(ctrl)
	cs.Project = NewMockProjectServiceIface(ctrl)
	cs.Quota = NewMockQuotaServiceIface(ctrl)
	cs.Region = NewMockRegionServiceIface(ctrl)
	cs.Registration = NewMockRegistrationServiceIface(ctrl)
	cs.Resourcemetadata = NewMockResourcemetadataServiceIface(ctrl)
	cs.Resourcetags = NewMockResourcetagsServiceIface(ctrl)
	cs.Role = NewMockRoleServiceIface(ctrl)
	cs.Router = NewMockRouterServiceIface(ctrl)
	cs.SSH = NewMockSSHServiceIface(ctrl)
	cs.SecurityGroup = NewMockSecurityGroupServiceIface(ctrl)
	cs.ServiceOffering = NewMockServiceOfferingServiceIface(ctrl)
	cs.Shutdown = NewMockShutdownServiceIface(ctrl)
	cs.Snapshot = NewMockSnapshotServiceIface(ctrl)
	cs.StoragePool = NewMockStoragePoolServiceIface(ctrl)
	cs.StratosphereSSP = NewMockStratosphereSSPServiceIface(ctrl)
	cs.Swift = NewMockSwiftServiceIface(ctrl)
	cs.SystemCapacity = NewMockSystemCapacityServiceIface(ctrl)
	cs.SystemVM = NewMockSystemVMServiceIface(ctrl)
	cs.Template = NewMockTemplateServiceIface(ctrl)
	cs.UCS = NewMockUCSServiceIface(ctrl)
	cs.Usage = NewMockUsageServiceIface(ctrl)
	cs.User = NewMockUserServiceIface(ctrl)
	cs.VLAN = NewMockVLANServiceIface(ctrl)
	cs.VMGroup = NewMockVMGroupServiceIface(ctrl)
	cs.VPC = NewMockVPCServiceIface(ctrl)
	cs.VPN = NewMockVPNServiceIface(ctrl)
	cs.VirtualMachine = NewMockVirtualMachineServiceIface(ctrl)
	cs.Volume = NewMockVolumeServiceIface(ctrl)
	cs.Zone = NewMockZoneServiceIface(ctrl)

	return cs
}

// Default non-async client. So for async calls you need to implement and check the async job result yourself. When using
// HTTPS with a self-signed certificate to connect to your CloudStack API, you would probably want to set 'verifyssl' to
// false so the call ignores the SSL errors/warnings.
func NewClient(apiurl string, apikey string, secret string, verifyssl bool, options ...ClientOption) *CloudStackClient {
	cs := newClient(apiurl, apikey, secret, false, verifyssl, options...)
	return cs
}

// For sync API calls this client behaves exactly the same as a standard client call, but for async API calls
// this client will wait until the async job is finished or until the configured AsyncTimeout is reached. When the async
// job finishes successfully it will return actual object received from the API and nil, but when the timeout is
// reached it will return the initial object containing the async job ID for the running job and a warning.
func NewAsyncClient(apiurl string, apikey string, secret string, verifyssl bool, options ...ClientOption) *CloudStackClient {
	cs := newClient(apiurl, apikey, secret, true, verifyssl, options...)
	return cs
}

// Creates a new mock client for communicating with CloudStack
func NewMockClient(ctrl *gomock.Controller) *CloudStackClient {
	cs := newMockClient(ctrl)
	return cs
}

// When using the async client an api call will wait for the async call to finish before returning. The default is to poll for 300 seconds
// seconds, to check if the async job is finished.
func (cs *CloudStackClient) AsyncTimeout(timeoutInSeconds int64) {
	cs.timeout = timeoutInSeconds
}

// Sets timeout when using sync api calls. Default is 60 seconds
func (cs *CloudStackClient) Timeout(timeout time.Duration) {
	cs.client.Timeout = timeout
}

// Set any default options that would be added to all API calls that support it.
func (cs *CloudStackClient) DefaultOptions(options ...OptionFunc) {
	if options != nil {
		cs.options = options
	} else {
		cs.options = []OptionFunc{}
	}
}

var AsyncTimeoutErr = errors.New("Timeout while waiting for async job to finish")

// A helper function that you can use to get the result of a running async job. If the job is not finished within the configured
// timeout, the async job returns a AsyncTimeoutErr.
func (cs *CloudStackClient) GetAsyncJobResult(jobid string, timeout int64) (json.RawMessage, error) {
	var timer time.Duration
	currentTime := time.Now().Unix()

	for {
		p := cs.Asyncjob.NewQueryAsyncJobResultParams(jobid)
		r, err := cs.Asyncjob.QueryAsyncJobResult(p)
		if err != nil {
			return nil, err
		}

		// Status 1 means the job is finished successfully
		if r.Jobstatus == 1 {
			return r.Jobresult, nil
		}

		// When the status is 2, the job has failed
		if r.Jobstatus == 2 {
			if r.Jobresulttype == "text" {
				return nil, fmt.Errorf(string(r.Jobresult))
			} else {
				return nil, fmt.Errorf("Undefined error: %s", string(r.Jobresult))
			}
		}

		if time.Now().Unix()-currentTime > timeout {
			return nil, AsyncTimeoutErr
		}

		// Add an (extremely simple) exponential backoff like feature to prevent
		// flooding the CloudStack API
		if timer < 15 {
			timer++
		}

		time.Sleep(timer * time.Second)
	}
}

// Execute the request against a CS API. Will return the raw JSON data returned by the API and nil if
// no error occurred. If the API returns an error the result will be nil and the HTTP error code and CS
// error details. If a processing (code) error occurs the result will be nil and the generated error
func (cs *CloudStackClient) newRequest(api string, params url.Values) (json.RawMessage, error) {
	return cs.newRawRequest(api, false, params)
}

// Execute the request against a CS API using POST. Will return the raw JSON data returned by the API and
// nil if no error occurred. If the API returns an error the result will be nil and the HTTP error code
// and CS error details. If a processing (code) error occurs the result will be nil and the generated error
func (cs *CloudStackClient) newPostRequest(api string, params url.Values) (json.RawMessage, error) {
	return cs.newRawRequest(api, true, params)
}

// Execute a raw request against a CS API. Will return the raw JSON data returned by the API and nil if
// no error occurred. If the API returns an error the result will be nil and the HTTP error code and CS
// error details. If a processing (code) error occurs the result will be nil and the generated error
func (cs *CloudStackClient) newRawRequest(api string, post bool, params url.Values) (json.RawMessage, error) {
	params.Set("apiKey", cs.apiKey)
	params.Set("command", api)
	params.Set("response", "json")

	// Generate signature for API call
	// * Serialize parameters, URL encoding only values and sort them by key, done by EncodeValues
	// * Convert the entire argument string to lowercase
	// * Replace all instances of '+' to '%20'
	// * Calculate HMAC SHA1 of argument string with CloudStack secret
	// * URL encode the string and convert to base64
	s := EncodeValues(params)
	s2 := strings.ToLower(s)
	mac := hmac.New(sha1.New, []byte(cs.secret))
	mac.Write([]byte(s2))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	var err error
	var resp *http.Response
	if !cs.HTTPGETOnly && post {
		// The deployVirtualMachine API should be called using a POST call
		// so we don't have to worry about the userdata size

		// Add the unescaped signature to the POST params
		params.Set("signature", signature)

		// Make a POST call
		resp, err = cs.client.PostForm(cs.baseURL, params)
	} else {
		// Create the final URL before we issue the request
		url := cs.baseURL + "?" + s + "&signature=" + url.QueryEscape(signature)

		// Make a GET call
		resp, err = cs.client.Get(url)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Need to get the raw value to make the result play nice
	b, err = getRawValue(b)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var e CSError
		if err := json.Unmarshal(b, &e); err != nil {
			return nil, err
		}
		return nil, e.Error()
	}
	return b, nil
}

// Custom version of net/url Encode that only URL escapes values
// Unmodified portions here remain under BSD license of The Go Authors: https://go.googlesource.com/go/+/master/LICENSE
func EncodeValues(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := k + "="
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			escaped := url.QueryEscape(v)
			// we need to ensure + (representing a space) is encoded as %20
			escaped = strings.Replace(escaped, "+", "%20", -1)
			// we need to ensure * is not escaped
			escaped = strings.Replace(escaped, "%2A", "*", -1)
			buf.WriteString(escaped)
		}
	}
	return buf.String()
}

// Generic function to get the first non-count raw value from a response as json.RawMessage
func getRawValue(b json.RawMessage) (json.RawMessage, error) {
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	getArrayResponse := false
	for k := range m {
		if k == "count" {
			getArrayResponse = true
		}
	}
	if getArrayResponse {
		var resp []json.RawMessage
		for k, v := range m {
			if k != "count" {
				if err := json.Unmarshal(v, &resp); err != nil {
					return nil, err
				}
				return resp[0], nil
			}
		}

	} else {
		for _, v := range m {
			return v, nil
		}
	}
	return nil, fmt.Errorf("Unable to extract the raw value from:\n\n%s\n\n", string(b))
}

// getSortedKeysFromMap returns the keys from m in increasing order.
func getSortedKeysFromMap(m map[string]string) (keys []string) {
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// WithAsyncTimeout takes a custom timeout to be used by the CloudStackClient
func WithAsyncTimeout(timeout int64) ClientOption {
	return func(cs *CloudStackClient) {
		if timeout != 0 {
			cs.timeout = timeout
		}
	}
}

// DomainIDSetter is an interface that every type that can set a domain ID must implement
type DomainIDSetter interface {
	SetDomainid(string)
}

// WithDomain takes either a domain name or ID and sets the `domainid` parameter
func WithDomain(domain string) OptionFunc {
	return func(cs *CloudStackClient, p interface{}) error {
		ps, ok := p.(DomainIDSetter)

		if !ok || domain == "" {
			return nil
		}

		if !IsID(domain) {
			id, _, err := cs.Domain.GetDomainID(domain)
			if err != nil {
				return err
			}
			domain = id
		}

		ps.SetDomainid(domain)

		return nil
	}
}

// WithHTTPClient takes a custom HTTP client to be used by the CloudStackClient
func WithHTTPClient(client *http.Client) ClientOption {
	return func(cs *CloudStackClient) {
		if client != nil {
			if client.Jar == nil {
				client.Jar = cs.client.Jar
			}
			cs.client = client
		}
	}
}

// ListallSetter is an interface that every type that can set listall must implement
type ListallSetter interface {
	SetListall(bool)
}

// WithListall takes either a project name or ID and sets the `listall` parameter
func WithListall(listall bool) OptionFunc {
	return func(cs *CloudStackClient, p interface{}) error {
		ps, ok := p.(ListallSetter)

		if !ok {
			return nil
		}

		ps.SetListall(listall)

		return nil
	}
}

// ProjectIDSetter is an interface that every type that can set a project ID must implement
type ProjectIDSetter interface {
	SetProjectid(string)
}

// WithProject takes either a project name or ID and sets the `projectid` parameter
func WithProject(project string) OptionFunc {
	return func(cs *CloudStackClient, p interface{}) error {
		ps, ok := p.(ProjectIDSetter)

		if !ok || project == "" {
			return nil
		}

		if !IsID(project) {
			id, _, err := cs.Project.GetProjectID(project)
			if err != nil {
				return err
			}
			project = id
		}

		ps.SetProjectid(project)

		return nil
	}
}

// VPCIDSetter is an interface that every type that can set a vpc ID must implement
type VPCIDSetter interface {
	SetVpcid(string)
}

// WithVPCID takes a vpc ID and sets the `vpcid` parameter
func WithVPCID(id string) OptionFunc {
	return func(cs *CloudStackClient, p interface{}) error {
		vs, ok := p.(VPCIDSetter)

		if !ok || id == "" {
			return nil
		}

		vs.SetVpcid(id)

		return nil
	}
}

// ZoneIDSetter is an interface that every type that can set a zone ID must implement
type ZoneIDSetter interface {
	SetZoneid(string)
}

// WithZone takes either a zone name or ID and sets the `zoneid` parameter
func WithZone(zone string) OptionFunc {
	return func(cs *CloudStackClient, p interface{}) error {
		zs, ok := p.(ZoneIDSetter)

		if !ok || zone == "" {
			return nil
		}

		if !IsID(zone) {
			id, _, err := cs.Zone.GetZoneID(zone)
			if err != nil {
				return err
			}
			zone = id
		}

		zs.SetZoneid(zone)

		return nil
	}
}

type APIDiscoveryService struct {
	cs *CloudStackClient
}

func NewAPIDiscoveryService(cs *CloudStackClient) APIDiscoveryServiceIface {
	return &APIDiscoveryService{cs: cs}
}

type AccountService struct {
	cs *CloudStackClient
}

func NewAccountService(cs *CloudStackClient) AccountServiceIface {
	return &AccountService{cs: cs}
}

type AddressService struct {
	cs *CloudStackClient
}

func NewAddressService(cs *CloudStackClient) AddressServiceIface {
	return &AddressService{cs: cs}
}

type AffinityGroupService struct {
	cs *CloudStackClient
}

func NewAffinityGroupService(cs *CloudStackClient) AffinityGroupServiceIface {
	return &AffinityGroupService{cs: cs}
}

type AlertService struct {
	cs *CloudStackClient
}

func NewAlertService(cs *CloudStackClient) AlertServiceIface {
	return &AlertService{cs: cs}
}

type AnnotationService struct {
	cs *CloudStackClient
}

func NewAnnotationService(cs *CloudStackClient) AnnotationServiceIface {
	return &AnnotationService{cs: cs}
}

type AsyncjobService struct {
	cs *CloudStackClient
}

func NewAsyncjobService(cs *CloudStackClient) AsyncjobServiceIface {
	return &AsyncjobService{cs: cs}
}

type AuthenticationService struct {
	cs *CloudStackClient
}

func NewAuthenticationService(cs *CloudStackClient) AuthenticationServiceIface {
	return &AuthenticationService{cs: cs}
}

type AutoScaleService struct {
	cs *CloudStackClient
}

func NewAutoScaleService(cs *CloudStackClient) AutoScaleServiceIface {
	return &AutoScaleService{cs: cs}
}

type BaremetalService struct {
	cs *CloudStackClient
}

func NewBaremetalService(cs *CloudStackClient) BaremetalServiceIface {
	return &BaremetalService{cs: cs}
}

type BigSwitchBCFService struct {
	cs *CloudStackClient
}

func NewBigSwitchBCFService(cs *CloudStackClient) BigSwitchBCFServiceIface {
	return &BigSwitchBCFService{cs: cs}
}

type BrocadeVCSService struct {
	cs *CloudStackClient
}

func NewBrocadeVCSService(cs *CloudStackClient) BrocadeVCSServiceIface {
	return &BrocadeVCSService{cs: cs}
}

type CertificateService struct {
	cs *CloudStackClient
}

func NewCertificateService(cs *CloudStackClient) CertificateServiceIface {
	return &CertificateService{cs: cs}
}

type CloudIdentifierService struct {
	cs *CloudStackClient
}

func NewCloudIdentifierService(cs *CloudStackClient) CloudIdentifierServiceIface {
	return &CloudIdentifierService{cs: cs}
}

type ClusterService struct {
	cs *CloudStackClient
}

func NewClusterService(cs *CloudStackClient) ClusterServiceIface {
	return &ClusterService{cs: cs}
}

type ConfigurationService struct {
	cs *CloudStackClient
}

func NewConfigurationService(cs *CloudStackClient) ConfigurationServiceIface {
	return &ConfigurationService{cs: cs}
}

type ConsoleEndpointService struct {
	cs *CloudStackClient
}

func NewConsoleEndpointService(cs *CloudStackClient) ConsoleEndpointServiceIface {
	return &ConsoleEndpointService{cs: cs}
}

type CustomService struct {
	cs *CloudStackClient
}

func NewCustomService(cs *CloudStackClient) CustomServiceIface {
	return &CustomService{cs: cs}
}

type DiagnosticsService struct {
	cs *CloudStackClient
}

func NewDiagnosticsService(cs *CloudStackClient) DiagnosticsServiceIface {
	return &DiagnosticsService{cs: cs}
}

type DiskOfferingService struct {
	cs *CloudStackClient
}

func NewDiskOfferingService(cs *CloudStackClient) DiskOfferingServiceIface {
	return &DiskOfferingService{cs: cs}
}

type DomainService struct {
	cs *CloudStackClient
}

func NewDomainService(cs *CloudStackClient) DomainServiceIface {
	return &DomainService{cs: cs}
}

type EventService struct {
	cs *CloudStackClient
}

func NewEventService(cs *CloudStackClient) EventServiceIface {
	return &EventService{cs: cs}
}

type FirewallService struct {
	cs *CloudStackClient
}

func NewFirewallService(cs *CloudStackClient) FirewallServiceIface {
	return &FirewallService{cs: cs}
}

type GuestOSService struct {
	cs *CloudStackClient
}

func NewGuestOSService(cs *CloudStackClient) GuestOSServiceIface {
	return &GuestOSService{cs: cs}
}

type HostService struct {
	cs *CloudStackClient
}

func NewHostService(cs *CloudStackClient) HostServiceIface {
	return &HostService{cs: cs}
}

type HypervisorService struct {
	cs *CloudStackClient
}

func NewHypervisorService(cs *CloudStackClient) HypervisorServiceIface {
	return &HypervisorService{cs: cs}
}

type ISOService struct {
	cs *CloudStackClient
}

func NewISOService(cs *CloudStackClient) ISOServiceIface {
	return &ISOService{cs: cs}
}

type ImageStoreService struct {
	cs *CloudStackClient
}

func NewImageStoreService(cs *CloudStackClient) ImageStoreServiceIface {
	return &ImageStoreService{cs: cs}
}

type InfrastructureUsageService struct {
	cs *CloudStackClient
}

func NewInfrastructureUsageService(cs *CloudStackClient) InfrastructureUsageServiceIface {
	return &InfrastructureUsageService{cs: cs}
}

type InternalLBService struct {
	cs *CloudStackClient
}

func NewInternalLBService(cs *CloudStackClient) InternalLBServiceIface {
	return &InternalLBService{cs: cs}
}

type KubernetesService struct {
	cs *CloudStackClient
}

func NewKubernetesService(cs *CloudStackClient) KubernetesServiceIface {
	return &KubernetesService{cs: cs}
}

type LDAPService struct {
	cs *CloudStackClient
}

func NewLDAPService(cs *CloudStackClient) LDAPServiceIface {
	return &LDAPService{cs: cs}
}

type LimitService struct {
	cs *CloudStackClient
}

func NewLimitService(cs *CloudStackClient) LimitServiceIface {
	return &LimitService{cs: cs}
}

type LoadBalancerService struct {
	cs *CloudStackClient
}

func NewLoadBalancerService(cs *CloudStackClient) LoadBalancerServiceIface {
	return &LoadBalancerService{cs: cs}
}

type ManagementService struct {
	cs *CloudStackClient
}

func NewManagementService(cs *CloudStackClient) ManagementServiceIface {
	return &ManagementService{cs: cs}
}

type MetricsService struct {
	cs *CloudStackClient
}

func NewMetricsService(cs *CloudStackClient) MetricsServiceIface {
	return &MetricsService{cs: cs}
}

type NATService struct {
	cs *CloudStackClient
}

func NewNATService(cs *CloudStackClient) NATServiceIface {
	return &NATService{cs: cs}
}

type NetworkACLService struct {
	cs *CloudStackClient
}

func NewNetworkACLService(cs *CloudStackClient) NetworkACLServiceIface {
	return &NetworkACLService{cs: cs}
}

type NetworkDeviceService struct {
	cs *CloudStackClient
}

func NewNetworkDeviceService(cs *CloudStackClient) NetworkDeviceServiceIface {
	return &NetworkDeviceService{cs: cs}
}

type NetworkOfferingService struct {
	cs *CloudStackClient
}

func NewNetworkOfferingService(cs *CloudStackClient) NetworkOfferingServiceIface {
	return &NetworkOfferingService{cs: cs}
}

type NetworkService struct {
	cs *CloudStackClient
}

func NewNetworkService(cs *CloudStackClient) NetworkServiceIface {
	return &NetworkService{cs: cs}
}

type NicService struct {
	cs *CloudStackClient
}

func NewNicService(cs *CloudStackClient) NicServiceIface {
	return &NicService{cs: cs}
}

type NiciraNVPService struct {
	cs *CloudStackClient
}

func NewNiciraNVPService(cs *CloudStackClient) NiciraNVPServiceIface {
	return &NiciraNVPService{cs: cs}
}

type OauthService struct {
	cs *CloudStackClient
}

func NewOauthService(cs *CloudStackClient) OauthServiceIface {
	return &OauthService{cs: cs}
}

type ObjectStoreService struct {
	cs *CloudStackClient
}

func NewObjectStoreService(cs *CloudStackClient) ObjectStoreServiceIface {
	return &ObjectStoreService{cs: cs}
}

type OutofbandManagementService struct {
	cs *CloudStackClient
}

func NewOutofbandManagementService(cs *CloudStackClient) OutofbandManagementServiceIface {
	return &OutofbandManagementService{cs: cs}
}

type OvsElementService struct {
	cs *CloudStackClient
}

func NewOvsElementService(cs *CloudStackClient) OvsElementServiceIface {
	return &OvsElementService{cs: cs}
}

type PodService struct {
	cs *CloudStackClient
}

func NewPodService(cs *CloudStackClient) PodServiceIface {
	return &PodService{cs: cs}
}

type PoolService struct {
	cs *CloudStackClient
}

func NewPoolService(cs *CloudStackClient) PoolServiceIface {
	return &PoolService{cs: cs}
}

type PortableIPService struct {
	cs *CloudStackClient
}

func NewPortableIPService(cs *CloudStackClient) PortableIPServiceIface {
	return &PortableIPService{cs: cs}
}

type ProjectService struct {
	cs *CloudStackClient
}

func NewProjectService(cs *CloudStackClient) ProjectServiceIface {
	return &ProjectService{cs: cs}
}

type QuotaService struct {
	cs *CloudStackClient
}

func NewQuotaService(cs *CloudStackClient) QuotaServiceIface {
	return &QuotaService{cs: cs}
}

type RegionService struct {
	cs *CloudStackClient
}

func NewRegionService(cs *CloudStackClient) RegionServiceIface {
	return &RegionService{cs: cs}
}

type RegistrationService struct {
	cs *CloudStackClient
}

func NewRegistrationService(cs *CloudStackClient) RegistrationServiceIface {
	return &RegistrationService{cs: cs}
}

type ResourcemetadataService struct {
	cs *CloudStackClient
}

func NewResourcemetadataService(cs *CloudStackClient) ResourcemetadataServiceIface {
	return &ResourcemetadataService{cs: cs}
}

type ResourcetagsService struct {
	cs *CloudStackClient
}

func NewResourcetagsService(cs *CloudStackClient) ResourcetagsServiceIface {
	return &ResourcetagsService{cs: cs}
}

type RoleService struct {
	cs *CloudStackClient
}

func NewRoleService(cs *CloudStackClient) RoleServiceIface {
	return &RoleService{cs: cs}
}

type RouterService struct {
	cs *CloudStackClient
}

func NewRouterService(cs *CloudStackClient) RouterServiceIface {
	return &RouterService{cs: cs}
}

type SSHService struct {
	cs *CloudStackClient
}

func NewSSHService(cs *CloudStackClient) SSHServiceIface {
	return &SSHService{cs: cs}
}

type SecurityGroupService struct {
	cs *CloudStackClient
}

func NewSecurityGroupService(cs *CloudStackClient) SecurityGroupServiceIface {
	return &SecurityGroupService{cs: cs}
}

type ServiceOfferingService struct {
	cs *CloudStackClient
}

func NewServiceOfferingService(cs *CloudStackClient) ServiceOfferingServiceIface {
	return &ServiceOfferingService{cs: cs}
}

type ShutdownService struct {
	cs *CloudStackClient
}

func NewShutdownService(cs *CloudStackClient) ShutdownServiceIface {
	return &ShutdownService{cs: cs}
}

type SnapshotService struct {
	cs *CloudStackClient
}

func NewSnapshotService(cs *CloudStackClient) SnapshotServiceIface {
	return &SnapshotService{cs: cs}
}

type StoragePoolService struct {
	cs *CloudStackClient
}

func NewStoragePoolService(cs *CloudStackClient) StoragePoolServiceIface {
	return &StoragePoolService{cs: cs}
}

type StratosphereSSPService struct {
	cs *CloudStackClient
}

func NewStratosphereSSPService(cs *CloudStackClient) StratosphereSSPServiceIface {
	return &StratosphereSSPService{cs: cs}
}

type SwiftService struct {
	cs *CloudStackClient
}

func NewSwiftService(cs *CloudStackClient) SwiftServiceIface {
	return &SwiftService{cs: cs}
}

type SystemCapacityService struct {
	cs *CloudStackClient
}

func NewSystemCapacityService(cs *CloudStackClient) SystemCapacityServiceIface {
	return &SystemCapacityService{cs: cs}
}

type SystemVMService struct {
	cs *CloudStackClient
}

func NewSystemVMService(cs *CloudStackClient) SystemVMServiceIface {
	return &SystemVMService{cs: cs}
}

type TemplateService struct {
	cs *CloudStackClient
}

func NewTemplateService(cs *CloudStackClient) TemplateServiceIface {
	return &TemplateService{cs: cs}
}

type UCSService struct {
	cs *CloudStackClient
}

func NewUCSService(cs *CloudStackClient) UCSServiceIface {
	return &UCSService{cs: cs}
}

type UsageService struct {
	cs *CloudStackClient
}

func NewUsageService(cs *CloudStackClient) UsageServiceIface {
	return &UsageService{cs: cs}
}

type UserService struct {
	cs *CloudStackClient
}

func NewUserService(cs *CloudStackClient) UserServiceIface {
	return &UserService{cs: cs}
}

type VLANService struct {
	cs *CloudStackClient
}

func NewVLANService(cs *CloudStackClient) VLANServiceIface {
	return &VLANService{cs: cs}
}

type VMGroupService struct {
	cs *CloudStackClient
}

func NewVMGroupService(cs *CloudStackClient) VMGroupServiceIface {
	return &VMGroupService{cs: cs}
}

type VPCService struct {
	cs *CloudStackClient
}

func NewVPCService(cs *CloudStackClient) VPCServiceIface {
	return &VPCService{cs: cs}
}

type VPNService struct {
	cs *CloudStackClient
}

func NewVPNService(cs *CloudStackClient) VPNServiceIface {
	return &VPNService{cs: cs}
}

type VirtualMachineService struct {
	cs *CloudStackClient
}

func NewVirtualMachineService(cs *CloudStackClient) VirtualMachineServiceIface {
	return &VirtualMachineService{cs: cs}
}

type VolumeService struct {
	cs *CloudStackClient
}

func NewVolumeService(cs *CloudStackClient) VolumeServiceIface {
	return &VolumeService{cs: cs}
}

type ZoneService struct {
	cs *CloudStackClient
}

func NewZoneService(cs *CloudStackClient) ZoneServiceIface {
	return &ZoneService{cs: cs}
}

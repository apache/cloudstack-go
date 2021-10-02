package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"createuserresponse": {
				"user": {
					"id": "cd2b5afa-db4d-4532-88ec-1356a273e534",
					"username": "dummyUser",
					"firstname": "firstname",
					"lastname": "lastname",
					"email": "user@xyz.com",
					"created": "2021-10-04T08:57:35+0000",
					"state": "enabled",
					"account": "admin",
					"accounttype": 1,
					"usersource": "native",
					"roleid": "16919363-5fe0-11ea-9a56-1e006800018c",
					"roletype": "Admin",
					"rolename": "Root Admin",
					"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
					"domain": "ROOT",
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"iscallerchilddomain": false,
					"isdefault": false
				}
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.User.NewCreateUserParams("admin", "user.xyz.com",
		"firstname", "lastname", "password", "dummyUser")
	resp, err := client.User.CreateUser(params)
	if err != nil {
		t.Errorf("Failed to create user due to %v", err)
		return
	}
	if resp == nil || resp.Username != "dummyUser" {
		t.Errorf("Failed to create user")
	}
}

func TestUserService_EnableUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"enableuserresponse": {
				  "user": {
					"account": "admin",
					"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					"accounttype": 1,
					"created": "2021-10-04T08:57:35+0000",
					"domain": "ROOT",
					"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
					"email": "user@xyz.com",
					"firstname": "firstname",
					"id": "cd2b5afa-db4d-4532-88ec-1356a273e534",
					"iscallerchilddomain": false,
					"isdefault": false,
					"lastname": "lastname",
					"roleid": "16919363-5fe0-11ea-9a56-1e006800018c",
					"rolename": "Root Admin",
					"roletype": "Admin",
					"state": "enabled",
					"username": "dummyUser",
					"usersource": "native"
				  }
				}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.User.NewEnableUserParams("cd2b5afa-db4d-4532-88ec-1356a273e534")
	resp, err := client.User.EnableUser(params)
	if err != nil {
		t.Errorf("Failed to enable user due to %v", err)
		return
	}
	fmt.Println(resp)
	if resp == nil || resp.State != "enabled" {
		t.Errorf("Failed to enable user")
	}
}

func TestUserService_DisableUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		responses := map[string]string{
			"disableUser": `{
				"disableResponse": {
					"jobid": "e4eb553f-6188-441c-beb8-09056b0a1e1f"
				}
			}`,
			"queryAsyncJobResult": `{
				"queryasyncjobresultresponse": {
					  "accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
					  "cmd": "org.apache.cloudstack.api.command.admin.user.DisableUserCmd",
					  "completed": "2021-10-04T09:06:53+0000",
					  "created": "2021-10-04T09:06:52+0000",
					  "jobid": "e4eb553f-6188-441c-beb8-09056b0a1e1f",
					  "jobprocstatus": 0,
					  "jobresult": {
						"user": {
						  "account": "admin",
						  "accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
						  "accounttype": 1,
						  "created": "2021-10-04T08:57:35+0000",
						  "domain": "ROOT",
						  "domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
						  "email": "user@xyz.com",
						  "firstname": "firstname",
						  "id": "cd2b5afa-db4d-4532-88ec-1356a273e534",
						  "iscallerchilddomain": false,
						  "isdefault": false,
						  "lastname": "lastname",
						  "roleid": "16919363-5fe0-11ea-9a56-1e006800018c",
						  "rolename": "Root Admin",
						  "roletype": "Admin",
						  "state": "disabled",
						  "username": "dummyUser",
						  "usersource": "native"
						}
					  },
					  "jobresultcode": 0,
					  "jobresulttype": "object",
					  "jobstatus": 1,
					  "userid": "27f2484f-5fe0-11ea-9a56-1e006800018c"
					}
				}`,
		}
		fmt.Fprintf(writer, responses[request.FormValue("command")])
	}))
	defer server.Close()
	client := NewAsyncClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.User.NewDisableUserParams("cd2b5afa-db4d-4532-88ec-1356a273e534")
	resp, err := client.User.DisableUser(params)
	if err != nil {
		t.Errorf("Failed to disable user due to %v", err)
		return
	}

	if resp == nil || resp.State != "disabled" {
		t.Errorf("Failed to disable user")
	}
}

func TestUserService_ListUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"listusersresponse": {
				"count": 1,
				"user": [
					{
						"id": "27f2484f-5fe0-11ea-9a56-1e006800018c",
						"username": "admin",
						"firstname": "admin",
						"lastname": "cloud",
						"created": "2020-03-06T20:24:53+0000",
						"state": "enabled",
						"account": "admin",
						"accounttype": 1,
						"usersource": "native",
						"roleid": "16919363-5fe0-11ea-9a56-1e006800018c",
						"roletype": "Admin",
						"rolename": "Root Admin",
						"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
						"domain": "ROOT",
						"timezone": "PST",
						"apikey": "DN7uIwbqAMASSB1GUGHWvXqUlDDBf7H6A3XI-kNeXahW5LMUoqgfDEDWMv8zWKaj51fDozg8fjqc7tGCSozScA",
						"accountid": "27ef5ba2-5fe0-11ea-9a56-1e006800018c",
						"iscallerchilddomain": false,
						"isdefault": true
					}
				]
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.User.NewListUsersParams()
	params.SetId("cd2b5afa-db4d-4532-88ec-1356a273e534")
	resp, err := client.User.ListUsers(params)
	if err != nil {
		t.Errorf("Failed to list user details due to %v", err)
		return
	}

	if resp == nil || resp.Count != 1 {
		t.Errorf("Failed to list user details")
	}
}

func TestUserService_LockUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"lockuserresponse": {
			  "user": {
				"account": "kittycat",
				"accountid": "2841efe4-aba4-4abf-bfdc-c6d0825c29a6",
				"accounttype": 0,
				"created": "2020-06-26T18:21:57+0000",
				"domain": "saml",
				"domainid": "1d4a08f6-5c37-4385-9d95-fa4e654ecf4e",
				"email": "cat@cat.com",
				"firstname": "Ashley",
				"id": "3b7325b3-849c-4314-ad19-dd3c483b6d1a",
				"iscallerchilddomain": false,
				"isdefault": false,
				"lastname": "Cat",
				"roleid": "1691ce2c-5fe0-11ea-9a56-1e006800018c",
				"rolename": "User",
				"roletype": "User",
				"state": "locked",
				"timezone": "IST",
				"username": "ashley",
				"usersource": "saml2disabled"
			  }
			}
		}`
		fmt.Fprintln(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.User.NewLockUserParams("3b7325b3-849c-4314-ad19-dd3c483b6d1a")
	resp, err := client.User.LockUser(params)
	if err != nil {
		t.Errorf("Failed to lock user due to %v", err)
		return
	}
	fmt.Println(resp)
	if resp == nil || resp.State != "locked" {
		t.Errorf("Failed to lock user")
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"deleteuserresponse": {
				"success": true
			}
		}`
		fmt.Fprintf(writer, response)
	}))
	defer server.Close()
	client := NewClient(server.URL, "APIKEY", "SECRETKEY", true)
	params := client.User.NewDeleteUserParams("cd2b5afa-db4d-4532-88ec-1356a273e534")
	resp, err := client.User.DeleteUser(params)
	if err != nil {
		t.Errorf("Failed to delete user due to %v", err)
		return
	}

	if resp == nil || !resp.Success {
		t.Errorf("Failed to delete user")
	}
}

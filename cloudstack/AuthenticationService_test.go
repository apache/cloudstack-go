package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticationService_Login(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		loginResponse := `{
	"loginresponse": {
		"username": "admin",
		"userid": "27f2484f-5fe0-11ea-9a56-1e006800018c",
		"domainid": "e4874e10-5fdf-11ea-9a56-1e006800018c",
		"timeout": 3600,
		"account": "admin",
		"firstname": "admin",
		"lastname": "cloud",
		"type": "1",
		"timezone": "PST",
		"timezoneoffset": "-7.0",
		"registered": "false",
		"sessionkey": "umn5ciBuEdorc784g4pJr-U5POM"
	}
}	
	`
		fmt.Fprintf(writer, loginResponse)
	}))
	defer server.Close()

	client := newClient(server.URL, "APIKEY", "SECRETKEY", true, true)
	loginParams := client.Authentication.NewLoginParams("admin", "password")
	resp, err := client.Authentication.Login(loginParams)

	if err != nil {
		t.Errorf("Failed to login user 'admin', due to: %v", err)
	}

	if resp == nil {
		t.Errorf("Failed to login user 'admin'")
	}

	if resp.Username != "admin" {
		t.Errorf("Failed to login user 'admin'")
	}
}

func TestAuthenticationService_Logout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		logoutResponse := `{
	"logoutresponse": {
		"description": "success"
	}
}
	`
		fmt.Fprintln(writer, logoutResponse)
	}))

	defer server.Close()

	client := newClient(server.URL, "APIKEY", "SECRETKEY", true, true)
	logoutParams := client.Authentication.NewLogoutParams()
	logoutResp, err := client.Authentication.Logout(logoutParams)

	if err != nil {
		t.Errorf("Failed to logout user 'admin', due to: %v", err)
	}

	if logoutResp.Description != "success" {
		t.Errorf("Failed to logout user 'admin'")
	}
}

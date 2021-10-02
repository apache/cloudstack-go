package cloudstack

import (
	"testing"
)

var (
	CS_API_URL    = "http://localhost:8080/client/api"
	CS_API_KEY    = "valid-api-key"
	CS_SECRET_KEY = "valid-secret-key"
)

func TestCreateAsyncClient(t *testing.T) {
	client := NewAsyncClient(CS_API_URL, CS_API_KEY, CS_SECRET_KEY, true)

	if client == nil {
		t.Errorf("Failed to create Cloudstack Async Client")
	}
}

func TestCreateSyncClient(t *testing.T) {
	client := NewClient(CS_API_URL, CS_API_KEY, CS_SECRET_KEY, true)

	if client == nil {
		t.Errorf("Failed to create Cloudstack Client")
	}
}

package cloudstack

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListApis_SetName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		response := `{
			"listapisresponse": {
  "api": [
    {
      "description": "List registered keypairs",
      "isasync": false,
      "name": "listSSHKeyPairs",
      "params": [
        {
          "description": "",
          "length": 255,
          "name": "page",
          "required": false,
          "type": "integer"
        },
        {
          "description": "A key pair name to look for",
          "length": 255,
          "name": "name",
          "required": false,
          "type": "string"
        },
        {
          "description": "If set to false, list only resources belonging to the command's caller; if set to true - list resources that the caller is authorized to see. Default value is false",
          "length": 255,
          "name": "listall",
          "required": false,
          "type": "boolean"
        },
        {
          "description": "list objects by project",
          "length": 255,
          "name": "projectid",
          "related": "activateProject",
          "required": false,
          "type": "uuid"
        },
        {
          "description": "list resources by account. Must be used with the domainId parameter.",
          "length": 255,
          "name": "account",
          "required": false,
          "type": "string"
        },
        {
          "description": "A public key fingerprint to look for",
          "length": 255,
          "name": "fingerprint",
          "required": false,
          "type": "string"
        },
        {
          "description": "List by keyword",
          "length": 255,
          "name": "keyword",
          "required": false,
          "type": "string"
        },
        {
          "description": "",
          "length": 255,
          "name": "pagesize",
          "required": false,
          "type": "integer"
        },
        {
          "description": "list only resources belonging to the domain specified",
          "length": 255,
          "name": "domainid",
          "related": "listDomains",
          "required": false,
          "type": "uuid"
        },
        {
          "description": "defaults to false, but if true, lists all resources from the parent specified by the domainId till leaves.",
          "length": 255,
          "name": "isrecursive",
          "required": false,
          "type": "boolean"
        }
      ],
      "related": "",
      "response": [
        {},
        {
          "description": "Name of the keypair",
          "name": "name",
          "type": "string"
        },
        {
          "description": "Fingerprint of the public key",
          "name": "fingerprint",
          "type": "string"
        },
        {
          "description": "the domain name of the keypair owner",
          "name": "domain",
          "type": "string"
        },
        {
          "description": "true if the entity/resource has annotations",
          "name": "hasannotations",
          "type": "boolean"
        },
        {
          "description": "the domain id of the keypair owner",
          "name": "domainid",
          "type": "string"
        },
        {
          "description": "the current status of the latest async job acting on this object",
          "name": "jobstatus",
          "type": "integer"
        },
        {
          "description": "the UUID of the latest async job acting on this object",
          "name": "jobid",
          "type": "string"
        },
        {
          "description": "the owner of the keypair",
          "name": "account",
          "type": "string"
        },
        {},
        {
          "description": "ID of the ssh keypair",
          "name": "id",
          "type": "string"
        }
      ]
    }
  ],
  "count": 1
	}
}`
		fmt.Fprintln(writer, response)
	}))

	defer server.Close()

	client := newClient(server.URL, "APIKEY", "SECRETKEY", true, true)
	p := client.APIDiscovery.NewListApisParams()
	p.SetName("listSSHKeyPairs")
	resp, err := client.APIDiscovery.ListApis(p)

	if err != nil {
		t.Errorf("Failed to fetch listSSHKeyPairs API details, due to: %s", err)
	}

	if resp.Count != 1 {
		t.Errorf("Failed to fetch listSSHKeyPairs API details")
	}

	if resp.Apis[0].Name != "listSSHKeyPairs" {
		t.Errorf("Failed to fetch listSSHKeyPairs API details")
	}
}

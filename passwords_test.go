package itglue

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPasswords(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": [
				{
					"id": "123456",
					"type": "passwords",
					"attributes": {
						"organization-id": 123456,
						"organization-name": "EXAMPLE ORG",
						"resource-url": "HTTPS://EXAMPLE.COM",
						"restricted": true,
						"my-glue": false,
						"name": "NAME",
						"autofill-selectors": null,
						"username": "USERNAME",
						"url": null,
						"notes": null,
						"password-updated-at": "2020-11-24T03:40:20.000Z",
						"updated-by": null,
						"updated-by-type": null,
						"resource-id": null,
						"cached-resource-type-name": null,
						"cached-resource-name": null,
						"password-category-id": 12345,
						"password-category-name": "CATEGORY NAME",
						"created-at": "2020-11-24T03:40:20.000Z",
						"updated-at": "2020-11-24T03:40:20.000Z",
						"otp-enabled": false,
						"password-folder-id": null,
						"resource-type": null,
						"parent-url": null,
						"vault-id": null,
						"archived": false
					},
					"relationships": {}
				},
				{
					"id": "123456",
					"type": "passwords",
					"attributes": {
						"organization-id": 123456,
						"organization-name": "EXAMPLE ORG",
						"resource-url": "HTTPS://EXAMPLE.COM",
						"restricted": true,
						"my-glue": false,
						"name": "NAME",
						"autofill-selectors": null,
						"username": "USERNAME",
						"url": null,
						"notes": null,
						"password-updated-at": "2020-11-24T03:40:20.000Z",
						"updated-by": null,
						"updated-by-type": null,
						"resource-id": null,
						"cached-resource-type-name": null,
						"cached-resource-name": null,
						"password-category-id": 12345,
						"password-category-name": "CATEGORY NAME",
						"created-at": "2020-11-24T03:40:20.000Z",
						"updated-at": "2020-11-24T03:40:20.000Z",
						"otp-enabled": false,
						"password-folder-id": null,
						"resource-type": null,
						"parent-url": null,
						"vault-id": null,
						"archived": false
					},
					"relationships": {}
				}
			]
		}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	_, err := ITGAPI.GetPasswords(1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetPasswordsByID(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": {
				"id": "123456",
				"type": "passwords",
				"attributes": {
					"organization-id": 123456,
					"organization-name": "EXAMPLE ORG",
					"resource-url": "https://EXAMPLE.COM",
					"restricted": true,
					"my-glue": false,
					"name": "NAME",
					"autofill-selectors": null,
					"username": "USERNAME",
					"url": null,
					"notes": null,
					"password-updated-at": "2020-11-24T03:40:20.000Z",
					"updated-by": null,
					"updated-by-type": null,
					"resource-id": null,
					"cached-resource-type-name": null,
					"cached-resource-name": null,
					"password-category-id": 12345,
					"password-category-name": "TEST CATEGORY",
					"created-at": "2020-11-24T03:40:20.000Z",
					"updated-at": "2020-11-24T03:40:20.000Z",
					"otp-enabled": false,
					"password-folder-id": null,
					"resource-type": null,
					"parent-url": null,
					"vault-id": null,
					"archived": false,
					"password": "EXAMPLE_PASSWORD"
				},
				"relationships": {}
			}
		}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	_, err := ITGAPI.GetPasswordsByID(1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPostPassword(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": {
				"id": "123456",
				"type": "passwords",
				"attributes": {
					"organization-id": 123456,
					"organization-name": "EXAMPLE ORG",
					"resource-url": "https://EXAMPLE.COM",
					"restricted": true,
					"my-glue": false,
					"name": "NAME",
					"autofill-selectors": null,
					"username": "USERNAME",
					"url": null,
					"notes": null,
					"password-updated-at": "2020-11-24T03:40:20.000Z",
					"updated-by": null,
					"updated-by-type": null,
					"resource-id": null,
					"cached-resource-type-name": null,
					"cached-resource-name": null,
					"password-category-id": 12345,
					"password-category-name": "TEST CATEGORY",
					"created-at": "2020-11-24T03:40:20.000Z",
					"updated-at": "2020-11-24T03:40:20.000Z",
					"otp-enabled": false,
					"password-folder-id": null,
					"resource-type": null,
					"parent-url": null,
					"vault-id": null,
					"archived": false,
					"password": "EXAMPLE_PASSWORD"
				},
				"relationships": {}
			}
		}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	p := &Password{}
	p.Data.Attributes.Username = "test"
	p.Data.Attributes.Password = "test2"
	_, err := ITGAPI.PostPassword(p)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPatchPassword(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": {
				"id": "123456",
				"type": "passwords",
				"attributes": {
					"organization-id": 123456,
					"organization-name": "EXAMPLE ORG",
					"resource-url": "https://EXAMPLE.COM",
					"restricted": true,
					"my-glue": false,
					"name": "NAME",
					"autofill-selectors": null,
					"username": "USERNAME",
					"url": null,
					"notes": null,
					"password-updated-at": "2020-11-24T03:40:20.000Z",
					"updated-by": null,
					"updated-by-type": null,
					"resource-id": null,
					"cached-resource-type-name": null,
					"cached-resource-name": null,
					"password-category-id": 12345,
					"password-category-name": "TEST CATEGORY",
					"created-at": "2020-11-24T03:40:20.000Z",
					"updated-at": "2020-11-24T03:40:20.000Z",
					"otp-enabled": false,
					"password-folder-id": null,
					"resource-type": null,
					"parent-url": null,
					"vault-id": null,
					"archived": false,
					"password": "EXAMPLE_PASSWORD"
				},
				"relationships": {}
			}
		}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	p := &Password{}
	p.Data.Attributes.Username = "test"
	p.Data.Attributes.Password = "test2"
	_, err := ITGAPI.PatchPassword(1, p)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestDeletePassword(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	_, err := ITGAPI.DeletePassword(1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

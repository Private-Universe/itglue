package itglue

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetOrganizationTypes(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": [
				{
					"id": "1",
					"type": "organization-types",
					"attributes": {
						"name": "Type 1",
						"created-at": "2019-06-21T03:18:36.000Z",
						"updated-at": "2019-06-21T03:18:36.000Z",
						"synced": true
					}
				},
				{
					"id": "2",
					"type": "Type 2",
					"attributes": {
						"name": "Competitor",
						"created-at": "2019-06-21T03:18:36.000Z",
						"updated-at": "2019-06-21T03:18:36.000Z",
						"synced": false
					}
				}
			],
			"meta": {
				"current-page": 1,
				"next-page": null,
				"prev-page": null,
				"total-pages": 1,
				"total-count": 2,
				"filters": {}
			}
		}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	_, err := ITGAPI.GetOrganizationTypes(1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

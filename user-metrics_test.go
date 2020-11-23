package itglue

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserMetrics(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": [
				{
					"id": "1",
					"type": "user-metrics",
					"attributes": {
						"user-id": 1,
						"organization-id": 1,
						"created": 0,
						"viewed": 1,
						"edited": 0,
						"deleted": 0,
						"date": "2020-11-23",
						"resource-type": "Resource 1"
					}
				},
				{
					"id": "1",
					"type": "user-metrics",
					"attributes": {
						"user-id": 1,
						"organization-id": 1,
						"created": 0,
						"viewed": 1,
						"edited": 0,
						"deleted": 0,
						"date": "2020-11-23",
						"resource-type": "Resource 2"
					}
				}
			],
			"meta": {
				"current-page": 1,
				"next-page": 1,
				"prev-page": null,
				"total-pages": 1,
				"total-count": 2,
				"filters": {}
			},
			"links": {
				"self": "",
				"next": "",
				"last": ""
			}
		}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	_, err := ITGAPI.GetUserMetrics(1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

package itglue

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFlexibleAssetTypes(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": [{
				"id": "1",
				"type": "flexible-asset-types",
				"attributes": {
					"name": "Dummy Name",
					"description": "Dummy Description",
					"created-at": "2019-10-20T03:07:52.000Z",
					"updated-at": "2019-11-20T03:07:52.000Z",
					"icon": "dummy_icon",
					"enabled": true
				},
				"relationships": {}
			}],
			"meta": {
				"current-page": 1,
				"next-page": null,
				"prev-page": null,
				"total-pages": 1,
				"total-count": 1,
				"filters": {}
			}
		}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	_, err := ITGAPI.GetFlexibleAssetTypes(1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

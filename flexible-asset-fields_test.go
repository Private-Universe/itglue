package itglue

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFlexibleFields(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": [
				{
					"id": "1",
					"type": "flexible-asset-fields",
					"attributes": {
						"order": 1,
						"name": "Name 1",
						"kind": "Text",
						"hint": null,
						"decimals": 0,
						"tag-type": null,
						"required": false,
						"use-for-title": true,
						"expiration": false,
						"show-in-list": true,
						"name-key": "name-1",
						"created-at": "2019-06-11T23:57:04.000Z",
						"updated-at": "2019-06-11T23:57:04.000Z",
						"flexible-asset-type-id": 1,
						"default-value": null
					},
					"relationships": {
						"flexible-asset-type": {
							"data": {
								"id": "1",
								"type": "flexible-asset-types"
							}
						}
					}
				},
				{
					"id": "2",
					"type": "flexible-asset-fields",
					"attributes": {
						"order": 2,
						"name": "Name 2",
						"kind": "Tag",
						"hint": null,
						"decimals": 0,
						"tag-type": "name-2",
						"required": false,
						"use-for-title": false,
						"expiration": false,
						"show-in-list": false,
						"name-key": "name-2",
						"created-at": "2019-06-11T23:57:04.000Z",
						"updated-at": "2019-06-11T23:57:04.000Z",
						"flexible-asset-type-id": 2,
						"default-value": null
					},
					"relationships": {
						"flexible-asset-type": {
							"data": {
								"id": "2",
								"type": "flexible-asset-types"
							}
						}
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
	_, err := ITGAPI.GetFlexibleAssetFields(1, 1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

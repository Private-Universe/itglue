package itglue

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFlexibleAssets(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": [
				{
					"id": "1",
					"type": "flexible-assets",
					"attributes": {
						"organization-id": 1,
						"organization-name": "Organization",
						"resource-url": "https://example.com",
						"restricted": false,
						"my-glue": false,
						"flexible-asset-type-id": 1,
						"flexible-asset-type-name": "Example 1",
						"name": "test",
						"traits": {
							"trait-1": "test",
							"trait-2": "https://example.com",
							"trait-3": "test 2",
							"trait-4": 1,
							"trait-5": {
								"type": "FlexibleAssetTypes",
								"values": [
									{
										"resource-url": "https://example.com",
										"id": 1,
										"name": "Example 2"
									}
								]
							}
						},
						"archived": false,
						"created-at": "2020-11-24T03:40:20.000Z",
						"updated-at": "2020-11-24T03:40:21.000Z"
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
	_, err := ITGAPI.GetFlexibleAssets(1, 1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetFlexibleAssetsByID(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": {
				"id": "1",
				"type": "flexible-assets",
				"attributes": {
					"organization-id": 1,
					"organization-name": "Organization",
					"resource-url": "https://example.com",
					"restricted": false,
					"my-glue": false,
					"flexible-asset-type-id": 1,
					"flexible-asset-type-name": "Example 1",
					"name": "test",
					"traits": {
						"trait-1": "test",
						"trait-2": "https://example.com",
						"trait-3": "test 2",
						"trait-4": 1,
						"trait-5": {
							"type": "FlexibleAssetTypes",
							"values": [
								{
									"resource-url": "https://example.com",
									"id": 1,
									"name": "Example 2"
								}
							]
						}
					},
					"archived": false,
					"created-at": "2020-11-24T03:40:20.000Z",
					"updated-at": "2020-11-24T03:40:21.000Z"
				},
				"relationships": {}
			}
		}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	_, err := ITGAPI.GetFlexibleAssetsByID(1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetFlexibleAssetsByName(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": [
				{
					"id": "1",
					"type": "flexible-assets",
					"attributes": {
						"organization-id": 1,
						"organization-name": "Organization",
						"resource-url": "https://example.com",
						"restricted": false,
						"my-glue": false,
						"flexible-asset-type-id": 1,
						"flexible-asset-type-name": "Example 1",
						"name": "test",
						"traits": {
							"trait-1": "test",
							"trait-2": "https://example.com",
							"trait-3": "test 2",
							"trait-4": 1,
							"trait-5": {
								"type": "FlexibleAssetTypes",
								"values": [
									{
										"resource-url": "https://example.com",
										"id": 1,
										"name": "Example 2"
									}
								]
							}
						},
						"archived": false,
						"created-at": "2020-11-24T03:40:20.000Z",
						"updated-at": "2020-11-24T03:40:21.000Z"
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
	_, err := ITGAPI.GetFlexibleAssetsByName(1, "name", 1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetFlexibleAssetsByOrganizationID(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": [
				{
					"id": "1",
					"type": "flexible-assets",
					"attributes": {
						"organization-id": 1,
						"organization-name": "Organization",
						"resource-url": "https://example.com",
						"restricted": false,
						"my-glue": false,
						"flexible-asset-type-id": 1,
						"flexible-asset-type-name": "Example 1",
						"name": "test",
						"traits": {
							"trait-1": "test",
							"trait-2": "https://example.com",
							"trait-3": "test 2",
							"trait-4": 1,
							"trait-5": {
								"type": "FlexibleAssetTypes",
								"values": [
									{
										"resource-url": "https://example.com",
										"id": 1,
										"name": "Example 2"
									}
								]
							}
						},
						"archived": false,
						"created-at": "2020-11-24T03:40:20.000Z",
						"updated-at": "2020-11-24T03:40:21.000Z"
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
	_, err := ITGAPI.GetFlexibleAssetsByOrganizationID(1, 1, 1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetFlexibleAssetsByOrganizationIDAndName(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": [
				{
					"id": "1",
					"type": "flexible-assets",
					"attributes": {
						"organization-id": 1,
						"organization-name": "Organization",
						"resource-url": "https://example.com",
						"restricted": false,
						"my-glue": false,
						"flexible-asset-type-id": 1,
						"flexible-asset-type-name": "Example 1",
						"name": "test",
						"traits": {
							"trait-1": "test",
							"trait-2": "https://example.com",
							"trait-3": "test 2",
							"trait-4": 1,
							"trait-5": {
								"type": "FlexibleAssetTypes",
								"values": [
									{
										"resource-url": "https://example.com",
										"id": 1,
										"name": "Example 2"
									}
								]
							}
						},
						"archived": false,
						"created-at": "2020-11-24T03:40:20.000Z",
						"updated-at": "2020-11-24T03:40:21.000Z"
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
	_, err := ITGAPI.GetFlexibleAssetsByOrganizationIDAndName(1, 1, "name", 1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPostFlexibleAsset(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": {
				"id": "1",
				"type": "flexible-assets",
				"attributes": {
					"organization-id": 1,
					"organization-name": "Organization",
					"resource-url": "https://example.com",
					"restricted": false,
					"my-glue": false,
					"flexible-asset-type-id": 1,
					"flexible-asset-type-name": "Example 1",
					"name": "test",
					"traits": {
						"trait-1": "test",
						"trait-2": "https://example.com",
						"trait-3": "test 2",
						"trait-4": 1,
						"trait-5": {
							"type": "FlexibleAssetTypes",
							"values": [
								{
									"resource-url": "https://example.com",
									"id": 1,
									"name": "Example 2"
								}
							]
						}
					},
					"archived": false,
					"created-at": "2020-11-24T03:40:20.000Z",
					"updated-at": "2020-11-24T03:40:21.000Z"
				},
				"relationships": {}
			}
		}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	a := &FlexibleAsset{}
	a.Data.Attributes.OrganizationID = 1
	a.Data.Attributes.FlexibleAssetTypeID = 1
	a.Data.Attributes.Traits = map[string]interface{}{
		"test":  "test",
		"test2": "test2",
	}
	_, err := ITGAPI.PostFlexibleAsset(a)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPatchFlexibleAsset(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": {
				"id": "1",
				"type": "flexible-assets",
				"attributes": {
					"organization-id": 1,
					"organization-name": "Organization",
					"resource-url": "https://example.com",
					"restricted": false,
					"my-glue": false,
					"flexible-asset-type-id": 1,
					"flexible-asset-type-name": "Example 1",
					"name": "test",
					"traits": {
						"trait-1": "test",
						"trait-2": "https://example.com",
						"trait-3": "test 2",
						"trait-4": 1,
						"trait-5": {
							"type": "FlexibleAssetTypes",
							"values": [
								{
									"resource-url": "https://example.com",
									"id": 1,
									"name": "Example 2"
								}
							]
						}
					},
					"archived": false,
					"created-at": "2020-11-24T03:40:20.000Z",
					"updated-at": "2020-11-24T03:40:21.000Z"
				},
				"relationships": {}
			}
		}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	a := &FlexibleAsset{}
	a.Data.Attributes.OrganizationID = 1
	a.Data.Attributes.FlexibleAssetTypeID = 1
	a.Data.Attributes.Traits = map[string]interface{}{
		"test":  "test",
		"test2": "test2",
	}
	_, err := ITGAPI.PatchFlexibleAsset(1, a)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestDeleteFlexibleAsset(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	_, err := ITGAPI.DeleteFlexibleAsset(1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

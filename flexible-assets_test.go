package itglue

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFlexibleAssetsJSON(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{"fake IT Glue json string"}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	_, err := ITGAPI.GetFlexibleAssetsJSON(1, 1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetFlexibleAssetsJSONByID(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{"fake IT Glue json string"}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	_, err := ITGAPI.GetFlexibleAssetsJSONByID(1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetFlexibleAssetsJSONByName(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{"fake IT Glue json string"}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	_, err := ITGAPI.GetFlexibleAssetsJSONByName(1, "name", 1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetFlexibleAssetsJSONByOrganizationID(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{"fake IT Glue json string"}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	_, err := ITGAPI.GetFlexibleAssetsJSONByOrganizationID(1, 1, 1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestGetFlexibleAssetsJSONByOrganizationIDAndName(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{"fake IT Glue json string"}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	_, err := ITGAPI.GetFlexibleAssetsJSONByOrganizationIDAndName(1, 1, "name", 1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPostFlexibleAsset(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{"fake IT Glue json string"}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	data := make([]byte, 5)
	_, err := ITGAPI.PostFlexibleAsset(data)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestPatchFlexibleAsset(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{"fake IT Glue json string"}`)
	}))
	// Close the server when test finishes
	defer server.Close()
	glueURL := server.URL
	ITGAPI := &ITGAPI{Site: glueURL, APIKey: "apiKey"}
	data := make([]byte, 5)
	_, err := ITGAPI.PatchFlexibleAsset(1, data)
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestDeleteFlexibleAsset(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{"fake IT Glue json string"}`)
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

package itglue

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, `{
			"data": [
				{
					"id": "1",
					"type": "users",
					"attributes": {
						"first-name": "Test",
						"last-name": "User 1",
						"name": "Test User 1",
						"role-name": "Lite",
						"email": "testuser1@example.com",
						"invitation-sent-at": "2019-06-11T23:57:04.000Z",
						"invitation-accepted-at": "2019-05-21T03:18:36.000Z",
						"current-sign-in-at": "2019-05-21T03:18:36.000Z",
						"current-sign-in-ip": "1.1.1.1",
						"last-sign-in-at": "2019-05-21T03:18:36.000Z",
						"last-sign-in-ip": "1.1.1.1",
						"reputation": 1,
						"created-at": "2019-05-21T03:18:36.000Z",
						"updated-at": "2019-05-21T03:18:36.000Z",
						"my-glue-account-id": null,
						"avatar": "",
						"my-glue": false
					}
				},
				{
					"id": "2",
					"type": "users",
					"attributes": {
						"first-name": "Test",
						"last-name": "User 2",
						"name": "Test User 2",
						"role-name": "Administrator",
						"email": "testuser2@example.com",
						"invitation-sent-at": "2019-06-20T06:59:18.000Z",
						"invitation-accepted-at": "2019-05-21T03:18:36.000Z",
						"current-sign-in-at": "2019-05-21T03:18:36.000Z",
						"current-sign-in-ip": "1.1.1.1",
						"last-sign-in-at": "2019-05-21T03:18:36.000Z",
						"last-sign-in-ip": "1.1.1.1",
						"reputation": 1,
						"created-at": "2019-05-21T03:18:36.000Z",
						"updated-at": "2019-05-21T03:18:36.000Z",
						"my-glue-account-id": null,
						"avatar": "",
						"my-glue": false
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
	_, err := ITGAPI.GetUsers(1)
	if err != nil {
		t.Errorf("%s", err)
	}
}

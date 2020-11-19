package itglue

import (
	"encoding/json"
	"fmt"
	"time"
)

type UserData struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		FirstName            string      `json:"first-name"`
		LastName             string      `json:"last-name"`
		Name                 string      `json:"name"`
		RoleName             string      `json:"role-name"`
		Email                string      `json:"email"`
		InvitationSentAt     time.Time   `json:"invitation-sent-at"`
		InvitationAcceptedAt time.Time   `json:"invitation-accepted-at"`
		CurrentSignInAt      time.Time   `json:"current-sign-in-at"`
		CurrentSignInIP      string      `json:"current-sign-in-ip"`
		LastSignInAt         time.Time   `json:"last-sign-in-at"`
		LastSignInIP         string      `json:"last-sign-in-ip"`
		Reputation           int         `json:"reputation"`
		CreatedAt            time.Time   `json:"created-at"`
		UpdatedAt            time.Time   `json:"updated-at"`
		MyGlueAccountID      interface{} `json:"my-glue-account-id"`
		Avatar               string      `json:"avatar"`
		MyGlue               bool        `json:"my-glue"`
	} `json:"attributes"`
}

type User struct {
	Data  struct{ UserData } `json:"data"`
	Meta  struct{ Metadata } `json:"meta"`
	Links struct{ Links }    `json:"links"`
}

type UserList struct {
	Data  []struct{ UserData } `json:"data"`
	Meta  struct{ Metadata }   `json:"meta"`
	Links struct{ Links }      `json:"links"`
}

func (itg *ITGAPI) GetUsers(pageNumber int) (*UserList, error) {
	req := itg.NewRequest("/users", "GET", nil)
	req.Page = pageNumber
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	users := &UserList{}
	err = json.Unmarshal(req.Body, users)
	if err != nil {
		return nil, fmt.Errorf("could not get users: %s", err)
	}

	return users, nil
}

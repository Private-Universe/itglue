package itglue

import (
	"encoding/json"
	"fmt"
	"time"
)

type PasswordData struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		OrganizationID         int         `json:"organization-id"`
		OrganizationName       string      `json:"organization-name"`
		ResourceURL            string      `json:"resource-url"`
		Restricted             bool        `json:"restricted"`
		MyGlue                 bool        `json:"my-glue"`
		Name                   string      `json:"name"`
		AutofillSelectors      interface{} `json:"autofill-selectors"`
		Username               string      `json:"username"`
		URL                    string      `json:"url"`
		Notes                  string      `json:"notes"`
		PasswordUpdatedAt      time.Time   `json:"password-updated-at"`
		UpdatedBy              int         `json:"updated-by"`
		UpdatedByType          string      `json:"updated-by-type"`
		ResourceID             int         `json:"resource-id"`
		CachedResourceTypeName string      `json:"cached-resource-type-name"`
		CachedResourceName     string      `json:"cached-resource-name"`
		PasswordCategoryID     int         `json:"password-category-id"`
		PasswordCategoryName   string      `json:"password-category-name"`
		CreatedAt              time.Time   `json:"created-at"`
		UpdatedAt              time.Time   `json:"updated-at"`
		OtpEnabled             bool        `json:"otp-enabled"`
		PasswordFolderID       int         `json:"password-folder-id"`
		ResourceType           string      `json:"resource-type"`
		ParentURL              string      `json:"parent-url"`
		VaultID                int         `json:"vault-id"`
		Archived               bool        `json:"archived"`
		Password               string      `json:"password"`
	} `json:"attributes"`
	Relationships struct {
	} `json:"relationships"`
}

type Password struct {
	Data  struct{ PasswordData } `json:"data"`
	Meta  struct{ Metadata }     `json:"meta"`
	Links struct{ Links }        `json:"links"`
}

type PasswordList struct {
	Data  []struct{ PasswordData } `json:"data"`
	Meta  struct{ Metadata }       `json:"meta"`
	Links struct{ Links }          `json:"links"`
}

//GetPasswords returns passwords
func (itg *ITGAPI) GetPasswords(pageNumber int) (*PasswordList, error) {
	req := itg.NewRequest("/passwords", "GET", nil)
	req.Page = pageNumber

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	password := &PasswordList{}
	err = json.Unmarshal(req.Body, password)
	if err != nil {
		return nil, fmt.Errorf("could not get password: %s", err)
	}

	if len((*password).Data) == 0 {
		return nil, fmt.Errorf("ITG returned no passwords")
	}

	return password, nil
}

//GetPasswordsByID returns passwords
func (itg *ITGAPI) GetPasswordsByID(passwordID int) (*Password, error) {
	req := itg.NewRequest(fmt.Sprintf("/passwords/%d", passwordID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	password := &Password{}
	err = json.Unmarshal(req.Body, password)
	if err != nil {
		return nil, fmt.Errorf("could not get password: %s", err)
	}

	return password, nil
}

//PostPassword create new asset
func (itg *ITGAPI) PostPassword(asset *Password) (*Password, error) {
	a, err := json.Marshal(asset)
	if err != nil {
		return nil, fmt.Errorf("could not decode asset: %s", err)
	}

	req := itg.NewRequest("/passwords", "POST", a)

	err = req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	password := &Password{}
	err = json.Unmarshal(req.Body, password)
	if err != nil {
		return nil, fmt.Errorf("could not get password: %s", err)
	}

	return password, nil
}

//PatchPassword update asset. Any attributes you don't specify will remain unchanged.
func (itg *ITGAPI) PatchPassword(passwordID int, asset *Password) (*Password, error) {
	a, err := json.Marshal(asset)
	if err != nil {
		return nil, fmt.Errorf("could not decode asset: %s", err)
	}

	req := itg.NewRequest(fmt.Sprintf("/passwords/%d", passwordID), "PATCH", a)

	err = req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	password := &Password{}
	err = json.Unmarshal(req.Body, password)
	if err != nil {
		return nil, fmt.Errorf("could not get password: %s", err)
	}

	return password, nil
}

//DeletePassword delete asset
func (itg *ITGAPI) DeletePassword(passwordID int) ([]byte, error) {
	req := itg.NewRequest(fmt.Sprintf("/passwords/%d", passwordID), "DELETE", nil)

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	return req.Body, nil
}

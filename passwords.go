package itglue

import (
	"encoding/json"
	"fmt"
	"time"
)

// PasswordData struct, pointers to have omitempty work for time.Time
type PasswordData struct {
	ID         string `json:"id,omitempty"`
	Type       string `json:"type,omitempty"`
	Attributes struct {
		OrganizationID         int         `json:"organization-id,omitempty"`
		OrganizationName       string      `json:"organization-name,omitempty"`
		ResourceURL            string      `json:"resource-url,omitempty"`
		Restricted             bool        `json:"restricted,omitempty"`
		MyGlue                 bool        `json:"my-glue,omitempty"`
		Name                   string      `json:"name,omitempty"`
		AutofillSelectors      interface{} `json:"autofill-selectors,omitempty"`
		Username               string      `json:"username,omitempty"`
		URL                    string      `json:"url,omitempty"`
		Notes                  string      `json:"notes,omitempty"`
		PasswordUpdatedAt      *time.Time  `json:"password-updated-at,omitempty"`
		UpdatedBy              int         `json:"updated-by,omitempty"`
		UpdatedByType          string      `json:"updated-by-type,omitempty"`
		ResourceID             int         `json:"resource-id,omitempty"`
		CachedResourceTypeName string      `json:"cached-resource-type-name,omitempty"`
		CachedResourceName     string      `json:"cached-resource-name,omitempty"`
		PasswordCategoryID     int         `json:"password-category-id,omitempty"`
		PasswordCategoryName   string      `json:"password-category-name,omitempty"`
		CreatedAt              *time.Time  `json:"created-at,omitempty"`
		UpdatedAt              *time.Time  `json:"updated-at,omitempty"`
		OtpEnabled             bool        `json:"otp-enabled,omitempty"`
		PasswordFolderID       int         `json:"password-folder-id,omitempty"`
		ResourceType           string      `json:"resource-type,omitempty"`
		ParentURL              string      `json:"parent-url,omitempty"`
		VaultID                int         `json:"vault-id,omitempty"`
		Archived               bool        `json:"archived,omitempty"`
		Password               string      `json:"password,omitempty"`
	} `json:"attributes,omitempty"`
	Relationships struct {
	} `json:"relationships,omitempty"`
}

type Password struct {
	Data  struct{ PasswordData } `json:"data,omitempty"`
	Meta  struct{ Metadata }     `json:"meta,omitempty"`
	Links struct{ Links }        `json:"links,omitempty"`
}

type PasswordList struct {
	Data  []struct{ PasswordData } `json:"data,omitempty"`
	Meta  struct{ Metadata }       `json:"meta,omitempty"`
	Links struct{ Links }          `json:"links,omitempty"`
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

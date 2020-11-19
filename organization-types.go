package itglue

import (
	"encoding/json"
	"fmt"
	"time"
)

type OrganizationTypeData struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created-at"`
		UpdatedAt time.Time `json:"updated-at"`
		Synced    bool      `json:"synced"`
	} `json:"attributes"`
}

type OrganizationType struct {
	Data  struct{ OrganizationTypeData } `json:"data"`
	Meta  struct{ Metadata }             `json:"meta"`
	Links struct{ Links }                `json:"links"`
}

type OrganizationTypeList struct {
	Data  []struct{ OrganizationTypeData } `json:"data"`
	Meta  struct{ Metadata }               `json:"meta"`
	Links struct{ Links }                  `json:"links"`
}

///organization_types
func (itg *ITGAPI) GetOrganizationTypes(pageNumber int) (*OrganizationTypeList, error) {
	req := itg.NewRequest("/organization_types", "GET", nil)
	req.Page = pageNumber
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	organizationTypes := &OrganizationTypeList{}
	err = json.Unmarshal(req.Body, organizationTypes)
	if err != nil {
		return nil, fmt.Errorf("could not get organization types: %s", err)
	}

	return organizationTypes, nil
}

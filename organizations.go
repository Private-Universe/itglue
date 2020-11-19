package itglue

import (
	"encoding/json"
	"fmt"
)

//OrganizationData contains the schema of an Organization in IT Glue without the "data" wrapper.
//This allows us to reuse the schema when data is either a JSON object or an array, depending on what results are returned
type OrganizationData struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name                   string `json:"name"`
		Description            string `json:"description"`
		OrganizationTypeID     int    `json:"organization-type-id"`
		OrganizationTypeName   string `json:"organization-type-name"`
		OrganizationStatusID   int    `json:"organization-status-id"`
		OrganizationStatusName string `json:"organization-status-name"`
		Logo                   string `json:"logo"`
		QuickNotes             string `json:"quick-notes"`
		ShortName              string `json:"short-name"`
		CreatedAt              string `json:"created-at"`
		UpdatedAt              string `json:"updated-at"`
	} `json:"attributes"`
}

//Organization contains a single Organization
type Organization struct {
	Data struct{ OrganizationData } `json:"data"`
	Meta struct{ Metadata }         `json:"meta"`
}

//OrganizationList contains a slice of Organizations
type OrganizationList struct {
	Data []struct{ OrganizationData } `json:"data"`
	Meta struct{ Metadata }           `json:"meta"`
}

//GetOrganizationByID expects an ITG organization ID
//Returns a pointer to an Organization struct
func (itg *ITGAPI) GetOrganizationByID(organizationID int) (*Organization, error) {
	req := itg.NewRequest(fmt.Sprintf("/organizations/%d", organizationID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	organization := &Organization{}
	err = json.Unmarshal(req.Body, organization)
	if err != nil {
		return nil, fmt.Errorf("could not get organization: %s", err)
	}

	return organization, nil
}

//GetOrganizationByName expects an exact matching organization name and returns an OrganizationList
func (itg *ITGAPI) GetOrganizationByName(organizationName string, pageNumber int) (*OrganizationList, error) {
	req := itg.NewRequest("/organizations", "GET", nil)
	req.Page = pageNumber
	req.URLValues.Add("filter[name]", organizationName)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	organization := &OrganizationList{}
	err = json.Unmarshal(req.Body, organization)
	if err != nil {
		return nil, fmt.Errorf("could not get organization: %s", err)
	}

	if len((*organization).Data) == 0 {
		return nil, fmt.Errorf("ITG returned no results for %s", organizationName)
	}

	return organization, nil
}

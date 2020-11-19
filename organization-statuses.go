package itglue

import (
	"encoding/json"
	"fmt"
	"time"
)

type OrganizationStatusData struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created-at"`
		UpdatedAt time.Time `json:"updated-at"`
		Synced    bool      `json:"synced"`
	} `json:"attributes"`
}

type OrganizationStatus struct {
	Data  struct{ OrganizationStatusData } `json:"data"`
	Meta  struct{ Metadata }               `json:"meta"`
	Links struct{ Links }                  `json:"links"`
}

type OrganizationStatusList struct {
	Data  []struct{ OrganizationStatusData } `json:"data"`
	Meta  struct{ Metadata }                 `json:"meta"`
	Links struct{ Links }                    `json:"links"`
}

func (itg *ITGAPI) GetOrganizationStatuses(pageNumber int) (*OrganizationStatusList, error) {
	req := itg.NewRequest("/organization_statuses", "GET", nil)
	req.Page = pageNumber
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	organizationStatuses := &OrganizationStatusList{}
	err = json.Unmarshal(req.Body, organizationStatuses)
	if err != nil {
		return nil, fmt.Errorf("could not get organization types: %s", err)
	}

	return organizationStatuses, nil

}

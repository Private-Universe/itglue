package itglue

import (
	"encoding/json"
	"fmt"
)

type UserMetricData struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		UserID         int    `json:"user-id"`
		OrganizationID int    `json:"organization-id"`
		Created        int    `json:"created"`
		Viewed         int    `json:"viewed"`
		Edited         int    `json:"edited"`
		Deleted        int    `json:"deleted"`
		Date           string `json:"date"`
		ResourceType   string `json:"resource-type"`
	} `json:"attributes"`
}

type UserMetric struct {
	Data  struct{ UserMetricData } `json:"data"`
	Meta  struct{ Metadata }       `json:"meta"`
	Links struct{ Links }          `json:"links"`
}

type UserMetricList struct {
	Data  []struct{ UserMetricData } `json:"data"`
	Meta  struct{ Metadata }         `json:"meta"`
	Links struct{ Links }            `json:"links"`
}

func (itg *ITGAPI) GetUserMetrics(pageNumber int) (*UserMetricList, error) {
	req := itg.NewRequest("/user_metrics", "GET", nil)
	req.Page = pageNumber
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	userMetrics := &UserMetricList{}
	err = json.Unmarshal(req.Body, userMetrics)
	if err != nil {
		return nil, fmt.Errorf("could not get users: %s", err)
	}

	return userMetrics, nil
}

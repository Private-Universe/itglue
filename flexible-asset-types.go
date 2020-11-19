package itglue

import (
	"encoding/json"
	"fmt"
	"time"
)

type FlexibleAssetTypeData struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created-at"`
		UpdatedAt   time.Time `json:"updated-at"`
		Icon        string    `json:"icon"`
		Enabled     bool      `json:"enabled"`
	} `json:"attributes"`
	Relationships struct {
	} `json:"relationships"`
}

type FlexibleAssetType struct {
	Data  struct{ FlexibleAssetTypeData } `json:"data"`
	Meta  struct{ Metadata }              `json:"meta"`
	Links struct{ Links }                 `json:"links"`
}

type FlexibleAssetTypeList struct {
	Data  []struct{ FlexibleAssetTypeData } `json:"data"`
	Meta  struct{ Metadata }                `json:"meta"`
	Links struct{ Links }                   `json:"links"`
}

func (itg *ITGAPI) GetFlexibleAssetTypes(pageNumber int) (*FlexibleAssetTypeList, error) {
	req := itg.NewRequest("/flexible_asset_types", "GET", nil)
	req.Page = pageNumber
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	flexibleAssetTypes := &FlexibleAssetTypeList{}
	err = json.Unmarshal(req.Body, flexibleAssetTypes)
	if err != nil {
		return nil, fmt.Errorf("could not get flexible asset types: %s", err)
	}

	return flexibleAssetTypes, nil
}

package itglue

import (
	"encoding/json"
	"fmt"
	"time"
)

type FlexibleAssetFieldData struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Order               int         `json:"order"`
		Name                string      `json:"name"`
		Kind                string      `json:"kind"`
		Hint                interface{} `json:"hint"`
		Decimals            int         `json:"decimals"`
		TagType             interface{} `json:"tag-type"`
		Required            bool        `json:"required"`
		UseForTitle         bool        `json:"use-for-title"`
		Expiration          bool        `json:"expiration"`
		ShowInList          bool        `json:"show-in-list"`
		NameKey             string      `json:"name-key"`
		CreatedAt           time.Time   `json:"created-at"`
		UpdatedAt           time.Time   `json:"updated-at"`
		FlexibleAssetTypeID int         `json:"flexible-asset-type-id"`
		DefaultValue        interface{} `json:"default-value"`
	} `json:"attributes"`
	Relationships struct {
		FlexibleAssetType struct {
			FlexibleAssetType
		} `json:"flexible-asset-type"`
	} `json:"relationships"`
}

type FlexibleAssetField struct {
	Data  struct{ FlexibleAssetFieldData } `json:"data"`
	Meta  struct{ Metadata }               `json:"meta"`
	Links struct{ Links }                  `json:"links"`
}

type FlexibleAssetFieldList struct {
	Data  []struct{ FlexibleAssetFieldData } `json:"data"`
	Meta  struct{ Metadata }                 `json:"meta"`
	Links struct{ Links }                    `json:"links"`
}

func (itg *ITGAPI) GetFlexibleAssetFields(flexibleAssetTypeID int, pageNumber int) (*FlexibleAssetFieldList, error) {
	req := itg.NewRequest(fmt.Sprintf("/flexible_asset_types/%d/relationships/flexible_asset_fields", flexibleAssetTypeID), "GET", nil)
	req.Page = pageNumber
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	flexibleAssetFields := &FlexibleAssetFieldList{}
	err = json.Unmarshal(req.Body, flexibleAssetFields)
	if err != nil {
		return nil, fmt.Errorf("could not get flexible asset fields: %s", err)
	}

	return flexibleAssetFields, nil
}

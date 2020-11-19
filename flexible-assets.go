package itglue

import (
	"fmt"
	"time"
)

type FlexibleAssetData struct {
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

type FlexibleAsset struct {
	Data  struct{ FlexibleAssetData } `json:"data"`
	Meta  struct{ Metadata }          `json:"meta"`
	Links struct{ Links }             `json:"links"`
}

type FlexibleAssetList struct {
	Data  []struct{ FlexibleAssetData } `json:"data"`
	Meta  struct{ Metadata }            `json:"meta"`
	Links struct{ Links }               `json:"links"`
}

//GetFlexibleAssetsJSON is a special function.  All flexible assets will return different data depending
//on which fields make up the flexible asset.  Because of this, this function simply returns the JSON bytes.
func (itg *ITGAPI) GetFlexibleAssetsJSON(flexibleAssetTypeID int, pageNumber int) ([]byte, error) {
	req := itg.NewRequest("/flexible_assets", "GET", nil)
	req.Page = pageNumber
	req.RawURLValues = fmt.Sprintf("filter[flexible_asset_type_id]=%d", flexibleAssetTypeID)

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	return req.Body, nil
}

//GetFlexibleAssetsJSONByID is a special function.  All flexible assets will return different data depending
//on which fields make up the flexible asset.  Because of this, this function simply returns the JSON bytes.
func (itg *ITGAPI) GetFlexibleAssetsJSONByID(flexibleAssetID int) ([]byte, error) {
	req := itg.NewRequest(fmt.Sprintf("/flexible_assets/%d", flexibleAssetID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	return req.Body, nil
}

//GetFlexibleAssetsJSONByName is a special function.  All flexible assets will return different data depending
//on which fields make up the flexible asset.  Because of this, this function simply returns the JSON bytes.
func (itg *ITGAPI) GetFlexibleAssetsJSONByName(flexibleAssetTypeID int, flexibleAssetName string, pageNumber int) ([]byte, error) {
	req := itg.NewRequest("/flexible_assets", "GET", nil)
	req.Page = pageNumber
	req.RawURLValues = fmt.Sprintf("filter[flexible_asset_type_id]=%d&filter[name]=%s", flexibleAssetTypeID, flexibleAssetName)

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	return req.Body, nil
}

//GetFlexibleAssetsJSONByOrganizationID is a special function.  All flexible assets will return different data depending
//on which fields make up the flexible asset.  Because of this, this function simply returns the JSON bytes.
func (itg *ITGAPI) GetFlexibleAssetsJSONByOrganizationID(flexibleAssetTypeID int, OrganizationID string, pageNumber int) ([]byte, error) {
	req := itg.NewRequest("/flexible_assets", "GET", nil)
	req.Page = pageNumber
	req.RawURLValues = fmt.Sprintf("filter[flexible_asset_type_id]=%d&filter[organization_id]=%d", flexibleAssetTypeID, OrganizationID)

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	return req.Body, nil
}

//GetFlexibleAssetsJSONByOrganizationIDAndName is a special function.  All flexible assets will return different data depending
//on which fields make up the flexible asset.  Because of this, this function simply returns the JSON bytes.
func (itg *ITGAPI) GetFlexibleAssetsJSONByOrganizationIDAndName(flexibleAssetTypeID int, OrganizationID string, flexibleAssetName string, pageNumber int) ([]byte, error) {
	req := itg.NewRequest("/flexible_assets", "GET", nil)
	req.Page = pageNumber
	req.RawURLValues = fmt.Sprintf("filter[flexible_asset_type_id]=%d&filter[organization_id]=%d&filter[name]=%s", flexibleAssetTypeID, OrganizationID, flexibleAssetName)

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	return req.Body, nil
}

package itglue

import (
	"fmt"
	"time"
)

type FlexibleAssetData struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		OrganizationID        int    `json:"organization-id"`
		OrganizationName      string `json:"organization-name"`
		ResourceURL           string `json:"resource-url"`
		Restricted            bool   `json:"restricted"`
		MyGlue                bool   `json:"my-glue"`
		FlexibleAssetTypeID   int    `json:"flexible-asset-type-id"`
		FlexibleAssetTypeName string `json:"flexible-asset-type-name"`
		Name                  string `json:"name"`
		Traits                struct {
		} `json:"traits"`
		Archived  bool      `json:"archived"`
		CreatedAt time.Time `json:"created-at"`
		UpdatedAt time.Time `json:"updated-at"`
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
func (itg *ITGAPI) GetFlexibleAssetsJSONByOrganizationID(flexibleAssetTypeID int, OrganizationID int, pageNumber int) ([]byte, error) {
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
func (itg *ITGAPI) GetFlexibleAssetsJSONByOrganizationIDAndName(flexibleAssetTypeID int, OrganizationID int, flexibleAssetName string, pageNumber int) ([]byte, error) {
	req := itg.NewRequest("/flexible_assets", "GET", nil)
	req.Page = pageNumber
	req.RawURLValues = fmt.Sprintf("filter[flexible_asset_type_id]=%d&filter[organization_id]=%d&filter[name]=%s", flexibleAssetTypeID, OrganizationID, flexibleAssetName)

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	return req.Body, nil
}

//PostFlexibleAsset create new asset
func (itg *ITGAPI) PostFlexibleAsset(asset []byte) ([]byte, error) {
	req := itg.NewRequest("/flexible_assets", "POST", asset)

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	return req.Body, nil
}

//PatchFlexibleAsset update asset. Any trait not specified will be deleted
func (itg *ITGAPI) PatchFlexibleAsset(flexibleAssetID int, asset []byte) ([]byte, error) {
	req := itg.NewRequest(fmt.Sprintf("/flexible_assets/%d", flexibleAssetID), "PATCH", asset)

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	return req.Body, nil
}

//DeleteFlexibleAsset delete asset
func (itg *ITGAPI) DeleteFlexibleAsset(flexibleAssetID int) ([]byte, error) {
	req := itg.NewRequest(fmt.Sprintf("/flexible_assets/%d", flexibleAssetID), "DELETE", nil)

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	return req.Body, nil
}

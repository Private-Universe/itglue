package itglue

import (
	"encoding/json"
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
		Traits                map[string]interface {
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

//GetFlexibleAssets returns flexible assets
func (itg *ITGAPI) GetFlexibleAssets(flexibleAssetTypeID int, pageNumber int) (*FlexibleAssetList, error) {
	req := itg.NewRequest("/flexible_assets", "GET", nil)
	req.Page = pageNumber
	req.RawURLValues = fmt.Sprintf("filter[flexible_asset_type_id]=%d", flexibleAssetTypeID)

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	flexibleAsset := &FlexibleAssetList{}
	err = json.Unmarshal(req.Body, flexibleAsset)
	if err != nil {
		return nil, fmt.Errorf("could not get flexibleAsset: %s", err)
	}

	if len((*flexibleAsset).Data) == 0 {
		return nil, fmt.Errorf("ITG returned no results for type %d", flexibleAssetTypeID)
	}

	return flexibleAsset, nil
}

//GetFlexibleAssetsByID returns flexible assets
func (itg *ITGAPI) GetFlexibleAssetsByID(flexibleAssetID int) (*FlexibleAsset, error) {
	req := itg.NewRequest(fmt.Sprintf("/flexible_assets/%d", flexibleAssetID), "GET", nil)
	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	flexibleAsset := &FlexibleAsset{}
	err = json.Unmarshal(req.Body, flexibleAsset)
	if err != nil {
		return nil, fmt.Errorf("could not get flexibleAsset: %s", err)
	}

	return flexibleAsset, nil
}

//GetFlexibleAssetsByName returns flexible assets
func (itg *ITGAPI) GetFlexibleAssetsByName(flexibleAssetTypeID int, flexibleAssetName string, pageNumber int) (*FlexibleAssetList, error) {
	req := itg.NewRequest("/flexible_assets", "GET", nil)
	req.Page = pageNumber
	req.RawURLValues = fmt.Sprintf("filter[flexible_asset_type_id]=%d&filter[name]=%s", flexibleAssetTypeID, flexibleAssetName)

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	flexibleAsset := &FlexibleAssetList{}
	err = json.Unmarshal(req.Body, flexibleAsset)
	if err != nil {
		return nil, fmt.Errorf("could not get flexibleAsset: %s", err)
	}

	if len((*flexibleAsset).Data) == 0 {
		return nil, fmt.Errorf("ITG returned no results for type %d", flexibleAssetTypeID)
	}

	return flexibleAsset, nil
}

//GetFlexibleAssetsByOrganizationID returns flexible assets
func (itg *ITGAPI) GetFlexibleAssetsByOrganizationID(flexibleAssetTypeID int, OrganizationID int, pageNumber int) (*FlexibleAssetList, error) {
	req := itg.NewRequest("/flexible_assets", "GET", nil)
	req.Page = pageNumber
	req.RawURLValues = fmt.Sprintf("filter[flexible_asset_type_id]=%d&filter[organization_id]=%d", flexibleAssetTypeID, OrganizationID)

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	flexibleAsset := &FlexibleAssetList{}
	err = json.Unmarshal(req.Body, flexibleAsset)
	if err != nil {
		return nil, fmt.Errorf("could not get flexibleAsset: %s", err)
	}

	if len((*flexibleAsset).Data) == 0 {
		return nil, fmt.Errorf("ITG returned no results for type %d", flexibleAssetTypeID)
	}

	return flexibleAsset, nil
}

//GetFlexibleAssetsByOrganizationIDAndName returns flexible assets
func (itg *ITGAPI) GetFlexibleAssetsByOrganizationIDAndName(flexibleAssetTypeID int, OrganizationID int, flexibleAssetName string, pageNumber int) (*FlexibleAssetList, error) {
	req := itg.NewRequest("/flexible_assets", "GET", nil)
	req.Page = pageNumber
	req.RawURLValues = fmt.Sprintf("filter[flexible_asset_type_id]=%d&filter[organization_id]=%d&filter[name]=%s", flexibleAssetTypeID, OrganizationID, flexibleAssetName)

	err := req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	flexibleAsset := &FlexibleAssetList{}
	err = json.Unmarshal(req.Body, flexibleAsset)
	if err != nil {
		return nil, fmt.Errorf("could not get flexibleAsset: %s", err)
	}

	if len((*flexibleAsset).Data) == 0 {
		return nil, fmt.Errorf("ITG returned no results for ID %s", flexibleAssetName)
	}

	return flexibleAsset, nil
}

//PostFlexibleAsset create new asset
func (itg *ITGAPI) PostFlexibleAsset(asset *FlexibleAsset) (*FlexibleAsset, error) {
	a, err := json.Marshal(asset)
	if err != nil {
		return nil, fmt.Errorf("could not decode asset: %s", err)
	}

	req := itg.NewRequest("/flexible_assets", "POST", a)

	err = req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	flexibleAsset := &FlexibleAsset{}
	err = json.Unmarshal(req.Body, flexibleAsset)
	if err != nil {
		return nil, fmt.Errorf("could not get flexibleAsset: %s", err)
	}

	return flexibleAsset, nil
}

//PatchFlexibleAsset update asset. Any trait not specified will be deleted
func (itg *ITGAPI) PatchFlexibleAsset(flexibleAssetID int, asset *FlexibleAsset) (*FlexibleAsset, error) {
	a, err := json.Marshal(asset)
	if err != nil {
		return nil, fmt.Errorf("could not decode asset: %s", err)
	}

	req := itg.NewRequest(fmt.Sprintf("/flexible_assets/%d", flexibleAssetID), "PATCH", a)

	err = req.Do()
	if err != nil {
		return nil, fmt.Errorf("request failed for %s: %s", req.RestAction, err)
	}

	flexibleAsset := &FlexibleAsset{}
	err = json.Unmarshal(req.Body, flexibleAsset)
	if err != nil {
		return nil, fmt.Errorf("could not get flexibleAsset: %s", err)
	}

	return flexibleAsset, nil
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

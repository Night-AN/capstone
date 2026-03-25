package usecase

import "github.com/google/uuid"

type IntangibleAssetsUpdateRequest struct {
	IntangibleAssetID                uuid.UUID      `json:"intangible_asset_id"`
	IntangibleAssetName              string         `json:"intangible_asset_name"`
	IntangibleAssetCode              string         `json:"intangible_asset_code"`
	IntangibleAssetDescription       string         `json:"intangible_asset_description"`
	IntangibleAssetFlag              string         `json:"intangible_asset_flag"`
	IntangibleAssetPurchasePrice     float64        `json:"intangible_asset_purchase_price"`
	IntangibleAssetDepreciationPrice float64        `json:"intangible_asset_depreciation_price"`
	IntangibleAssetManufacturer      string         `json:"intangible_asset_manufacturer"`
	IntangibleAssetModel             string         `json:"intangible_asset_model"`
	OtherMetadata                    map[string]any `json:"other_metadata"`
	IntangibleAssetTypeID            *uuid.UUID     `json:"intangible_asset_type_id"`
	OrganizationID                   *uuid.UUID     `json:"organization_id"`
}
type IntangibleAssetsUpdateResponse struct {
	Success bool `json:"success"`
}

package usecase

import "github.com/google/uuid"

// FixedAssetsUpdateRequest represents the request for updating a fixed asset
type FixedAssetUpdateRequest struct {
	FixedAssetID                uuid.UUID              `json:"fixed_asset_id"`
	FixedAssetName              string                 `json:"fixed_asset_name"`
	FixedAssetCode              string                 `json:"fixed_asset_code"`
	FixedAssetDescription       string                 `json:"fixed_asset_description"`
	FixedAssetFlag              string                 `json:"fixed_asset_flag"`
	FixedAssetQuantity          int                    `json:"fixed_asset_quantity"`
	FixedAssetLocation          string                 `json:"fixed_asset_location"`
	FixedAssetPurchasePrice     float64                `json:"fixed_asset_purchase_price"`
	FixedAssetDepreciationPrice float64                `json:"fixed_asset_depreciation_price"`
	FixedAssetManufacturer      string                 `json:"fixed_asset_manufacturer"`
	FixedAssetModel             string                 `json:"fixed_asset_model"`
	OtherMetadata               map[string]interface{} `json:"other_metadata"`
	FixedAssetTypeID            *uuid.UUID             `json:"fixed_asset_type_id"`
	OrganizationID              *uuid.UUID             `json:"organization_id"`
}

// FixedAssetsUpdateResponse represents the response for updating a fixed asset
type FixedAssetUpdateResponse struct {
	Success bool `json:"success"`
}

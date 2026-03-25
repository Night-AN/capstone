package usecase

import (
	"time"

	"github.com/google/uuid"
)

type FixedAssetItem struct {
	FixedAssetID                uuid.UUID              `json:"fixed_asset_id"`
	FixedAssetName              string                 `json:"fixed_asset_name"`
	FixedAssetCode              string                 `json:"fixed_asset_code"`
	FixedAssetDescription       string                 `json:"fixed_asset_description"`
	FixedAssetFlag              string                 `json:"fixed_asset_flag"`
	FixedAssetQuantity          int                    `json:"fixed_asset_quantity"`
	FixedAssetLocation          string                 `json:"fixed_asset_location"`
	FixedAssetPurchasePrice     float64                `json:"fixed_asset_purchase_price"`
	FixedAssetDepreciationPrice float64                `json:"fixed_asset_depreciation_price"`
	FixedAssetPurchaseDate      *time.Time             `json:"fixed_asset_purchase_date"`
	FixedAssetManufacturer      string                 `json:"fixed_asset_manufacturer"`
	FixedAssetModel             string                 `json:"fixed_asset_model"`
	OtherMetadata               map[string]interface{} `json:"other_metadata"`
	CreatedAt                   time.Time              `json:"created_at"`
	UpdatedAt                   *time.Time             `json:"updated_at"`
	FixedAssetTypeID            *uuid.UUID             `json:"fixed_asset_type_id"`
	OrganizationID              *uuid.UUID             `json:"organization_id"`
}

type FixedAssetGetRequest struct {
	FixedAssetID uuid.UUID `json:"fixed_asset_id"`
}

type FixedAssetGetResponse = FixedAssetItem

package usecase

import (
	"time"

	"github.com/google/uuid"
)

// IntangibleAssetsCreateRequest represents the request for creating an intangible asset
type IntangibleAssetsCreateRequest struct {
	IntangibleAssetName              string         `json:"intangible_asset_name"`
	IntangibleAssetCode              string         `json:"intangible_asset_code"`
	IntangibleAssetDescription       string         `json:"intangible_asset_description"`
	IntangibleAssetFlag              string         `json:"intangible_asset_flag"`
	IntangibleAssetPurchasePrice     float64        `json:"intangible_asset_purchase_price"`
	IntangibleAssetDepreciationPrice float64        `json:"intangible_asset_depreciation_price"`
	IntangibleAssetPurchaseDate      *time.Time     `json:"intangible_asset_purchase_date"`
	OtherMetadata                    map[string]any `json:"other_metadata"`
	IntangibleAssetTypeID            *uuid.UUID     `json:"intangible_asset_type_id"`
	OrganizationID                   *uuid.UUID     `json:"organization_id"`
}

// IntangibleAssetsCreateResponse represents the response for creating an intangible asset
type IntangibleAssetsCreateResponse struct {
	Success bool `json:"success"`
}

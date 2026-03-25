package usecase

import (
	"time"

	"github.com/google/uuid"
)

type IntangibleAssetsItem struct {
	IntangibleAssetID                uuid.UUID      `json:"intangible_asset_id"`
	IntangibleAssetName              string         `json:"intangible_asset_name"`
	IntangibleAssetCode              string         `json:"intangible_asset_code"`
	IntangibleAssetDescription       string         `json:"intangible_asset_description"`
	IntangibleAssetFlag              string         `json:"intangible_asset_flag"`
	IntangibleAssetPurchasePrice     float64        `json:"intangible_asset_purchase_price"`
	IntangibleAssetDepreciationPrice float64        `json:"intangible_asset_depreciation_price"`
	IntangibleAssetPurchaseDate      *time.Time     `json:"intangible_asset_purchase_date"`
	OtherMetadata                    map[string]any `json:"other_metadata"`
	CreatedAt                        time.Time      `json:"created_at"`
	UpdatedAt                        *time.Time     `json:"updated_at"`
	IntangibleAssetTypeID            *uuid.UUID     `json:"intangible_asset_type_id"`
	OrganizationID                   *uuid.UUID     `json:"organization_id"`
}

type IntangibleAssetsGetRequest struct {
	IntangibleAssetID uuid.UUID `json:"intangible_asset_id"`
}

type IntangibleAssetsGetResponse = IntangibleAssetsItem

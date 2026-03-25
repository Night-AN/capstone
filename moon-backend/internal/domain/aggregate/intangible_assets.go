package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type IntangibleAssets struct {
	IntangibleAssetID                uuid.UUID      `gorm:"column:intangible_asset_id"`
	IntangibleAssetName              string         `gorm:"column:intangible_asset_name"`
	IntangibleAssetCode              string         `gorm:"column:intangible_asset_code"`
	IntangibleAssetDescription       string         `gorm:"column:intangible_asset_description"`
	IntangibleAssetFlag              string         `gorm:"column:intangible_asset_flag"`
	IntangibleAssetPurchasePrice     float64        `gorm:"column:intangible_asset_purchase_price"`
	IntangibleAssetDepreciationPrice float64        `gorm:"column:intangible_asset_depreciation_price"`
	IntangibleAssetPurchaseDate      *time.Time     `gorm:"column:intangible_asset_purchase_date"`
	OtherMetadata                    map[string]any `gorm:"column:other_metadata"`
	CreatedAt                        time.Time      `gorm:"column:created_at"`
	UpdatedAt                        *time.Time     `gorm:"column:updated_at"`

	IntangibleAssetTypeID *uuid.UUID           `gorm:"column:int_asset_type_id"`
	IntangibleAssetType   *IntangibleAssetType `gorm:"foreignKey:IntangibleAssetTypeID;references:IntangibleAssetTypeID"`

	OrganizationID *uuid.UUID    `gorm:"column:organization_id"`
	Organization   *Organization `gorm:"foreignKey:OrganizationID;references:OrganizationID"`
}

func (IntangibleAssets) TableName() string {
	return "biz.intangible_assets"
}

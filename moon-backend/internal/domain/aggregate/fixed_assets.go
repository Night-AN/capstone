package aggregate

import (
	"time"

	"github.com/google/uuid"
)

// FixedAssets represents the fixed assets aggregate
// This struct maps to the biz.fixed_assets table in the database
type FixedAssets struct {
	// FixedAssetID is the primary key for the fixed asset
	FixedAssetID uuid.UUID `gorm:"column:fixed_asset_id"`

	// FixedAssetName is the name of the fixed asset
	FixedAssetName string `gorm:"column:fixed_asset_name"`

	// FixedAssetCode is the unique code for the fixed asset
	FixedAssetCode string `gorm:"column:fixed_asset_code"`

	// FixedAssetDescription is the description of the fixed asset
	FixedAssetDescription string `gorm:"column:fixed_asset_description"`

	// FixedAssetFlag indicates the status of the fixed asset (e.g., ACTIVE, INACTIVE)
	FixedAssetFlag string `gorm:"column:fixed_asset_flag"`

	// FixedAssetQuantity is the quantity of the fixed asset
	FixedAssetQuantity int `gorm:"column:fixed_asset_quantity"`

	// FixedAssetLocation is the physical location of the fixed asset
	FixedAssetLocation string `gorm:"column:fixed_asset_location"`

	// FixedAssetPurchasePrice is the purchase price of the fixed asset
	FixedAssetPurchasePrice float64 `gorm:"column:fixed_asset_purchase_price"`

	// FixedAssetDepreciationPrice is the depreciation price of the fixed asset
	FixedAssetDepreciationPrice float64 `gorm:"column:fixed_asset_depreciation_price"`

	// FixedAssetPurchaseDate is the date when the fixed asset was purchased
	FixedAssetPurchaseDate *time.Time `gorm:"column:fixed_asset_purchase_date"`

	// FixedAssetManufacturer is the manufacturer of the fixed asset
	FixedAssetManufacturer string `gorm:"column:fixed_asset_manufacturer"`

	// FixedAssetModel is the model number of the fixed asset
	FixedAssetModel string `gorm:"column:fixed_asset_model"`

	// OtherMetadata stores additional metadata as a JSON object
	// Using map[string]interface{} to handle the jsonb type from PostgreSQL
	OtherMetadata map[string]interface{} `gorm:"column:other_metadata;type:jsonb"`

	// CreatedAt is the timestamp when the fixed asset was created
	CreatedAt time.Time `gorm:"column:created_at"`

	// UpdatedAt is the timestamp when the fixed asset was last updated
	UpdatedAt *time.Time `gorm:"column:updated_at"`

	// FixedAssetTypeID is the foreign key referencing biz.fixed_assets_type
	FixedAssetTypeID *uuid.UUID      `gorm:"column:fixed_asset_type_id"`
	FixedAssetType   *FixedAssetType `gorm:"foreignKey:FixedAssetTypeID;references:FixedAssetTypeID"`

	// OrganizationID is the foreign key referencing systems.organization
	OrganizationID *uuid.UUID    `gorm:"column:organization_id"`
	Organization   *Organization `gorm:"foreignKey:OrganizationID;references:OrganizationID"`
}

// TableName specifies the table name for the FixedAssets struct
func (FixedAssets) TableName() string {
	return "biz.fixed_assets"
}

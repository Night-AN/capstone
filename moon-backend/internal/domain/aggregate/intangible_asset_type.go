package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type IntangibleAssetType struct {
	IntangibleAssetTypeID   uuid.UUID  `gorm:"column:intangible_asset_type_id"`
	IntangibleAssetTypeName string     `gorm:"column:intangible_asset_type_name"`
	IntangibleAssetTypeCode string     `gorm:"column:intangible_asset_type_code"`
	IntangibleAssetTypeFlag string     `gorm:"column:intangible_asset_type_flag"`
	CreatedAt               time.Time  `gorm:"column:created_at"`
	UpdatedAt               *time.Time `gorm:"column:updated_at"`
}

func (IntangibleAssetType) TableName() string {
	return "biz.intangible_asset_types"
}

package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type FixedAssetType struct {
	FixedAssetTypeID   uuid.UUID `gorm:"column:fixed_asset_type_id"`
	FixedAssetTypeName string    `gorm:"column:fixed_asset_type_name"`
	FixedAssetTypeCode string    `gorm:"column:fixed_asset_type_code"`
	FixedAssetTypeFlag string    `gorm:"column:fixed_asset_type_flag"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
}

func (FixedAssetType) TableName() string {
	return "biz.fixed_asset_type"
}

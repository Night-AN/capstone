package usecase

import "github.com/google/uuid"

type FixedAssetTypeUpdateRequest struct {
	FixedAssetTypeID   uuid.UUID `json:"fixed_asset_type_id"`
	FixedAssetTypeName string    `json:"fixed_asset_type_name"`
	FixedAssetTypeCode string    `json:"fixed_asset_type_code"`
	FixedAssetTypeFlag string    `json:"fixed_asset_type_flag"`
}

type FixedAssetTypeUpdateResponse struct {
	Success bool `json:"success"`
}

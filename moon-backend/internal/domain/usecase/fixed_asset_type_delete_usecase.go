package usecase

import "github.com/google/uuid"

type FixedAssetTypeDeleteRequest struct {
	FixedAssetsTypeID uuid.UUID `json:"fixed_assets_type_id"`
}

type FixedAssetTypeDeleteResponse struct {
	Success bool `json:"success"`
}

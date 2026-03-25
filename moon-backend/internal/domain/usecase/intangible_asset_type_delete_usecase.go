package usecase

import "github.com/google/uuid"

type IntangibleAssetTypeDeleteRequest struct {
	IntangibleAssetTypeID uuid.UUID `json:"intangible_asset_type_id"`
}

type IntangibleAssetTypeDeleteResponse struct {
	Success bool `json:"success"`
}

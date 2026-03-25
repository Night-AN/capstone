package usecase

import "github.com/google/uuid"

type IntangibleAssetsDeleteRequest struct {
	IntangibleAssetID uuid.UUID `json:"intangible_asset_id"`
}

type IntangibleAssetsDeleteResponse struct {
	Success bool `json:"success"`
}

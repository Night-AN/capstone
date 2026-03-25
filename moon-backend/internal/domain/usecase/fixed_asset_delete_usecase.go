package usecase

import "github.com/google/uuid"

// FixedAssetsDeleteRequest represents the request for deleting a fixed asset
type FixedAssetDeleteRequest struct {
	FixedAssetID uuid.UUID `json:"fixed_asset_id"`
}

// FixedAssetsDeleteResponse represents the response for deleting a fixed asset
type FixedAssetDeleteResponse struct {
	Success bool `json:"success"`
}

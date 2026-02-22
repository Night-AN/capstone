package usecase

import (
	"github.com/google/uuid"
)

// AssetDeleteRequest defines the request structure for deleting an asset
type AssetDeleteRequest struct {
	AssetID uuid.UUID `json:"asset_id"`
}

// AssetDeleteResponse defines the response structure for deleting an asset
type AssetDeleteResponse struct {
	Success bool `json:"success"`
}

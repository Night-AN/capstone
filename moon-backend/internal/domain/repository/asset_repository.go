package repository

import (
	"context"
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

// AssetRepository defines the interface for asset repository operations
type AssetRepository interface {
	// SaveAsset saves an asset to the repository
	SaveAsset(ctx *context.Context, asset aggregate.Asset) error

	// UpdateAsset updates an existing asset in the repository
	UpdateAsset(ctx *context.Context, asset aggregate.Asset) error

	// DeleteAsset deletes an asset from the repository by ID
	DeleteAsset(ctx *context.Context, assetID uuid.UUID) error

	// FindAssetByID finds an asset by ID
	FindAssetByID(ctx *context.Context, assetID uuid.UUID) (aggregate.Asset, error)

	// ListAllAssets returns all assets from the repository
	ListAllAssets(ctx *context.Context) ([]aggregate.Asset, error)

	// FindAssetsByType finds assets by type
	FindAssetsByType(ctx *context.Context, assetType string) ([]aggregate.Asset, error)

	// FindAssetsByStatus finds assets by status
	FindAssetsByStatus(ctx *context.Context, status string) ([]aggregate.Asset, error)

	// FindAssetsByIPAddress finds assets by IP address
	FindAssetsByIPAddress(ctx *context.Context, ipAddress string) ([]aggregate.Asset, error)

	// FindAssetsByMACAddress finds assets by MAC address
	FindAssetsByMACAddress(ctx *context.Context, macAddress string) ([]aggregate.Asset, error)

	// FindAssetsByDepartment finds assets by department
	FindAssetsByDepartment(ctx *context.Context, department string) ([]aggregate.Asset, error)

	// FindAssetsByOwner finds assets by owner
	FindAssetsByOwner(ctx *context.Context, owner string) ([]aggregate.Asset, error)

	// FindAssetsByOrganizationID finds assets by organization ID
	FindAssetsByOrganizationID(ctx *context.Context, organizationID uuid.UUID) ([]aggregate.Asset, error)

	// CountAssets returns the total number of assets
	CountAssets(ctx *context.Context) (int, error)

	// CountAssetsByType returns the number of assets by type
	CountAssetsByType(ctx *context.Context, assetType string) (int, error)

	// CountAssetsByStatus returns the number of assets by status
	CountAssetsByStatus(ctx *context.Context, status string) (int, error)

	// CountAssetsByOrganizationID returns the number of assets by organization ID
	CountAssetsByOrganizationID(ctx *context.Context, organizationID uuid.UUID) (int, error)
}

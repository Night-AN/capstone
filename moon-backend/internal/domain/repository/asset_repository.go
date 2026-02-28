package repository

import (
	"context"
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

type AssetRepository interface {
	SaveAsset(ctx context.Context, asset aggregate.Asset) error
	UpdateAsset(ctx context.Context, asset aggregate.Asset) error
	DeleteAsset(ctx context.Context, assetID uuid.UUID) error
	FindAssetByID(ctx context.Context, assetID uuid.UUID) (aggregate.Asset, error)
	ListAllAssets(ctx context.Context) ([]aggregate.Asset, error)
	FindAssetsByType(ctx context.Context, assetType string) ([]aggregate.Asset, error)
	FindAssetsByStatus(ctx context.Context, status string) ([]aggregate.Asset, error)
	FindAssetsByIPAddress(ctx context.Context, ipAddress string) ([]aggregate.Asset, error)
	FindAssetsByMACAddress(ctx context.Context, macAddress string) ([]aggregate.Asset, error)
	FindAssetsByDepartment(ctx context.Context, department string) ([]aggregate.Asset, error)
	FindAssetsByOwner(ctx context.Context, owner string) ([]aggregate.Asset, error)
	FindAssetsByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]aggregate.Asset, error)
	CountAssets(ctx context.Context) (int, error)
	CountAssetsByType(ctx context.Context, assetType string) (int, error)
	CountAssetsByStatus(ctx context.Context, status string) (int, error)
	CountAssetsByOrganizationID(ctx context.Context, organizationID uuid.UUID) (int, error)
}

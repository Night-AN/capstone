package postgres

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type assetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) repository.AssetRepository {
	return &assetRepository{db: db}
}

func (ar *assetRepository) SaveAsset(ctx *context.Context, asset aggregate.Asset) error {
	// 检查资产是否已存在
	var existingAsset aggregate.Asset
	err := ar.db.WithContext(*ctx).Where("asset_id = ?", asset.AssetID).First(&existingAsset).Error

	if err == gorm.ErrRecordNotFound {
		// 资产不存在，创建新记录
		return ar.db.WithContext(*ctx).Create(&asset).Error
	} else if err != nil {
		// 其他错误
		return err
	} else {
		// 资产存在，更新记录
		return ar.db.WithContext(*ctx).Model(&aggregate.Asset{}).Where("asset_id = ?", asset.AssetID).Updates(&asset).Error
	}
}

func (ar *assetRepository) UpdateAsset(ctx *context.Context, asset aggregate.Asset) error {
	return ar.db.WithContext(*ctx).Save(&asset).Error
}

func (ar *assetRepository) DeleteAsset(ctx *context.Context, assetID uuid.UUID) error {
	return ar.db.WithContext(*ctx).Delete(&aggregate.Asset{}, "asset_id = ?", assetID).Error
}

func (ar *assetRepository) FindAssetByID(ctx *context.Context, assetID uuid.UUID) (aggregate.Asset, error) {
	var asset = aggregate.Asset{}
	err := ar.db.WithContext(*ctx).Model(&asset).Where("asset_id =?", assetID).First(&asset).Error
	return asset, err
}

func (ar *assetRepository) ListAllAssets(ctx *context.Context) ([]aggregate.Asset, error) {
	var assets []aggregate.Asset
	err := ar.db.WithContext(*ctx).Find(&assets).Error
	return assets, err
}

func (ar *assetRepository) FindAssetsByType(ctx *context.Context, assetType string) ([]aggregate.Asset, error) {
	var assets []aggregate.Asset
	err := ar.db.WithContext(*ctx).Where("asset_type = ?", assetType).Find(&assets).Error
	return assets, err
}

func (ar *assetRepository) FindAssetsByStatus(ctx *context.Context, status string) ([]aggregate.Asset, error) {
	var assets []aggregate.Asset
	err := ar.db.WithContext(*ctx).Where("status = ?", status).Find(&assets).Error
	return assets, err
}

func (ar *assetRepository) FindAssetsByIPAddress(ctx *context.Context, ipAddress string) ([]aggregate.Asset, error) {
	var assets []aggregate.Asset
	err := ar.db.WithContext(*ctx).Where("ip_address = ?", ipAddress).Find(&assets).Error
	return assets, err
}

func (ar *assetRepository) FindAssetsByMACAddress(ctx *context.Context, macAddress string) ([]aggregate.Asset, error) {
	var assets []aggregate.Asset
	err := ar.db.WithContext(*ctx).Where("mac_address = ?", macAddress).Find(&assets).Error
	return assets, err
}

func (ar *assetRepository) FindAssetsByDepartment(ctx *context.Context, department string) ([]aggregate.Asset, error) {
	var assets []aggregate.Asset
	err := ar.db.WithContext(*ctx).Where("department = ?", department).Find(&assets).Error
	return assets, err
}

func (ar *assetRepository) FindAssetsByOwner(ctx *context.Context, owner string) ([]aggregate.Asset, error) {
	var assets []aggregate.Asset
	err := ar.db.WithContext(*ctx).Where("owner = ?", owner).Find(&assets).Error
	return assets, err
}

func (ar *assetRepository) CountAssets(ctx *context.Context) (int, error) {
	var count int64
	err := ar.db.WithContext(*ctx).Model(&aggregate.Asset{}).Count(&count).Error
	return int(count), err
}

func (ar *assetRepository) CountAssetsByType(ctx *context.Context, assetType string) (int, error) {
	var count int64
	err := ar.db.WithContext(*ctx).Model(&aggregate.Asset{}).Where("asset_type = ?", assetType).Count(&count).Error
	return int(count), err
}

func (ar *assetRepository) CountAssetsByStatus(ctx *context.Context, status string) (int, error) {
	var count int64
	err := ar.db.WithContext(*ctx).Model(&aggregate.Asset{}).Where("status = ?", status).Count(&count).Error
	return int(count), err
}

func (ar *assetRepository) FindAssetsByOrganizationID(ctx *context.Context, organizationID uuid.UUID) ([]aggregate.Asset, error) {
	var assets []aggregate.Asset
	err := ar.db.WithContext(*ctx).Where("organization_id = ?", organizationID).Find(&assets).Error
	return assets, err
}

func (ar *assetRepository) CountAssetsByOrganizationID(ctx *context.Context, organizationID uuid.UUID) (int, error) {
	var count int64
	err := ar.db.WithContext(*ctx).Model(&aggregate.Asset{}).Where("organization_id = ?", organizationID).Count(&count).Error
	return int(count), err
}

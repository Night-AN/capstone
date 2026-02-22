package service

import (
	"context"
	"moon/internal/domain/errors"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"
)

// AssetService defines the interface for asset service operations
type AssetService interface {
	CreateAsset(ctx *context.Context, req usecase.AssetCreateRequest) (usecase.AssetCreateResponse, errors.DomainError)
	UpdateAsset(ctx *context.Context, req usecase.AssetUpdateRequest) (usecase.AssetUpdateResponse, errors.DomainError)
	DeleteAsset(ctx *context.Context, req usecase.AssetDeleteRequest) (usecase.AssetDeleteResponse, errors.DomainError)
	GetAssetByID(ctx *context.Context, req usecase.AssetGetRequest) (usecase.AssetGetResponse, errors.DomainError)
	ListAllAssets(ctx *context.Context, req usecase.AssetListRequest) (usecase.AssetListResponse, errors.DomainError)
}

// NewAssetService creates a new asset service instance
func NewAssetService(assetRepo repository.AssetRepository) AssetService {
	return &assetService{assetRepo}
}

// assetService implements the AssetService interface
type assetService struct {
	AssetRepository repository.AssetRepository
}

// CreateAsset creates a new asset
func (as *assetService) CreateAsset(ctx *context.Context, req usecase.AssetCreateRequest) (usecase.AssetCreateResponse, errors.DomainError) {
	asset := usecase.ConvertAssetCreateRequestToAssetAggregate(req)
	err := as.AssetRepository.SaveAsset(ctx, asset)
	if err != nil {
		return usecase.AssetCreateResponse{}, errors.NewDomainWithError("401", "Create Asset Err", err)
	}
	return usecase.AssetCreateResponse{
		AssetID:   asset.AssetID,
		AssetName: asset.AssetName,
		AssetCode: asset.AssetCode,
	}, errors.DomainError{}
}

// UpdateAsset updates an existing asset
func (as *assetService) UpdateAsset(ctx *context.Context, req usecase.AssetUpdateRequest) (usecase.AssetUpdateResponse, errors.DomainError) {
	// Check if the asset exists
	_, err := as.AssetRepository.FindAssetByID(ctx, req.AssetID)
	if err != nil {
		return usecase.AssetUpdateResponse{}, errors.NewDomainWithError("401", "Get Asset Err", err)
	}

	asset := usecase.ConvertAssetUpdateRequestToAssetAggregate(req)
	err = as.AssetRepository.UpdateAsset(ctx, asset)
	if err != nil {
		return usecase.AssetUpdateResponse{}, errors.NewDomainWithError("401", "Update Asset Err", err)
	}
	return usecase.AssetUpdateResponse{
		AssetID:   asset.AssetID,
		AssetName: asset.AssetName,
		AssetCode: asset.AssetCode,
	}, errors.DomainError{}
}

// DeleteAsset deletes an asset
func (as *assetService) DeleteAsset(ctx *context.Context, req usecase.AssetDeleteRequest) (usecase.AssetDeleteResponse, errors.DomainError) {
	// Check if the asset exists
	_, err := as.AssetRepository.FindAssetByID(ctx, req.AssetID)
	if err != nil {
		return usecase.AssetDeleteResponse{Success: false}, errors.NewDomainWithError("401", "Get Asset Err", err)
	}

	err = as.AssetRepository.DeleteAsset(ctx, req.AssetID)
	if err != nil {
		return usecase.AssetDeleteResponse{Success: false}, errors.NewDomainWithError("401", "Delete Asset Err", err)
	}
	return usecase.AssetDeleteResponse{Success: true}, errors.DomainError{}
}

// GetAssetByID gets an asset by ID
func (as *assetService) GetAssetByID(ctx *context.Context, req usecase.AssetGetRequest) (usecase.AssetGetResponse, errors.DomainError) {
	asset, err := as.AssetRepository.FindAssetByID(ctx, req.AssetID)
	if err != nil {
		return usecase.AssetGetResponse{}, errors.NewDomainWithError("401", "Get Asset Err", err)
	}
	return usecase.ConvertAssetAggregateToAssetGetResponse(asset), errors.DomainError{}
}

// ListAllAssets gets all assets
func (as *assetService) ListAllAssets(ctx *context.Context, req usecase.AssetListRequest) (usecase.AssetListResponse, errors.DomainError) {
	assets, err := as.AssetRepository.ListAllAssets(ctx)
	if err != nil {
		return usecase.AssetListResponse{}, errors.NewDomainWithError("401", "List Assets Err", err)
	}
	return usecase.ConvertAssetAggregatesToAssetListResponse(assets), errors.DomainError{}
}

package service

import (
	"context"
	"moon/internal/domain/errors"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"
)

type AssetService interface {
	CreateAsset(ctx context.Context, req usecase.AssetCreateRequest) (usecase.AssetCreateResponse, errors.DomainError)
	UpdateAsset(ctx context.Context, req usecase.AssetUpdateRequest) (usecase.AssetUpdateResponse, errors.DomainError)
	DeleteAsset(ctx context.Context, req usecase.AssetDeleteRequest) (usecase.AssetDeleteResponse, errors.DomainError)
	GetAssetByID(ctx context.Context, req usecase.AssetGetRequest) (usecase.AssetGetResponse, errors.DomainError)
	ListAllAssets(ctx context.Context, req usecase.AssetListRequest) (usecase.AssetListResponse, errors.DomainError)
	BatchCreateAsset(ctx context.Context, req usecase.BatchAssetCreateRequest, enableAIClassification bool) (usecase.BatchAssetCreateResponse, errors.DomainError)
}

func NewAssetService(assetRepo repository.AssetRepository, classificationService AssetClassificationService) AssetService {
	return &assetService{
		AssetRepository:       assetRepo,
		ClassificationService: classificationService,
	}
}

type assetService struct {
	AssetRepository       repository.AssetRepository
	ClassificationService AssetClassificationService
}

func (as *assetService) CreateAsset(ctx context.Context, req usecase.AssetCreateRequest) (usecase.AssetCreateResponse, errors.DomainError) {
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

func (as *assetService) UpdateAsset(ctx context.Context, req usecase.AssetUpdateRequest) (usecase.AssetUpdateResponse, errors.DomainError) {
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

func (as *assetService) DeleteAsset(ctx context.Context, req usecase.AssetDeleteRequest) (usecase.AssetDeleteResponse, errors.DomainError) {
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

func (as *assetService) GetAssetByID(ctx context.Context, req usecase.AssetGetRequest) (usecase.AssetGetResponse, errors.DomainError) {
	asset, err := as.AssetRepository.FindAssetByID(ctx, req.AssetID)
	if err != nil {
		return usecase.AssetGetResponse{}, errors.NewDomainWithError("401", "Get Asset Err", err)
	}
	return usecase.ConvertAssetAggregateToAssetGetResponse(asset), errors.DomainError{}
}

func (as *assetService) ListAllAssets(ctx context.Context, req usecase.AssetListRequest) (usecase.AssetListResponse, errors.DomainError) {
	assets, err := as.AssetRepository.ListAllAssets(ctx)
	if err != nil {
		return usecase.AssetListResponse{}, errors.NewDomainWithError("401", "List Assets Err", err)
	}
	return usecase.ConvertAssetAggregatesToAssetListResponse(assets), errors.DomainError{}
}

func (as *assetService) BatchCreateAsset(ctx context.Context, req usecase.BatchAssetCreateRequest, enableAIClassification bool) (usecase.BatchAssetCreateResponse, errors.DomainError) {
	results := make([]usecase.AssetCreateResult, 0, len(req.Assets))
	classifications := make([]usecase.AssetClassificationResult, 0)
	successCount := 0
	failedCount := 0

	for _, assetReq := range req.Assets {
		asset := usecase.ConvertAssetCreateRequestToAssetAggregate(assetReq)
		err := as.AssetRepository.SaveAsset(ctx, asset)

		if err != nil {
			failedCount++
			results = append(results, usecase.AssetCreateResult{
				AssetID:   asset.AssetID,
				AssetName: asset.AssetName,
				AssetCode: asset.AssetCode,
				Success:   false,
				Error:     err.Error(),
			})
			continue
		}

		successCount++
		results = append(results, usecase.AssetCreateResult{
			AssetID:   asset.AssetID,
			AssetName: asset.AssetName,
			AssetCode: asset.AssetCode,
			Success:   true,
		})

		if enableAIClassification {
			classifyReq := usecase.AssetClassifyRequest{
				AssetID:          asset.AssetID.String(),
				PromptTemplateID: req.PromptTemplateID,
			}
			classifyResp, classifyErr := as.ClassificationService.ClassifyAsset(ctx, classifyReq)
			if classifyErr.Code == "" {
				classifications = append(classifications, usecase.AssetClassificationResult{
					AssetID:           asset.AssetID,
					ClassificationID:  classifyResp.ClassificationID,
					PredictedCategory: classifyResp.PredictedCategory,
					Confidence:        classifyResp.Confidence,
					Reasoning:         classifyResp.Reasoning,
					Success:           true,
				})
			} else {
				classifications = append(classifications, usecase.AssetClassificationResult{
					AssetID: asset.AssetID,
					Success: false,
					Error:   classifyErr.Message,
				})
			}
		}
	}

	return usecase.BatchAssetCreateResponse{
		TotalCount:      len(req.Assets),
		SuccessCount:    successCount,
		FailedCount:     failedCount,
		Results:         results,
		Classifications: classifications,
	}, errors.DomainError{}
}

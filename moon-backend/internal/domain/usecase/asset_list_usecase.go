package usecase

import (
	"moon/internal/domain/aggregate"
)

// AssetListRequest defines the request structure for listing assets
type AssetListRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// AssetListResponse defines the response structure for listing assets
type AssetListResponse struct {
	Total  int              `json:"total"`
	Assets []AssetGetResponse `json:"assets"`
}

// ConvertAssetAggregatesToAssetListResponse converts asset aggregates to an asset list response
func ConvertAssetAggregatesToAssetListResponse(assets []aggregate.Asset) AssetListResponse {
	assetResponses := make([]AssetGetResponse, len(assets))
	for i, asset := range assets {
		assetResponses[i] = ConvertAssetAggregateToAssetGetResponse(asset)
	}

	return AssetListResponse{
		Total:  len(assets),
		Assets: assetResponses,
	}
}

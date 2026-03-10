package usecase

import (
	"moon/internal/domain/aggregate"
	"time"

	"github.com/google/uuid"
)

type BatchAssetCreateRequest struct {
	Assets                 []AssetCreateRequest `json:"assets"`
	EnableAIClassification bool                 `json:"enable_ai_classification"`
	PromptTemplateID       *string              `json:"prompt_template_id,omitempty"`
}

type BatchAssetCreateResponse struct {
	TotalCount      int                         `json:"total_count"`
	SuccessCount    int                         `json:"success_count"`
	FailedCount     int                         `json:"failed_count"`
	Results         []AssetCreateResult         `json:"results"`
	Classifications []AssetClassificationResult `json:"classifications"`
}

type AssetCreateResult struct {
	AssetID   uuid.UUID `json:"asset_id"`
	AssetName string    `json:"asset_name"`
	AssetCode string    `json:"asset_code"`
	Success   bool      `json:"success"`
	Error     string    `json:"error,omitempty"`
}

type AssetClassificationResult struct {
	AssetID           uuid.UUID `json:"asset_id"`
	ClassificationID  uuid.UUID `json:"classification_id"`
	PredictedCategory string    `json:"predicted_category"`
	Confidence        float64   `json:"confidence"`
	Reasoning         string    `json:"reasoning"`
	Success           bool      `json:"success"`
	Error             string    `json:"error,omitempty"`
}

func ConvertBatchAssetCreateRequestToAssetAggregates(req BatchAssetCreateRequest) []aggregate.Asset {
	assets := make([]aggregate.Asset, len(req.Assets))
	for i, assetReq := range req.Assets {
		assets[i] = aggregate.Asset{
			AssetID:         uuid.New(),
			AssetName:       assetReq.AssetName,
			AssetCode:       assetReq.AssetCode,
			Description:     assetReq.AssetDescription,
			AssetType:       assetReq.AssetType,
			Manufacturer:    assetReq.Manufacturer,
			Model:           assetReq.Model,
			SerialNumber:    assetReq.SerialNumber,
			IPAddress:       assetReq.IPAddress,
			MACAddress:      assetReq.MACAddress,
			Location:        assetReq.Location,
			Department:      assetReq.Department,
			Owner:           assetReq.Owner,
			ContactInfo:     assetReq.ContactInfo,
			Status:          assetReq.Status,
			PurchaseDate:    assetReq.PurchaseDate,
			WarrantyEndDate: assetReq.WarrantyEndDate,
			Value:           assetReq.Value,
			Notes:           assetReq.Notes,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		}
	}
	return assets
}

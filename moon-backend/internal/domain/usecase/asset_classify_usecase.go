package usecase

import (
	"github.com/google/uuid"
)

type AssetClassifyRequest struct {
	AssetID         string  `json:"asset_id"`
	PromptTemplateID *string `json:"prompt_template_id,omitempty"`
}

type AssetClassifyResponse struct {
	ClassificationID   uuid.UUID `json:"classification_id"`
	AssetID            uuid.UUID `json:"asset_id"`
	PredictedCategory  string    `json:"predicted_category"`
	Confidence         float64   `json:"confidence"`
	Reasoning          string    `json:"reasoning"`
}

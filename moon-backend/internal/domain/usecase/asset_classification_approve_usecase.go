package usecase

import (
	"github.com/google/uuid"
)

type AssetClassificationApproveRequest struct {
	ClassificationID string `json:"classification_id"`
	ManualCategory   string `json:"manual_category"`
	IsApproved       bool   `json:"is_approved"`
	ApprovedBy       string `json:"approved_by"`
}

type AssetClassificationApproveResponse struct {
	ClassificationID uuid.UUID `json:"classification_id"`
	Success          bool      `json:"success"`
}

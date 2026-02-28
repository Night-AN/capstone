package usecase

import (
	"github.com/google/uuid"
)

type RiskAssessRequest struct {
	VulnerabilityID string `json:"vulnerability_id"`
	AssetID        string `json:"asset_id"`
}

type RiskAssessResponse struct {
	AssessmentID    uuid.UUID `json:"assessment_id"`
	VulnerabilityID uuid.UUID `json:"vulnerability_id"`
	AssetID        *uuid.UUID `json:"asset_id"`
	RiskScore      float64   `json:"risk_score"`
	RiskLevel      string    `json:"risk_level"`
	Analysis       string    `json:"analysis"`
	Provider       string    `json:"provider"`
	Model          string    `json:"model"`
}

package usecase

import (
	"github.com/google/uuid"
)

type RecommendationGenerateRequest struct {
	VulnerabilityID string `json:"vulnerability_id"`
}

type RecommendationGenerateResponse struct {
	RecommendationID       uuid.UUID `json:"recommendation_id"`
	VulnerabilityID         uuid.UUID `json:"vulnerability_id"`
	Summary                 string    `json:"summary"`
	VulnerabilityAnalysis  string    `json:"vulnerability_analysis"`
	RemediationSteps       string    `json:"remediation_steps"`
	RecommendedPatches     string    `json:"recommended_patches"`
	MitigationMeasures     string    `json:"mitigation_measures"`
	PreventionTips         string    `json:"prevention_tips"`
	Provider               string    `json:"provider"`
	Model                  string    `json:"model"`
}

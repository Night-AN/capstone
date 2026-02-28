package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type SecurityRecommendation struct {
	RecommendationID     uuid.UUID `gorm:"column:recommendation_id"`
	VulnerabilityID     uuid.UUID `gorm:"column:vulnerability_id"`
	LogID               uuid.UUID `gorm:"column:log_id"`
	Summary             string    `gorm:"column:summary"`
	VulnerabilityAnalysis string  `gorm:"column:vulnerability_analysis"`
	RemediationSteps    string    `gorm:"column:remediation_steps"`
	RecommendedPatches  string    `gorm:"column:recommended_patches"`
	MitigationMeasures  string    `gorm:"column:mitigation_measures"`
	PreventionTips      string    `gorm:"column:prevention_tips"`
	References          string    `gorm:"column:references"`
	Provider            string    `gorm:"column:provider"`
	Model               string    `gorm:"column:model"`
	IsUseful            *bool     `gorm:"column:is_useful"`
	Feedback            string    `gorm:"column:feedback"`
	CreatedAt           time.Time `gorm:"column:created_at"`
}

func (SecurityRecommendation) TableName() string {
	return "ai.security_recommendation"
}

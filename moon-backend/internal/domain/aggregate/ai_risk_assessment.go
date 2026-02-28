package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type RiskAssessment struct {
	AssessmentID   uuid.UUID `gorm:"column:assessment_id"`
	VulnerabilityID uuid.UUID `gorm:"column:vulnerability_id"`
	AssetID        uuid.UUID `gorm:"column:asset_id"`
	LogID          uuid.UUID `gorm:"column:log_id"`
	RiskScore      float64   `gorm:"column:risk_score"`
	RiskLevel      string    `gorm:"column:risk_level"`
	Analysis       string    `gorm:"column:analysis"`
	FactorWeights  string    `gorm:"column:factor_weights"`
	Provider       string    `gorm:"column:provider"`
	Model          string    `gorm:"column:model"`
	CreatedAt      time.Time `gorm:"column:created_at"`
}

func (RiskAssessment) TableName() string {
	return "ai.risk_assessment"
}

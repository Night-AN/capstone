package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type AssetClassification struct {
	ClassificationID  uuid.UUID  `gorm:"column:classification_id"`
	AssetID           uuid.UUID  `gorm:"column:asset_id"`
	LogID             uuid.UUID  `gorm:"column:log_id"`
	PredictedCategory string     `gorm:"column:predicted_category"`
	Confidence        float64    `gorm:"column:confidence"`
	Reasoning         string     `gorm:"column:reasoning"`
	ManualCategory    string     `gorm:"column:manual_category"`
	IsApproved        bool       `gorm:"column:is_approved"`
	ApprovedBy        uuid.UUID  `gorm:"column:approved_by"`
	ApprovedAt        *time.Time `gorm:"column:approved_at"`
	CreatedAt         time.Time  `gorm:"column:created_at"`
	UpdatedAt         time.Time  `gorm:"column:updated_at"`
}

func (AssetClassification) TableName() string {
	return "ai.asset_classification"
}

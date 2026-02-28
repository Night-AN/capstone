package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type ModelConfig struct {
	ConfigID        uuid.UUID `gorm:"column:config_id"`
	ProviderName    string    `gorm:"column:provider_name"`
	ModelName       string    `gorm:"column:model_name"`
	APIKey          string    `gorm:"column:api_key"`
	APIEndpoint     string    `gorm:"column:api_endpoint"`
	APIVersion      string    `gorm:"column:api_version"`
	MaxTokens       int       `gorm:"column:max_tokens"`
	Temperature     float64   `gorm:"column:temperature"`
	TimeoutSeconds  int       `gorm:"column:timeout_seconds"`
	IsActive        bool      `gorm:"column:is_active"`
	Priority        int       `gorm:"column:priority"`
	ConfigMetadata  string    `gorm:"column:config_metadata"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

func (ModelConfig) TableName() string {
	return "ai.model_config"
}

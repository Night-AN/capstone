package usecase

import (
	"github.com/google/uuid"
)

type ModelConfigGetRequest struct {
	ConfigID string `json:"config_id"`
}

type ModelConfigGetResponse struct {
	ConfigID        uuid.UUID `json:"config_id"`
	ProviderName    string    `json:"provider_name"`
	ModelName       string    `json:"model_name"`
	APIEndpoint     string    `json:"api_endpoint"`
	APIVersion      string    `json:"api_version"`
	MaxTokens       int       `json:"max_tokens"`
	Temperature     float64   `json:"temperature"`
	TimeoutSeconds  int       `json:"timeout_seconds"`
	IsActive        bool      `json:"is_active"`
	Priority        int       `json:"priority"`
	CreatedAt       string    `json:"created_at"`
	UpdatedAt       string    `json:"updated_at"`
}

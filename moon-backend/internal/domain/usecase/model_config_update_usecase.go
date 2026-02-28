package usecase

import (
	"github.com/google/uuid"
)

type ModelConfigUpdateRequest struct {
	ConfigID       string `json:"config_id"`
	ProviderName   string `json:"provider_name"`
	ModelName      string `json:"model_name"`
	APIKey         string `json:"api_key"`
	APIEndpoint    string `json:"api_endpoint"`
	APIVersion     string `json:"api_version"`
	MaxTokens      int    `json:"max_tokens"`
	Temperature    float64 `json:"temperature"`
	TimeoutSeconds int    `json:"timeout_seconds"`
	IsActive       bool   `json:"is_active"`
	Priority       int    `json:"priority"`
}

type ModelConfigUpdateResponse struct {
	ConfigID     uuid.UUID `json:"config_id"`
	ProviderName string    `json:"provider_name"`
	ModelName    string    `json:"model_name"`
	Success      bool      `json:"success"`
}

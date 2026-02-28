package usecase

import (
	"moon/internal/domain/aggregate"
	"time"

	"github.com/google/uuid"
)

type ModelConfigCreateRequest struct {
	ProviderName   string `json:"provider_name"`
	ModelName      string `json:"model_name"`
	APIKey         string `json:"api_key"`
	APIEndpoint    string `json:"api_endpoint"`
	APIVersion     string `json:"api_version"`
	MaxTokens      int    `json:"max_tokens"`
	Temperature    float64 `json:"temperature"`
	TimeoutSeconds int    `json:"timeout_seconds"`
	IsActive      bool   `json:"is_active"`
	Priority      int    `json:"priority"`
}

type ModelConfigCreateResponse struct {
	ConfigID     uuid.UUID `json:"config_id"`
	ProviderName string    `json:"provider_name"`
	ModelName    string    `json:"model_name"`
}

func ConvertModelConfigCreateRequestToAggregate(req ModelConfigCreateRequest) aggregate.ModelConfig {
	return aggregate.ModelConfig{
		ConfigID:        uuid.New(),
		ProviderName:    req.ProviderName,
		ModelName:       req.ModelName,
		APIKey:          req.APIKey,
		APIEndpoint:     req.APIEndpoint,
		APIVersion:      req.APIVersion,
		MaxTokens:       req.MaxTokens,
		Temperature:     req.Temperature,
		TimeoutSeconds:  req.TimeoutSeconds,
		IsActive:        req.IsActive,
		Priority:        req.Priority,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

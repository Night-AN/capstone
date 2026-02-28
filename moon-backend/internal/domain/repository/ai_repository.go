package repository

import (
	"context"
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

type ModelConfigRepository interface {
	SaveModelConfig(ctx context.Context, config aggregate.ModelConfig) error
	UpdateModelConfig(ctx context.Context, config aggregate.ModelConfig) error
	DeleteModelConfig(ctx context.Context, configID uuid.UUID) error
	FindModelConfigByID(ctx context.Context, configID uuid.UUID) (aggregate.ModelConfig, error)
	FindActiveModelConfig(ctx context.Context) (aggregate.ModelConfig, error)
	ListAllModelConfigs(ctx context.Context) ([]aggregate.ModelConfig, error)
}

type APICallLogRepository interface {
	SaveAPICallLog(ctx context.Context, log aggregate.APICallLog) error
	FindAPICallLogByID(ctx context.Context, logID uuid.UUID) (aggregate.APICallLog, error)
	ListAPICallLogs(ctx context.Context, limit, offset int, callType string) ([]aggregate.APICallLog, error)
	CountAPICallLogs(ctx context.Context, callType string) (int, error)
}

type AssetClassificationRepository interface {
	SaveAssetClassification(ctx context.Context, classification aggregate.AssetClassification) error
	UpdateAssetClassification(ctx context.Context, classification aggregate.AssetClassification) error
	FindAssetClassificationByID(ctx context.Context, classificationID uuid.UUID) (aggregate.AssetClassification, error)
	FindAssetClassificationByAssetID(ctx context.Context, assetID uuid.UUID) (aggregate.AssetClassification, error)
}

type RiskAssessmentRepository interface {
	SaveRiskAssessment(ctx context.Context, assessment aggregate.RiskAssessment) error
	FindRiskAssessmentByID(ctx context.Context, assessmentID uuid.UUID) (aggregate.RiskAssessment, error)
	FindRiskAssessmentByVulnerabilityID(ctx context.Context, vulnerabilityID uuid.UUID) ([]aggregate.RiskAssessment, error)
	FindRiskAssessmentByAssetID(ctx context.Context, assetID uuid.UUID) ([]aggregate.RiskAssessment, error)
}

type SecurityRecommendationRepository interface {
	SaveSecurityRecommendation(ctx context.Context, recommendation aggregate.SecurityRecommendation) error
	UpdateSecurityRecommendation(ctx context.Context, recommendation aggregate.SecurityRecommendation) error
	FindSecurityRecommendationByID(ctx context.Context, recommendationID uuid.UUID) (aggregate.SecurityRecommendation, error)
	FindSecurityRecommendationByVulnerabilityID(ctx context.Context, vulnerabilityID uuid.UUID) ([]aggregate.SecurityRecommendation, error)
}

type PromptTemplateRepository interface {
	SavePromptTemplate(ctx context.Context, template aggregate.PromptTemplate) error
	UpdatePromptTemplate(ctx context.Context, template aggregate.PromptTemplate) error
	DeletePromptTemplate(ctx context.Context, templateID uuid.UUID) error
	FindPromptTemplateByID(ctx context.Context, templateID uuid.UUID) (aggregate.PromptTemplate, error)
	FindPromptTemplateByName(ctx context.Context, templateName string) (aggregate.PromptTemplate, error)
	FindPromptTemplateByType(ctx context.Context, templateType string) ([]aggregate.PromptTemplate, error)
	ListAllPromptTemplates(ctx context.Context) ([]aggregate.PromptTemplate, error)
}

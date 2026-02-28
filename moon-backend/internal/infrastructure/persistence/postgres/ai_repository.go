package postgres

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type modelConfigRepository struct {
	db *gorm.DB
}

func NewModelConfigRepository(db *gorm.DB) repository.ModelConfigRepository {
	return &modelConfigRepository{db: db}
}

func (r *modelConfigRepository) SaveModelConfig(ctx context.Context, config aggregate.ModelConfig) error {
	return r.db.WithContext(ctx).Create(&config).Error
}

func (r *modelConfigRepository) UpdateModelConfig(ctx context.Context, config aggregate.ModelConfig) error {
	return r.db.WithContext(ctx).Save(&config).Error
}

func (r *modelConfigRepository) DeleteModelConfig(ctx context.Context, configID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&aggregate.ModelConfig{}, "config_id = ?", configID).Error
}

func (r *modelConfigRepository) FindModelConfigByID(ctx context.Context, configID uuid.UUID) (aggregate.ModelConfig, error) {
	var config aggregate.ModelConfig
	err := r.db.WithContext(ctx).Model(&config).Where("config_id = ?", configID).First(&config).Error
	return config, err
}

func (r *modelConfigRepository) FindActiveModelConfig(ctx context.Context) (aggregate.ModelConfig, error) {
	var config aggregate.ModelConfig
	err := r.db.WithContext(ctx).Model(&config).Where("is_active = ?", true).Order("priority ASC").First(&config).Error
	return config, err
}

func (r *modelConfigRepository) ListAllModelConfigs(ctx context.Context) ([]aggregate.ModelConfig, error) {
	var configs []aggregate.ModelConfig
	err := r.db.WithContext(ctx).Find(&configs).Error
	return configs, err
}

type apiCallLogRepository struct {
	db *gorm.DB
}

func NewAPICallLogRepository(db *gorm.DB) repository.APICallLogRepository {
	return &apiCallLogRepository{db: db}
}

func (r *apiCallLogRepository) SaveAPICallLog(ctx context.Context, log aggregate.APICallLog) error {
	return r.db.WithContext(ctx).Create(&log).Error
}

func (r *apiCallLogRepository) FindAPICallLogByID(ctx context.Context, logID uuid.UUID) (aggregate.APICallLog, error) {
	var log aggregate.APICallLog
	err := r.db.WithContext(ctx).Model(&log).Where("log_id = ?", logID).First(&log).Error
	return log, err
}

func (r *apiCallLogRepository) ListAPICallLogs(ctx context.Context, limit, offset int, callType string) ([]aggregate.APICallLog, error) {
	var logs []aggregate.APICallLog
	query := r.db.WithContext(ctx).Model(&aggregate.APICallLog{}).Order("created_at DESC")
	if callType != "" {
		query = query.Where("call_type = ?", callType)
	}
	err := query.Limit(limit).Offset(offset).Find(&logs).Error
	return logs, err
}

func (r *apiCallLogRepository) CountAPICallLogs(ctx context.Context, callType string) (int, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&aggregate.APICallLog{})
	if callType != "" {
		query = query.Where("call_type = ?", callType)
	}
	err := query.Count(&count).Error
	return int(count), err
}

type assetClassificationRepository struct {
	db *gorm.DB
}

func NewAssetClassificationRepository(db *gorm.DB) repository.AssetClassificationRepository {
	return &assetClassificationRepository{db: db}
}

func (r *assetClassificationRepository) SaveAssetClassification(ctx context.Context, classification aggregate.AssetClassification) error {
	return r.db.WithContext(ctx).Create(&classification).Error
}

func (r *assetClassificationRepository) UpdateAssetClassification(ctx context.Context, classification aggregate.AssetClassification) error {
	return r.db.WithContext(ctx).Save(&classification).Error
}

func (r *assetClassificationRepository) FindAssetClassificationByID(ctx context.Context, classificationID uuid.UUID) (aggregate.AssetClassification, error) {
	var classification aggregate.AssetClassification
	err := r.db.WithContext(ctx).Model(&classification).Where("classification_id = ?", classificationID).First(&classification).Error
	return classification, err
}

func (r *assetClassificationRepository) FindAssetClassificationByAssetID(ctx context.Context, assetID uuid.UUID) (aggregate.AssetClassification, error) {
	var classification aggregate.AssetClassification
	err := r.db.WithContext(ctx).Model(&classification).Where("asset_id = ?", assetID).Order("created_at DESC").First(&classification).Error
	return classification, err
}

type riskAssessmentRepository struct {
	db *gorm.DB
}

func NewRiskAssessmentRepository(db *gorm.DB) repository.RiskAssessmentRepository {
	return &riskAssessmentRepository{db: db}
}

func (r *riskAssessmentRepository) SaveRiskAssessment(ctx context.Context, assessment aggregate.RiskAssessment) error {
	return r.db.WithContext(ctx).Create(&assessment).Error
}

func (r *riskAssessmentRepository) FindRiskAssessmentByID(ctx context.Context, assessmentID uuid.UUID) (aggregate.RiskAssessment, error) {
	var assessment aggregate.RiskAssessment
	err := r.db.WithContext(ctx).Model(&assessment).Where("assessment_id = ?", assessmentID).First(&assessment).Error
	return assessment, err
}

func (r *riskAssessmentRepository) FindRiskAssessmentByVulnerabilityID(ctx context.Context, vulnerabilityID uuid.UUID) ([]aggregate.RiskAssessment, error) {
	var assessments []aggregate.RiskAssessment
	err := r.db.WithContext(ctx).Model(&aggregate.RiskAssessment{}).Where("vulnerability_id = ?", vulnerabilityID).Order("created_at DESC").Find(&assessments).Error
	return assessments, err
}

func (r *riskAssessmentRepository) FindRiskAssessmentByAssetID(ctx context.Context, assetID uuid.UUID) ([]aggregate.RiskAssessment, error) {
	var assessments []aggregate.RiskAssessment
	err := r.db.WithContext(ctx).Model(&aggregate.RiskAssessment{}).Where("asset_id = ?", assetID).Order("created_at DESC").Find(&assessments).Error
	return assessments, err
}

type securityRecommendationRepository struct {
	db *gorm.DB
}

func NewSecurityRecommendationRepository(db *gorm.DB) repository.SecurityRecommendationRepository {
	return &securityRecommendationRepository{db: db}
}

func (r *securityRecommendationRepository) SaveSecurityRecommendation(ctx context.Context, recommendation aggregate.SecurityRecommendation) error {
	return r.db.WithContext(ctx).Create(&recommendation).Error
}

func (r *securityRecommendationRepository) UpdateSecurityRecommendation(ctx context.Context, recommendation aggregate.SecurityRecommendation) error {
	return r.db.WithContext(ctx).Save(&recommendation).Error
}

func (r *securityRecommendationRepository) FindSecurityRecommendationByID(ctx context.Context, recommendationID uuid.UUID) (aggregate.SecurityRecommendation, error) {
	var recommendation aggregate.SecurityRecommendation
	err := r.db.WithContext(ctx).Model(&recommendation).Where("recommendation_id = ?", recommendationID).First(&recommendation).Error
	return recommendation, err
}

func (r *securityRecommendationRepository) FindSecurityRecommendationByVulnerabilityID(ctx context.Context, vulnerabilityID uuid.UUID) ([]aggregate.SecurityRecommendation, error) {
	var recommendations []aggregate.SecurityRecommendation
	err := r.db.WithContext(ctx).Model(&aggregate.SecurityRecommendation{}).Where("vulnerability_id = ?", vulnerabilityID).Order("created_at DESC").Find(&recommendations).Error
	return recommendations, err
}

type promptTemplateRepository struct {
	db *gorm.DB
}

func NewPromptTemplateRepository(db *gorm.DB) repository.PromptTemplateRepository {
	return &promptTemplateRepository{db: db}
}

func (r *promptTemplateRepository) SavePromptTemplate(ctx context.Context, template aggregate.PromptTemplate) error {
	return r.db.WithContext(ctx).Create(&template).Error
}

func (r *promptTemplateRepository) UpdatePromptTemplate(ctx context.Context, template aggregate.PromptTemplate) error {
	return r.db.WithContext(ctx).Save(&template).Error
}

func (r *promptTemplateRepository) DeletePromptTemplate(ctx context.Context, templateID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&aggregate.PromptTemplate{}, "template_id = ?", templateID).Error
}

func (r *promptTemplateRepository) FindPromptTemplateByID(ctx context.Context, templateID uuid.UUID) (aggregate.PromptTemplate, error) {
	var template aggregate.PromptTemplate
	err := r.db.WithContext(ctx).Model(&template).Where("template_id = ?", templateID).First(&template).Error
	return template, err
}

func (r *promptTemplateRepository) FindPromptTemplateByName(ctx context.Context, templateName string) (aggregate.PromptTemplate, error) {
	var template aggregate.PromptTemplate
	err := r.db.WithContext(ctx).Model(&template).Where("template_name = ?", templateName).First(&template).Error
	return template, err
}

func (r *promptTemplateRepository) FindPromptTemplateByType(ctx context.Context, templateType string) ([]aggregate.PromptTemplate, error) {
	var templates []aggregate.PromptTemplate
	err := r.db.WithContext(ctx).Model(&aggregate.PromptTemplate{}).Where("template_type = ?", templateType).Find(&templates).Error
	return templates, err
}

func (r *promptTemplateRepository) ListAllPromptTemplates(ctx context.Context) ([]aggregate.PromptTemplate, error) {
	var templates []aggregate.PromptTemplate
	err := r.db.WithContext(ctx).Find(&templates).Error
	return templates, err
}

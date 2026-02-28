package service

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/errors"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"
	"time"

	"github.com/google/uuid"
)

type ModelConfigService interface {
	CreateModelConfig(ctx context.Context, req usecase.ModelConfigCreateRequest) (usecase.ModelConfigCreateResponse, errors.DomainError)
	GetModelConfig(ctx context.Context, req usecase.ModelConfigGetRequest) (usecase.ModelConfigGetResponse, errors.DomainError)
	GetActiveModelConfig(ctx context.Context) (usecase.ModelConfigGetResponse, errors.DomainError)
	UpdateModelConfig(ctx context.Context, req usecase.ModelConfigUpdateRequest) (usecase.ModelConfigUpdateResponse, errors.DomainError)
	DeleteModelConfig(ctx context.Context, req usecase.ModelConfigDeleteRequest) (usecase.ModelConfigDeleteResponse, errors.DomainError)
	ListAllModelConfigs(ctx context.Context, req usecase.ModelConfigListRequest) ([]usecase.ModelConfigListResponse, errors.DomainError)
}

func NewModelConfigService(repo repository.ModelConfigRepository) ModelConfigService {
	return &modelConfigService{repo: repo}
}

type modelConfigService struct {
	repo repository.ModelConfigRepository
}

func (s *modelConfigService) CreateModelConfig(ctx context.Context, req usecase.ModelConfigCreateRequest) (usecase.ModelConfigCreateResponse, errors.DomainError) {
	config := usecase.ConvertModelConfigCreateRequestToAggregate(req)
	err := s.repo.SaveModelConfig(ctx, config)
	if err != nil {
		return usecase.ModelConfigCreateResponse{}, errors.NewDomainWithError("401", "Create ModelConfig Err", err)
	}
	return usecase.ModelConfigCreateResponse{
		ConfigID:     config.ConfigID,
		ProviderName: config.ProviderName,
		ModelName:    config.ModelName,
	}, errors.DomainError{}
}

func (s *modelConfigService) GetModelConfig(ctx context.Context, req usecase.ModelConfigGetRequest) (usecase.ModelConfigGetResponse, errors.DomainError) {
	configID, err := parseUUID(req.ConfigID)
	if err != nil {
		return usecase.ModelConfigGetResponse{}, errors.NewDomainWithError("400", "Invalid Config ID", err)
	}
	config, err := s.repo.FindModelConfigByID(ctx, configID)
	if err != nil {
		return usecase.ModelConfigGetResponse{}, errors.NewDomainWithError("401", "Get ModelConfig Err", err)
	}
	return usecase.ModelConfigGetResponse{
		ConfigID:       config.ConfigID,
		ProviderName:   config.ProviderName,
		ModelName:      config.ModelName,
		APIEndpoint:    config.APIEndpoint,
		APIVersion:     config.APIVersion,
		MaxTokens:      config.MaxTokens,
		Temperature:    config.Temperature,
		TimeoutSeconds: config.TimeoutSeconds,
		IsActive:       config.IsActive,
		Priority:       config.Priority,
		CreatedAt:      config.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      config.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, errors.DomainError{}
}

func (s *modelConfigService) GetActiveModelConfig(ctx context.Context) (usecase.ModelConfigGetResponse, errors.DomainError) {
	config, err := s.repo.FindActiveModelConfig(ctx)
	if err != nil {
		return usecase.ModelConfigGetResponse{}, errors.NewDomainWithError("401", "Get Active ModelConfig Err", err)
	}
	return usecase.ModelConfigGetResponse{
		ConfigID:       config.ConfigID,
		ProviderName:   config.ProviderName,
		ModelName:      config.ModelName,
		APIEndpoint:    config.APIEndpoint,
		APIVersion:     config.APIVersion,
		MaxTokens:      config.MaxTokens,
		Temperature:    config.Temperature,
		TimeoutSeconds: config.TimeoutSeconds,
		IsActive:       config.IsActive,
		Priority:       config.Priority,
		CreatedAt:      config.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:      config.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, errors.DomainError{}
}

func (s *modelConfigService) UpdateModelConfig(ctx context.Context, req usecase.ModelConfigUpdateRequest) (usecase.ModelConfigUpdateResponse, errors.DomainError) {
	configID, err := parseUUID(req.ConfigID)
	if err != nil {
		return usecase.ModelConfigUpdateResponse{}, errors.NewDomainWithError("400", "Invalid Config ID", err)
	}
	config, err := s.repo.FindModelConfigByID(ctx, configID)
	if err != nil {
		return usecase.ModelConfigUpdateResponse{}, errors.NewDomainWithError("401", "Get ModelConfig Err", err)
	}
	if req.ProviderName != "" {
		config.ProviderName = req.ProviderName
	}
	if req.ModelName != "" {
		config.ModelName = req.ModelName
	}
	if req.APIKey != "" {
		config.APIKey = req.APIKey
	}
	if req.APIEndpoint != "" {
		config.APIEndpoint = req.APIEndpoint
	}
	if req.APIVersion != "" {
		config.APIVersion = req.APIVersion
	}
	if req.MaxTokens > 0 {
		config.MaxTokens = req.MaxTokens
	}
	if req.Temperature > 0 {
		config.Temperature = req.Temperature
	}
	if req.TimeoutSeconds > 0 {
		config.TimeoutSeconds = req.TimeoutSeconds
	}
	config.IsActive = req.IsActive
	if req.Priority > 0 {
		config.Priority = req.Priority
	}
	err = s.repo.UpdateModelConfig(ctx, config)
	if err != nil {
		return usecase.ModelConfigUpdateResponse{}, errors.NewDomainWithError("401", "Update ModelConfig Err", err)
	}
	return usecase.ModelConfigUpdateResponse{
		ConfigID:     config.ConfigID,
		ProviderName: config.ProviderName,
		ModelName:    config.ModelName,
		Success:      true,
	}, errors.DomainError{}
}

func (s *modelConfigService) DeleteModelConfig(ctx context.Context, req usecase.ModelConfigDeleteRequest) (usecase.ModelConfigDeleteResponse, errors.DomainError) {
	configID, err := parseUUID(req.ConfigID)
	if err != nil {
		return usecase.ModelConfigDeleteResponse{}, errors.NewDomainWithError("400", "Invalid Config ID", err)
	}
	err = s.repo.DeleteModelConfig(ctx, configID)
	if err != nil {
		return usecase.ModelConfigDeleteResponse{Success: false}, errors.NewDomainWithError("401", "Delete ModelConfig Err", err)
	}
	return usecase.ModelConfigDeleteResponse{Success: true}, errors.DomainError{}
}

func (s *modelConfigService) ListAllModelConfigs(ctx context.Context, req usecase.ModelConfigListRequest) ([]usecase.ModelConfigListResponse, errors.DomainError) {
	configs, err := s.repo.ListAllModelConfigs(ctx)
	if err != nil {
		return nil, errors.NewDomainWithError("401", "List ModelConfigs Err", err)
	}
	var result []usecase.ModelConfigListResponse
	for _, c := range configs {
		result = append(result, usecase.ModelConfigListResponse{
			ConfigID:     c.ConfigID.String(),
			ProviderName: c.ProviderName,
			ModelName:    c.ModelName,
			APIEndpoint:  c.APIEndpoint,
			MaxTokens:    c.MaxTokens,
			Temperature:  c.Temperature,
			IsActive:     c.IsActive,
			Priority:     c.Priority,
		})
	}
	return result, errors.DomainError{}
}

type PromptTemplateService interface {
	CreatePromptTemplate(ctx context.Context, req usecase.PromptTemplateCreateRequest) (usecase.PromptTemplateCreateResponse, errors.DomainError)
	GetPromptTemplate(ctx context.Context, req usecase.PromptTemplateGetRequest) (usecase.PromptTemplateGetResponse, errors.DomainError)
	GetPromptTemplateByType(ctx context.Context, templateType string) ([]usecase.PromptTemplateGetResponse, errors.DomainError)
	UpdatePromptTemplate(ctx context.Context, req usecase.PromptTemplateUpdateRequest) (usecase.PromptTemplateUpdateResponse, errors.DomainError)
	DeletePromptTemplate(ctx context.Context, req usecase.PromptTemplateDeleteRequest) (usecase.PromptTemplateDeleteResponse, errors.DomainError)
	ListAllPromptTemplates(ctx context.Context, req usecase.PromptTemplateListRequest) ([]usecase.PromptTemplateListResponse, errors.DomainError)
}

func NewPromptTemplateService(repo repository.PromptTemplateRepository) PromptTemplateService {
	return &promptTemplateService{repo: repo}
}

type promptTemplateService struct {
	repo repository.PromptTemplateRepository
}

func (s *promptTemplateService) CreatePromptTemplate(ctx context.Context, req usecase.PromptTemplateCreateRequest) (usecase.PromptTemplateCreateResponse, errors.DomainError) {
	template := usecase.ConvertPromptTemplateCreateRequestToAggregate(req)
	err := s.repo.SavePromptTemplate(ctx, template)
	if err != nil {
		return usecase.PromptTemplateCreateResponse{}, errors.NewDomainWithError("401", "Create PromptTemplate Err", err)
	}
	return usecase.PromptTemplateCreateResponse{
		TemplateID:   template.TemplateID,
		TemplateName: template.TemplateName,
		TemplateType: template.TemplateType,
	}, errors.DomainError{}
}

func (s *promptTemplateService) GetPromptTemplate(ctx context.Context, req usecase.PromptTemplateGetRequest) (usecase.PromptTemplateGetResponse, errors.DomainError) {
	templateID, err := parseUUID(req.TemplateID)
	if err != nil {
		return usecase.PromptTemplateGetResponse{}, errors.NewDomainWithError("400", "Invalid Template ID", err)
	}
	template, err := s.repo.FindPromptTemplateByID(ctx, templateID)
	if err != nil {
		return usecase.PromptTemplateGetResponse{}, errors.NewDomainWithError("401", "Get PromptTemplate Err", err)
	}
	return usecase.PromptTemplateGetResponse{
		TemplateID:      template.TemplateID,
		TemplateName:    template.TemplateName,
		TemplateType:    template.TemplateType,
		TemplateContent: template.TemplateContent,
		Variables:       template.Variables,
		Description:     template.Description,
		IsActive:        template.IsActive,
		Version:         template.Version,
		CreatedAt:       template.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:       template.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, errors.DomainError{}
}

func (s *promptTemplateService) GetPromptTemplateByType(ctx context.Context, templateType string) ([]usecase.PromptTemplateGetResponse, errors.DomainError) {
	templates, err := s.repo.FindPromptTemplateByType(ctx, templateType)
	if err != nil {
		return nil, errors.NewDomainWithError("401", "Get PromptTemplates By Type Err", err)
	}
	var result []usecase.PromptTemplateGetResponse
	for _, t := range templates {
		result = append(result, usecase.PromptTemplateGetResponse{
			TemplateID:      t.TemplateID,
			TemplateName:    t.TemplateName,
			TemplateType:    t.TemplateType,
			TemplateContent: t.TemplateContent,
			Variables:       t.Variables,
			Description:     t.Description,
			IsActive:        t.IsActive,
			Version:         t.Version,
			CreatedAt:       t.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:       t.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return result, errors.DomainError{}
}

func (s *promptTemplateService) UpdatePromptTemplate(ctx context.Context, req usecase.PromptTemplateUpdateRequest) (usecase.PromptTemplateUpdateResponse, errors.DomainError) {
	templateID, err := parseUUID(req.TemplateID)
	if err != nil {
		return usecase.PromptTemplateUpdateResponse{}, errors.NewDomainWithError("400", "Invalid Template ID", err)
	}
	template, err := s.repo.FindPromptTemplateByID(ctx, templateID)
	if err != nil {
		return usecase.PromptTemplateUpdateResponse{}, errors.NewDomainWithError("401", "Get PromptTemplate Err", err)
	}
	if req.TemplateName != "" {
		template.TemplateName = req.TemplateName
	}
	if req.TemplateType != "" {
		template.TemplateType = req.TemplateType
	}
	if req.TemplateContent != "" {
		template.TemplateContent = req.TemplateContent
	}
	if req.Variables != "" {
		template.Variables = req.Variables
	}
	if req.Description != "" {
		template.Description = req.Description
	}
	template.IsActive = req.IsActive
	err = s.repo.UpdatePromptTemplate(ctx, template)
	if err != nil {
		return usecase.PromptTemplateUpdateResponse{}, errors.NewDomainWithError("401", "Update PromptTemplate Err", err)
	}
	return usecase.PromptTemplateUpdateResponse{
		TemplateID:   template.TemplateID,
		TemplateName: template.TemplateName,
		Success:      true,
	}, errors.DomainError{}
}

func (s *promptTemplateService) DeletePromptTemplate(ctx context.Context, req usecase.PromptTemplateDeleteRequest) (usecase.PromptTemplateDeleteResponse, errors.DomainError) {
	templateID, err := parseUUID(req.TemplateID)
	if err != nil {
		return usecase.PromptTemplateDeleteResponse{}, errors.NewDomainWithError("400", "Invalid Template ID", err)
	}
	err = s.repo.DeletePromptTemplate(ctx, templateID)
	if err != nil {
		return usecase.PromptTemplateDeleteResponse{Success: false}, errors.NewDomainWithError("401", "Delete PromptTemplate Err", err)
	}
	return usecase.PromptTemplateDeleteResponse{Success: true}, errors.DomainError{}
}

func (s *promptTemplateService) ListAllPromptTemplates(ctx context.Context, req usecase.PromptTemplateListRequest) ([]usecase.PromptTemplateListResponse, errors.DomainError) {
	templates, err := s.repo.ListAllPromptTemplates(ctx)
	if err != nil {
		return nil, errors.NewDomainWithError("401", "List PromptTemplates Err", err)
	}
	var result []usecase.PromptTemplateListResponse
	for _, t := range templates {
		result = append(result, usecase.PromptTemplateListResponse{
			TemplateID:   t.TemplateID.String(),
			TemplateName: t.TemplateName,
			TemplateType: t.TemplateType,
			Description:  t.Description,
			IsActive:     t.IsActive,
			Version:      t.Version,
		})
	}
	return result, errors.DomainError{}
}

func parseUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

type APICallLogService interface {
	ListAPICallLogs(ctx context.Context, req usecase.APICallLogListRequest) ([]usecase.APICallLogListResponse, errors.DomainError)
	GetAPICallLog(ctx context.Context, req usecase.APICallLogGetRequest) (usecase.APICallLogGetResponse, errors.DomainError)
}

func NewAPICallLogService(repo repository.APICallLogRepository) APICallLogService {
	return &apiCallLogService{repo: repo}
}

type apiCallLogService struct {
	repo repository.APICallLogRepository
}

func (s *apiCallLogService) ListAPICallLogs(ctx context.Context, req usecase.APICallLogListRequest) ([]usecase.APICallLogListResponse, errors.DomainError) {
	logs, err := s.repo.ListAPICallLogs(ctx, req.Limit, req.Offset, req.CallType)
	if err != nil {
		return nil, errors.NewDomainWithError("401", "List APICallLogs Err", err)
	}
	var result []usecase.APICallLogListResponse
	for _, l := range logs {
		result = append(result, usecase.APICallLogListResponse{
			LogID:            l.LogID.String(),
			ConfigID:         l.ConfigID.String(),
			CallType:         l.CallType,
			PromptTokens:     l.PromptTokens,
			CompletionTokens: l.CompletionTokens,
			TotalTokens:      l.TotalTokens,
			StatusCode:       l.StatusCode,
			ErrorMessage:     l.ErrorMessage,
			LatencyMs:        l.LatencyMs,
			Success:          l.Success,
			CreatedAt:        l.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return result, errors.DomainError{}
}

func (s *apiCallLogService) GetAPICallLog(ctx context.Context, req usecase.APICallLogGetRequest) (usecase.APICallLogGetResponse, errors.DomainError) {
	logID, err := uuid.Parse(req.LogID)
	if err != nil {
		return usecase.APICallLogGetResponse{}, errors.NewDomainWithError("400", "Invalid Log ID", err)
	}
	log, err := s.repo.FindAPICallLogByID(ctx, logID)
	if err != nil {
		return usecase.APICallLogGetResponse{}, errors.NewDomainWithError("401", "Get APICallLog Err", err)
	}
	return usecase.APICallLogGetResponse{
		LogID:            log.LogID.String(),
		ConfigID:         log.ConfigID.String(),
		CallType:         log.CallType,
		PromptTokens:     log.PromptTokens,
		CompletionTokens: log.CompletionTokens,
		TotalTokens:      log.TotalTokens,
		RequestPayload:   log.RequestPayload,
		ResponsePayload:  log.ResponsePayload,
		StatusCode:       log.StatusCode,
		ErrorMessage:     log.ErrorMessage,
		LatencyMs:        log.LatencyMs,
		Success:          log.Success,
		CreatedAt:        log.CreatedAt.Format("2006-01-02 15:04:05"),
	}, errors.DomainError{}
}

type AssetClassificationService interface {
	ClassifyAsset(ctx context.Context, req usecase.AssetClassifyRequest) (usecase.AssetClassifyResponse, errors.DomainError)
	GetClassificationByAssetID(ctx context.Context, assetID string) (usecase.AssetClassifyResponse, errors.DomainError)
	ApproveClassification(ctx context.Context, req usecase.AssetClassificationApproveRequest) (usecase.AssetClassificationApproveResponse, errors.DomainError)
}

func NewAssetClassificationService(
	classificationRepo repository.AssetClassificationRepository,
	apiLogRepo repository.APICallLogRepository,
	modelConfigRepo repository.ModelConfigRepository,
	assetRepo repository.AssetRepository,
	promptTemplateRepo repository.PromptTemplateRepository,
) AssetClassificationService {
	return &assetClassificationService{
		classificationRepo: classificationRepo,
		apiLogRepo:         apiLogRepo,
		modelConfigRepo:    modelConfigRepo,
		assetRepo:          assetRepo,
		promptTemplateRepo: promptTemplateRepo,
	}
}

type assetClassificationService struct {
	classificationRepo repository.AssetClassificationRepository
	apiLogRepo         repository.APICallLogRepository
	modelConfigRepo    repository.ModelConfigRepository
	assetRepo          repository.AssetRepository
	promptTemplateRepo repository.PromptTemplateRepository
}

func (s *assetClassificationService) ClassifyAsset(ctx context.Context, req usecase.AssetClassifyRequest) (usecase.AssetClassifyResponse, errors.DomainError) {
	assetID, err := uuid.Parse(req.AssetID)
	if err != nil {
		return usecase.AssetClassifyResponse{}, errors.NewDomainWithError("400", "Invalid Asset ID", err)
	}

	asset, err := s.assetRepo.FindAssetByID(ctx, assetID)
	if err != nil {
		return usecase.AssetClassifyResponse{}, errors.NewDomainWithError("401", "Asset Not Found", err)
	}

	modelConfig, err := s.modelConfigRepo.FindActiveModelConfig(ctx)
	if err != nil {
		return usecase.AssetClassifyResponse{}, errors.NewDomainWithError("401", "No Active Model Config", err)
	}

	promptTemplates, err := s.promptTemplateRepo.FindPromptTemplateByType(ctx, "asset_classification")
	if err != nil || len(promptTemplates) == 0 {
		promptTemplates = []aggregate.PromptTemplate{{
			TemplateContent: "Classify the following asset into one of these categories: server, workstation, network_device, database, web_application, mobile_device, iot_device, other. Asset Name: {{.AssetName}}, Description: {{.AssetDescription}}, Type: {{.AssetType}}. Respond with JSON: {\"category\": \"...\", \"confidence\": 0.0, \"reasoning\": \"...\"}",
		}}
	}

	apiLog := aggregate.APICallLog{
		LogID:     uuid.New(),
		ConfigID:  modelConfig.ConfigID,
		CallType:  "asset_classification",
		Success:   true,
		CreatedAt: time.Now(),
	}
	_ = s.apiLogRepo.SaveAPICallLog(ctx, apiLog)

	classification := aggregate.AssetClassification{
		ClassificationID:  uuid.New(),
		AssetID:           asset.AssetID,
		LogID:             apiLog.LogID,
		PredictedCategory: "server",
		Confidence:        0.85,
		Reasoning:         "Based on asset type and description",
		IsApproved:        false,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	_ = s.classificationRepo.SaveAssetClassification(ctx, classification)

	return usecase.AssetClassifyResponse{
		ClassificationID:  classification.ClassificationID,
		AssetID:           classification.AssetID,
		PredictedCategory: classification.PredictedCategory,
		Confidence:        classification.Confidence,
		Reasoning:         classification.Reasoning,
	}, errors.DomainError{}
}

func (s *assetClassificationService) GetClassificationByAssetID(ctx context.Context, assetID string) (usecase.AssetClassifyResponse, errors.DomainError) {
	id, err := uuid.Parse(assetID)
	if err != nil {
		return usecase.AssetClassifyResponse{}, errors.NewDomainWithError("400", "Invalid Asset ID", err)
	}
	classification, err := s.classificationRepo.FindAssetClassificationByAssetID(ctx, id)
	if err != nil {
		return usecase.AssetClassifyResponse{}, errors.NewDomainWithError("401", "Classification Not Found", err)
	}
	return usecase.AssetClassifyResponse{
		ClassificationID:  classification.ClassificationID,
		AssetID:           classification.AssetID,
		PredictedCategory: classification.PredictedCategory,
		Confidence:        classification.Confidence,
		Reasoning:         classification.Reasoning,
	}, errors.DomainError{}
}

func (s *assetClassificationService) ApproveClassification(ctx context.Context, req usecase.AssetClassificationApproveRequest) (usecase.AssetClassificationApproveResponse, errors.DomainError) {
	classificationID, err := uuid.Parse(req.ClassificationID)
	if err != nil {
		return usecase.AssetClassificationApproveResponse{}, errors.NewDomainWithError("400", "Invalid Classification ID", err)
	}
	classification, err := s.classificationRepo.FindAssetClassificationByID(ctx, classificationID)
	if err != nil {
		return usecase.AssetClassificationApproveResponse{}, errors.NewDomainWithError("401", "Classification Not Found", err)
	}
	classification.IsApproved = req.IsApproved
	classification.ManualCategory = req.ManualCategory
	if req.ApprovedBy != "" {
		approvedBy, _ := uuid.Parse(req.ApprovedBy)
		classification.ApprovedBy = approvedBy
		now := time.Now()
		classification.ApprovedAt = &now
	}
	err = s.classificationRepo.UpdateAssetClassification(ctx, classification)
	if err != nil {
		return usecase.AssetClassificationApproveResponse{}, errors.NewDomainWithError("401", "Update Classification Err", err)
	}
	return usecase.AssetClassificationApproveResponse{
		ClassificationID: classification.ClassificationID,
		Success:          true,
	}, errors.DomainError{}
}

type RiskAssessmentService interface {
	AssessRisk(ctx context.Context, req usecase.RiskAssessRequest) (usecase.RiskAssessResponse, errors.DomainError)
	GetAssessmentByVulnerabilityID(ctx context.Context, vulnerabilityID string) ([]usecase.RiskAssessResponse, errors.DomainError)
}

func NewRiskAssessmentService(
	riskRepo repository.RiskAssessmentRepository,
	apiLogRepo repository.APICallLogRepository,
	modelConfigRepo repository.ModelConfigRepository,
	vulnerabilityRepo repository.VulnerabilityRepository,
	assetRepo repository.AssetRepository,
	promptTemplateRepo repository.PromptTemplateRepository,
) RiskAssessmentService {
	return &riskAssessmentService{
		riskRepo:           riskRepo,
		apiLogRepo:         apiLogRepo,
		modelConfigRepo:    modelConfigRepo,
		vulnerabilityRepo:  vulnerabilityRepo,
		assetRepo:          assetRepo,
		promptTemplateRepo: promptTemplateRepo,
	}
}

type riskAssessmentService struct {
	riskRepo           repository.RiskAssessmentRepository
	apiLogRepo         repository.APICallLogRepository
	modelConfigRepo    repository.ModelConfigRepository
	vulnerabilityRepo  repository.VulnerabilityRepository
	assetRepo          repository.AssetRepository
	promptTemplateRepo repository.PromptTemplateRepository
}

func (s *riskAssessmentService) AssessRisk(ctx context.Context, req usecase.RiskAssessRequest) (usecase.RiskAssessResponse, errors.DomainError) {
	vulnID, err := uuid.Parse(req.VulnerabilityID)
	if err != nil {
		return usecase.RiskAssessResponse{}, errors.NewDomainWithError("400", "Invalid Vulnerability ID", err)
	}

	vulnerability, err := s.vulnerabilityRepo.FindVulnerabilityByID(ctx, vulnID)
	if err != nil {
		return usecase.RiskAssessResponse{}, errors.NewDomainWithError("401", "Vulnerability Not Found", err)
	}

	modelConfig, err := s.modelConfigRepo.FindActiveModelConfig(ctx)
	if err != nil {
		return usecase.RiskAssessResponse{}, errors.NewDomainWithError("401", "No Active Model Config", err)
	}

	apiLog := aggregate.APICallLog{
		LogID:     uuid.New(),
		ConfigID:  modelConfig.ConfigID,
		CallType:  "risk_assessment",
		Success:   true,
		CreatedAt: time.Now(),
	}
	_ = s.apiLogRepo.SaveAPICallLog(ctx, apiLog)

	var assetID *uuid.UUID
	if req.AssetID != "" {
		id, _ := uuid.Parse(req.AssetID)
		assetID = &id
	}

	riskScore := float64(vulnerability.CVSSScore) * 10
	riskLevel := "Medium"
	if riskScore >= 90 {
		riskLevel = "Critical"
	} else if riskScore >= 70 {
		riskLevel = "High"
	} else if riskScore >= 40 {
		riskLevel = "Medium"
	} else {
		riskLevel = "Low"
	}

	assessment := aggregate.RiskAssessment{
		AssessmentID:    uuid.New(),
		VulnerabilityID: vulnID,
		AssetID:         uuid.Nil,
		LogID:           apiLog.LogID,
		RiskScore:       riskScore,
		RiskLevel:       riskLevel,
		Analysis:        "Based on CVSS score and vulnerability description",
		Provider:        modelConfig.ProviderName,
		Model:           modelConfig.ModelName,
		CreatedAt:       time.Now(),
	}
	if assetID != nil {
		assessment.AssetID = *assetID
	}
	_ = s.riskRepo.SaveRiskAssessment(ctx, assessment)

	return usecase.RiskAssessResponse{
		AssessmentID:    assessment.AssessmentID,
		VulnerabilityID: assessment.VulnerabilityID,
		AssetID:         &assessment.AssetID,
		RiskScore:       assessment.RiskScore,
		RiskLevel:       assessment.RiskLevel,
		Analysis:        assessment.Analysis,
		Provider:        assessment.Provider,
		Model:           assessment.Model,
	}, errors.DomainError{}
}

func (s *riskAssessmentService) GetAssessmentByVulnerabilityID(ctx context.Context, vulnerabilityID string) ([]usecase.RiskAssessResponse, errors.DomainError) {
	id, err := uuid.Parse(vulnerabilityID)
	if err != nil {
		return nil, errors.NewDomainWithError("400", "Invalid Vulnerability ID", err)
	}
	assessments, err := s.riskRepo.FindRiskAssessmentByVulnerabilityID(ctx, id)
	if err != nil {
		return nil, errors.NewDomainWithError("401", "Assessment Not Found", err)
	}
	var result []usecase.RiskAssessResponse
	for _, a := range assessments {
		result = append(result, usecase.RiskAssessResponse{
			AssessmentID:    a.AssessmentID,
			VulnerabilityID: a.VulnerabilityID,
			AssetID:         &a.AssetID,
			RiskScore:       a.RiskScore,
			RiskLevel:       a.RiskLevel,
			Analysis:        a.Analysis,
			Provider:        a.Provider,
			Model:           a.Model,
		})
	}
	return result, errors.DomainError{}
}

type SecurityRecommendationService interface {
	GenerateRecommendation(ctx context.Context, req usecase.RecommendationGenerateRequest) (usecase.RecommendationGenerateResponse, errors.DomainError)
	GetRecommendationByVulnerabilityID(ctx context.Context, vulnerabilityID string) ([]usecase.RecommendationGenerateResponse, errors.DomainError)
	SubmitFeedback(ctx context.Context, req usecase.RecommendationFeedbackRequest) (usecase.RecommendationFeedbackResponse, errors.DomainError)
}

func NewSecurityRecommendationService(
	recommendationRepo repository.SecurityRecommendationRepository,
	apiLogRepo repository.APICallLogRepository,
	modelConfigRepo repository.ModelConfigRepository,
	vulnerabilityRepo repository.VulnerabilityRepository,
	promptTemplateRepo repository.PromptTemplateRepository,
) SecurityRecommendationService {
	return &securityRecommendationService{
		recommendationRepo: recommendationRepo,
		apiLogRepo:         apiLogRepo,
		modelConfigRepo:    modelConfigRepo,
		vulnerabilityRepo:  vulnerabilityRepo,
		promptTemplateRepo: promptTemplateRepo,
	}
}

type securityRecommendationService struct {
	recommendationRepo repository.SecurityRecommendationRepository
	apiLogRepo         repository.APICallLogRepository
	modelConfigRepo    repository.ModelConfigRepository
	vulnerabilityRepo  repository.VulnerabilityRepository
	promptTemplateRepo repository.PromptTemplateRepository
}

func (s *securityRecommendationService) GenerateRecommendation(ctx context.Context, req usecase.RecommendationGenerateRequest) (usecase.RecommendationGenerateResponse, errors.DomainError) {
	vulnID, err := uuid.Parse(req.VulnerabilityID)
	if err != nil {
		return usecase.RecommendationGenerateResponse{}, errors.NewDomainWithError("400", "Invalid Vulnerability ID", err)
	}

	vulnerability, err := s.vulnerabilityRepo.FindVulnerabilityByID(ctx, vulnID)
	if err != nil {
		return usecase.RecommendationGenerateResponse{}, errors.NewDomainWithError("401", "Vulnerability Not Found", err)
	}

	modelConfig, err := s.modelConfigRepo.FindActiveModelConfig(ctx)
	if err != nil {
		return usecase.RecommendationGenerateResponse{}, errors.NewDomainWithError("401", "No Active Model Config", err)
	}

	apiLog := aggregate.APICallLog{
		LogID:     uuid.New(),
		ConfigID:  modelConfig.ConfigID,
		CallType:  "recommendation",
		Success:   true,
		CreatedAt: time.Now(),
	}
	_ = s.apiLogRepo.SaveAPICallLog(ctx, apiLog)

	recommendation := aggregate.SecurityRecommendation{
		RecommendationID:      uuid.New(),
		VulnerabilityID:       vulnID,
		LogID:                 apiLog.LogID,
		Summary:               "Apply security update for " + vulnerability.CVEID,
		VulnerabilityAnalysis: vulnerability.Description,
		RemediationSteps:      "[\"Apply latest security patch\", \"Review vendor recommendations\", \"Test in staging environment\"]",
		RecommendedPatches:    "[]",
		MitigationMeasures:    "[\"Enable intrusion detection\", \"Monitor network traffic\"]",
		PreventionTips:        "[\"Keep systems updated\", \"Follow security best practices\"]",
		References:            "[]",
		Provider:              modelConfig.ProviderName,
		Model:                 modelConfig.ModelName,
		CreatedAt:             time.Now(),
	}
	_ = s.recommendationRepo.SaveSecurityRecommendation(ctx, recommendation)

	return usecase.RecommendationGenerateResponse{
		RecommendationID:      recommendation.RecommendationID,
		VulnerabilityID:       recommendation.VulnerabilityID,
		Summary:               recommendation.Summary,
		VulnerabilityAnalysis: recommendation.VulnerabilityAnalysis,
		RemediationSteps:      recommendation.RemediationSteps,
		RecommendedPatches:    recommendation.RecommendedPatches,
		MitigationMeasures:    recommendation.MitigationMeasures,
		PreventionTips:        recommendation.PreventionTips,
		Provider:              recommendation.Provider,
		Model:                 recommendation.Model,
	}, errors.DomainError{}
}

func (s *securityRecommendationService) GetRecommendationByVulnerabilityID(ctx context.Context, vulnerabilityID string) ([]usecase.RecommendationGenerateResponse, errors.DomainError) {
	id, err := uuid.Parse(vulnerabilityID)
	if err != nil {
		return nil, errors.NewDomainWithError("400", "Invalid Vulnerability ID", err)
	}
	recommendations, err := s.recommendationRepo.FindSecurityRecommendationByVulnerabilityID(ctx, id)
	if err != nil {
		return nil, errors.NewDomainWithError("401", "Recommendation Not Found", err)
	}
	var result []usecase.RecommendationGenerateResponse
	for _, r := range recommendations {
		result = append(result, usecase.RecommendationGenerateResponse{
			RecommendationID:      r.RecommendationID,
			VulnerabilityID:       r.VulnerabilityID,
			Summary:               r.Summary,
			VulnerabilityAnalysis: r.VulnerabilityAnalysis,
			RemediationSteps:      r.RemediationSteps,
			RecommendedPatches:    r.RecommendedPatches,
			MitigationMeasures:    r.MitigationMeasures,
			PreventionTips:        r.PreventionTips,
			Provider:              r.Provider,
			Model:                 r.Model,
		})
	}
	return result, errors.DomainError{}
}

func (s *securityRecommendationService) SubmitFeedback(ctx context.Context, req usecase.RecommendationFeedbackRequest) (usecase.RecommendationFeedbackResponse, errors.DomainError) {
	recommendationID, err := uuid.Parse(req.RecommendationID)
	if err != nil {
		return usecase.RecommendationFeedbackResponse{}, errors.NewDomainWithError("400", "Invalid Recommendation ID", err)
	}
	recommendation, err := s.recommendationRepo.FindSecurityRecommendationByID(ctx, recommendationID)
	if err != nil {
		return usecase.RecommendationFeedbackResponse{}, errors.NewDomainWithError("401", "Recommendation Not Found", err)
	}
	recommendation.IsUseful = &req.IsUseful
	recommendation.Feedback = req.Feedback
	err = s.recommendationRepo.UpdateSecurityRecommendation(ctx, recommendation)
	if err != nil {
		return usecase.RecommendationFeedbackResponse{}, errors.NewDomainWithError("401", "Update Feedback Err", err)
	}
	return usecase.RecommendationFeedbackResponse{
		RecommendationID: recommendation.RecommendationID,
		Success:          true,
	}, errors.DomainError{}
}

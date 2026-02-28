package service_test

import (
	"moon/internal/domain/service"
	"moon/internal/domain/usecase"
	"moon/internal/infrastructure/persistence/postgres"
	"testing"

	"github.com/google/uuid"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	modelConfigSvc    service.ModelConfigService
	promptTemplateSvc service.PromptTemplateService
	apiCallLogSvc     service.APICallLogService
)

func init() {
	dsn := "host=localhost user=capstone password=capstone dbname=capstone port=5432 sslmode=disable"
	db, err := gorm.Open(driver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	modelConfigRepo := postgres.NewModelConfigRepository(db)
	promptTemplateRepo := postgres.NewPromptTemplateRepository(db)
	apiCallLogRepo := postgres.NewAPICallLogRepository(db)

	modelConfigSvc = service.NewModelConfigService(modelConfigRepo)
	promptTemplateSvc = service.NewPromptTemplateService(promptTemplateRepo)
	apiCallLogSvc = service.NewAPICallLogService(apiCallLogRepo)
}

func TestCreateModelConfig(t *testing.T) {
	req := usecase.ModelConfigCreateRequest{
		ProviderName:   "openai",
		ModelName:      "gpt-4",
		APIKey:         "test-api-key",
		APIEndpoint:    "",
		APIVersion:     "v1",
		MaxTokens:      4096,
		Temperature:    0.7,
		TimeoutSeconds: 30,
		IsActive:       true,
		Priority:       1,
	}

	resp, err := modelConfigSvc.CreateModelConfig(testCtx, req)
	if err.Code != "" {
		t.Errorf("Expected no error, got: %s", err.Message)
	}

	if resp.ConfigID == uuid.Nil {
		t.Errorf("Expected non-nil ConfigID, got %v", resp.ConfigID)
	}
	if resp.ProviderName != req.ProviderName {
		t.Errorf("Expected ProviderName %s, got %s", req.ProviderName, resp.ProviderName)
	}
}

func TestGetModelConfig(t *testing.T) {
	createReq := usecase.ModelConfigCreateRequest{
		ProviderName:   "anthropic",
		ModelName:      "claude-3",
		APIKey:         "test-key",
		MaxTokens:      4096,
		Temperature:    0.7,
		TimeoutSeconds: 30,
		IsActive:       true,
		Priority:       1,
	}
	createResp, err := modelConfigSvc.CreateModelConfig(testCtx, createReq)
	if err.Code != "" {
		t.Errorf("Expected no error when creating, got: %s", err.Message)
	}

	getReq := usecase.ModelConfigGetRequest{
		ConfigID: createResp.ConfigID.String(),
	}
	getResp, err := modelConfigSvc.GetModelConfig(testCtx, getReq)
	if err.Code != "" {
		t.Errorf("Expected no error when getting, got: %s", err.Message)
	}

	if getResp.ConfigID != createResp.ConfigID {
		t.Errorf("Expected ConfigID %s, got %s", createResp.ConfigID, getResp.ConfigID)
	}
}

func TestUpdateModelConfig(t *testing.T) {
	createReq := usecase.ModelConfigCreateRequest{
		ProviderName:   "qwen",
		ModelName:      "qwen-max",
		APIKey:         "original-key",
		MaxTokens:      4096,
		Temperature:    0.7,
		TimeoutSeconds: 30,
		IsActive:       true,
		Priority:       1,
	}
	createResp, _ := modelConfigSvc.CreateModelConfig(testCtx, createReq)

	updateReq := usecase.ModelConfigUpdateRequest{
		ConfigID: createResp.ConfigID.String(),
		APIKey:   "updated-key",
		IsActive: false,
	}
	updateResp, err := modelConfigSvc.UpdateModelConfig(testCtx, updateReq)
	if err.Code != "" {
		t.Errorf("Expected no error, got: %s", err.Message)
	}

	if !updateResp.Success {
		t.Errorf("Expected Success to be true")
	}
}

func TestDeleteModelConfig(t *testing.T) {
	createReq := usecase.ModelConfigCreateRequest{
		ProviderName:   "ernie",
		ModelName:      "ernie-bot",
		APIKey:         "delete-test-key",
		MaxTokens:      4096,
		Temperature:    0.7,
		TimeoutSeconds: 30,
		IsActive:       true,
		Priority:       1,
	}
	createResp, _ := modelConfigSvc.CreateModelConfig(testCtx, createReq)

	deleteReq := usecase.ModelConfigDeleteRequest{
		ConfigID: createResp.ConfigID.String(),
	}
	deleteResp, err := modelConfigSvc.DeleteModelConfig(testCtx, deleteReq)
	if err.Code != "" {
		t.Errorf("Expected no error, got: %s", err.Message)
	}

	if !deleteResp.Success {
		t.Errorf("Expected Success to be true")
	}
}

func TestListModelConfigs(t *testing.T) {
	_, _ = modelConfigSvc.CreateModelConfig(testCtx, usecase.ModelConfigCreateRequest{
		ProviderName:   "openai",
		ModelName:      "gpt-4",
		APIKey:         "key1",
		MaxTokens:      4096,
		Temperature:    0.7,
		TimeoutSeconds: 30,
		IsActive:       true,
		Priority:       1,
	})

	listResp, err := modelConfigSvc.ListAllModelConfigs(testCtx, usecase.ModelConfigListRequest{})
	if err.Code != "" {
		t.Errorf("Expected no error, got: %s", err.Message)
	}

	if len(listResp) == 0 {
		t.Errorf("Expected at least one config")
	}
}

func TestCreatePromptTemplate(t *testing.T) {
	req := usecase.PromptTemplateCreateRequest{
		TemplateName:    "test_template_" + uuid.New().String()[:8],
		TemplateType:    "asset_classification",
		TemplateContent: "Classify the asset: {{.AssetName}}",
		Variables:       "{\"AssetName\": \"string\"}",
		Description:     "Test template",
		IsActive:        true,
	}

	resp, err := promptTemplateSvc.CreatePromptTemplate(testCtx, req)
	if err.Code != "" {
		t.Errorf("Expected no error, got: %s", err.Message)
	}

	if resp.TemplateID == uuid.Nil {
		t.Errorf("Expected non-nil TemplateID, got %v", resp.TemplateID)
	}
}

func TestGetPromptTemplate(t *testing.T) {
	createReq := usecase.PromptTemplateCreateRequest{
		TemplateName:    "test_get_template_" + uuid.New().String()[:8],
		TemplateType:    "risk_assessment",
		TemplateContent: "Assess risk for: {{.VulnerabilityID}}",
		IsActive:        true,
	}
	createResp, _ := promptTemplateSvc.CreatePromptTemplate(testCtx, createReq)

	getReq := usecase.PromptTemplateGetRequest{
		TemplateID: createResp.TemplateID.String(),
	}
	getResp, err := promptTemplateSvc.GetPromptTemplate(testCtx, getReq)
	if err.Code != "" {
		t.Errorf("Expected no error, got: %s", err.Message)
	}

	if getResp.TemplateID != createResp.TemplateID {
		t.Errorf("Expected TemplateID %s, got %s", createResp.TemplateID, getResp.TemplateID)
	}
}

func TestListPromptTemplates(t *testing.T) {
	listResp, err := promptTemplateSvc.ListAllPromptTemplates(testCtx, usecase.PromptTemplateListRequest{})
	if err.Code != "" {
		t.Errorf("Expected no error, got: %s", err.Message)
	}

	if listResp == nil {
		t.Errorf("Expected non-nil list")
	}
}

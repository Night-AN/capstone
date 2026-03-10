package http

import (
	"moon/internal/domain/service"
	"moon/internal/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AIHandler struct {
	modelConfigService         service.ModelConfigService
	promptTemplateService      service.PromptTemplateService
	assetClassificationService service.AssetClassificationService
	riskAssessmentService      service.RiskAssessmentService
	recommendationService      service.SecurityRecommendationService
	apiCallLogService          service.APICallLogService
	chatService                service.ChatService
}

func NewAIHandler(
	modelConfigService service.ModelConfigService,
	promptTemplateService service.PromptTemplateService,
	assetClassificationService service.AssetClassificationService,
	riskAssessmentService service.RiskAssessmentService,
	recommendationService service.SecurityRecommendationService,
	apiCallLogService service.APICallLogService,
	chatService service.ChatService,
) *AIHandler {
	return &AIHandler{
		modelConfigService:         modelConfigService,
		promptTemplateService:      promptTemplateService,
		assetClassificationService: assetClassificationService,
		riskAssessmentService:      riskAssessmentService,
		recommendationService:      recommendationService,
		apiCallLogService:          apiCallLogService,
		chatService:                chatService,
	}
}

func (h *AIHandler) CreateModelConfig(c *gin.Context) {
	var req usecase.ModelConfigCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "invalid request: " + err.Error()})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.modelConfigService.CreateModelConfig(ctx, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"code": "201", "message": "success", "data": resp})
}

func (h *AIHandler) GetModelConfig(c *gin.Context) {
	configID := c.Query("config_id")
	if configID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "config_id is required"})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.modelConfigService.GetModelConfig(ctx, usecase.ModelConfigGetRequest{ConfigID: configID})
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) GetActiveModelConfig(c *gin.Context) {
	ctx := c.Request.Context()
	resp, domainErr := h.modelConfigService.GetActiveModelConfig(ctx)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) UpdateModelConfig(c *gin.Context) {
	var req usecase.ModelConfigUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "invalid request: " + err.Error()})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.modelConfigService.UpdateModelConfig(ctx, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) DeleteModelConfig(c *gin.Context) {
	var req usecase.ModelConfigDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "invalid request: " + err.Error()})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.modelConfigService.DeleteModelConfig(ctx, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) ListModelConfigs(c *gin.Context) {
	ctx := c.Request.Context()
	resp, domainErr := h.modelConfigService.ListAllModelConfigs(ctx, usecase.ModelConfigListRequest{})
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) CreatePromptTemplate(c *gin.Context) {
	var req usecase.PromptTemplateCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "invalid request: " + err.Error()})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.promptTemplateService.CreatePromptTemplate(ctx, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"code": "201", "message": "success", "data": resp})
}

func (h *AIHandler) GetPromptTemplate(c *gin.Context) {
	templateID := c.Query("template_id")
	if templateID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "template_id is required"})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.promptTemplateService.GetPromptTemplate(ctx, usecase.PromptTemplateGetRequest{TemplateID: templateID})
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) GetPromptTemplateByType(c *gin.Context) {
	templateType := c.Query("template_type")
	if templateType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "template_type is required"})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.promptTemplateService.GetPromptTemplateByType(ctx, templateType)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) UpdatePromptTemplate(c *gin.Context) {
	var req usecase.PromptTemplateUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "invalid request: " + err.Error()})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.promptTemplateService.UpdatePromptTemplate(ctx, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) DeletePromptTemplate(c *gin.Context) {
	var req usecase.PromptTemplateDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "invalid request: " + err.Error()})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.promptTemplateService.DeletePromptTemplate(ctx, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) ListPromptTemplates(c *gin.Context) {
	ctx := c.Request.Context()
	resp, domainErr := h.promptTemplateService.ListAllPromptTemplates(ctx, usecase.PromptTemplateListRequest{})
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) ClassifyAsset(c *gin.Context) {
	var req usecase.AssetClassifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "invalid request: " + err.Error()})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.assetClassificationService.ClassifyAsset(ctx, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) GetClassificationByAssetID(c *gin.Context) {
	assetID := c.Query("asset_id")
	if assetID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "asset_id is required"})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.assetClassificationService.GetClassificationByAssetID(ctx, assetID)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) ApproveClassification(c *gin.Context) {
	var req usecase.AssetClassificationApproveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "invalid request: " + err.Error()})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.assetClassificationService.ApproveClassification(ctx, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) AssessRisk(c *gin.Context) {
	var req usecase.RiskAssessRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "invalid request: " + err.Error()})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.riskAssessmentService.AssessRisk(ctx, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) GetAssessmentByVulnerabilityID(c *gin.Context) {
	vulnerabilityID := c.Query("vulnerability_id")
	if vulnerabilityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "vulnerability_id is required"})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.riskAssessmentService.GetAssessmentByVulnerabilityID(ctx, vulnerabilityID)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) GenerateRecommendation(c *gin.Context) {
	var req usecase.RecommendationGenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "invalid request: " + err.Error()})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.recommendationService.GenerateRecommendation(ctx, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) GetRecommendationByVulnerabilityID(c *gin.Context) {
	vulnerabilityID := c.Query("vulnerability_id")
	if vulnerabilityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "vulnerability_id is required"})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.recommendationService.GetRecommendationByVulnerabilityID(ctx, vulnerabilityID)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) SubmitFeedback(c *gin.Context) {
	var req usecase.RecommendationFeedbackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "invalid request: " + err.Error()})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.recommendationService.SubmitFeedback(ctx, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) ListAPICallLogs(c *gin.Context) {
	limit := 10
	offset := 0
	callType := c.Query("call_type")
	if l := c.Query("limit"); l != "" {
		_, _ = uuid.Parse(l)
	}
	ctx := c.Request.Context()
	resp, domainErr := h.apiCallLogService.ListAPICallLogs(ctx, usecase.APICallLogListRequest{Limit: limit, Offset: offset, CallType: callType})
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) GetAPICallLog(c *gin.Context) {
	logID := c.Param("id")
	if logID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "log_id is required"})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.apiCallLogService.GetAPICallLog(ctx, usecase.APICallLogGetRequest{LogID: logID})
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) Chat(c *gin.Context) {
	var req usecase.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "invalid request: " + err.Error()})
		return
	}
	if req.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "message is required"})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.chatService.Chat(ctx, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) GetConversation(c *gin.Context) {
	conversationID := c.Query("conversation_id")
	if conversationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "conversation_id is required"})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.chatService.GetConversation(ctx, usecase.GetConversationRequest{ConversationID: conversationID})
	if domainErr.Code != "" {
		c.JSON(http.StatusNotFound, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) ListConversations(c *gin.Context) {
	ctx := c.Request.Context()
	resp, domainErr := h.chatService.ListConversations(ctx, usecase.ConversationListRequest{
		Limit:  50,
		Offset: 0,
	})
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

func (h *AIHandler) DeleteConversation(c *gin.Context) {
	var req usecase.DeleteConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "invalid request: " + err.Error()})
		return
	}
	if req.ConversationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "400", "message": "conversation_id is required"})
		return
	}
	ctx := c.Request.Context()
	resp, domainErr := h.chatService.DeleteConversation(ctx, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{"code": domainErr.Code, "message": domainErr.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "message": "success", "data": resp})
}

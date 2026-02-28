package usecase

import (
	"moon/internal/domain/aggregate"
	"time"

	"github.com/google/uuid"
)

type PromptTemplateCreateRequest struct {
	TemplateName   string `json:"template_name"`
	TemplateType   string `json:"template_type"`
	TemplateContent string `json:"template_content"`
	Variables      string `json:"variables"`
	Description    string `json:"description"`
	IsActive       bool   `json:"is_active"`
}

type PromptTemplateCreateResponse struct {
	TemplateID   uuid.UUID `json:"template_id"`
	TemplateName string    `json:"template_name"`
	TemplateType string    `json:"template_type"`
}

func ConvertPromptTemplateCreateRequestToAggregate(req PromptTemplateCreateRequest) aggregate.PromptTemplate {
	return aggregate.PromptTemplate{
		TemplateID:      uuid.New(),
		TemplateName:   req.TemplateName,
		TemplateType:   req.TemplateType,
		TemplateContent: req.TemplateContent,
		Variables:      req.Variables,
		Description:    req.Description,
		IsActive:       req.IsActive,
		Version:        1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

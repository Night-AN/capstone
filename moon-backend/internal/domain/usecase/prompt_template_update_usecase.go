package usecase

import (
	"github.com/google/uuid"
)

type PromptTemplateUpdateRequest struct {
	TemplateID     string `json:"template_id"`
	TemplateName   string `json:"template_name"`
	TemplateType   string `json:"template_type"`
	TemplateContent string `json:"template_content"`
	Variables      string `json:"variables"`
	Description    string `json:"description"`
	IsActive       bool   `json:"is_active"`
}

type PromptTemplateUpdateResponse struct {
	TemplateID   uuid.UUID `json:"template_id"`
	TemplateName string    `json:"template_name"`
	Success      bool      `json:"success"`
}

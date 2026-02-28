package usecase

import (
	"github.com/google/uuid"
)

type PromptTemplateGetRequest struct {
	TemplateID string `json:"template_id"`
}

type PromptTemplateGetResponse struct {
	TemplateID      uuid.UUID `json:"template_id"`
	TemplateName    string    `json:"template_name"`
	TemplateType    string    `json:"template_type"`
	TemplateContent string    `json:"template_content"`
	Variables       string    `json:"variables"`
	Description     string    `json:"description"`
	IsActive        bool      `json:"is_active"`
	Version         int       `json:"version"`
	CreatedAt       string    `json:"created_at"`
	UpdatedAt       string    `json:"updated_at"`
}

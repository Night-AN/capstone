package usecase

type PromptTemplateDeleteRequest struct {
	TemplateID string `json:"template_id"`
}

type PromptTemplateDeleteResponse struct {
	Success bool `json:"success"`
}

package usecase

type PromptTemplateListRequest struct {
}

type PromptTemplateListResponse struct {
	TemplateID    string `json:"template_id"`
	TemplateName  string `json:"template_name"`
	TemplateType  string `json:"template_type"`
	Description   string `json:"description"`
	IsActive      bool   `json:"is_active"`
	Version       int    `json:"version"`
}

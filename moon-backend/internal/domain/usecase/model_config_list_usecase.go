package usecase

type ModelConfigListRequest struct {
}

type ModelConfigListResponse struct {
	ConfigID        string `json:"config_id"`
	ProviderName    string `json:"provider_name"`
	ModelName       string `json:"model_name"`
	APIEndpoint     string `json:"api_endpoint"`
	MaxTokens       int    `json:"max_tokens"`
	Temperature     float64 `json:"temperature"`
	IsActive        bool    `json:"is_active"`
	Priority        int     `json:"priority"`
}

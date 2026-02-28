package usecase

type APICallLogListRequest struct {
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	CallType string `json:"call_type"`
}

type APICallLogListResponse struct {
	LogID            string `json:"log_id"`
	ConfigID         string `json:"config_id"`
	CallType         string `json:"call_type"`
	PromptTokens     int    `json:"prompt_tokens"`
	CompletionTokens int    `json:"completion_tokens"`
	TotalTokens      int    `json:"total_tokens"`
	StatusCode       int    `json:"status_code"`
	ErrorMessage     string `json:"error_message"`
	LatencyMs        int    `json:"latency_ms"`
	Success          bool   `json:"success"`
	CreatedAt        string `json:"created_at"`
}

type APICallLogGetRequest struct {
	LogID string `json:"log_id"`
}

type APICallLogGetResponse struct {
	LogID            string `json:"log_id"`
	ConfigID         string `json:"config_id"`
	CallType         string `json:"call_type"`
	PromptTokens     int    `json:"prompt_tokens"`
	CompletionTokens int    `json:"completion_tokens"`
	TotalTokens      int    `json:"total_tokens"`
	RequestPayload   string `json:"request_payload"`
	ResponsePayload  string `json:"response_payload"`
	StatusCode       int    `json:"status_code"`
	ErrorMessage     string `json:"error_message"`
	LatencyMs        int    `json:"latency_ms"`
	Success          bool   `json:"success"`
	CreatedAt        string `json:"created_at"`
}

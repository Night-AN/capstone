package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type APICallLog struct {
	LogID             uuid.UUID `gorm:"column:log_id"`
	ConfigID          uuid.UUID `gorm:"column:config_id"`
	CallType          string    `gorm:"column:call_type"`
	PromptTokens      int       `gorm:"column:prompt_tokens"`
	CompletionTokens  int       `gorm:"column:completion_tokens"`
	TotalTokens       int       `gorm:"column:total_tokens"`
	RequestPayload    string    `gorm:"column:request_payload"`
	ResponsePayload   string    `gorm:"column:response_payload"`
	StatusCode        int       `gorm:"column:status_code"`
	ErrorMessage      string    `gorm:"column:error_message"`
	LatencyMs         int       `gorm:"column:latency_ms"`
	Success           bool      `gorm:"column:success"`
	CreatedAt         time.Time `gorm:"column:created_at"`
}

func (APICallLog) TableName() string {
	return "ai.api_call_log"
}

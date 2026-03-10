package usecase

import (
	"time"

	"github.com/google/uuid"
)

type ChatRequest struct {
	Message           string  `json:"message"`
	ModelConfigID     *string `json:"model_config_id,omitempty"`
	PromptTemplateID  *string `json:"prompt_template_id,omitempty"`
	ConversationID    *string `json:"conversation_id,omitempty"`
}

type ChatResponse struct {
	ConversationID uuid.UUID       `json:"conversation_id"`
	Message        string          `json:"message"`
	Response       string          `json:"response"`
	ModelUsed      string          `json:"model_used"`
	TokensUsed     int             `json:"tokens_used"`
	CreatedAt      time.Time       `json:"created_at"`
	Messages       []ChatMessage   `json:"messages,omitempty"`
}

type ChatMessage struct {
	MessageID      uuid.UUID `json:"message_id"`
	ConversationID uuid.UUID `json:"conversation_id"`
	Role           string    `json:"role"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}

type ConversationListRequest struct {
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
	ModelID  string `json:"model_id,omitempty"`
}

type ConversationListResponse struct {
	Conversations []ConversationSummary `json:"conversations"`
	Total         int                    `json:"total"`
}

type ConversationSummary struct {
	ConversationID uuid.UUID `json:"conversation_id"`
	Title          string    `json:"title"`
	ModelUsed      string    `json:"model_used"`
	MessageCount   int       `json:"message_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type GetConversationRequest struct {
	ConversationID string `json:"conversation_id"`
}

type DeleteConversationRequest struct {
	ConversationID string `json:"conversation_id"`
}

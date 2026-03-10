package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type ChatMessage struct {
	MessageID      uuid.UUID `gorm:"column:message_id"`
	ConversationID uuid.UUID `gorm:"column:conversation_id"`
	Role           string    `gorm:"column:role"`
	Content        string    `gorm:"column:content"`
	CreatedAt      time.Time `gorm:"column:created_at"`
}

func (ChatMessage) TableName() string {
	return "ai.chat_message"
}

type Conversation struct {
	ConversationID uuid.UUID `gorm:"column:conversation_id"`
	Title          string    `gorm:"column:title"`
	ModelUsed      string    `gorm:"column:model_used"`
	MessageCount   int       `gorm:"column:message_count"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (Conversation) TableName() string {
	return "ai.conversation"
}

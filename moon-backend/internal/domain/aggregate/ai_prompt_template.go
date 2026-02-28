package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type PromptTemplate struct {
	TemplateID     uuid.UUID `gorm:"column:template_id"`
	TemplateName   string    `gorm:"column:template_name"`
	TemplateType   string    `gorm:"column:template_type"`
	TemplateContent string   `gorm:"column:template_content"`
	Variables      string    `gorm:"column:variables"`
	Description    string    `gorm:"column:description"`
	IsActive       bool      `gorm:"column:is_active"`
	Version        int       `gorm:"column:version"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (PromptTemplate) TableName() string {
	return "ai.prompt_template"
}

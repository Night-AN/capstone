package aggregate

import (
	"time"

	"github.com/google/uuid"
)

// Organization represents an organizational entity within the system.
type Organization struct {

	// OrganizationID is the unique identifier for the organization, using UUID to ensure global uniqueness.
	OrganizationID uuid.UUID `gorm:"column:organization_id;type:uuid"`

	// OrganizationName is the display name of the organization, used for UI presentation and logging.
	OrganizationName string `gorm:"column:organization_name;type:text"`

	// OrganizationCode is the unique code for the organization, used for identification and path construction.
	// Constraint: Must be a URL/path-safe string; using lowercase letters, numbers, and hyphens.
	OrganizationCode string `gorm:"column:organization_code;type:text"`

	// OrganizationDescription provides additional details about the organization's purpose and scope.
	OrganizationDescription string `gorm:"column:organization_description;type:text"`

	// OrganizationFlag is a generic flag field for extensible organization attributes or state indicators.
	// Usage depends on business requirements (e.g., type classification, status markers).
	OrganizationFlag string `gorm:"column:organization_flag;type:text"`

	// SensitiveFlag indicates whether this organization contains sensitive data or requires special handling.
	// When true, additional access controls and audit logging may be applied.
	SensitiveFlag bool `gorm:"column:sensitive_flag;type:boolean"`

	// CreatedAt records the creation timestamp for auditing, sorting, and statistics.
	CreatedAt time.Time `gorm:"column:created_at;type:timestamptz"`

	// UpdatedAt records the last modification timestamp for optimistic locking and change tracking.
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamptz"`
}

func (Organization) TableName() string {
	return "systems.organization"
}

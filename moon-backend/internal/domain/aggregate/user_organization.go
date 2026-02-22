package aggregate

import (
	"time"

	"github.com/google/uuid"
)

// UserOrganization represents the relationship between users and organizations.
// This aggregate entity manages the association between users and the organizations they belong to.
type UserOrganization struct {
	// ID is the unique identifier for the user-organization relationship.
	ID uuid.UUID `gorm:"column:id;type:uuid;primaryKey"`

	// UserID is the unique identifier for the user.
	UserID uuid.UUID `gorm:"column:user_id;type:uuid"`

	// OrganizationID is the unique identifier for the organization.
	OrganizationID uuid.UUID `gorm:"column:organization_id;type:uuid"`

	// CreatedAt records the timestamp when the user-organization relationship was created.
	CreatedAt time.Time `gorm:"column:created_at;type:timestamptz"`

	// UpdatedAt records the timestamp when the user-organization relationship was last modified.
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamptz"`
}

// TableName specifies the table name for the UserOrganization struct
func (UserOrganization) TableName() string {
	return "systems.user_organization"
}

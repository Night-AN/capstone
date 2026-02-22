package aggregate

import (
	"github.com/google/uuid"
)

// OrganizationRole represents the relationship between organizations and roles
// This aggregate entity manages the association between organizations and the roles available to them
type OrganizationRole struct {
	// ID is the unique identifier for the organization-role relationship
	ID uuid.UUID `gorm:"column:id;type:uuid;primaryKey"`

	// OrganizationID is the unique identifier for the organization
	OrganizationID uuid.UUID `gorm:"column:organization_id;type:uuid"`

	// RoleID is the unique identifier for the role
	RoleID uuid.UUID `gorm:"column:role_id;type:uuid"`
}

// TableName specifies the table name for the OrganizationRole struct
func (OrganizationRole) TableName() string {
	return "systems.organization_role"
}
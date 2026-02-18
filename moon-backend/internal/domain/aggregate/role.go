package aggregate

import (
	"time"

	"github.com/google/uuid"
)

// Role represents a collection of permissions grouped together for easy assignment to users.
// Roles provide a way to manage permissions at a higher level of abstraction than individual permissions.
// System roles are pre-defined and cannot be modified, while custom roles can be created and customized.
// Roles are assigned to users through user-role relationships.
// The RoleCode follows the format: {system|custom}:{name}, e.g., "system:admin", "custom:finance-manager".
type Role struct {
	// RoleID is the unique identifier for the role, using UUID to ensure global uniqueness
	// and prevent collision across distributed systems.
	RoleID uuid.UUID `gorm:"column:role_id;primaryKey"`

	// RoleName is the human-readable display name for UI, logging, and documentation.
	// Examples: "System Administrator", "Finance Manager", "Read-Only User"
	RoleName string `gorm:"column:role_name"`

	// Description provides detailed explanation of what this role is intended for and the permissions it grants.
	// Can be nil if no detailed description is needed.
	// Usage: Shown in role management UI, audit reports, and access review workflows
	Description *string `gorm:"column:role_description"`

	// RoleCode is the programmatic identifier used in code for role checks, following the format: {system|custom}:{name}.
	// Examples: "system:admin", "custom:finance-manager"
	// Constraints: Must be unique, URL-safe, lowercase recommended
	RoleCode string `gorm:"column:role_code"`

	// RoleFlag indicates the status or type of the role.
	// Examples: "active", "inactive", "system", "custom"
	RoleFlag string `gorm:"column:role_flag"`

	// SensitiveFlag marks roles that grant access to sensitive data or critical operations.
	// When true, additional audit logging and approval workflows may be triggered.
	// Examples: Roles accessing PII, financial data, or administrative actions should have this flag set to true
	// Usage: Drives enhanced logging, requires additional approval, and flags in compliance reports
	SensitiveFlag bool `gorm:"column:sensitive_flag"`

	// Permissions is the collection of permissions assigned to this role.
	// A role's permissions define what actions users with this role can perform.
	// Permissions are managed through role-permission relationships.
	// Permissions []Permission `gorm:"many2many:systems.permission_role;"`

	// CreatedAt records the timestamp when the role was first defined in the system.
	// Used for auditing, role lifecycle management, and compliance reporting.
	CreatedAt time.Time `gorm:"column:created_at"`

	// UpdatedAt records the timestamp when the role was last modified.
	// Used for auditing, role lifecycle management, and compliance reporting.
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName specifies the table name for the Role struct
func (Role) TableName() string {
	return "systems.role"
}

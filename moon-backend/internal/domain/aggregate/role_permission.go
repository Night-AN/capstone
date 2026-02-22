package aggregate

import (
	"github.com/google/uuid"
)

// RolePermission represents the relationship between roles and permissions
// This aggregate entity manages the association between roles and the permissions they grant
type RolePermission struct {
	// ID is the unique identifier for the role-permission relationship
	ID uuid.UUID `gorm:"column:id;type:uuid;primaryKey"`

	// RoleID is the unique identifier for the role
	RoleID uuid.UUID `gorm:"column:role_id;type:uuid"`

	// PermissionID is the unique identifier for the permission
	PermissionID uuid.UUID `gorm:"column:permission_id;type:uuid"`
}

// TableName specifies the table name for the RolePermission struct
func (RolePermission) TableName() string {
	return "systems.permission_role"
}
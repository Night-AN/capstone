package aggregate

import (
	"github.com/google/uuid"
)

// PermissionResource represents the relationship between permissions and resources
// This aggregate entity manages the association between permissions and the resources they protect
type PermissionResource struct {
	// ID is the unique identifier for the permission-resource relationship
	ID uuid.UUID `gorm:"column:id;type:uuid;primaryKey"`

	// PermissionID is the unique identifier for the permission
	PermissionID uuid.UUID `gorm:"column:permission_id;type:uuid"`

	// ResourceID is the unique identifier for the resource
	ResourceID uuid.UUID `gorm:"column:resource_id;type:uuid"`
}

// TableName specifies the table name for the PermissionResource struct
func (PermissionResource) TableName() string {
	return "systems.permission_resource"
}
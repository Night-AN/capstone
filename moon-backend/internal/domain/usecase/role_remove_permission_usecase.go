package usecase

import (
	"github.com/google/uuid"
)

// RoleRemovePermissionRequest represents the request for removing a permission from a role
// It contains the IDs of the role and permission
type RoleRemovePermissionRequest struct {
	// RoleID is the unique identifier of the role to remove the permission from
	RoleID uuid.UUID `json:"role_id"`

	// PermissionID is the unique identifier of the permission to remove
	PermissionID uuid.UUID `json:"permission_id"`
}

// RoleRemovePermissionResponse represents the response for removing a permission from a role
// It indicates whether the removal was successful
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields
type RoleRemovePermissionResponse struct {
	// Success indicates whether the removal was successful
	Success bool `json:"success"`
}
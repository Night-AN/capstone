package usecase

import (
	"github.com/google/uuid"
)

// RoleAssignPermissionRequest represents the request for assigning a permission to a role
// It contains the IDs of the role and permission
type RoleAssignPermissionRequest struct {
	// RoleID is the unique identifier of the role to assign the permission to
	RoleID uuid.UUID `json:"role_id"`

	// PermissionID is the unique identifier of the permission to assign
	PermissionID uuid.UUID `json:"permission_id"`
}

// RoleAssignPermissionResponse represents the response for assigning a permission to a role
// It indicates whether the assignment was successful
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields
type RoleAssignPermissionResponse struct {
	// Success indicates whether the assignment was successful
	Success bool `json:"success"`
}
package usecase

import (
	"github.com/google/uuid"
)

// RolePermissionsRequest represents the request for retrieving a role's permissions
// It contains the ID of the role
type RolePermissionsRequest struct {
	// RoleID is the unique identifier of the role to retrieve permissions for
	RoleID uuid.UUID `json:"role_id"`
}

// RolePermissionsResponse represents the response for retrieving a role's permissions
// It contains the role's ID and a list of its permissions
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields
type RolePermissionsResponse struct {
	// RoleID is the unique identifier of the role
	RoleID uuid.UUID `json:"role_id"`

	// Permissions is the list of permissions assigned to the role
	Permissions []PermissionResponse `json:"permissions"`
}
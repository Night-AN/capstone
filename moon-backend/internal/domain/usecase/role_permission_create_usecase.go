package usecase

import (
	"github.com/google/uuid"
)

// RolePermissionCreateRequest represents the request for creating a role-permission relationship
// It contains the IDs of the role and permission to associate
type RolePermissionCreateRequest struct {
	// RoleID is the unique identifier of the role
	RoleID uuid.UUID `json:"role_id"`

	// PermissionID is the unique identifier of the permission
	PermissionID uuid.UUID `json:"permission_id"`
}

// RolePermissionCreateResponse represents the response for creating a role-permission relationship
// It indicates whether the creation was successful

type RolePermissionCreateResponse struct {
	// Success indicates whether the creation was successful
	Success bool `json:"success"`
}

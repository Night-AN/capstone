package usecase

import (
	"github.com/google/uuid"
)

// RolePermissionDeleteRequest represents the request for deleting a role-permission relationship
// It contains the IDs of the role and permission to disassociate
type RolePermissionDeleteRequest struct {
	// RoleID is the unique identifier of the role
	RoleID uuid.UUID `json:"role_id"`

	// PermissionID is the unique identifier of the permission
	PermissionID uuid.UUID `json:"permission_id"`
}

// RolePermissionDeleteResponse represents the response for deleting a role-permission relationship
// It indicates whether the deletion was successful

type RolePermissionDeleteResponse struct {
	// Success indicates whether the deletion was successful
	Success bool `json:"success"`
}
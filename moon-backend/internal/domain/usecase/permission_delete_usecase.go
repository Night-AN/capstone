package usecase

import (
	"github.com/google/uuid"
)

// PermissionDeleteRequest represents the request for deleting a permission
// It contains the ID of the permission to delete
type PermissionDeleteRequest struct {
	// PermissionID is the unique identifier of the permission to delete
	PermissionID uuid.UUID `json:"permission_id"`
}

// PermissionDeleteResponse represents the response for deleting a permission
// It indicates whether the deletion was successful
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields
type PermissionDeleteResponse struct {
	// Success indicates whether the deletion was successful
	Success bool `json:"success"`
}
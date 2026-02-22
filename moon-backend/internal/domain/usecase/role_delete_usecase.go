package usecase

import (
	"github.com/google/uuid"
)

// RoleDeleteRequest represents the request for deleting a role
// It contains the ID of the role to delete
type RoleDeleteRequest struct {
	// RoleID is the unique identifier of the role to delete
	RoleID uuid.UUID `json:"role_id"`
}

// RoleDeleteResponse represents the response for deleting a role
// It indicates whether the deletion was successful
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields
type RoleDeleteResponse struct {
	// Success indicates whether the deletion was successful
	Success bool `json:"success"`
}
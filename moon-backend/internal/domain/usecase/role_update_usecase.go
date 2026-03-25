package usecase

import (
	"github.com/google/uuid"
)

// RoleUpdateRequest represents the request for updating a role
// It contains all the necessary information to update an existing role
// The RoleCode should follow the format: {system|custom}:{name}
// The RoleFlag indicates the status or type of the role
type RoleUpdateRequest struct {
	// RoleID is the unique identifier of the role to update
	RoleID uuid.UUID `json:"role_id"`

	// RoleName is the human-readable display name for UI, logging, and documentation
	RoleName string `json:"role_name"`

	// RoleCode is the programmatic identifier used in code for role checks
	RoleCode string `json:"role_code"`

	// RoleDescription provides detailed explanation of what this role is intended for and the permissions it grants
	RoleDescription *string `json:"role_description"`

	// RoleFlag indicates the status or type of the role
	RoleFlag string `json:"role_flag"`
}

// RoleUpdateResponse represents the response for updating a role
// It contains the basic information of the updated role
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields
type RoleUpdateResponse struct {
	Success bool `json:"success"`
}

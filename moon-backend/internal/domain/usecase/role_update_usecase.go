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

	// Description provides detailed explanation of what this role is intended for and the permissions it grants
	Description *string `json:"description"`

	// RoleCode is the programmatic identifier used in code for role checks
	RoleCode string `json:"role_code"`

	// RoleFlag indicates the status or type of the role
	RoleFlag string `json:"role_flag"`

	// SensitiveFlag marks roles that grant access to sensitive data or critical operations
	SensitiveFlag bool `json:"sensitive_flag"`
}

// RoleUpdateResponse represents the response for updating a role
// It contains the basic information of the updated role
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields
type RoleUpdateResponse struct {
	// RoleID is the unique identifier of the updated role
	RoleID uuid.UUID `json:"role_id"`

	// RoleName is the human-readable display name of the updated role
	RoleName string `json:"role_name"`

	// RoleCode is the programmatic identifier of the updated role
	RoleCode string `json:"role_code"`
}
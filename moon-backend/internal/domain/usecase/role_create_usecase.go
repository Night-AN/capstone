package usecase

import (
	"github.com/google/uuid"
)

// RoleCreateRequest represents the request for creating a role
// It contains all the necessary information to create a new role
// The RoleCode should follow the format: {system|custom}:{name}
// The RoleFlag indicates the status or type of the role
type RoleCreateRequest struct {
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

// RoleCreateResponse represents the response for creating a role
// It contains the basic information of the created role
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields
type RoleCreateResponse struct {
	// RoleID is the unique identifier of the created role
	RoleID uuid.UUID `json:"role_id"`

	// RoleName is the human-readable display name of the created role
	RoleName string `json:"role_name"`

	// RoleCode is the programmatic identifier of the created role
	RoleCode string `json:"role_code"`
}
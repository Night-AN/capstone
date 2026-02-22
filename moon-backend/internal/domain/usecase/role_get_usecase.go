package usecase

import (
	"time"

	"github.com/google/uuid"
)

// RoleGetRequest represents the request for retrieving a role
// It contains the ID of the role to retrieve
type RoleGetRequest struct {
	// RoleID is the unique identifier of the role to retrieve
	RoleID uuid.UUID `json:"role_id"`
}

// RoleGetResponse represents the response for retrieving a role
// It contains all the information of the retrieved role
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields
type RoleGetResponse struct {
	// RoleID is the unique identifier of the retrieved role
	RoleID uuid.UUID `json:"role_id"`

	// RoleName is the human-readable display name of the retrieved role
	RoleName string `json:"role_name"`

	// Description provides detailed explanation of what this role is intended for and the permissions it grants
	Description *string `json:"description"`

	// RoleCode is the programmatic identifier of the retrieved role
	RoleCode string `json:"role_code"`

	// RoleFlag indicates the status or type of the role
	RoleFlag string `json:"role_flag"`

	// SensitiveFlag marks roles that grant access to sensitive data or critical operations
	SensitiveFlag bool `json:"sensitive_flag"`

	// CreatedAt records the timestamp when the role was first defined in the system
	CreatedAt time.Time `json:"created_at"`
}
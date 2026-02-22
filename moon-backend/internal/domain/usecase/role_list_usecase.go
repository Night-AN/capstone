package usecase

import (
	"time"

	"github.com/google/uuid"
)

// RoleListRequest represents the request for listing roles
// It can be extended with pagination, filtering, and sorting parameters in the future
type RoleListRequest struct {
	// Empty for now, can be extended with pagination, filtering, and sorting parameters
}

// RoleListItem represents a single role in the list response
// It contains the basic information of a role
type RoleListItem struct {
	// RoleID is the unique identifier of the role
	RoleID uuid.UUID `json:"role_id"`

	// RoleName is the human-readable display name of the role
	RoleName string `json:"role_name"`

	// RoleCode is the programmatic identifier of the role
	RoleCode string `json:"role_code"`

	// RoleFlag indicates the status or type of the role
	RoleFlag string `json:"role_flag"`

	// SensitiveFlag marks roles that grant access to sensitive data or critical operations
	SensitiveFlag bool `json:"sensitive_flag"`

	// CreatedAt records the timestamp when the role was first defined in the system
	CreatedAt time.Time `json:"created_at"`
}

// RoleListResponse represents the response for listing roles
// It contains a list of roles
type RoleListResponse struct {
	// Roles is the list of roles
	Roles []RoleListItem `json:"roles"`
}
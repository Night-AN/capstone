package usecase

import (
	"time"

	"github.com/google/uuid"
)

// PermissionGetRequest represents the request for retrieving a permission
// It contains the ID of the permission to retrieve
type PermissionGetRequest struct {
	// PermissionID is the unique identifier of the permission to retrieve
	PermissionID uuid.UUID `json:"permission_id"`
}

// PermissionGetResponse represents the response for retrieving a permission
// It contains all the information of the retrieved permission
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields
type PermissionGetResponse struct {
	// PermissionID is the unique identifier of the retrieved permission
	PermissionID uuid.UUID `json:"permission_id"`

	// PermissionName is the human-readable display name of the retrieved permission
	PermissionName string `json:"permission_name"`

	// Description provides detailed explanation of what this permission grants and its security implications
	Description *string `json:"description"`

	// PermissionCode is the programmatic identifier of the retrieved permission
	PermissionCode string `json:"permission_code"`

	// SensitiveFlag marks permissions that grant access to sensitive data or critical operations
	SensitiveFlag bool `json:"sensitive_flag"`

	// CreatedAt records the timestamp when the permission was first defined in the system
	CreatedAt time.Time `json:"created_at"`
}
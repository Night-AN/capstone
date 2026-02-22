package usecase

import (
	"github.com/google/uuid"
)

// PermissionUpdateRequest represents the request for updating a permission
// It contains all the necessary information to update an existing permission
// The PermissionCode should follow the standard format: {resource}:{type}:{action}:{subaction?}:{scope?}
// The SensitiveFlag should be set to true for permissions that grant access to sensitive data or critical operations
type PermissionUpdateRequest struct {
	// PermissionID is the unique identifier of the permission to update
	PermissionID uuid.UUID `json:"permission_id"`

	// PermissionName is the human-readable display name for UI, logging, and documentation
	PermissionName string `json:"permission_name"`

	// Description provides detailed explanation of what this permission grants and its security implications
	Description *string `json:"description"`

	// PermissionCode is the programmatic identifier used in code for permission checks
	PermissionCode string `json:"permission_code"`

	// SensitiveFlag marks permissions that grant access to sensitive data or critical operations
	SensitiveFlag bool `json:"sensitive_flag"`
}

// PermissionUpdateResponse represents the response for updating a permission
// It contains the basic information of the updated permission
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields
type PermissionUpdateResponse struct {
	// PermissionID is the unique identifier of the updated permission
	PermissionID uuid.UUID `json:"permission_id"`

	// PermissionName is the human-readable display name of the updated permission
	PermissionName string `json:"permission_name"`

	// PermissionCode is the programmatic identifier of the updated permission
	PermissionCode string `json:"permission_code"`
}
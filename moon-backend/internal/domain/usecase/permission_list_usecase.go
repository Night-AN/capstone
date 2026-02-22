package usecase

import (
	"github.com/google/uuid"
)

// PermissionListRequest represents the request for listing permissions
// It contains pagination information for retrieving permissions
// The Limit field specifies the maximum number of permissions to retrieve
// The Offset field specifies the number of permissions to skip
type PermissionListRequest struct {
	// Limit specifies the maximum number of permissions to retrieve
	Limit int `json:"limit"`

	// Offset specifies the number of permissions to skip
	Offset int `json:"offset"`
}

// PermissionResponse represents the basic information of a permission
// It is used in lists of permissions
type PermissionResponse struct {
	// PermissionID is the unique identifier of the permission
	PermissionID uuid.UUID `json:"permission_id"`

	// PermissionName is the human-readable display name of the permission
	PermissionName string `json:"permission_name"`

	// PermissionCode is the programmatic identifier of the permission
	PermissionCode string `json:"permission_code"`

	// SensitiveFlag marks permissions that grant access to sensitive data or critical operations
	SensitiveFlag bool `json:"sensitive_flag"`
}

// PermissionListResponse represents the response for listing permissions
// It contains a list of permissions
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields
type PermissionListResponse struct {
	// Permissions is the list of permissions retrieved
	Permissions []PermissionResponse `json:"permissions"`
}
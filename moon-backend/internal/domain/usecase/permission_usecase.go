package usecase

import (
	"time"

	"github.com/google/uuid"
)

// PermissionCreateRequest represents the request for creating a permission
// It contains all the necessary information to create a new permission
// The PermissionCode should follow the standard format: {resource}:{type}:{action}:{subaction?}:{scope?}
// The SensitiveFlag should be set to true for permissions that grant access to sensitive data or critical operations
type PermissionCreateRequest struct {
	// PermissionName is the human-readable display name for UI, logging, and documentation
	PermissionName string `json:"permission_name"`

	// Description provides detailed explanation of what this permission grants and its security implications
	Description *string `json:"description"`

	// PermissionCode is the programmatic identifier used in code for permission checks
	PermissionCode string `json:"permission_code"`

	// SensitiveFlag marks permissions that grant access to sensitive data or critical operations
	SensitiveFlag bool `json:"sensitive_flag"`
}

// PermissionCreateResponse represents the response for creating a permission
// It contains the basic information of the created permission
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields
type PermissionCreateResponse struct {
	// PermissionID is the unique identifier of the created permission
	PermissionID uuid.UUID `json:"permission_id"`

	// PermissionName is the human-readable display name of the created permission
	PermissionName string `json:"permission_name"`

	// PermissionCode is the programmatic identifier of the created permission
	PermissionCode string `json:"permission_code"`
}

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

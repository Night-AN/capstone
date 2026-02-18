package usecase

import (
	"time"

	"github.com/google/uuid"
)

// RoleCreateRequest represents the request for creating a role
// It contains all the necessary information to create a new role
// The RoleCode should follow the format: {system|custom}:{name}
// The RoleFlag indicates the status or type of the role

type RoleCreateRequest struct {
	// RoleName is the human-readable display name for UI, logging, and documentation
	RoleName string

	// Description provides detailed explanation of what this role is intended for and the permissions it grants
	Description *string

	// RoleCode is the programmatic identifier used in code for role checks
	RoleCode string

	// RoleFlag indicates the status or type of the role
	RoleFlag string

	// SensitiveFlag marks roles that grant access to sensitive data or critical operations
	SensitiveFlag bool
}

// RoleCreateResponse represents the response for creating a role
// It contains the basic information of the created role
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields

type RoleCreateResponse struct {
	// RoleID is the unique identifier of the created role
	RoleID uuid.UUID

	// RoleName is the human-readable display name of the created role
	RoleName string

	// RoleCode is the programmatic identifier of the created role
	RoleCode string
}

// RoleGetRequest represents the request for retrieving a role
// It contains the ID of the role to retrieve

type RoleGetRequest struct {
	// RoleID is the unique identifier of the role to retrieve
	RoleID uuid.UUID
}

// RoleGetResponse represents the response for retrieving a role
// It contains all the information of the retrieved role
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields

type RoleGetResponse struct {
	// RoleID is the unique identifier of the retrieved role
	RoleID uuid.UUID

	// RoleName is the human-readable display name of the retrieved role
	RoleName string

	// Description provides detailed explanation of what this role is intended for and the permissions it grants
	Description *string

	// RoleCode is the programmatic identifier of the retrieved role
	RoleCode string

	// RoleFlag indicates the status or type of the role
	RoleFlag string

	// SensitiveFlag marks roles that grant access to sensitive data or critical operations
	SensitiveFlag bool

	// CreatedAt records the timestamp when the role was first defined in the system
	CreatedAt time.Time
}

// RoleUpdateRequest represents the request for updating a role
// It contains all the necessary information to update an existing role
// The RoleCode should follow the format: {system|custom}:{name}
// The RoleFlag indicates the status or type of the role

type RoleUpdateRequest struct {
	// RoleID is the unique identifier of the role to update
	RoleID uuid.UUID

	// RoleName is the human-readable display name for UI, logging, and documentation
	RoleName string

	// Description provides detailed explanation of what this role is intended for and the permissions it grants
	Description *string

	// RoleCode is the programmatic identifier used in code for role checks
	RoleCode string

	// RoleFlag indicates the status or type of the role
	RoleFlag string

	// SensitiveFlag marks roles that grant access to sensitive data or critical operations
	SensitiveFlag bool
}

// RoleUpdateResponse represents the response for updating a role
// It contains the basic information of the updated role
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields

type RoleUpdateResponse struct {
	// RoleID is the unique identifier of the updated role
	RoleID uuid.UUID

	// RoleName is the human-readable display name of the updated role
	RoleName string

	// RoleCode is the programmatic identifier of the updated role
	RoleCode string
}

// RoleDeleteRequest represents the request for deleting a role
// It contains the ID of the role to delete

type RoleDeleteRequest struct {
	// RoleID is the unique identifier of the role to delete
	RoleID uuid.UUID
}

// RoleDeleteResponse represents the response for deleting a role
// It indicates whether the deletion was successful
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields

type RoleDeleteResponse struct {
	// Success indicates whether the deletion was successful
	Success bool
}

// RoleAssignPermissionRequest represents the request for assigning a permission to a role
// It contains the IDs of the role and permission

type RoleAssignPermissionRequest struct {
	// RoleID is the unique identifier of the role to assign the permission to
	RoleID uuid.UUID

	// PermissionID is the unique identifier of the permission to assign
	PermissionID uuid.UUID
}

// RoleAssignPermissionResponse represents the response for assigning a permission to a role
// It indicates whether the assignment was successful
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields

type RoleAssignPermissionResponse struct {
	// Success indicates whether the assignment was successful
	Success bool
}

// RoleRemovePermissionRequest represents the request for removing a permission from a role
// It contains the IDs of the role and permission

type RoleRemovePermissionRequest struct {
	// RoleID is the unique identifier of the role to remove the permission from
	RoleID uuid.UUID

	// PermissionID is the unique identifier of the permission to remove
	PermissionID uuid.UUID
}

// RoleRemovePermissionResponse represents the response for removing a permission from a role
// It indicates whether the removal was successful
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields

type RoleRemovePermissionResponse struct {
	// Success indicates whether the removal was successful
	Success bool
}

// RolePermissionsRequest represents the request for retrieving a role's permissions
// It contains the ID of the role

type RolePermissionsRequest struct {
	// RoleID is the unique identifier of the role to retrieve permissions for
	RoleID uuid.UUID
}

// RolePermissionsResponse represents the response for retrieving a role's permissions
// It contains the role's ID and a list of its permissions
// The response format follows the LoginResponse pattern, directly containing entity information
// without Status and Message fields

type RolePermissionsResponse struct {
	// RoleID is the unique identifier of the role
	RoleID uuid.UUID

	// Permissions is the list of permissions assigned to the role
	Permissions []PermissionResponse
}

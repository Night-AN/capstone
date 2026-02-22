package usecase

import (
	"github.com/google/uuid"
)

// PermissionResourceCreateRequest represents the request for creating a permission-resource relationship
// It contains the IDs of the permission and resource to associate
type PermissionResourceCreateRequest struct {
	// PermissionID is the unique identifier of the permission
	PermissionID uuid.UUID `json:"permission_id"`

	// ResourceID is the unique identifier of the resource
	ResourceID uuid.UUID `json:"resource_id"`
}

// PermissionResourceCreateResponse represents the response for creating a permission-resource relationship
// It indicates whether the creation was successful

type PermissionResourceCreateResponse struct {
	// Success indicates whether the creation was successful
	Success bool `json:"success"`
}
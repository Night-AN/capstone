package usecase

import (
	"github.com/google/uuid"
)

// PermissionResourceDeleteRequest represents the request for deleting a permission-resource relationship
// It contains the IDs of the permission and resource to disassociate
type PermissionResourceDeleteRequest struct {
	// PermissionID is the unique identifier of the permission
	PermissionID uuid.UUID `json:"permission_id"`

	// ResourceID is the unique identifier of the resource
	ResourceID uuid.UUID `json:"resource_id"`
}

// PermissionResourceDeleteResponse represents the response for deleting a permission-resource relationship
// It indicates whether the deletion was successful

type PermissionResourceDeleteResponse struct {
	// Success indicates whether the deletion was successful
	Success bool `json:"success"`
}
package usecase

import (
	"github.com/google/uuid"
)

// ResourceMoveRequest represents the request for moving a resource
// This is used when transferring a resource from one parent or location to another
type ResourceMoveRequest struct {
	// ResourceID is the ID of the resource to be moved
	ResourceID uuid.UUID `json:"resource_id"`
	
	// NewParentResourceID is the ID of the new parent resource
	// If this is empty, the resource will be moved to the root level
	NewParentResourceID *uuid.UUID `json:"new_parent_resource_id"`
	
	// NewOrganizationID is the ID of the new organization
	// If this is provided, the resource will be transferred to a different organization
	NewOrganizationID *uuid.UUID `json:"new_organization_id"`
}

// ResourceMoveResponse represents the response after moving a resource
type ResourceMoveResponse struct {
	// Success indicates whether the resource was moved successfully
	Success bool `json:"success"`
	
	// ResourceID is the ID of the moved resource
	ResourceID uuid.UUID `json:"resource_id"`
	
	// NewParentResourceID is the new parent resource ID
	NewParentResourceID *uuid.UUID `json:"new_parent_resource_id"`
	
	// NewOrganizationID is the new organization ID (if changed)
	NewOrganizationID *uuid.UUID `json:"new_organization_id"`
}

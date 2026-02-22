package usecase

import (
	"github.com/google/uuid"
)

// OrganizationRoleDeleteRequest represents the request for deleting an organization-role relationship
// It contains the IDs of the organization and role to disassociate
type OrganizationRoleDeleteRequest struct {
	// OrganizationID is the unique identifier of the organization
	OrganizationID uuid.UUID `json:"organization_id"`

	// RoleID is the unique identifier of the role
	RoleID uuid.UUID `json:"role_id"`
}

// OrganizationRoleDeleteResponse represents the response for deleting an organization-role relationship
// It indicates whether the deletion was successful

type OrganizationRoleDeleteResponse struct {
	// Success indicates whether the deletion was successful
	Success bool `json:"success"`
}
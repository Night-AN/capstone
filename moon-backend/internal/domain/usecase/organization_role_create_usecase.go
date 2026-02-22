package usecase

import (
	"github.com/google/uuid"
)

// OrganizationRoleCreateRequest represents the request for creating an organization-role relationship
// It contains the IDs of the organization and role to associate
type OrganizationRoleCreateRequest struct {
	// OrganizationID is the unique identifier of the organization
	OrganizationID uuid.UUID `json:"organization_id"`

	// RoleID is the unique identifier of the role
	RoleID uuid.UUID `json:"role_id"`
}

// OrganizationRoleCreateResponse represents the response for creating an organization-role relationship
// It indicates whether the creation was successful

type OrganizationRoleCreateResponse struct {
	// Success indicates whether the creation was successful
	Success bool `json:"success"`
}
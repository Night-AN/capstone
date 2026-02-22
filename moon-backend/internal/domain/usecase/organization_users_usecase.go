package usecase

import (
	"github.com/google/uuid"
)

// OrganizationUsersRequest represents the request for getting users by organization ID
type OrganizationUsersRequest struct {
	OrganizationID uuid.UUID `json:"organization_id" binding:"required"`
}

// OrganizationUsersResponse represents the response for getting users by organization ID
type OrganizationUsersResponse struct {
	Users []UserListItem `json:"users"`
}

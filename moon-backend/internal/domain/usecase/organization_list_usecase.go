package usecase

import (
	"time"

	"github.com/google/uuid"
)

// OrganizationListRequest represents the request for listing organizations
// It can be extended with pagination, filtering, and sorting parameters in the future
type OrganizationListRequest struct {
	// Empty for now, can be extended with pagination, filtering, and sorting parameters
}

// OrganizationListItem represents a single organization in the list response
// It contains the basic information of an organization
type OrganizationListItem struct {
	// OrganizationID is the unique identifier of the organization
	OrganizationID uuid.UUID `json:"organization_id"`

	// OrganizationName is the human-readable display name of the organization
	OrganizationName string `json:"organization_name"`

	// OrganizationCode is the programmatic identifier of the organization
	OrganizationCode string `json:"organization_code"`

	// OrganizationFlag indicates the status or type of the organization
	OrganizationFlag string `json:"organization_flag"`

	// CreatedAt records the timestamp when the organization was first defined in the system
	CreatedAt time.Time `json:"created_at"`
}

// OrganizationListResponse represents the response for listing organizations
// It contains a list of organizations
type OrganizationListResponse struct {
	// Organizations is the list of organizations
	Organizations []OrganizationListItem `json:"organizations"`
}

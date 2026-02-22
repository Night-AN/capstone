package usecase

import (
	"github.com/google/uuid"
)

// OrganizationMoveRequest represents the request for moving an organization
// It contains the ID of the organization to move and the ID of the new parent organization
// If NewParentID is nil, the organization will become a root organization
// The request is validated by the handler before being passed to the service
// The service will ensure that the move is valid (e.g., no circular references)
// The repository will perform the actual update in the database
// The response will indicate whether the move was successful
// The handler will return an appropriate HTTP status code based on the response
// Example: Move organization with ID "org-123" to be a child of organization with ID "org-456"
//
//	{
//	  "organization_id": "org-123",
//	  "new_parent_id": "org-456"
//	}
//
// Example: Move organization with ID "org-123" to be a root organization
//
//	{
//	  "organization_id": "org-123",
//	  "new_parent_id": null
//	}
type OrganizationMoveRequest struct {
	OrganizationID uuid.UUID  `json:"organization_id" binding:"required"`
	NewParentID    *uuid.UUID `json:"new_parent_id"`
}

// OrganizationMoveResponse represents the response for moving an organization
// It contains a success flag indicating whether the move was successful
// If the move was not successful, the success flag will be false
// The handler will use this response to determine the HTTP status code
// Example: Successful move
//
//	{
//	  "success": true
//	}
//
// Example: Failed move
//
//	{
//	  "success": false
//	}
type OrganizationMoveResponse struct {
	Success bool `json:"success"`
}

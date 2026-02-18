package usecase

import "github.com/google/uuid"

type OrganizationDeleteRequest struct {
	OrganizationID uuid.UUID `json:"organization_id"`
}

type OrganizationDeleteResponse struct {
	Success bool `json:"success"`
}

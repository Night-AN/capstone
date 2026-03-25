package usecase

import "github.com/google/uuid"

type OrganizationUpdateRequest struct {
	OrganizationID          uuid.UUID `json:"organization_id"`
	OrganizationName        string    `json:"organization_name"`
	OrganizationCode        string    `json:"organization_code"`
	OrganizationDescription string    `json:"organization_description"`
	OrganizationFlag        string    `json:"organization_flag"`
}

type OrganizationUpdateResponse struct {
	Success bool `json:"success"`
}

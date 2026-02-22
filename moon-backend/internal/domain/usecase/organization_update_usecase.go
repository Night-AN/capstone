package usecase

import "github.com/google/uuid"

type OrganizationUpdateRequest struct {
	OrganizationID          uuid.UUID  `json:"organization_id"`
	OrganizationName        string     `json:"organization_name"`
	OrganizationCode        string     `json:"organization_code"`
	OrganizationDescription string     `json:"organization_description"`
	OrganizationFlag        string     `json:"organization_flag"`
	SensitiveFlag           bool       `json:"sensitive_flag"`
	ParentID                *uuid.UUID `json:"parent_id"`
}

type OrganizationUpdateResponse struct {
	OrganizationID   uuid.UUID `json:"organization_id"`
	OrganizationName string    `json:"organization_name"`
	OrganizationCode string    `json:"organization_code"`
}

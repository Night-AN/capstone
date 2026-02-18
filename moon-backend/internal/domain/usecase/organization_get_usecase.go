package usecase

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationGetRequest struct {
	OrganizationID uuid.UUID `json:"organization_id"`
}

type OrganizationGetResponse struct {
	OrganizationID          uuid.UUID `json:"organization_id"`
	OrganizationName        string    `json:"organization_name"`
	OrganizationCode        string    `json:"organization_code"`
	OrganizationDescription string    `json:"organization_description"`
	OrganizationFlag        string    `json:"organization_flag"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

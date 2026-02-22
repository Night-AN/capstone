package usecase

import (
	"moon/internal/domain/aggregate"
	"time"

	"github.com/google/uuid"
)

type OrganizationCreateRequest struct {
	OrganizationName string `json:"organization_name"`

	OrganizationCode string `json:"organization_code"`

	OrganizationDescription string `json:"organization_description"`

	OrganizationFlag string `json:"organization_flag"`

	SensitiveFlag bool `json:"sensitive_flag"`

	ParentID *uuid.UUID `json:"parent_id"`
}

type OrganizationCreateResponse struct {
	OrganizationID   uuid.UUID `json:"organization_id"`
	OrganizationName string    `json:"organization_name"`
	OrganizationCode string    `json:"organization_code"`
}

func ConvertOrganizationCreateRequestToOrganizationAggregate(req OrganizationCreateRequest) aggregate.Organization {
	return aggregate.Organization{
		OrganizationID:          uuid.New(),
		OrganizationName:        req.OrganizationName,
		OrganizationCode:        req.OrganizationCode,
		OrganizationDescription: req.OrganizationDescription,
		OrganizationFlag:        req.OrganizationFlag,
		SensitiveFlag:           req.SensitiveFlag,
		ParentID:                req.ParentID,
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
	}
}

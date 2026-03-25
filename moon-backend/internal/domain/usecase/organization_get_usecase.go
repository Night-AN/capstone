package usecase

import (
	"time"

	"github.com/google/uuid"
)

type OrganizationItem struct {
	OrganizationID          uuid.UUID `json:"organization_id"`
	OrganizationName        string    `json:"organization_name"`
	OrganizationCode        string    `json:"organization_code"`
	OrganizationDescription string    `json:"organization_description"`
	OrganizationFlag        string    `json:"organization_flag"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

type OrganizationTreeNode struct {
	OrganizationID       uuid.UUID                        `json:"organization_id"`
	OrganizationName     string                           `json:"organization_name"`
	OrganizationCode     string                           `json:"organization_code"`
	OrganizationCodePart string                           `json:"organization_code_part"`
	Children             map[string]*OrganizationTreeNode `json:"children"`
}

type OrganizationGetRequest struct {
	OrganizationID uuid.UUID `json:"organization_id"`
}

type OrganizationGetResponse = OrganizationItem

type OrganizationTreeRequest struct {
	OrganizationID uuid.UUID `json:"organization_id"`
}

type OrganizationTreeResponse struct {
	OrganizationTree *OrganizationTreeNode `json:"organization_tree"`
}

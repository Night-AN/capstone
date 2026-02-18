package repository

import (
	"context"
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

type OrganizationRepository interface {
	SaveOrganization(ctx *context.Context, org aggregate.Organization) error
	FindOrganizationByID(ctx *context.Context, org_id uuid.UUID) (aggregate.Organization, error)
	FindOrganizationByName(ctx *context.Context, org_name string) ([]aggregate.Organization, error)
	FindOrganizationByCode(ctx *context.Context, org_code string) ([]aggregate.Organization, error)
}

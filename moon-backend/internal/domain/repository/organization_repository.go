package repository

import (
	"context"
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

type OrganizationRepository interface {
	SaveOrganization(ctx context.Context, org aggregate.Organization) error
	FindOrganizationByID(ctx context.Context, org_id uuid.UUID) (aggregate.Organization, error)
	FindOrganizationByName(ctx context.Context, org_name string) ([]aggregate.Organization, error)
	FindOrganizationByCode(ctx context.Context, org_code string) ([]aggregate.Organization, error)
	FindAllOrganizations(ctx context.Context) ([]aggregate.Organization, error)
	FindOrganizationsByParentID(ctx context.Context, parent_id uuid.UUID) ([]aggregate.Organization, error)
	UpdateOrganizationParent(ctx context.Context, org_id uuid.UUID, parent_id *uuid.UUID) error
	AssignRoleToOrganization(ctx context.Context, org_id uuid.UUID, role_id uuid.UUID) error
	RemoveRoleFromOrganization(ctx context.Context, org_id uuid.UUID, role_id uuid.UUID) error
	FindRolesByOrganizationID(ctx context.Context, org_id uuid.UUID) ([]aggregate.Role, error)
	DeleteOrganization(ctx context.Context, org_id uuid.UUID) error
}

package service

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"

	"github.com/google/uuid"
)

// OrganizationRoleService defines the service interface for organization-role operations
type OrganizationRoleService interface {
	// Create creates a new organization-role relationship
	Create(ctx context.Context, req usecase.OrganizationRoleCreateRequest) (usecase.OrganizationRoleCreateResponse, error)

	// Delete deletes an existing organization-role relationship
	Delete(ctx context.Context, req usecase.OrganizationRoleDeleteRequest) (usecase.OrganizationRoleDeleteResponse, error)

	// GetRolesByOrganizationID gets all roles for an organization
	GetRolesByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]aggregate.Role, error)

	// GetOrganizationsByRoleID gets all organizations for a role
	GetOrganizationsByRoleID(ctx context.Context, roleID uuid.UUID) ([]aggregate.Organization, error)
}

// organizationRoleService implements OrganizationRoleService
type organizationRoleService struct {
	organizationRoleRepo repository.OrganizationRoleRepository
	organizationRepo     repository.OrganizationRepository
	roleRepo             repository.RoleRepository
}

// NewOrganizationRoleService creates a new OrganizationRoleService
func NewOrganizationRoleService(
	organizationRoleRepo repository.OrganizationRoleRepository,
	organizationRepo repository.OrganizationRepository,
	roleRepo repository.RoleRepository,
) OrganizationRoleService {
	return &organizationRoleService{
		organizationRoleRepo: organizationRoleRepo,
		organizationRepo:     organizationRepo,
		roleRepo:             roleRepo,
	}
}

// Create creates a new organization-role relationship
func (s *organizationRoleService) Create(ctx context.Context, req usecase.OrganizationRoleCreateRequest) (usecase.OrganizationRoleCreateResponse, error) {
	// Check if organization exists
	_, err := s.organizationRepo.GetByID(ctx, req.OrganizationID)
	if err != nil {
		return usecase.OrganizationRoleCreateResponse{Success: false}, err
	}

	// Check if role exists
	_, err = s.roleRepo.GetByID(ctx, req.RoleID)
	if err != nil {
		return usecase.OrganizationRoleCreateResponse{Success: false}, err
	}

	// Create organization-role relationship
	or := aggregate.OrganizationRole{
		ID:             uuid.New(),
		OrganizationID: req.OrganizationID,
		RoleID:         req.RoleID,
	}

	err = s.organizationRoleRepo.Create(ctx, or)
	if err != nil {
		return usecase.OrganizationRoleCreateResponse{Success: false}, err
	}

	return usecase.OrganizationRoleCreateResponse{Success: true}, nil
}

// Delete deletes an existing organization-role relationship
func (s *organizationRoleService) Delete(ctx context.Context, req usecase.OrganizationRoleDeleteRequest) (usecase.OrganizationRoleDeleteResponse, error) {
	err := s.organizationRoleRepo.Delete(ctx, req.OrganizationID, req.RoleID)
	if err != nil {
		return usecase.OrganizationRoleDeleteResponse{Success: false}, err
	}

	return usecase.OrganizationRoleDeleteResponse{Success: true}, nil
}

// GetRolesByOrganizationID gets all roles for an organization
func (s *organizationRoleService) GetRolesByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]aggregate.Role, error) {
	return s.organizationRoleRepo.GetRolesByOrganizationID(ctx, organizationID)
}

// GetOrganizationsByRoleID gets all organizations for a role
func (s *organizationRoleService) GetOrganizationsByRoleID(ctx context.Context, roleID uuid.UUID) ([]aggregate.Organization, error) {
	return s.organizationRoleRepo.GetOrganizationsByRoleID(ctx, roleID)
}

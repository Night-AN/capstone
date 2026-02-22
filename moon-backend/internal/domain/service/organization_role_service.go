package service

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"

	"github.com/google/uuid"
)

// OrganizationRoleService defines the interface for organization-role relationship operations
// It provides methods for creating, deleting, and retrieving organization-role relationships
// All methods accept a context and return the requested data or an error
// The context is used for cancellation, timeouts, and passing request-scoped values
// The service layer is responsible for business logic and orchestration
// It should coordinate calls to multiple repositories and handle business rules
// The service layer should be tested in isolation from the repository layer
// using mock repositories
// The service layer should be stateless and thread-safe
// It should not maintain any internal state between method calls
// All dependencies should be injected via the constructor
// The service layer should not contain any database-specific logic
// It should only be concerned with business logic and orchestration
type OrganizationRoleService interface {
	// Create creates an organization-role relationship
	// It returns a response indicating success or failure and an error if the operation fails
	// If the organization or role does not exist, it returns an error
	// If the relationship already exists, it returns an error
	Create(ctx context.Context, req usecase.OrganizationRoleCreateRequest) (usecase.OrganizationRoleCreateResponse, error)

	// Delete deletes an organization-role relationship
	// It returns a response indicating success or failure and an error if the operation fails
	// If the relationship does not exist, it returns an error
	Delete(ctx context.Context, req usecase.OrganizationRoleDeleteRequest) (usecase.OrganizationRoleDeleteResponse, error)

	// GetRolesByOrganizationID retrieves all roles assigned to an organization
	// It returns a list of roles and an error if the operation fails
	// If no roles are found, it returns an empty list and no error
	GetRolesByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]aggregate.Role, error)

	// GetOrganizationsByRoleID retrieves all organizations assigned to a role
	// It returns a list of organizations and an error if the operation fails
	// If no organizations are found, it returns an empty list and no error
	GetOrganizationsByRoleID(ctx context.Context, roleID uuid.UUID) ([]aggregate.Organization, error)
}

// organizationRoleService implements the OrganizationRoleService interface
// It provides methods for creating, deleting, and retrieving organization-role relationships
// All methods accept a context and return the requested data or an error
// The context is used for cancellation, timeouts, and passing request-scoped values
// The service layer is responsible for business logic and orchestration
// It should coordinate calls to multiple repositories and handle business rules
// The service layer should be tested in isolation from the repository layer
// using mock repositories
// The service layer should be stateless and thread-safe
// It should not maintain any internal state between method calls
// All dependencies should be injected via the constructor
// The service layer should not contain any database-specific logic
// It should only be concerned with business logic and orchestration
type organizationRoleService struct {
	// organizationRoleRepo is the repository for organization-role relationships
	organizationRoleRepo repository.OrganizationRoleRepository
	// organizationRepo is the repository for organizations
	organizationRepo repository.OrganizationRepository
	// roleRepo is the repository for roles
	roleRepo repository.RoleRepository
}

// NewOrganizationRoleService creates a new OrganizationRoleService instance
// It injects the required dependencies
// It returns an OrganizationRoleService instance
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

// Create creates an organization-role relationship
// It returns a response indicating success or failure and an error if the operation fails
// If the organization or role does not exist, it returns an error
// If the relationship already exists, it returns an error
func (s *organizationRoleService) Create(ctx context.Context, req usecase.OrganizationRoleCreateRequest) (usecase.OrganizationRoleCreateResponse, error) {
	// Create the organization-role relationship
	err := s.organizationRoleRepo.Create(ctx, aggregate.OrganizationRole{
		OrganizationID: req.OrganizationID,
		RoleID:         req.RoleID,
	})
	if err != nil {
		return usecase.OrganizationRoleCreateResponse{Success: false}, err
	}

	// Return success response
	return usecase.OrganizationRoleCreateResponse{Success: true}, nil
}

// Delete deletes an organization-role relationship
// It returns a response indicating success or failure and an error if the operation fails
// If the relationship does not exist, it returns an error
func (s *organizationRoleService) Delete(ctx context.Context, req usecase.OrganizationRoleDeleteRequest) (usecase.OrganizationRoleDeleteResponse, error) {
	// Delete the organization-role relationship
	err := s.organizationRoleRepo.Delete(ctx, req.OrganizationID, req.RoleID)
	if err != nil {
		return usecase.OrganizationRoleDeleteResponse{Success: false}, err
	}

	// Return success response
	return usecase.OrganizationRoleDeleteResponse{Success: true}, nil
}

// GetRolesByOrganizationID retrieves all roles assigned to an organization
// It returns a list of roles and an error if the operation fails
// If no roles are found, it returns an empty list and no error
func (s *organizationRoleService) GetRolesByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]aggregate.Role, error) {
	// Get roles by organization ID
	return s.organizationRoleRepo.GetRolesByOrganizationID(ctx, organizationID)
}

// GetOrganizationsByRoleID retrieves all organizations assigned to a role
// It returns a list of organizations and an error if the operation fails
// If no organizations are found, it returns an empty list and no error
func (s *organizationRoleService) GetOrganizationsByRoleID(ctx context.Context, roleID uuid.UUID) ([]aggregate.Organization, error) {
	// Get organizations by role ID
	return s.organizationRoleRepo.GetOrganizationsByRoleID(ctx, roleID)
}

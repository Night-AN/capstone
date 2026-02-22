package service

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"

	"github.com/google/uuid"
)

// RolePermissionService defines the interface for role-permission relationship operations
// It provides methods for creating, deleting, and retrieving role-permission relationships
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
type RolePermissionService interface {
	// Create creates a role-permission relationship
	// It returns a response indicating success or failure and an error if the operation fails
	// If the role or permission does not exist, it returns an error
	// If the relationship already exists, it returns an error
	Create(ctx context.Context, req usecase.RolePermissionCreateRequest) (usecase.RolePermissionCreateResponse, error)

	// Delete deletes a role-permission relationship
	// It returns a response indicating success or failure and an error if the operation fails
	// If the relationship does not exist, it returns an error
	Delete(ctx context.Context, req usecase.RolePermissionDeleteRequest) (usecase.RolePermissionDeleteResponse, error)

	// GetPermissionsByRoleID retrieves all permissions assigned to a role
	// It returns a list of permissions and an error if the operation fails
	// If no permissions are found, it returns an empty list and no error
	GetPermissionsByRoleID(ctx context.Context, roleID uuid.UUID) ([]aggregate.Permission, error)

	// GetRolesByPermissionID retrieves all roles assigned to a permission
	// It returns a list of roles and an error if the operation fails
	// If no roles are found, it returns an empty list and no error
	GetRolesByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]aggregate.Role, error)
}

// rolePermissionService implements the RolePermissionService interface
// It provides methods for creating, deleting, and retrieving role-permission relationships
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
type rolePermissionService struct {
	// rolePermissionRepo is the repository for role-permission relationships
	rolePermissionRepo repository.RolePermissionRepository
	// roleRepo is the repository for roles
	roleRepo repository.RoleRepository
	// permissionRepo is the repository for permissions
	permissionRepo repository.PermissionRepository
}

// NewRolePermissionService creates a new RolePermissionService instance
// It injects the required dependencies
// It returns a RolePermissionService instance
func NewRolePermissionService(
	rolePermissionRepo repository.RolePermissionRepository,
	roleRepo repository.RoleRepository,
	permissionRepo repository.PermissionRepository,
) RolePermissionService {
	return &rolePermissionService{
		rolePermissionRepo: rolePermissionRepo,
		roleRepo:           roleRepo,
		permissionRepo:     permissionRepo,
	}
}

// Create creates a role-permission relationship
// It returns a response indicating success or failure and an error if the operation fails
// If the role or permission does not exist, it returns an error
// If the relationship already exists, it returns an error
func (s *rolePermissionService) Create(ctx context.Context, req usecase.RolePermissionCreateRequest) (usecase.RolePermissionCreateResponse, error) {
	// Create the role-permission relationship
	err := s.rolePermissionRepo.Create(ctx, aggregate.RolePermission{
		RoleID:       req.RoleID,
		PermissionID: req.PermissionID,
	})
	if err != nil {
		return usecase.RolePermissionCreateResponse{Success: false}, err
	}

	// Return success response
	return usecase.RolePermissionCreateResponse{Success: true}, nil
}

// Delete deletes a role-permission relationship
// It returns a response indicating success or failure and an error if the operation fails
// If the relationship does not exist, it returns an error
func (s *rolePermissionService) Delete(ctx context.Context, req usecase.RolePermissionDeleteRequest) (usecase.RolePermissionDeleteResponse, error) {
	// Delete the role-permission relationship
	err := s.rolePermissionRepo.Delete(ctx, req.RoleID, req.PermissionID)
	if err != nil {
		return usecase.RolePermissionDeleteResponse{Success: false}, err
	}

	// Return success response
	return usecase.RolePermissionDeleteResponse{Success: true}, nil
}

// GetPermissionsByRoleID retrieves all permissions assigned to a role
// It returns a list of permissions and an error if the operation fails
// If no permissions are found, it returns an empty list and no error
func (s *rolePermissionService) GetPermissionsByRoleID(ctx context.Context, roleID uuid.UUID) ([]aggregate.Permission, error) {
	// Get permissions by role ID
	return s.rolePermissionRepo.GetPermissionsByRoleID(ctx, roleID)
}

// GetRolesByPermissionID retrieves all roles assigned to a permission
// It returns a list of roles and an error if the operation fails
// If no roles are found, it returns an empty list and no error
func (s *rolePermissionService) GetRolesByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]aggregate.Role, error) {
	// Get roles by permission ID
	return s.rolePermissionRepo.GetRolesByPermissionID(ctx, permissionID)
}

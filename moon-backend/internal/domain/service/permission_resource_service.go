package service

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"

	"github.com/google/uuid"
)

// PermissionResourceService defines the interface for permission-resource relationship operations
// It provides methods for creating, deleting, and retrieving permission-resource relationships
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
type PermissionResourceService interface {
	// Create creates a permission-resource relationship
	// It returns a response indicating success or failure and an error if the operation fails
	// If the permission or resource does not exist, it returns an error
	// If the relationship already exists, it returns an error
	Create(ctx context.Context, req usecase.PermissionResourceCreateRequest) (usecase.PermissionResourceCreateResponse, error)

	// Delete deletes a permission-resource relationship
	// It returns a response indicating success or failure and an error if the operation fails
	// If the relationship does not exist, it returns an error
	Delete(ctx context.Context, req usecase.PermissionResourceDeleteRequest) (usecase.PermissionResourceDeleteResponse, error)

	// GetPermissionsByResourceID retrieves all permissions assigned to a resource
	// It returns a list of permissions and an error if the operation fails
	// If no permissions are found, it returns an empty list and no error
	GetPermissionsByResourceID(ctx context.Context, resourceID uuid.UUID) ([]aggregate.Permission, error)

	// GetResourcesByPermissionID retrieves all resources assigned to a permission
	// It returns a list of resources and an error if the operation fails
	// If no resources are found, it returns an empty list and no error
	GetResourcesByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]aggregate.Resource, error)
}

// permissionResourceService implements the PermissionResourceService interface
// It provides methods for creating, deleting, and retrieving permission-resource relationships
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
type permissionResourceService struct {
	// permissionResourceRepo is the repository for permission-resource relationships
	permissionResourceRepo repository.PermissionResourceRepository
	// permissionRepo is the repository for permissions
	permissionRepo repository.PermissionRepository
	// resourceRepo is the repository for resources
	resourceRepo repository.ResourceRepository
}

// NewPermissionResourceService creates a new PermissionResourceService instance
// It injects the required dependencies
// It returns a PermissionResourceService instance
func NewPermissionResourceService(
	permissionResourceRepo repository.PermissionResourceRepository,
	permissionRepo repository.PermissionRepository,
	resourceRepo repository.ResourceRepository,
) PermissionResourceService {
	return &permissionResourceService{
		permissionResourceRepo: permissionResourceRepo,
		permissionRepo:         permissionRepo,
		resourceRepo:           resourceRepo,
	}
}

// Create creates a permission-resource relationship
// It returns a response indicating success or failure and an error if the operation fails
// If the permission or resource does not exist, it returns an error
// If the relationship already exists, it returns an error
func (s *permissionResourceService) Create(ctx context.Context, req usecase.PermissionResourceCreateRequest) (usecase.PermissionResourceCreateResponse, error) {
	// Create the permission-resource relationship
	err := s.permissionResourceRepo.Create(ctx, aggregate.PermissionResource{
		PermissionID: req.PermissionID,
		ResourceID:   req.ResourceID,
	})
	if err != nil {
		return usecase.PermissionResourceCreateResponse{Success: false}, err
	}

	// Return success response
	return usecase.PermissionResourceCreateResponse{Success: true}, nil
}

// Delete deletes a permission-resource relationship
// It returns a response indicating success or failure and an error if the operation fails
// If the relationship does not exist, it returns an error
func (s *permissionResourceService) Delete(ctx context.Context, req usecase.PermissionResourceDeleteRequest) (usecase.PermissionResourceDeleteResponse, error) {
	// Delete the permission-resource relationship
	err := s.permissionResourceRepo.Delete(ctx, req.PermissionID, req.ResourceID)
	if err != nil {
		return usecase.PermissionResourceDeleteResponse{Success: false}, err
	}

	// Return success response
	return usecase.PermissionResourceDeleteResponse{Success: true}, nil
}

// GetPermissionsByResourceID retrieves all permissions assigned to a resource
// It returns a list of permissions and an error if the operation fails
// If no permissions are found, it returns an empty list and no error
func (s *permissionResourceService) GetPermissionsByResourceID(ctx context.Context, resourceID uuid.UUID) ([]aggregate.Permission, error) {
	// Get permissions by resource ID
	return s.permissionResourceRepo.GetPermissionsByResourceID(ctx, resourceID)
}

// GetResourcesByPermissionID retrieves all resources assigned to a permission
// It returns a list of resources and an error if the operation fails
// If no resources are found, it returns an empty list and no error
func (s *permissionResourceService) GetResourcesByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]aggregate.Resource, error) {
	// Get resources by permission ID
	return s.permissionResourceRepo.GetResourcesByPermissionID(ctx, permissionID)
}

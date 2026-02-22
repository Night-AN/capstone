package service

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"

	"github.com/google/uuid"
)

// PermissionResourceService defines the service interface for permission-resource operations
type PermissionResourceService interface {
	// Create creates a new permission-resource relationship
	Create(ctx context.Context, req usecase.PermissionResourceCreateRequest) (usecase.PermissionResourceCreateResponse, error)

	// Delete deletes an existing permission-resource relationship
	Delete(ctx context.Context, req usecase.PermissionResourceDeleteRequest) (usecase.PermissionResourceDeleteResponse, error)

	// GetPermissionsByResourceID gets all permissions for a resource
	GetPermissionsByResourceID(ctx context.Context, resourceID uuid.UUID) ([]aggregate.Permission, error)

	// GetResourcesByPermissionID gets all resources for a permission
	GetResourcesByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]aggregate.Resource, error)
}

// permissionResourceService implements PermissionResourceService
type permissionResourceService struct {
	permissionResourceRepo repository.PermissionResourceRepository
	permissionRepo         repository.PermissionRepository
	resourceRepo           repository.ResourceRepository
}

// NewPermissionResourceService creates a new PermissionResourceService
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

// Create creates a new permission-resource relationship
func (s *permissionResourceService) Create(ctx context.Context, req usecase.PermissionResourceCreateRequest) (usecase.PermissionResourceCreateResponse, error) {
	// Check if permission exists
	_, err := s.permissionRepo.GetByID(ctx, req.PermissionID)
	if err != nil {
		return usecase.PermissionResourceCreateResponse{Success: false}, err
	}

	// Check if resource exists
	_, err = s.resourceRepo.GetByID(ctx, req.ResourceID)
	if err != nil {
		return usecase.PermissionResourceCreateResponse{Success: false}, err
	}

	// Create permission-resource relationship
	pr := aggregate.PermissionResource{
		ID:           uuid.New(),
		PermissionID: req.PermissionID,
		ResourceID:   req.ResourceID,
	}

	err = s.permissionResourceRepo.Create(ctx, pr)
	if err != nil {
		return usecase.PermissionResourceCreateResponse{Success: false}, err
	}

	return usecase.PermissionResourceCreateResponse{Success: true}, nil
}

// Delete deletes an existing permission-resource relationship
func (s *permissionResourceService) Delete(ctx context.Context, req usecase.PermissionResourceDeleteRequest) (usecase.PermissionResourceDeleteResponse, error) {
	err := s.permissionResourceRepo.Delete(ctx, req.PermissionID, req.ResourceID)
	if err != nil {
		return usecase.PermissionResourceDeleteResponse{Success: false}, err
	}

	return usecase.PermissionResourceDeleteResponse{Success: true}, nil
}

// GetPermissionsByResourceID gets all permissions for a resource
func (s *permissionResourceService) GetPermissionsByResourceID(ctx context.Context, resourceID uuid.UUID) ([]aggregate.Permission, error) {
	return s.permissionResourceRepo.GetPermissionsByResourceID(ctx, resourceID)
}

// GetResourcesByPermissionID gets all resources for a permission
func (s *permissionResourceService) GetResourcesByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]aggregate.Resource, error) {
	return s.permissionResourceRepo.GetResourcesByPermissionID(ctx, permissionID)
}

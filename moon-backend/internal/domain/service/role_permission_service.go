package service

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"

	"github.com/google/uuid"
)

// RolePermissionService defines the service interface for role-permission operations
type RolePermissionService interface {
	// Create creates a new role-permission relationship
	Create(ctx context.Context, req usecase.RolePermissionCreateRequest) (usecase.RolePermissionCreateResponse, error)

	// Delete deletes an existing role-permission relationship
	Delete(ctx context.Context, req usecase.RolePermissionDeleteRequest) (usecase.RolePermissionDeleteResponse, error)

	// GetPermissionsByRoleID gets all permissions for a role
	GetPermissionsByRoleID(ctx context.Context, roleID uuid.UUID) ([]aggregate.Permission, error)

	// GetRolesByPermissionID gets all roles for a permission
	GetRolesByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]aggregate.Role, error)
}

// rolePermissionService implements RolePermissionService
type rolePermissionService struct {
	rolePermissionRepo repository.RolePermissionRepository
	roleRepo           repository.RoleRepository
	permissionRepo     repository.PermissionRepository
}

// NewRolePermissionService creates a new RolePermissionService
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

// Create creates a new role-permission relationship
func (s *rolePermissionService) Create(ctx context.Context, req usecase.RolePermissionCreateRequest) (usecase.RolePermissionCreateResponse, error) {
	// Check if role exists
	_, err := s.roleRepo.GetByID(ctx, req.RoleID)
	if err != nil {
		return usecase.RolePermissionCreateResponse{Success: false}, err
	}

	// Check if permission exists
	_, err = s.permissionRepo.GetByID(ctx, req.PermissionID)
	if err != nil {
		return usecase.RolePermissionCreateResponse{Success: false}, err
	}

	// Create role-permission relationship
	rp := aggregate.RolePermission{
		ID:           uuid.New(),
		RoleID:       req.RoleID,
		PermissionID: req.PermissionID,
	}

	err = s.rolePermissionRepo.Create(ctx, rp)
	if err != nil {
		return usecase.RolePermissionCreateResponse{Success: false}, err
	}

	return usecase.RolePermissionCreateResponse{Success: true}, nil
}

// Delete deletes an existing role-permission relationship
func (s *rolePermissionService) Delete(ctx context.Context, req usecase.RolePermissionDeleteRequest) (usecase.RolePermissionDeleteResponse, error) {
	err := s.rolePermissionRepo.Delete(ctx, req.RoleID, req.PermissionID)
	if err != nil {
		return usecase.RolePermissionDeleteResponse{Success: false}, err
	}

	return usecase.RolePermissionDeleteResponse{Success: true}, nil
}

// GetPermissionsByRoleID gets all permissions for a role
func (s *rolePermissionService) GetPermissionsByRoleID(ctx context.Context, roleID uuid.UUID) ([]aggregate.Permission, error) {
	return s.rolePermissionRepo.GetPermissionsByRoleID(ctx, roleID)
}

// GetRolesByPermissionID gets all roles for a permission
func (s *rolePermissionService) GetRolesByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]aggregate.Role, error) {
	return s.rolePermissionRepo.GetRolesByPermissionID(ctx, permissionID)
}

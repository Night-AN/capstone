package postgres

import (
	"context"
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// permissionResourceRepository implements the repository.PermissionResourceRepository interface
type permissionResourceRepository struct {
	db *gorm.DB
}

// NewPermissionResourceRepository creates a new PermissionResourceRepository
func NewPermissionResourceRepository(db *gorm.DB) *permissionResourceRepository {
	return &permissionResourceRepository{db: db}
}

// Create creates a new permission-resource relationship
func (r *permissionResourceRepository) Create(ctx context.Context, pr aggregate.PermissionResource) error {
	return r.db.WithContext(ctx).Create(&pr).Error
}

// Delete deletes an existing permission-resource relationship
func (r *permissionResourceRepository) Delete(ctx context.Context, permissionID, resourceID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("permission_id = ? AND resource_id = ?", permissionID, resourceID).Delete(&aggregate.PermissionResource{}).Error
}

// GetByPermissionAndResource gets a permission-resource relationship by permission and resource IDs
func (r *permissionResourceRepository) GetByPermissionAndResource(ctx context.Context, permissionID, resourceID uuid.UUID) (aggregate.PermissionResource, error) {
	var pr aggregate.PermissionResource
	err := r.db.WithContext(ctx).Where("permission_id = ? AND resource_id = ?", permissionID, resourceID).First(&pr).Error
	return pr, err
}

// GetPermissionsByResourceID gets all permissions for a resource
func (r *permissionResourceRepository) GetPermissionsByResourceID(ctx context.Context, resourceID uuid.UUID) ([]aggregate.Permission, error) {
	var permissions []aggregate.Permission
	err := r.db.WithContext(ctx).Table("systems.permissions").Select("systems.permissions.*").Joins("INNER JOIN systems.permission_resource ON systems.permissions.id = systems.permission_resource.permission_id").Where("systems.permission_resource.resource_id = ?", resourceID).Find(&permissions).Error
	return permissions, err
}

// GetResourcesByPermissionID gets all resources for a permission
func (r *permissionResourceRepository) GetResourcesByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]aggregate.Resource, error) {
	var resources []aggregate.Resource
	err := r.db.WithContext(ctx).Table("systems.resources").Select("systems.resources.*").Joins("INNER JOIN systems.permission_resource ON systems.resources.id = systems.permission_resource.resource_id").Where("systems.permission_resource.permission_id = ?", permissionID).Find(&resources).Error
	return resources, err
}

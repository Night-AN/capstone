package postgres

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// permissionRepository implements the repository.PermissionRepository interface
// It provides methods for creating, retrieving, updating, and deleting permissions
// All methods use GORM to interact with the database
// The permissionRepository is initialized with a GORM DB instance
// It is stateless and thread-safe
// It should be tested with an actual PostgreSQL database
type permissionRepository struct {
	// db is the GORM DB instance used for database operations
	db *gorm.DB
}

// NewPermissionRepository creates a new instance of permissionRepository
// It takes a GORM DB instance as a parameter
// It returns a repository.PermissionRepository
func NewPermissionRepository(db *gorm.DB) repository.PermissionRepository {
	return &permissionRepository{db}
}

// SavePermission saves a permission to the database
// If the permission already exists, it updates the existing record
// If the permission does not exist, it creates a new record
// It returns an error if the operation fails
func (pr *permissionRepository) SavePermission(ctx *context.Context, permission aggregate.Permission) error {
	// Check if the permission already exists
	var existingPermission aggregate.Permission
	err := pr.db.WithContext(*ctx).Where("permission_id = ?", permission.PermissionID).First(&existingPermission).Error

	if err == gorm.ErrRecordNotFound {
		// Permission does not exist, create a new record
		return pr.db.WithContext(*ctx).Create(&permission).Error
	} else if err != nil {
		// Other error occurred
		return err
	} else {
		// Permission exists, update the existing record
		return pr.db.WithContext(*ctx).Model(&aggregate.Permission{}).Where("permission_id = ?", permission.PermissionID).Updates(&permission).Error
	}
}

// FindPermissionByID retrieves a permission by its ID
// It returns the permission and an error if the operation fails
// If the permission is not found, it returns an error
func (pr *permissionRepository) FindPermissionByID(ctx *context.Context, permissionID uuid.UUID) (aggregate.Permission, error) {
	var permission aggregate.Permission
	err := pr.db.WithContext(*ctx).Where("permission_id = ?", permissionID).First(&permission).Error
	return permission, err
}

// FindPermissionByName retrieves permissions by their name
// It returns a list of permissions and an error if the operation fails
// If no permissions are found, it returns an empty list and no error
func (pr *permissionRepository) FindPermissionByName(ctx *context.Context, permissionName string) ([]aggregate.Permission, error) {
	var permissions []aggregate.Permission
	pattern := "%" + permissionName + "%"
	err := pr.db.WithContext(*ctx).Where("permission_name LIKE ?", pattern).Find(&permissions).Error
	return permissions, err
}

// FindPermissionByCode retrieves permissions by their code
// It returns a list of permissions and an error if the operation fails
// If no permissions are found, it returns an empty list and no error
func (pr *permissionRepository) FindPermissionByCode(ctx *context.Context, permissionCode string) ([]aggregate.Permission, error) {
	var permissions []aggregate.Permission
	pattern := "%" + permissionCode + "%"
	err := pr.db.WithContext(*ctx).Where("permission_code LIKE ?", pattern).Find(&permissions).Error
	return permissions, err
}

// DeletePermission deletes a permission by its ID
// It returns an error if the operation fails
func (pr *permissionRepository) DeletePermission(ctx *context.Context, permissionID uuid.UUID) error {
	return pr.db.WithContext(*ctx).Where("permission_id = ?", permissionID).Delete(&aggregate.Permission{}).Error
}

// ListPermissions retrieves all permissions matching the specified criteria
// It returns a list of permissions and an error if the operation fails
// If no permissions are found, it returns an empty list and no error
func (pr *permissionRepository) ListPermissions(ctx *context.Context, limit, offset int) ([]aggregate.Permission, error) {
	var permissions []aggregate.Permission
	query := pr.db.WithContext(*ctx)

	if limit > 0 {
		query = query.Limit(limit)
	}

	if offset > 0 {
		query = query.Offset(offset)
	}

	err := query.Find(&permissions).Error
	return permissions, err
}

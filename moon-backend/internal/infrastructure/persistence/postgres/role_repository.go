package postgres

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// roleRepository implements the repository.RoleRepository interface
// It provides methods for creating, retrieving, updating, and deleting roles
// as well as managing role-permission relationships
// All methods use GORM to interact with the database
// The roleRepository is initialized with a GORM DB instance
// It is stateless and thread-safe
// It should be tested with an actual PostgreSQL database
type roleRepository struct {
	// db is the GORM DB instance used for database operations
	db *gorm.DB
}

// NewRoleRepository creates a new instance of roleRepository
// It takes a GORM DB instance as a parameter
// It returns a repository.RoleRepository
func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &roleRepository{db}
}

// SaveRole saves a role to the database
// If the role already exists, it updates the existing record
// If the role does not exist, it creates a new record
// It returns an error if the operation fails
func (rr *roleRepository) SaveRole(ctx *context.Context, role aggregate.Role) error {
	// Check if the role already exists
	var existingRole aggregate.Role
	err := rr.db.WithContext(*ctx).Where("role_id = ?", role.RoleID).First(&existingRole).Error

	if err == gorm.ErrRecordNotFound {
		// Role does not exist, create a new record
		return rr.db.WithContext(*ctx).Create(&role).Error
	} else if err != nil {
		// Other error occurred
		return err
	} else {
		// Role exists, update the existing record
		return rr.db.WithContext(*ctx).Model(&aggregate.Role{}).Where("role_id = ?", role.RoleID).Updates(&role).Error
	}
}

// FindRoleByID retrieves a role by its ID
// It returns the role and an error if the operation fails
// If the role is not found, it returns an error
func (rr *roleRepository) FindRoleByID(ctx *context.Context, roleID uuid.UUID) (aggregate.Role, error) {
	var role aggregate.Role
	err := rr.db.WithContext(*ctx).Where("role_id = ?", roleID).First(&role).Error
	return role, err
}

// FindRoleByName retrieves roles by their name
// It returns a list of roles and an error if the operation fails
// If no roles are found, it returns an empty list and no error
func (rr *roleRepository) FindRoleByName(ctx *context.Context, roleName string) ([]aggregate.Role, error) {
	var roles []aggregate.Role
	pattern := "%" + roleName + "%"
	err := rr.db.WithContext(*ctx).Where("role_name LIKE ?", pattern).Find(&roles).Error
	return roles, err
}

// FindRoleByCode retrieves roles by their code
// It returns a list of roles and an error if the operation fails
// If no roles are found, it returns an empty list and no error
func (rr *roleRepository) FindRoleByCode(ctx *context.Context, roleCode string) ([]aggregate.Role, error) {
	var roles []aggregate.Role
	pattern := "%" + roleCode + "%"
	err := rr.db.WithContext(*ctx).Where("role_code LIKE ?", pattern).Find(&roles).Error
	return roles, err
}

// FindAllRoles retrieves all roles
// It returns a list of roles and an error if the operation fails
// If no roles are found, it returns an empty list and no error
func (rr *roleRepository) FindAllRoles(ctx *context.Context) ([]aggregate.Role, error) {
	var roles []aggregate.Role
	err := rr.db.WithContext(*ctx).Find(&roles).Error
	return roles, err
}

// DeleteRole deletes a role by its ID
// It returns an error if the operation fails
func (rr *roleRepository) DeleteRole(ctx *context.Context, roleID uuid.UUID) error {
	return rr.db.WithContext(*ctx).Where("role_id = ?", roleID).Delete(&aggregate.Role{}).Error
}

// AssignPermission assigns a permission to a role
// It returns an error if the operation fails
func (rr *roleRepository) AssignPermission(ctx *context.Context, roleID uuid.UUID, permissionID uuid.UUID) error {
	// Create a role-permission relationship
	rolePermission := struct {
		RoleID       uuid.UUID `gorm:"column:role_id"`
		PermissionID uuid.UUID `gorm:"column:permission_id"`
	}{
		RoleID:       roleID,
		PermissionID: permissionID,
	}

	return rr.db.WithContext(*ctx).Table("systems.permission_role").Create(&rolePermission).Error
}

// RemovePermission removes a permission from a role
// It returns an error if the operation fails
func (rr *roleRepository) RemovePermission(ctx *context.Context, roleID uuid.UUID, permissionID uuid.UUID) error {
	// Delete the role-permission relationship
	return rr.db.WithContext(*ctx).Table("systems.permission_role").Where("role_id = ? AND permission_id = ?", roleID, permissionID).Delete(nil).Error
}

// GetRolePermissions retrieves all permissions assigned to a role
// It returns a list of permissions and an error if the operation fails
// If no permissions are found, it returns an empty list and no error
func (rr *roleRepository) GetRolePermissions(ctx *context.Context, roleID uuid.UUID) ([]aggregate.Permission, error) {
	var permissions []aggregate.Permission
	err := rr.db.WithContext(*ctx).Table("systems.permission").Joins("JOIN systems.permission_role ON systems.permission.permission_id = systems.permission_role.permission_id").Where("systems.permission_role.role_id = ?", roleID).Find(&permissions).Error
	return permissions, err
}

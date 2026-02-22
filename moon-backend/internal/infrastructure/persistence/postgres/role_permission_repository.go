package postgres

import (
	"context"
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// rolePermissionRepository implements the repository.RolePermissionRepository interface
type rolePermissionRepository struct {
	db *gorm.DB
}

// NewRolePermissionRepository creates a new RolePermissionRepository
func NewRolePermissionRepository(db *gorm.DB) *rolePermissionRepository {
	return &rolePermissionRepository{db: db}
}

// Create creates a new role-permission relationship
func (r *rolePermissionRepository) Create(ctx context.Context, rp aggregate.RolePermission) error {
	return r.db.WithContext(ctx).Create(&rp).Error
}

// Delete deletes an existing role-permission relationship
func (r *rolePermissionRepository) Delete(ctx context.Context, roleID, permissionID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("role_id = ? AND permission_id = ?", roleID, permissionID).Delete(&aggregate.RolePermission{}).Error
}

// GetByRoleAndPermission gets a role-permission relationship by role and permission IDs
func (r *rolePermissionRepository) GetByRoleAndPermission(ctx context.Context, roleID, permissionID uuid.UUID) (aggregate.RolePermission, error) {
	var rp aggregate.RolePermission
	err := r.db.WithContext(ctx).Where("role_id = ? AND permission_id = ?", roleID, permissionID).First(&rp).Error
	return rp, err
}

// GetPermissionsByRoleID gets all permissions for a role
func (r *rolePermissionRepository) GetPermissionsByRoleID(ctx context.Context, roleID uuid.UUID) ([]aggregate.Permission, error) {
	var permissions []aggregate.Permission
	err := r.db.WithContext(ctx).Table("systems.permissions").Select("systems.permissions.*").Joins("INNER JOIN systems.permission_role ON systems.permissions.id = systems.permission_role.permission_id").Where("systems.permission_role.role_id = ?", roleID).Find(&permissions).Error
	return permissions, err
}

// GetRolesByPermissionID gets all roles for a permission
func (r *rolePermissionRepository) GetRolesByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]aggregate.Role, error) {
	var roles []aggregate.Role
	err := r.db.WithContext(ctx).Table("systems.roles").Select("systems.roles.*").Joins("INNER JOIN systems.permission_role ON systems.roles.id = systems.permission_role.role_id").Where("systems.permission_role.permission_id = ?", permissionID).Find(&roles).Error
	return roles, err
}

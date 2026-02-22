package postgres

import (
	"context"
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// organizationRoleRepository implements the repository.OrganizationRoleRepository interface
type organizationRoleRepository struct {
	db *gorm.DB
}

// NewOrganizationRoleRepository creates a new OrganizationRoleRepository
func NewOrganizationRoleRepository(db *gorm.DB) *organizationRoleRepository {
	return &organizationRoleRepository{db: db}
}

// Create creates a new organization-role relationship
func (r *organizationRoleRepository) Create(ctx context.Context, or aggregate.OrganizationRole) error {
	return r.db.WithContext(ctx).Create(&or).Error
}

// Delete deletes an existing organization-role relationship
func (r *organizationRoleRepository) Delete(ctx context.Context, organizationID, roleID uuid.UUID) error {
	return r.db.WithContext(ctx).Where("organization_id = ? AND role_id = ?", organizationID, roleID).Delete(&aggregate.OrganizationRole{}).Error
}

// GetByOrganizationAndRole gets an organization-role relationship by organization and role IDs
func (r *organizationRoleRepository) GetByOrganizationAndRole(ctx context.Context, organizationID, roleID uuid.UUID) (aggregate.OrganizationRole, error) {
	var or aggregate.OrganizationRole
	err := r.db.WithContext(ctx).Where("organization_id = ? AND role_id = ?", organizationID, roleID).First(&or).Error
	return or, err
}

// GetRolesByOrganizationID gets all roles for an organization
func (r *organizationRoleRepository) GetRolesByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]aggregate.Role, error) {
	var roles []aggregate.Role
	err := r.db.WithContext(ctx).Table("systems.roles").Select("systems.roles.*").Joins("INNER JOIN systems.organization_role ON systems.roles.id = systems.organization_role.role_id").Where("systems.organization_role.organization_id = ?", organizationID).Find(&roles).Error
	return roles, err
}

// GetOrganizationsByRoleID gets all organizations for a role
func (r *organizationRoleRepository) GetOrganizationsByRoleID(ctx context.Context, roleID uuid.UUID) ([]aggregate.Organization, error) {
	var organizations []aggregate.Organization
	err := r.db.WithContext(ctx).Table("systems.organizations").Select("systems.organizations.*").Joins("INNER JOIN systems.organization_role ON systems.organizations.id = systems.organization_role.organization_id").Where("systems.organization_role.role_id = ?", roleID).Find(&organizations).Error
	return organizations, err
}
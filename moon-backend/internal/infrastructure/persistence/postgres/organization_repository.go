package postgres

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type organizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) repository.OrganizationRepository {
	return &organizationRepository{db}
}

func (or *organizationRepository) SaveOrganization(ctx context.Context, org aggregate.Organization) error {
	// 检查组织是否已存在
	var existingOrg aggregate.Organization
	err := or.db.WithContext(ctx).Table("systems.organization").Select("organization_id").Where("organization_id = ?", org.OrganizationID).First(&existingOrg).Error

	if err == gorm.ErrRecordNotFound {
		// 组织不存在，创建新记录
		// Create a map without ParentID to avoid database error
		orgData := map[string]interface{}{
			"organization_id":          org.OrganizationID,
			"organization_name":        org.OrganizationName,
			"organization_code":        org.OrganizationCode,
			"organization_description": org.OrganizationDescription,
			"organization_flag":        org.OrganizationFlag,
			"sensitive_flag":           org.SensitiveFlag,
			"created_at":               org.CreatedAt,
			"updated_at":               org.UpdatedAt,
		}
		return or.db.WithContext(ctx).Table("systems.organization").Create(orgData).Error
	} else if err != nil {
		// 其他错误
		return err
	} else {
		// 组织存在，更新记录
		org.UpdatedAt = time.Now()
		// Create a map without ParentID to avoid database error
		orgData := map[string]interface{}{
			"organization_name":        org.OrganizationName,
			"organization_code":        org.OrganizationCode,
			"organization_description": org.OrganizationDescription,
			"organization_flag":        org.OrganizationFlag,
			"sensitive_flag":           org.SensitiveFlag,
			"updated_at":               org.UpdatedAt,
		}
		return or.db.WithContext(ctx).Table("systems.organization").Where("organization_id = ?", org.OrganizationID).Updates(orgData).Error
	}
}

func (or *organizationRepository) FindOrganizationByID(ctx context.Context, org_id uuid.UUID) (aggregate.Organization, error) {
	var org = aggregate.Organization{}
	err := or.db.WithContext(ctx).Table("systems.organization").Select("organization_id, organization_name, organization_code, organization_description, organization_flag, sensitive_flag, created_at, updated_at").Where("organization_id = ?", org_id).First(&org).Error
	return org, err
}

func (or *organizationRepository) FindOrganizationByName(ctx context.Context, org_name string) ([]aggregate.Organization, error) {
	var org = []aggregate.Organization{}
	pattern := "%" + org_name + "%"
	err := or.db.WithContext(ctx).Table("systems.organization").Select("organization_id, organization_name, organization_code, organization_description, organization_flag, sensitive_flag, created_at, updated_at").Where("organization_name LIKE ?", pattern).Find(&org).Error
	return org, err
}

func (or *organizationRepository) FindOrganizationByCode(ctx context.Context, org_code string) ([]aggregate.Organization, error) {
	var org = []aggregate.Organization{}
	err := or.db.WithContext(ctx).Table("systems.organization").Select("organization_id, organization_name, organization_code, organization_description, organization_flag, sensitive_flag, created_at, updated_at").Where("organization_code LIKE ?", org_code+"%").Find(&org).Error
	return org, err
}

func (or *organizationRepository) FindAllOrganizations(ctx context.Context) ([]aggregate.Organization, error) {
	var org = []aggregate.Organization{}
	err := or.db.WithContext(ctx).Table("systems.organization").Select("organization_id, organization_name, organization_code, organization_description, organization_flag, sensitive_flag, created_at, updated_at").Find(&org).Error
	return org, err
}

func (or *organizationRepository) FindOrganizationsByParentID(ctx context.Context, parent_id uuid.UUID) ([]aggregate.Organization, error) {
	var orgs = []aggregate.Organization{}
	err := or.db.WithContext(ctx).Table("systems.organization").Select("organization_id, organization_name, organization_code, organization_description, organization_flag, sensitive_flag, created_at, updated_at").Where("parent_id = ?", parent_id).Find(&orgs).Error
	return orgs, err
}

func (or *organizationRepository) UpdateOrganizationParent(ctx context.Context, org_id uuid.UUID, parent_id *uuid.UUID) error {
	// 暂时不实现此方法，因为数据库中没有 parent_id 字段
	return nil
}

func (or *organizationRepository) AssignRoleToOrganization(ctx context.Context, org_id uuid.UUID, role_id uuid.UUID) error {
	// 检查关联是否已存在
	var count int64
	err := or.db.WithContext(ctx).Table("systems.organization_role").Where("organization_id = ? AND role_id = ?", org_id, role_id).Count(&count).Error
	if err != nil {
		return err
	}

	// 如果关联不存在，创建新关联
	if count == 0 {
		organizationRole := map[string]interface{}{
			"organization_id": org_id,
			"role_id":         role_id,
		}
		return or.db.WithContext(ctx).Table("systems.organization_role").Create(organizationRole).Error
	}

	return nil
}

func (or *organizationRepository) RemoveRoleFromOrganization(ctx context.Context, org_id uuid.UUID, role_id uuid.UUID) error {
	return or.db.WithContext(ctx).Table("systems.organization_role").Where("organization_id = ? AND role_id = ?", org_id, role_id).Delete(nil).Error
}

func (or *organizationRepository) FindRolesByOrganizationID(ctx context.Context, org_id uuid.UUID) ([]aggregate.Role, error) {
	var roles = []aggregate.Role{}
	err := or.db.WithContext(ctx).Table("systems.role").Select("role_id, role_name, role_code, role_description, role_flag, sensitive_flag, created_at, updated_at").Joins("JOIN systems.organization_role ON systems.role.role_id = systems.organization_role.role_id").Where("systems.organization_role.organization_id = ?", org_id).Find(&roles).Error
	return roles, err
}

func (or *organizationRepository) DeleteOrganization(ctx context.Context, org_id uuid.UUID) error {
	// First, delete any organization-role associations
	err := or.db.WithContext(ctx).Table("systems.organization_role").Where("organization_id = ?", org_id).Delete(nil).Error
	if err != nil {
		return err
	}
	
	// Then, delete the organization itself
	return or.db.WithContext(ctx).Table("systems.organization").Where("organization_id = ?", org_id).Delete(nil).Error
}

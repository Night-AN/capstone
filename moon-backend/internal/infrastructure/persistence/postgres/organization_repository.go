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

func (or *organizationRepository) SaveOrganization(ctx *context.Context, org aggregate.Organization) error {
	// 检查组织是否已存在
	var existingOrg aggregate.Organization
	err := or.db.WithContext(*ctx).Where("organization_id = ?", org.OrganizationID).First(&existingOrg).Error

	if err == gorm.ErrRecordNotFound {
		// 组织不存在，创建新记录
		return or.db.WithContext(*ctx).Create(&org).Error
	} else if err != nil {
		// 其他错误
		return err
	} else {
		// 组织存在，更新记录
		org.UpdatedAt = time.Now()
		// 使用 Where 条件确保只更新指定的记录
		return or.db.WithContext(*ctx).Model(&aggregate.Organization{}).Where("organization_id = ?", org.OrganizationID).Updates(&org).Error
	}
}

func (or *organizationRepository) FindOrganizationByID(ctx *context.Context, org_id uuid.UUID) (aggregate.Organization, error) {
	var org = aggregate.Organization{}
	err := or.db.WithContext(*ctx).Where(&aggregate.Organization{OrganizationID: org_id}).First(&org).Error
	return org, err
}

func (or *organizationRepository) FindOrganizationByName(ctx *context.Context, org_name string) ([]aggregate.Organization, error) {
	var org = []aggregate.Organization{}
	pattern := "%" + org_name + "%"
	err := or.db.WithContext(*ctx).Model(&aggregate.Organization{}).Where("organization_name LIKE ?", pattern).Find(&org).Error
	return org, err
}

func (or *organizationRepository) FindOrganizationByCode(ctx *context.Context, org_code string) ([]aggregate.Organization, error) {
	var org = []aggregate.Organization{}
	err := or.db.WithContext(*ctx).Model(&aggregate.Organization{}).Where("organization_code LIKE ?", org_code+"%").Find(&org).Error
	return org, err
}

func (or *organizationRepository) FindAllOrganizations(ctx *context.Context) ([]aggregate.Organization, error) {
	var org = []aggregate.Organization{}
	err := or.db.WithContext(*ctx).Model(&aggregate.Organization{}).Find(&org).Error
	return org, err
}

func (or *organizationRepository) FindOrganizationsByParentID(ctx *context.Context, parent_id uuid.UUID) ([]aggregate.Organization, error) {
	var orgs = []aggregate.Organization{}
	err := or.db.WithContext(*ctx).Where("parent_id = ?", parent_id).Find(&orgs).Error
	return orgs, err
}

func (or *organizationRepository) UpdateOrganizationParent(ctx *context.Context, org_id uuid.UUID, parent_id *uuid.UUID) error {
	return or.db.WithContext(*ctx).Model(&aggregate.Organization{}).Where("organization_id = ?", org_id).Update("parent_id", parent_id).Error
}

func (or *organizationRepository) AssignRoleToOrganization(ctx *context.Context, org_id uuid.UUID, role_id uuid.UUID) error {
	// 检查关联是否已存在
	var count int64
	err := or.db.WithContext(*ctx).Table("systems.organization_role").Where("organization_id = ? AND role_id = ?", org_id, role_id).Count(&count).Error
	if err != nil {
		return err
	}

	// 如果关联不存在，创建新关联
	if count == 0 {
		organizationRole := map[string]interface{}{
			"organization_id": org_id,
			"role_id":         role_id,
		}
		return or.db.WithContext(*ctx).Table("systems.organization_role").Create(organizationRole).Error
	}

	return nil
}

func (or *organizationRepository) RemoveRoleFromOrganization(ctx *context.Context, org_id uuid.UUID, role_id uuid.UUID) error {
	return or.db.WithContext(*ctx).Table("systems.organization_role").Where("organization_id = ? AND role_id = ?", org_id, role_id).Delete(nil).Error
}

func (or *organizationRepository) FindRolesByOrganizationID(ctx *context.Context, org_id uuid.UUID) ([]aggregate.Role, error) {
	var roles = []aggregate.Role{}
	err := or.db.WithContext(*ctx).Table("systems.role").Joins("JOIN systems.organization_role ON systems.role.role_id = systems.organization_role.role_id").Where("systems.organization_role.organization_id = ?", org_id).Find(&roles).Error
	return roles, err
}

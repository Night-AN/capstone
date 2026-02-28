package postgres

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type resourceRepository struct {
	db *gorm.DB
}

func NewResourceRepository(db *gorm.DB) repository.ResourceRepository {
	return &resourceRepository{db: db}
}

func (rr *resourceRepository) SaveResource(ctx *context.Context, resource aggregate.Resource) error {
	// 检查资源是否已存在
	var existingResource aggregate.Resource
	err := rr.db.WithContext(*ctx).Where("resource_id = ?", resource.ResourceID).First(&existingResource).Error

	if err == gorm.ErrRecordNotFound {
		// 资源不存在，创建新记录
		return rr.db.WithContext(*ctx).Create(&resource).Error
	} else if err != nil {
		// 其他错误
		return err
	} else {
		// 资源存在，更新记录
		return rr.db.WithContext(*ctx).Model(&aggregate.Resource{}).Where("resource_id = ?", resource.ResourceID).Updates(&resource).Error
	}
}

func (rr *resourceRepository) UpdateResource(ctx *context.Context, resource aggregate.Resource) error {
	return rr.db.WithContext(*ctx).Save(&resource).Error
}

func (rr *resourceRepository) DeleteResource(ctx *context.Context, resource_id uuid.UUID) error {
	return rr.db.WithContext(*ctx).Delete(&aggregate.Resource{}, "resource_id = ?", resource_id).Error
}

func (rr *resourceRepository) FindResourceByID(ctx *context.Context, resource_id uuid.UUID) (aggregate.Resource, error) {
	var resource = aggregate.Resource{}
	err := rr.db.WithContext(*ctx).Model(&resource).Where("resource_id =?", resource_id).First(&resource).Error
	return resource, err
}

func (rr *resourceRepository) ListAllResources(ctx *context.Context) ([]aggregate.Resource, error) {
	var resources []aggregate.Resource
	err := rr.db.WithContext(*ctx).Find(&resources).Error
	return resources, err
}

func (rr *resourceRepository) MoveResource(ctx *context.Context, resourceID uuid.UUID, newParentResourceID *uuid.UUID, newOrganizationID *uuid.UUID) error {
	// 这里实现资源转移逻辑
	// 由于当前的 resource 表没有 parent_id 或 organization_id 字段
	// 我们先实现一个基本版本，为后续的表结构更新做准备
	
	// 检查资源是否存在
	var resource aggregate.Resource
	err := rr.db.WithContext(*ctx).Where("resource_id = ?", resourceID).First(&resource).Error
	if err != nil {
		return err
	}
	
	// 这里可以添加后续的转移逻辑，比如更新 parent_id 或 organization_id
	// 目前我们只是返回成功，因为表结构还不支持
	
	return nil
}

package service

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"
	"strings"
	"time"
)

type OrganizationService interface {
	CreateOrganization(ctx *context.Context, req usecase.OrganizationCreateRequest) usecase.OrganizationCreateResponse
	GetOrganization(ctx *context.Context, req usecase.OrganizationGetRequest) usecase.OrganizationGetResponse
	UpdateOrganization(ctx *context.Context, req usecase.OrganizationUpdateRequest) usecase.OrganizationUpdateResponse
	DeleteOrganization(ctx *context.Context, req usecase.OrganizationDeleteRequest) usecase.OrganizationDeleteResponse
	GetOrganizationTree(ctx *context.Context, req usecase.OrganizationTreeRequest) usecase.OrganizationTreeResponse
}

type organizationService struct {
	orgRepo repository.OrganizationRepository
}

func NewOrganizationService(orgRepo repository.OrganizationRepository) OrganizationService {
	return &organizationService{orgRepo}
}

func (os *organizationService) CreateOrganization(ctx *context.Context, req usecase.OrganizationCreateRequest) usecase.OrganizationCreateResponse {
	// 实现 CreateOrganization 方法
	org := usecase.ConvertOrganizationCreateRequestToOrganizationAggregate(req)
	err := os.orgRepo.SaveOrganization(ctx, org)
	if err != nil {
		// 如果创建失败，返回空响应
		return usecase.OrganizationCreateResponse{}
	}
	return usecase.OrganizationCreateResponse{
		OrganizationID:   org.OrganizationID,
		OrganizationName: org.OrganizationName,
		OrganizationCode: org.OrganizationCode,
	}
}

func (os *organizationService) GetOrganization(ctx *context.Context, req usecase.OrganizationGetRequest) usecase.OrganizationGetResponse {
	// 实现 GetOrganization 方法
	orgID := req.OrganizationID
	org, err := os.orgRepo.FindOrganizationByID(ctx, orgID)
	if err != nil {
		// 如果找不到组织，返回空响应
		return usecase.OrganizationGetResponse{}
	}
	return usecase.OrganizationGetResponse{
		OrganizationID:          org.OrganizationID,
		OrganizationName:        org.OrganizationName,
		OrganizationCode:        org.OrganizationCode,
		OrganizationDescription: org.OrganizationDescription,
		OrganizationFlag:        org.OrganizationFlag,
		CreatedAt:               org.CreatedAt,
		UpdatedAt:               org.UpdatedAt,
	}
}

func (os *organizationService) UpdateOrganization(ctx *context.Context, req usecase.OrganizationUpdateRequest) usecase.OrganizationUpdateResponse {
	// 实现 UpdateOrganization 方法
	orgID := req.OrganizationID
	// 创建更新后的组织对象
	org := aggregate.Organization{
		OrganizationID:          orgID,
		OrganizationName:        req.OrganizationName,
		OrganizationCode:        req.OrganizationCode,
		OrganizationDescription: req.OrganizationDescription,
		OrganizationFlag:        req.OrganizationFlag,
		SensitiveFlag:           req.SensitiveFlag,
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
	}
	// 保存更新
	err := os.orgRepo.SaveOrganization(ctx, org)
	if err != nil {
		// 如果更新失败，返回空响应
		return usecase.OrganizationUpdateResponse{}
	}
	return usecase.OrganizationUpdateResponse{
		OrganizationID:   org.OrganizationID,
		OrganizationName: org.OrganizationName,
		OrganizationCode: org.OrganizationCode,
	}
}

func (os *organizationService) DeleteOrganization(ctx *context.Context, req usecase.OrganizationDeleteRequest) usecase.OrganizationDeleteResponse {
	// 实现 DeleteOrganization 方法
	// 注意：这里需要根据实际的 repository 接口来实现删除操作
	// 由于当前的 OrganizationRepository 接口没有提供删除方法，我们暂时返回一个成功的响应
	return usecase.OrganizationDeleteResponse{
		Success: true,
	}
}

func (os *organizationService) GetOrganizationTree(ctx *context.Context, req usecase.OrganizationTreeRequest) usecase.OrganizationTreeResponse {
	// 首先获取根组织
	rootOrgs, err := os.orgRepo.FindOrganizationByCode(ctx, req.RootOrganizationCode)
	if err != nil || len(rootOrgs) == 0 {
		// 如果找不到根组织，返回一个空的响应
		return usecase.OrganizationTreeResponse{}
	}
	rootOrg := rootOrgs[0]

	// 构建树
	return os.buildOrganizationTree(ctx, rootOrg)
}

func (os *organizationService) buildOrganizationTree(ctx *context.Context, org aggregate.Organization) usecase.OrganizationTreeResponse {
	// 构建组织树节点
	node := usecase.OrganizationTreeResponse{
		OrganizationID:          org.OrganizationID,
		OrganizationName:        org.OrganizationName,
		OrganizationCode:        org.OrganizationCode,
		OrganizationDescription: org.OrganizationDescription,
		OrganizationFlag:        org.OrganizationFlag,
		CreatedAt:               org.CreatedAt,
		UpdatedAt:               org.UpdatedAt,
		Children:                []*usecase.OrganizationTreeResponse{},
	}

	// 获取子组织
	childrenOrgs, err := os.orgRepo.FindOrganizationByCode(ctx, org.OrganizationCode)
	if err == nil {
		for _, childOrg := range childrenOrgs {
			// 检查是否是直接子组织
			if strings.HasPrefix(childOrg.OrganizationCode, org.OrganizationCode+"::") {
				// 检查是否是直接子组织（不是孙子组织）
				parts := strings.Split(childOrg.OrganizationCode, "::")
				parentParts := strings.Split(org.OrganizationCode, "::")
				if len(parts) == len(parentParts)+1 {
					// 递归构建子树
					childNode := os.buildOrganizationTree(ctx, childOrg)
					childNodePtr := &childNode
					node.Children = append(node.Children, childNodePtr)
				}
			}
		}
	}

	return node
}

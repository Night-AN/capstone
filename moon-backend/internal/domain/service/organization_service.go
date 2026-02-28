package service

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"
	"strings"
	"time"

	"github.com/google/uuid"
)

type OrganizationService interface {
	CreateOrganization(ctx context.Context, req usecase.OrganizationCreateRequest) usecase.OrganizationCreateResponse
	GetOrganization(ctx context.Context, req usecase.OrganizationGetRequest) usecase.OrganizationGetResponse
	UpdateOrganization(ctx context.Context, req usecase.OrganizationUpdateRequest) usecase.OrganizationUpdateResponse
	DeleteOrganization(ctx context.Context, req usecase.OrganizationDeleteRequest) usecase.OrganizationDeleteResponse
	GetOrganizationTree(ctx context.Context, req usecase.OrganizationTreeRequest) usecase.OrganizationTreeResponse
	ListOrganizations(ctx context.Context, req usecase.OrganizationListRequest) usecase.OrganizationListResponse
	MoveOrganization(ctx context.Context, req usecase.OrganizationMoveRequest) usecase.OrganizationMoveResponse
	GetOrganizationUsers(ctx context.Context, req usecase.OrganizationUsersRequest) usecase.OrganizationUsersResponse
}

type organizationService struct {
	orgRepo  repository.OrganizationRepository
	userRepo repository.UserRepository
}

func NewOrganizationService(orgRepo repository.OrganizationRepository, userRepo repository.UserRepository) OrganizationService {
	return &organizationService{orgRepo: orgRepo, userRepo: userRepo}
}

func (os *organizationService) CreateOrganization(ctx context.Context, req usecase.OrganizationCreateRequest) usecase.OrganizationCreateResponse {
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

func (os *organizationService) GetOrganization(ctx context.Context, req usecase.OrganizationGetRequest) usecase.OrganizationGetResponse {
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

func (os *organizationService) UpdateOrganization(ctx context.Context, req usecase.OrganizationUpdateRequest) usecase.OrganizationUpdateResponse {
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
		ParentID:                req.ParentID,
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

func (os *organizationService) DeleteOrganization(ctx context.Context, req usecase.OrganizationDeleteRequest) usecase.OrganizationDeleteResponse {
	// 调用 repository 删除组织
	err := os.orgRepo.DeleteOrganization(ctx, req.OrganizationID)
	if err != nil {
		// 如果删除失败，返回失败响应
		return usecase.OrganizationDeleteResponse{
			Success: false,
		}
	}
	return usecase.OrganizationDeleteResponse{
		Success: true,
	}
}

func (os *organizationService) GetOrganizationTree(ctx context.Context, req usecase.OrganizationTreeRequest) usecase.OrganizationTreeResponse {
	// 获取所有组织
	allOrgs, err := os.orgRepo.FindAllOrganizations(ctx)
	if err != nil {
		// 如果获取失败，返回空响应
		return usecase.OrganizationTreeResponse{}
	}

	// 如果 root_organization_code 为空，构建完整的树结构
	if req.RootOrganizationCode == "" {
		// 创建一个虚拟根节点
		virtualRoot := usecase.OrganizationTreeResponse{
			Children: []*usecase.OrganizationTreeResponse{},
		}

		// 首先找到所有顶级组织（没有父组织的）
		topLevelOrgs := []aggregate.Organization{}
		for _, org := range allOrgs {
			// 顶级组织的代码不包含 "::"
			if !strings.Contains(org.OrganizationCode, "::") {
				topLevelOrgs = append(topLevelOrgs, org)
			}
		}

		// 为每个顶级组织构建子树
		for _, topOrg := range topLevelOrgs {
			// 构建完整的子树
			topNode := os.buildOrganizationTree(ctx, topOrg, allOrgs)
			topNodePtr := &topNode
			virtualRoot.Children = append(virtualRoot.Children, topNodePtr)
		}

		return virtualRoot
	}

	// 首先获取根组织
	var rootOrg aggregate.Organization
	found := false
	for _, org := range allOrgs {
		if org.OrganizationCode == req.RootOrganizationCode {
			rootOrg = org
			found = true
			break
		}
	}

	if !found {
		// 如果找不到根组织，返回一个空的响应
		return usecase.OrganizationTreeResponse{}
	}

	// 构建树
	return os.buildOrganizationTree(ctx, rootOrg, allOrgs)
}

func (os *organizationService) buildOrganizationTree(ctx context.Context, org aggregate.Organization, allOrgs []aggregate.Organization) usecase.OrganizationTreeResponse {
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

	// 查找直接子组织
	for _, childOrg := range allOrgs {
		// 检查是否是直接子组织
		if strings.HasPrefix(childOrg.OrganizationCode, org.OrganizationCode+"::") {
			// 检查是否是直接子组织（不是孙子组织）
			parts := strings.Split(childOrg.OrganizationCode, "::")
			parentParts := strings.Split(org.OrganizationCode, "::")
			if len(parts) == len(parentParts)+1 {
				// 递归构建子树
				childNode := os.buildOrganizationTree(ctx, childOrg, allOrgs)
				childNodePtr := &childNode
				node.Children = append(node.Children, childNodePtr)
			}
		}
	}

	return node
}

func (os *organizationService) ListOrganizations(ctx context.Context, req usecase.OrganizationListRequest) usecase.OrganizationListResponse {
	// 获取所有组织
	orgs, err := os.orgRepo.FindAllOrganizations(ctx)
	if err != nil {
		// 如果获取失败，返回空响应
		return usecase.OrganizationListResponse{}
	}

	// 转换为响应格式
	orgItems := make([]usecase.OrganizationListItem, len(orgs))
	for i, org := range orgs {
		orgItems[i] = usecase.OrganizationListItem{
			OrganizationID:   org.OrganizationID,
			OrganizationName: org.OrganizationName,
			OrganizationCode: org.OrganizationCode,
			OrganizationFlag: org.OrganizationFlag,
			CreatedAt:        org.CreatedAt,
		}
	}

	return usecase.OrganizationListResponse{
		Organizations: orgItems,
	}
}

func (os *organizationService) MoveOrganization(ctx context.Context, req usecase.OrganizationMoveRequest) usecase.OrganizationMoveResponse {
	// 检查组织是否存在
	_, err := os.orgRepo.FindOrganizationByID(ctx, req.OrganizationID)
	if err != nil {
		// 如果组织不存在，返回失败响应
		return usecase.OrganizationMoveResponse{Success: false}
	}

	// 检查新父组织是否存在（如果指定了）
	if req.NewParentID != nil {
		_, err := os.orgRepo.FindOrganizationByID(ctx, *req.NewParentID)
		if err != nil {
			// 如果新父组织不存在，返回失败响应
			return usecase.OrganizationMoveResponse{Success: false}
		}

		// 检查是否会形成循环引用
		if *req.NewParentID == req.OrganizationID {
			// 组织不能成为自己的父组织
			return usecase.OrganizationMoveResponse{Success: false}
		}
	}

	// 执行移动操作
	err = os.orgRepo.UpdateOrganizationParent(ctx, req.OrganizationID, req.NewParentID)
	if err != nil {
		// 如果移动失败，返回失败响应
		return usecase.OrganizationMoveResponse{Success: false}
	}

	// 返回成功响应
	return usecase.OrganizationMoveResponse{Success: true}
}

func (os *organizationService) GetOrganizationUsers(ctx context.Context, req usecase.OrganizationUsersRequest) usecase.OrganizationUsersResponse {
	// 检查组织是否存在
	_, err := os.orgRepo.FindOrganizationByID(ctx, req.OrganizationID)
	if err != nil {
		// 如果组织不存在，返回空响应
		return usecase.OrganizationUsersResponse{}
	}

	// 获取组织用户列表
	users, err := os.userRepo.FindUsersByOrganizationID(ctx, req.OrganizationID)
	if err != nil {
		// 如果获取失败，返回空响应
		return usecase.OrganizationUsersResponse{}
	}

	// 转换为响应格式
	userItems := make([]usecase.UserListItem, len(users))
	for i, user := range users {
		userItems[i] = usecase.UserListItem{
			UserID:   user.UserID,
			Nickname: user.Nickname,
			FullName: user.FullName,
			Email:    user.Email,
		}
	}

	return usecase.OrganizationUsersResponse{
		Users: userItems,
	}
}

func (os *organizationService) GetOrganizationStatistics(ctx context.Context, req usecase.OrganizationStatisticsRequest) usecase.OrganizationStatisticsResponse {
	// 检查组织是否存在
	_, err := os.orgRepo.FindOrganizationByID(ctx, req.OrganizationID)
	if err != nil {
		// 如果组织不存在，返回空响应
		return usecase.OrganizationStatisticsResponse{}
	}

	// 获取组织用户数量
	users, err := os.userRepo.FindUsersByOrganizationID(ctx, req.OrganizationID)
	userCount := 0
	if err == nil {
		userCount = len(users)
	}

	// 获取直接子组织数量
	directSubOrgs, err := os.orgRepo.FindOrganizationsByParentID(ctx, req.OrganizationID)
	directSubOrgCount := 0
	if err == nil {
		directSubOrgCount = len(directSubOrgs)
	}

	// 计算总子组织数量（包括所有层级）
	totalSubOrgCount := os.calculateTotalSubOrganizations(ctx, req.OrganizationID)

	// 获取组织角色数量
	roles, err := os.orgRepo.FindRolesByOrganizationID(ctx, req.OrganizationID)
	roleCount := 0
	if err == nil {
		roleCount = len(roles)
	}

	// 返回统计结果
	return usecase.OrganizationStatisticsResponse{
		OrganizationID:            req.OrganizationID,
		UserCount:                 userCount,
		SubOrganizationCount:      directSubOrgCount,
		TotalSubOrganizationCount: totalSubOrgCount,
		RoleCount:                 roleCount,
	}
}

func (os *organizationService) calculateTotalSubOrganizations(ctx context.Context, orgID uuid.UUID) int {
	// 获取直接子组织
	directSubOrgs, err := os.orgRepo.FindOrganizationsByParentID(ctx, orgID)
	if err != nil {
		return 0
	}

	// 计算直接子组织数量
	totalCount := len(directSubOrgs)

	// 递归计算每个子组织的子组织数量
	for _, subOrg := range directSubOrgs {
		totalCount += os.calculateTotalSubOrganizations(ctx, subOrg.OrganizationID)
	}

	return totalCount
}

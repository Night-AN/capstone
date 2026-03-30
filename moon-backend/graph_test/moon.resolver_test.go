package graph_test

import (
	"context"
	"moon/graph"
	"moon/pkg/graphqltest"
	"testing"

	"github.com/google/uuid"
)

// createTestResolver 创建一个测试用的 Resolver
func createTestResolver() *graph.Resolver {
	// 获取测试 ent 客户端
	entClient := graphqltest.NewTestEntClient()

	// 获取测试 S3 客户端
	s3Client := graphqltest.NewTestS3Client()

	// 创建并返回 Resolver
	return &graph.Resolver{
		Client:   entClient,
		S3Client: s3Client,
	}
}

func TestMutationResolver_Login(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试登录功能
	// 这里需要先创建用户，然后再测试登录
	// 由于这是单元测试，我们可以使用mock数据或直接测试错误情况
	_, err := resolver.Mutation().Login(context.Background(), "user3@example.com", "user3")
	if err == nil {
		t.Error("Login should fail for non-existent user")
	}
}

func TestMutationResolver_Register(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试注册功能
	token, err := resolver.Mutation().Register(context.Background(), "test_register@example.com", "password")
	if err != nil {
		t.Errorf("Register failed: %v", err)
	}
	if token == "" {
		t.Error("Register should return a token")
	}
}

func TestQueryResolver_Profile(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取用户资料
	// 由于需要有效的用户ID和认证，我们可以测试错误情况
	_, err := resolver.Query().Profile(context.Background(), uuid.New())
	if err == nil {
		t.Error("Profile should fail for non-authenticated user")
	}
}

func TestQueryResolver_Organizations(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取组织列表
	_, err := resolver.Query().Organizations(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("Organizations query failed: %v", err)
	}
}

func TestQueryResolver_Users(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取用户列表
	_, err := resolver.Query().Users(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("Users query failed: %v", err)
	}
}

func TestQueryResolver_Roles(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取角色列表
	_, err := resolver.Query().Roles(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("Roles query failed: %v", err)
	}
}

func TestQueryResolver_Permissions(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取权限列表
	_, err := resolver.Query().Permissions(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("Permissions query failed: %v", err)
	}
}

func TestQueryResolver_Files(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取文件列表
	_, err := resolver.Query().Files(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("Files query failed: %v", err)
	}
}

func TestQueryResolver_ProcurementPlanTypes(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取采购计划类型列表
	_, err := resolver.Query().ProcurementPlanTypes(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("ProcurementPlanTypes query failed: %v", err)
	}
}

func TestQueryResolver_ProcurementPlans(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取采购计划列表
	_, err := resolver.Query().ProcurementPlans(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("ProcurementPlans query failed: %v", err)
	}
}

func TestQueryResolver_ProcurementImplementations(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取采购实施列表
	_, err := resolver.Query().ProcurementImplementations(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("ProcurementImplementations query failed: %v", err)
	}
}

func TestQueryResolver_ProcurementExperts(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取采购专家列表
	_, err := resolver.Query().ProcurementExperts(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("ProcurementExperts query failed: %v", err)
	}
}

func TestQueryResolver_ProcurementAcceptances(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取采购验收列表
	_, err := resolver.Query().ProcurementAcceptances(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("ProcurementAcceptances query failed: %v", err)
	}
}

func TestQueryResolver_ProcurementReviews(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取采购评审列表
	_, err := resolver.Query().ProcurementReviews(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("ProcurementReviews query failed: %v", err)
	}
}

func TestQueryResolver_Assets(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取资产列表
	_, err := resolver.Query().Assets(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("Assets query failed: %v", err)
	}
}

func TestQueryResolver_AssetCategories(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取资产分类列表
	_, err := resolver.Query().AssetCategories(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("AssetCategories query failed: %v", err)
	}
}

func TestQueryResolver_AssetTypes(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取资产类型列表
	_, err := resolver.Query().AssetTypes(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("AssetTypes query failed: %v", err)
	}
}

func TestQueryResolver_LlmModels(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取LLM模型列表
	_, err := resolver.Query().LlmModels(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("LlmModels query failed: %v", err)
	}
}

func TestQueryResolver_LlmTokenUsages(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取LLM Token使用记录列表
	_, err := resolver.Query().LlmTokenUsages(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("LlmTokenUsages query failed: %v", err)
	}
}

func TestQueryResolver_ProcurementFraudRisks(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取采购欺诈风险记录列表
	_, err := resolver.Query().ProcurementFraudRisks(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("ProcurementFraudRisks query failed: %v", err)
	}
}

func TestQueryResolver_LlmProcurementAnalyses(t *testing.T) {
	// 获取测试客户端
	resolver := createTestResolver()

	// 测试获取LLM采购分析记录列表
	_, err := resolver.Query().LlmProcurementAnalyses(context.Background(), nil, nil, nil, nil, nil)
	if err != nil {
		t.Errorf("LlmProcurementAnalyses query failed: %v", err)
	}
}

package service_test

import (
	"moon/internal/domain/usecase"
	"testing"

	"github.com/google/uuid"
)

func TestCreateOrganization(t *testing.T) {
	// 创建测试数据
	req := usecase.OrganizationCreateRequest{
		OrganizationName:        "Test Organization",
		OrganizationCode:        "test",
		OrganizationDescription: "A test organization",
		OrganizationFlag:        "active",
		SensitiveFlag:           false,
	}

	// 调用 CreateOrganization 方法
	resp := organizationSvc.CreateOrganization(testCtx, req)

	// 验证响应
	if resp.OrganizationID == uuid.Nil {
		t.Errorf("Expected non-nil OrganizationID, got %v", resp.OrganizationID)
	}
	if resp.OrganizationName != req.OrganizationName {
		t.Errorf("Expected OrganizationName %s, got %s", req.OrganizationName, resp.OrganizationName)
	}
	if resp.OrganizationCode != req.OrganizationCode {
		t.Errorf("Expected OrganizationCode %s, got %s", req.OrganizationCode, resp.OrganizationCode)
	}
}

func TestGetOrganization(t *testing.T) {
	// 首先创建一个组织
	createReq := usecase.OrganizationCreateRequest{
		OrganizationName:        "Test Organization",
		OrganizationCode:        "test",
		OrganizationDescription: "A test organization",
		OrganizationFlag:        "active",
		SensitiveFlag:           false,
	}
	createResp := organizationSvc.CreateOrganization(testCtx, createReq)

	// 然后获取这个组织
	getReq := usecase.OrganizationGetRequest{
		OrganizationID: createResp.OrganizationID,
	}
	getResp := organizationSvc.GetOrganization(testCtx, getReq)

	// 验证响应
	if getResp.OrganizationID != createResp.OrganizationID {
		t.Errorf("Expected OrganizationID %v, got %v", createResp.OrganizationID, getResp.OrganizationID)
	}
	if getResp.OrganizationName != createReq.OrganizationName {
		t.Errorf("Expected OrganizationName %s, got %s", createReq.OrganizationName, getResp.OrganizationName)
	}
	if getResp.OrganizationCode != createReq.OrganizationCode {
		t.Errorf("Expected OrganizationCode %s, got %s", createReq.OrganizationCode, getResp.OrganizationCode)
	}
}

func TestUpdateOrganization(t *testing.T) {
	// 首先创建一个组织
	createReq := usecase.OrganizationCreateRequest{
		OrganizationName:        "Test Organization",
		OrganizationCode:        "test",
		OrganizationDescription: "A test organization",
		OrganizationFlag:        "active",
		SensitiveFlag:           false,
	}
	createResp := organizationSvc.CreateOrganization(testCtx, createReq)

	// 然后更新这个组织
	updateReq := usecase.OrganizationUpdateRequest{
		OrganizationID:          createResp.OrganizationID,
		OrganizationName:        "Updated Test Organization",
		OrganizationCode:        "test-updated",
		OrganizationDescription: "An updated test organization",
		OrganizationFlag:        "inactive",
		SensitiveFlag:           true,
	}
	updateResp := organizationSvc.UpdateOrganization(testCtx, updateReq)

	// 验证响应
	if updateResp.OrganizationID != createResp.OrganizationID {
		t.Errorf("Expected OrganizationID %v, got %v", createResp.OrganizationID, updateResp.OrganizationID)
	}
	if updateResp.OrganizationName != updateReq.OrganizationName {
		t.Errorf("Expected OrganizationName %s, got %s", updateReq.OrganizationName, updateResp.OrganizationName)
	}
	if updateResp.OrganizationCode != updateReq.OrganizationCode {
		t.Errorf("Expected OrganizationCode %s, got %s", updateReq.OrganizationCode, updateResp.OrganizationCode)
	}
}

func TestDeleteOrganization(t *testing.T) {
	// 首先创建一个组织
	createReq := usecase.OrganizationCreateRequest{
		OrganizationName:        "Test Organization",
		OrganizationCode:        "test",
		OrganizationDescription: "A test organization",
		OrganizationFlag:        "active",
		SensitiveFlag:           false,
	}
	createResp := organizationSvc.CreateOrganization(testCtx, createReq)

	// 然后删除这个组织
	deleteReq := usecase.OrganizationDeleteRequest{
		OrganizationID: createResp.OrganizationID,
	}
	deleteResp := organizationSvc.DeleteOrganization(testCtx, deleteReq)

	// 验证响应
	if !deleteResp.Success {
		t.Errorf("Expected Success to be true, got false")
	}
}

func TestGetOrganizationTree(t *testing.T) {
	// 首先创建根组织
	rootReq := usecase.OrganizationCreateRequest{
		OrganizationName:        "Root Organization",
		OrganizationCode:        "root",
		OrganizationDescription: "Root organization",
		OrganizationFlag:        "active",
		SensitiveFlag:           false,
	}
	rootResp := organizationSvc.CreateOrganization(testCtx, rootReq)

	// 然后创建子组织
	childReq := usecase.OrganizationCreateRequest{
		OrganizationName:        "Child Organization",
		OrganizationCode:        "root::child",
		OrganizationDescription: "Child organization",
		OrganizationFlag:        "active",
		SensitiveFlag:           false,
	}
	childResp := organizationSvc.CreateOrganization(testCtx, childReq)

	// 然后获取组织树
	treeReq := usecase.OrganizationTreeRequest{
		RootOrganizationCode: "root",
	}
	treeResp := organizationSvc.GetOrganizationTree(testCtx, treeReq)

	// 验证响应
	if treeResp.OrganizationID != rootResp.OrganizationID {
		t.Errorf("Expected OrganizationID %v, got %v", rootResp.OrganizationID, treeResp.OrganizationID)
	}
	if len(treeResp.Children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(treeResp.Children))
	}
	if len(treeResp.Children) > 0 && treeResp.Children[0].OrganizationID != childResp.OrganizationID {
		t.Errorf("Expected Child OrganizationID %v, got %v", childResp.OrganizationID, treeResp.Children[0].OrganizationID)
	}
}

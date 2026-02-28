package service_test

import (
	"moon/internal/domain/usecase"
	"testing"

	"github.com/google/uuid"
)

func TestCreateUser(t *testing.T) {
	// 创建测试数据
	req := usecase.UserCreateRequest{
		Nickname: "Test User 1",
		FullName: "Test User Full Name",
		Email:    "testuser@example.com",
		Password: "password123",
	}

	// 调用 CreateUser 方法
	resp, err := userSvc.CreateUser(testCtx, req)
	if err.Code != "" {
		t.Errorf("Expected no error, got: %s", err.Message)
	}

	// 验证响应
	if resp.UserID == uuid.Nil {
		t.Errorf("Expected non-nil UserID, got %v", resp.UserID)
	}
	if resp.Nickname != req.Nickname {
		t.Errorf("Expected Nickname %s, got %s", req.Nickname, resp.Nickname)
	}
	if resp.Email != req.Email {
		t.Errorf("Expected Email %s, got %s", req.Email, resp.Email)
	}
}

func TestGetUserByID(t *testing.T) {
	// 首先创建一个用户
	createReq := usecase.UserCreateRequest{
		Nickname: "Test User 2",
		FullName: "Test User Full Name",
		Email:    "testuser2@example.com",
		Password: "password123",
	}
	createResp, err := userSvc.CreateUser(testCtx, createReq)
	if err.Code != "" {
		t.Errorf("Expected no error when creating user, got: %s", err.Message)
	}

	// 然后获取这个用户
	getReq := usecase.UserGetRequest{
		UserID: createResp.UserID,
	}
	getResp, err := userSvc.GetUserByID(testCtx, getReq)
	if err.Code != "" {
		t.Errorf("Expected no error when getting user, got: %s", err.Message)
	}

	// 验证响应
	if getResp.UserID != createResp.UserID {
		t.Errorf("Expected UserID %v, got %v", createResp.UserID, getResp.UserID)
	}
	if getResp.Nickname != createReq.Nickname {
		t.Errorf("Expected Nickname %s, got %s", createReq.Nickname, getResp.Nickname)
	}
	if getResp.Email != createReq.Email {
		t.Errorf("Expected Email %s, got %s", createReq.Email, getResp.Email)
	}
}

func TestUpdateUser(t *testing.T) {
	// 首先创建一个用户
	createReq := usecase.UserCreateRequest{
		Nickname: "Test User 3",
		FullName: "Test User Full Name",
		Email:    "testuser3@example.com",
		Password: "password123",
	}
	createResp, err := userSvc.CreateUser(testCtx, createReq)
	if err.Code != "" {
		t.Errorf("Expected no error when creating user, got: %s", err.Message)
	}

	// 然后更新这个用户
	updateReq := usecase.UserUpdateRequest{
		UserID:   createResp.UserID,
		Nickname: "Updated Test User",
		FullName: "Updated Test User Full Name",
		Email:    "updatedtestuser@example.com",
	}
	updateResp, err := userSvc.UpdateUser(testCtx, updateReq)
	if err.Code != "" {
		t.Errorf("Expected no error when updating user, got: %s", err.Message)
	}

	// 验证响应
	if updateResp.UserID != createResp.UserID {
		t.Errorf("Expected UserID %v, got %v", createResp.UserID, updateResp.UserID)
	}
	if updateResp.Nickname != updateReq.Nickname {
		t.Errorf("Expected Nickname %s, got %s", updateReq.Nickname, updateResp.Nickname)
	}
	if updateResp.Email != updateReq.Email {
		t.Errorf("Expected Email %s, got %s", updateReq.Email, updateResp.Email)
	}
}

func TestDeleteUser(t *testing.T) {
	// 首先创建一个用户
	createReq := usecase.UserCreateRequest{
		Nickname: "Test User 4",
		FullName: "Test User Full Name",
		Email:    "testuser4@example.com",
		Password: "password123",
	}
	createResp, err := userSvc.CreateUser(testCtx, createReq)
	if err.Code != "" {
		t.Errorf("Expected no error when creating user, got: %s", err.Message)
	}

	// 然后删除这个用户
	deleteReq := usecase.UserDeleteRequest{
		UserID: createResp.UserID,
	}
	deleteResp, err := userSvc.DeleteUser(testCtx, deleteReq)
	if err.Code != "" {
		t.Errorf("Expected no error when deleting user, got: %s", err.Message)
	}

	// 验证响应
	if !deleteResp.Success {
		t.Errorf("Expected Success to be true, got false")
	}
}

func TestListUsers(t *testing.T) {
	// 调用 ListUsers 方法
	listReq := usecase.UserListRequest{}
	listResp, err := userSvc.ListUsers(testCtx, listReq)
	if err.Code != "" {
		t.Errorf("Expected no error when listing users, got: %s", err.Message)
	}

	// 验证响应
	if listResp.Users == nil {
		t.Errorf("Expected Users to be non-nil, got nil")
	}
}

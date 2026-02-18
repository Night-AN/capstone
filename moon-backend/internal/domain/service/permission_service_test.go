package service_test

import (
	"moon/internal/domain/usecase"
	"testing"

	"github.com/google/uuid"
)

// TestCreatePermission tests the CreatePermission method
// It creates a new permission and verifies that it is created successfully
func TestCreatePermission(t *testing.T) {
	// Create test data
	description := "Test permission description"
	req := usecase.PermissionCreateRequest{
		PermissionName: "Test Permission",
		Description:    &description,
		PermissionCode: "test:api:create",
		SensitiveFlag:  false,
	}

	// Create the permission
	resp := permissionSvc.CreatePermission(&testCtx, req)

	// Verify the response
	if resp.PermissionID == uuid.Nil {
		t.Errorf("Expected non-nil PermissionID, got %v", resp.PermissionID)
	}
	if resp.PermissionName != req.PermissionName {
		t.Errorf("Expected PermissionName %s, got %s", req.PermissionName, resp.PermissionName)
	}
	if resp.PermissionCode != req.PermissionCode {
		t.Errorf("Expected PermissionCode %s, got %s", req.PermissionCode, resp.PermissionCode)
	}
}

// TestGetPermission tests the GetPermission method
// It creates a new permission and then retrieves it
func TestGetPermission(t *testing.T) {
	// Create test data
	description := "Test permission description"
	createReq := usecase.PermissionCreateRequest{
		PermissionName: "Test Permission",
		Description:    &description,
		PermissionCode: "test:api:read",
		SensitiveFlag:  false,
	}

	// Create the permission
	createResp := permissionSvc.CreatePermission(&testCtx, createReq)

	// Retrieve the permission
	getReq := usecase.PermissionGetRequest{
		PermissionID: createResp.PermissionID,
	}
	getResp := permissionSvc.GetPermission(&testCtx, getReq)

	// Verify the response
	if getResp.PermissionID != createResp.PermissionID {
		t.Errorf("Expected PermissionID %v, got %v", createResp.PermissionID, getResp.PermissionID)
	}
	if getResp.PermissionName != createReq.PermissionName {
		t.Errorf("Expected PermissionName %s, got %s", createReq.PermissionName, getResp.PermissionName)
	}
	if getResp.PermissionCode != createReq.PermissionCode {
		t.Errorf("Expected PermissionCode %s, got %s", createReq.PermissionCode, getResp.PermissionCode)
	}
	if getResp.SensitiveFlag != createReq.SensitiveFlag {
		t.Errorf("Expected SensitiveFlag %v, got %v", createReq.SensitiveFlag, getResp.SensitiveFlag)
	}
}

// TestUpdatePermission tests the UpdatePermission method
// It creates a new permission, updates it, and verifies the changes
func TestUpdatePermission(t *testing.T) {
	// Create test data
	description := "Test permission description"
	createReq := usecase.PermissionCreateRequest{
		PermissionName: "Test Permission",
		Description:    &description,
		PermissionCode: "test:api:update",
		SensitiveFlag:  false,
	}

	// Create the permission
	createResp := permissionSvc.CreatePermission(&testCtx, createReq)

	// Update the permission
	updatedDescription := "Updated test permission description"
	updateReq := usecase.PermissionUpdateRequest{
		PermissionID:   createResp.PermissionID,
		PermissionName: "Updated Test Permission",
		Description:    &updatedDescription,
		PermissionCode: "test:api:update:updated",
		SensitiveFlag:  true,
	}
	updateResp := permissionSvc.UpdatePermission(&testCtx, updateReq)

	// Verify the response
	if updateResp.PermissionID != createResp.PermissionID {
		t.Errorf("Expected PermissionID %v, got %v", createResp.PermissionID, updateResp.PermissionID)
	}
	if updateResp.PermissionName != updateReq.PermissionName {
		t.Errorf("Expected PermissionName %s, got %s", updateReq.PermissionName, updateResp.PermissionName)
	}
	if updateResp.PermissionCode != updateReq.PermissionCode {
		t.Errorf("Expected PermissionCode %s, got %s", updateReq.PermissionCode, updateResp.PermissionCode)
	}
}

// TestDeletePermission tests the DeletePermission method
// It creates a new permission, deletes it, and verifies the deletion
func TestDeletePermission(t *testing.T) {
	// Create test data
	description := "Test permission description"
	createReq := usecase.PermissionCreateRequest{
		PermissionName: "Test Permission",
		Description:    &description,
		PermissionCode: "test:api:delete",
		SensitiveFlag:  false,
	}

	// Create the permission
	createResp := permissionSvc.CreatePermission(&testCtx, createReq)

	// Delete the permission
	deleteReq := usecase.PermissionDeleteRequest{
		PermissionID: createResp.PermissionID,
	}
	deleteResp := permissionSvc.DeletePermission(&testCtx, deleteReq)

	// Verify the response
	if !deleteResp.Success {
		t.Errorf("Expected deletion to be successful, got %v", deleteResp.Success)
	}
}

// TestListPermissions tests the ListPermissions method
// It creates multiple permissions and then lists them
func TestListPermissions(t *testing.T) {
	// Create test data
	description := "Test permission description"

	// Create first permission
	createReq1 := usecase.PermissionCreateRequest{
		PermissionName: "Test Permission 1",
		Description:    &description,
		PermissionCode: "test:api:list:1",
		SensitiveFlag:  false,
	}
	permissionSvc.CreatePermission(&testCtx, createReq1)

	// Create second permission
	createReq2 := usecase.PermissionCreateRequest{
		PermissionName: "Test Permission 2",
		Description:    &description,
		PermissionCode: "test:api:list:2",
		SensitiveFlag:  true,
	}
	permissionSvc.CreatePermission(&testCtx, createReq2)

	// List permissions
	listReq := usecase.PermissionListRequest{
		Limit:  10,
		Offset: 0,
	}
	listResp := permissionSvc.ListPermissions(&testCtx, listReq)

	// Verify the response
	if len(listResp.Permissions) < 2 {
		t.Errorf("Expected at least 2 permissions, got %d", len(listResp.Permissions))
	}
}

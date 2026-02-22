package service_test

import (
	"moon/internal/domain/usecase"
	"testing"

	"github.com/google/uuid"
)

// TestCreatePermissionResource tests the Create method of PermissionResourceService
// It creates a permission-resource relationship
func TestCreatePermissionResource(t *testing.T) {
	// Create permission-resource relationship
	req := usecase.PermissionResourceCreateRequest{
		PermissionID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		ResourceID:   uuid.MustParse("00000000-0000-0000-0000-000000000002"),
	}

	// Execute
	resp, err := permissionResourceSvc.Create(testCtx, req)

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !resp.Success {
		t.Errorf("Expected success to be true, got %v", resp.Success)
	}
}

// TestDeletePermissionResource tests the Delete method of PermissionResourceService
// It deletes a permission-resource relationship
func TestDeletePermissionResource(t *testing.T) {
	// Delete permission-resource relationship
	req := usecase.PermissionResourceDeleteRequest{
		PermissionID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		ResourceID:   uuid.MustParse("00000000-0000-0000-0000-000000000002"),
	}

	// Execute
	resp, err := permissionResourceSvc.Delete(testCtx, req)

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !resp.Success {
		t.Errorf("Expected success to be true, got %v", resp.Success)
	}
}

// TestGetPermissionsByResourceID tests the GetPermissionsByResourceID method of PermissionResourceService
// It retrieves permissions by resource ID
func TestGetPermissionsByResourceID(t *testing.T) {
	// Get permissions by resource ID
	_, err := permissionResourceSvc.GetPermissionsByResourceID(testCtx, uuid.MustParse("00000000-0000-0000-0000-000000000002"))

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// We don't expect any permissions since we didn't create any
}

// TestGetResourcesByPermissionID tests the GetResourcesByPermissionID method of PermissionResourceService
// It retrieves resources by permission ID
func TestGetResourcesByPermissionID(t *testing.T) {
	// Get resources by permission ID
	_, err := permissionResourceSvc.GetResourcesByPermissionID(testCtx, uuid.MustParse("00000000-0000-0000-0000-000000000001"))

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// We don't expect any resources since we didn't create any
}

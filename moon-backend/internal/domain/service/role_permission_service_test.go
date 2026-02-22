package service_test

import (
	"context"
	"moon/internal/domain/service"
	"moon/internal/domain/usecase"
	"moon/internal/infrastructure/persistence/postgres"
	"testing"

	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// db is the database connection used for testing
	db *gorm.DB
	// rolePermissionSvc is the RolePermissionService instance used for testing
	rolePermissionSvc service.RolePermissionService
	// testCtx is the context used for testing
	testCtx = context.Background()
)

// TestMain initializes the test environment
// It sets up the database connection, migrates the schema, and initializes the services
// It runs the tests and cleans up the test data
func TestMain(m *testing.M) {
	// Initialize the database connection
	var err error
	db, err = gorm.Open(driver.Open("host=localhost user=capstone password=capstone dbname=capstone port=5432 sslmode=disable"))
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Clear test data
	err = db.Exec("DELETE FROM systems.permission_role").Error
	if err != nil {
		panic("Failed to clear permission_role data: " + err.Error())
	}
	err = db.Exec("DELETE FROM systems.role").Error
	if err != nil {
		panic("Failed to clear role data: " + err.Error())
	}
	err = db.Exec("DELETE FROM systems.permission").Error
	if err != nil {
		panic("Failed to clear permission data: " + err.Error())
	}

	// Initialize the repositories and services
	rolePermissionRepo := postgres.NewRolePermissionRepository(db)
	roleRepo := postgres.NewRoleRepository(db)
	permissionRepo := postgres.NewPermissionRepository(db)

	rolePermissionSvc = service.NewRolePermissionService(rolePermissionRepo, roleRepo, permissionRepo)

	// Run the tests
	m.Run()
}

// TestCreateRolePermission tests the Create method of RolePermissionService
// It creates a new role and permission, then creates a role-permission relationship
func TestCreateRolePermission(t *testing.T) {
	// Create test role
	roleReq := usecase.RoleCreateRequest{
		RoleName:      "Test Role",
		RoleCode:      "custom:test-role",
		RoleFlag:      "active",
		SensitiveFlag: false,
	}

	// Create test permission
	permissionReq := usecase.PermissionCreateRequest{
		PermissionName: "Test Permission",
		PermissionCode: "test:api:test",
		SensitiveFlag:  false,
	}

	// Create role and permission (using existing services from role_service_test)
	roleResp := roleSvc.CreateRole(&testCtx, roleReq)
	permissionResp := permissionSvc.CreatePermission(&testCtx, permissionReq)

	// Create role-permission relationship
	req := usecase.RolePermissionCreateRequest{
		RoleID:       roleResp.RoleID,
		PermissionID: permissionResp.PermissionID,
	}

	// Execute
	resp, err := rolePermissionSvc.Create(testCtx, req)

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !resp.Success {
		t.Errorf("Expected success to be true, got %v", resp.Success)
	}
}

// TestDeleteRolePermission tests the Delete method of RolePermissionService
// It creates a new role and permission, creates a role-permission relationship, then deletes it
func TestDeleteRolePermission(t *testing.T) {
	// Create test role
	roleReq := usecase.RoleCreateRequest{
		RoleName:      "Test Role Delete",
		RoleCode:      "custom:test-role-delete",
		RoleFlag:      "active",
		SensitiveFlag: false,
	}

	// Create test permission
	permissionReq := usecase.PermissionCreateRequest{
		PermissionName: "Test Permission Delete",
		PermissionCode: "test:api:test-delete",
		SensitiveFlag:  false,
	}

	// Create role and permission
	roleResp := roleSvc.CreateRole(&testCtx, roleReq)
	permissionResp := permissionSvc.CreatePermission(&testCtx, permissionReq)

	// Create role-permission relationship
	createReq := usecase.RolePermissionCreateRequest{
		RoleID:       roleResp.RoleID,
		PermissionID: permissionResp.PermissionID,
	}
	rolePermissionSvc.Create(testCtx, createReq)

	// Delete role-permission relationship
	deleteReq := usecase.RolePermissionDeleteRequest{
		RoleID:       roleResp.RoleID,
		PermissionID: permissionResp.PermissionID,
	}

	// Execute
	resp, err := rolePermissionSvc.Delete(testCtx, deleteReq)

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !resp.Success {
		t.Errorf("Expected success to be true, got %v", resp.Success)
	}
}

// TestGetPermissionsByRoleID tests the GetPermissionsByRoleID method of RolePermissionService
// It creates a new role and permission, creates a role-permission relationship, then retrieves permissions by role ID
func TestGetPermissionsByRoleID(t *testing.T) {
	// Create test role
	roleReq := usecase.RoleCreateRequest{
		RoleName:      "Test Role Permissions",
		RoleCode:      "custom:test-role-permissions",
		RoleFlag:      "active",
		SensitiveFlag: false,
	}

	// Create test permission
	permissionReq := usecase.PermissionCreateRequest{
		PermissionName: "Test Permission Permissions",
		PermissionCode: "test:api:test-permissions",
		SensitiveFlag:  false,
	}

	// Create role and permission
	roleResp := roleSvc.CreateRole(&testCtx, roleReq)
	permissionResp := permissionSvc.CreatePermission(&testCtx, permissionReq)

	// Create role-permission relationship
	createReq := usecase.RolePermissionCreateRequest{
		RoleID:       roleResp.RoleID,
		PermissionID: permissionResp.PermissionID,
	}
	rolePermissionSvc.Create(testCtx, createReq)

	// Get permissions by role ID
	permissions, err := rolePermissionSvc.GetPermissionsByRoleID(testCtx, roleResp.RoleID)

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(permissions) < 1 {
		t.Errorf("Expected at least 1 permission, got %d", len(permissions))
	}
}

// TestGetRolesByPermissionID tests the GetRolesByPermissionID method of RolePermissionService
// It creates a new role and permission, creates a role-permission relationship, then retrieves roles by permission ID
func TestGetRolesByPermissionID(t *testing.T) {
	// Create test role
	roleReq := usecase.RoleCreateRequest{
		RoleName:      "Test Role Roles",
		RoleCode:      "custom:test-role-roles",
		RoleFlag:      "active",
		SensitiveFlag: false,
	}

	// Create test permission
	permissionReq := usecase.PermissionCreateRequest{
		PermissionName: "Test Permission Roles",
		PermissionCode: "test:api:test-roles",
		SensitiveFlag:  false,
	}

	// Create role and permission
	roleResp := roleSvc.CreateRole(&testCtx, roleReq)
	permissionResp := permissionSvc.CreatePermission(&testCtx, permissionReq)

	// Create role-permission relationship
	createReq := usecase.RolePermissionCreateRequest{
		RoleID:       roleResp.RoleID,
		PermissionID: permissionResp.PermissionID,
	}
	rolePermissionSvc.Create(testCtx, createReq)

	// Get roles by permission ID
	roles, err := rolePermissionSvc.GetRolesByPermissionID(testCtx, permissionResp.PermissionID)

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(roles) < 1 {
		t.Errorf("Expected at least 1 role, got %d", len(roles))
	}
}

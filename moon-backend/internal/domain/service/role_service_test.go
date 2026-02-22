package service_test

import (
	"context"
	"moon/internal/domain/service"
	"moon/internal/domain/usecase"
	"moon/internal/infrastructure/persistence/postgres"
	"testing"

	"github.com/google/uuid"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// db is the database connection used for testing
	db *gorm.DB
	// roleSvc is the RoleService instance used for testing
	roleSvc service.RoleService
	// permissionSvc is the PermissionService instance used for testing
	permissionSvc service.PermissionService
	// organizationSvc is the OrganizationService instance used for testing
	organizationSvc service.OrganizationService
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

	// Skip auto migration because tables are already created by setup.sql
	// err = db.AutoMigrate(&aggregate.Role{})
	// if err != nil {
	// 	panic("Failed to migrate database: " + err.Error())
	// }

	// // Auto migrate the Permission schema
	// err = db.AutoMigrate(&aggregate.Permission{})
	// if err != nil {
	// 	panic("Failed to migrate database: " + err.Error())
	// }

	// Create permission_role table if it doesn't exist
	err = db.Exec(`CREATE TABLE IF NOT EXISTS systems.permission_role (
		id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
		permission_id UUID NOT NULL,
		role_id UUID NOT NULL,
		PRIMARY KEY (id)
	)`).Error
	if err != nil {
		panic("Failed to create permission_role table: " + err.Error())
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
	err = db.Exec("DELETE FROM systems.organization").Error
	if err != nil {
		panic("Failed to clear organization data: " + err.Error())
	}

	// Initialize the repositories and services
	roleRepo := postgres.NewRoleRepository(db)
	roleSvc = service.NewRoleService(roleRepo)

	// Initialize permission service
	permissionRepo := postgres.NewPermissionRepository(db)
	permissionSvc = service.NewPermissionService(permissionRepo)

	// Initialize organization service
	organizationRepo := postgres.NewOrganizationRepository(db)
	organizationSvc = service.NewOrganizationService(organizationRepo, nil)

	// Run the tests
	m.Run()
}

// TestCreateRole tests the CreateRole method
// It creates a new role and verifies that it is created successfully
func TestCreateRole(t *testing.T) {
	// Create test data
	description := "Test role description"
	req := usecase.RoleCreateRequest{
		RoleName:      "Test Role",
		Description:   &description,
		RoleCode:      "custom:test-role",
		RoleFlag:      "active",
		SensitiveFlag: false,
	}

	// Create the role
	resp := roleSvc.CreateRole(&testCtx, req)

	// Verify the response
	if resp.RoleID == uuid.Nil {
		t.Errorf("Expected non-nil RoleID, got %v", resp.RoleID)
	}
	if resp.RoleName != req.RoleName {
		t.Errorf("Expected RoleName %s, got %s", req.RoleName, resp.RoleName)
	}
	if resp.RoleCode != req.RoleCode {
		t.Errorf("Expected RoleCode %s, got %s", req.RoleCode, resp.RoleCode)
	}
}

// TestGetRole tests the GetRole method
// It creates a new role and then retrieves it
func TestGetRole(t *testing.T) {
	// Create test data
	description := "Test role description"
	createReq := usecase.RoleCreateRequest{
		RoleName:      "Test Role",
		Description:   &description,
		RoleCode:      "custom:test-role-read",
		RoleFlag:      "active",
		SensitiveFlag: false,
	}

	// Create the role
	createResp := roleSvc.CreateRole(&testCtx, createReq)

	// Retrieve the role
	getReq := usecase.RoleGetRequest{
		RoleID: createResp.RoleID,
	}
	getResp := roleSvc.GetRole(&testCtx, getReq)

	// Verify the response
	if getResp.RoleID != createResp.RoleID {
		t.Errorf("Expected RoleID %v, got %v", createResp.RoleID, getResp.RoleID)
	}
	if getResp.RoleName != createReq.RoleName {
		t.Errorf("Expected RoleName %s, got %s", createReq.RoleName, getResp.RoleName)
	}
	if getResp.RoleCode != createReq.RoleCode {
		t.Errorf("Expected RoleCode %s, got %s", createReq.RoleCode, getResp.RoleCode)
	}
	if getResp.RoleFlag != createReq.RoleFlag {
		t.Errorf("Expected RoleFlag %s, got %s", createReq.RoleFlag, getResp.RoleFlag)
	}
	if getResp.SensitiveFlag != createReq.SensitiveFlag {
		t.Errorf("Expected SensitiveFlag %v, got %v", createReq.SensitiveFlag, getResp.SensitiveFlag)
	}
}

// TestUpdateRole tests the UpdateRole method
// It creates a new role, updates it, and verifies the changes
func TestUpdateRole(t *testing.T) {
	// Create test data
	description := "Test role description"
	createReq := usecase.RoleCreateRequest{
		RoleName:      "Test Role",
		Description:   &description,
		RoleCode:      "custom:test-role-update",
		RoleFlag:      "active",
		SensitiveFlag: false,
	}

	// Create the role
	createResp := roleSvc.CreateRole(&testCtx, createReq)

	// Update the role
	updatedDescription := "Updated test role description"
	updateReq := usecase.RoleUpdateRequest{
		RoleID:        createResp.RoleID,
		RoleName:      "Updated Test Role",
		Description:   &updatedDescription,
		RoleCode:      "custom:test-role-updated",
		RoleFlag:      "active",
		SensitiveFlag: false,
	}
	updateResp := roleSvc.UpdateRole(&testCtx, updateReq)

	// Verify the response
	if updateResp.RoleID != createResp.RoleID {
		t.Errorf("Expected RoleID %v, got %v", createResp.RoleID, updateResp.RoleID)
	}
	if updateResp.RoleName != updateReq.RoleName {
		t.Errorf("Expected RoleName %s, got %s", updateReq.RoleName, updateResp.RoleName)
	}
	if updateResp.RoleCode != updateReq.RoleCode {
		t.Errorf("Expected RoleCode %s, got %s", updateReq.RoleCode, updateResp.RoleCode)
	}
}

// TestDeleteRole tests the DeleteRole method
// It creates a new role, deletes it, and verifies the deletion
func TestDeleteRole(t *testing.T) {
	// Create test data
	description := "Test role description"
	createReq := usecase.RoleCreateRequest{
		RoleName:      "Test Role",
		Description:   &description,
		RoleCode:      "custom:test-role-delete",
		RoleFlag:      "active",
		SensitiveFlag: false,
	}

	// Create the role
	createResp := roleSvc.CreateRole(&testCtx, createReq)

	// Delete the role
	deleteReq := usecase.RoleDeleteRequest{
		RoleID: createResp.RoleID,
	}
	deleteResp := roleSvc.DeleteRole(&testCtx, deleteReq)

	// Verify the response
	if !deleteResp.Success {
		t.Errorf("Expected deletion to be successful, got %v", deleteResp.Success)
	}
}

// TestAssignPermission tests the AssignPermission method
// It creates a new role and permission, assigns the permission to the role, and verifies the assignment
func TestAssignPermission(t *testing.T) {
	// Create a role
	description := "Test role description"
	createRoleReq := usecase.RoleCreateRequest{
		RoleName:      "Test Role",
		Description:   &description,
		RoleCode:      "custom:test-role-assign",
		RoleFlag:      "active",
		SensitiveFlag: false,
	}
	createRoleResp := roleSvc.CreateRole(&testCtx, createRoleReq)

	// Create a permission
	description2 := "Test permission description"
	createPermissionReq := usecase.PermissionCreateRequest{
		PermissionName: "Test Permission",
		Description:    &description2,
		PermissionCode: "test:api:assign",
		SensitiveFlag:  false,
	}

	// Get the permission service (using the global permissionSvc variable)

	// Create the permission
	createPermissionResp := permissionSvc.CreatePermission(&testCtx, createPermissionReq)

	// Assign the permission to the role
	assignReq := usecase.RoleAssignPermissionRequest{
		RoleID:       createRoleResp.RoleID,
		PermissionID: createPermissionResp.PermissionID,
	}
	assignResp := roleSvc.AssignPermission(&testCtx, assignReq)

	// Verify the response
	if !assignResp.Success {
		t.Errorf("Expected assignment to be successful, got %v", assignResp.Success)
	}
}

// TestRemovePermission tests the RemovePermission method
// It creates a new role and permission, assigns the permission to the role, removes it, and verifies the removal
func TestRemovePermission(t *testing.T) {
	// Create a role
	description := "Test role description"
	createRoleReq := usecase.RoleCreateRequest{
		RoleName:      "Test Role",
		Description:   &description,
		RoleCode:      "custom:test-role-remove",
		RoleFlag:      "active",
		SensitiveFlag: false,
	}
	createRoleResp := roleSvc.CreateRole(&testCtx, createRoleReq)

	// Create a permission
	description2 := "Test permission description"
	createPermissionReq := usecase.PermissionCreateRequest{
		PermissionName: "Test Permission",
		Description:    &description2,
		PermissionCode: "test:api:remove",
		SensitiveFlag:  false,
	}

	// Get the permission service (using the global permissionSvc variable)

	// Create the permission
	createPermissionResp := permissionSvc.CreatePermission(&testCtx, createPermissionReq)

	// Assign the permission to the role
	assignReq := usecase.RoleAssignPermissionRequest{
		RoleID:       createRoleResp.RoleID,
		PermissionID: createPermissionResp.PermissionID,
	}
	roleSvc.AssignPermission(&testCtx, assignReq)

	// Remove the permission from the role
	removeReq := usecase.RoleRemovePermissionRequest{
		RoleID:       createRoleResp.RoleID,
		PermissionID: createPermissionResp.PermissionID,
	}
	removeResp := roleSvc.RemovePermission(&testCtx, removeReq)

	// Verify the response
	if !removeResp.Success {
		t.Errorf("Expected removal to be successful, got %v", removeResp.Success)
	}
}

// TestGetRolePermissions tests the GetRolePermissions method
// It creates a new role and permission, assigns the permission to the role, and retrieves the role's permissions
func TestGetRolePermissions(t *testing.T) {
	// Create a role
	description := "Test role description"
	createRoleReq := usecase.RoleCreateRequest{
		RoleName:      "Test Role",
		Description:   &description,
		RoleCode:      "custom:test-role-permissions",
		RoleFlag:      "active",
		SensitiveFlag: false,
	}
	createRoleResp := roleSvc.CreateRole(&testCtx, createRoleReq)

	// Create a permission
	description2 := "Test permission description"
	createPermissionReq := usecase.PermissionCreateRequest{
		PermissionName: "Test Permission",
		Description:    &description2,
		PermissionCode: "test:api:permissions",
		SensitiveFlag:  false,
	}

	// Get the permission service (using the global permissionSvc variable)

	// Create the permission
	createPermissionResp := permissionSvc.CreatePermission(&testCtx, createPermissionReq)

	// Assign the permission to the role
	assignReq := usecase.RoleAssignPermissionRequest{
		RoleID:       createRoleResp.RoleID,
		PermissionID: createPermissionResp.PermissionID,
	}
	roleSvc.AssignPermission(&testCtx, assignReq)

	// Get the role's permissions
	getPermissionsReq := usecase.RolePermissionsRequest{
		RoleID: createRoleResp.RoleID,
	}
	getPermissionsResp := roleSvc.GetRolePermissions(&testCtx, getPermissionsReq)

	// Verify the response
	if getPermissionsResp.RoleID != createRoleResp.RoleID {
		t.Errorf("Expected RoleID %v, got %v", createRoleResp.RoleID, getPermissionsResp.RoleID)
	}
	if len(getPermissionsResp.Permissions) < 1 {
		t.Errorf("Expected at least 1 permission, got %d", len(getPermissionsResp.Permissions))
	}
}

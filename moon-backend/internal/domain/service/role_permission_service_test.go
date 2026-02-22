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
	// rolePermissionSvc is the RolePermissionService instance used for testing
	rolePermissionSvc service.RolePermissionService
	// organizationRoleSvc is the OrganizationRoleService instance used for testing
	organizationRoleSvc service.OrganizationRoleService
	// permissionResourceSvc is the PermissionResourceService instance used for testing
	permissionResourceSvc service.PermissionResourceService
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

	// Create organization_role table if it doesn't exist
	err = db.Exec(`CREATE TABLE IF NOT EXISTS systems.organization_role (
		id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
		organization_id UUID NOT NULL,
		role_id UUID NOT NULL,
		PRIMARY KEY (id)
	)`).Error
	if err != nil {
		panic("Failed to create organization_role table: " + err.Error())
	}

	// Create permission_resource table if it doesn't exist
	err = db.Exec(`CREATE TABLE IF NOT EXISTS systems.permission_resource (
		id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
		permission_id UUID NOT NULL,
		resource_id UUID NOT NULL,
		PRIMARY KEY (id)
	)`).Error
	if err != nil {
		panic("Failed to create permission_resource table: " + err.Error())
	}

	// Clear test data
	clearTestData()

	// Initialize the repositories and services
	initializeServices()

	// Run the tests
	m.Run()
}

// clearTestData clears all test data from the database
func clearTestData() {
	// Clear join tables first
	if err := db.Exec("DELETE FROM systems.permission_role").Error; err != nil {
		panic("Failed to clear permission_role data: " + err.Error())
	}
	if err := db.Exec("DELETE FROM systems.organization_role").Error; err != nil {
		panic("Failed to clear organization_role data: " + err.Error())
	}
	if err := db.Exec("DELETE FROM systems.permission_resource").Error; err != nil {
		panic("Failed to clear permission_resource data: " + err.Error())
	}
	if err := db.Exec("DELETE FROM systems.role").Error; err != nil {
		panic("Failed to clear role data: " + err.Error())
	}
	if err := db.Exec("DELETE FROM systems.permission").Error; err != nil {
		panic("Failed to clear permission data: " + err.Error())
	}
	if err := db.Exec("DELETE FROM systems.organization").Error; err != nil {
		panic("Failed to clear organization data: " + err.Error())
	}
}

// initializeServices initializes all the repositories and services
func initializeServices() {
	// Initialize repositories
	rolePermissionRepo := postgres.NewRolePermissionRepository(db)
	organizationRoleRepo := postgres.NewOrganizationRoleRepository(db)
	permissionResourceRepo := postgres.NewPermissionResourceRepository(db)
	roleRepo := postgres.NewRoleRepository(db)
	permissionRepo := postgres.NewPermissionRepository(db)
	organizationRepo := postgres.NewOrganizationRepository(db)
	resourceRepo := postgres.NewResourceRepository(db)

	// Initialize services
	rolePermissionSvc = service.NewRolePermissionService(rolePermissionRepo, roleRepo, permissionRepo)
	organizationRoleSvc = service.NewOrganizationRoleService(organizationRoleRepo, organizationRepo, roleRepo)
	permissionResourceSvc = service.NewPermissionResourceService(permissionResourceRepo, permissionRepo, resourceRepo)
	roleSvc = service.NewRoleService(roleRepo)
	permissionSvc = service.NewPermissionService(permissionRepo)
	organizationSvc = service.NewOrganizationService(organizationRepo, nil)
}

// TestCreateRolePermission tests the Create method of RolePermissionService
// It creates a role-permission relationship
func TestCreateRolePermission(t *testing.T) {
	// Create role-permission relationship
	req := usecase.RolePermissionCreateRequest{
		RoleID:       uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		PermissionID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
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
// It deletes a role-permission relationship
func TestDeleteRolePermission(t *testing.T) {
	// Delete role-permission relationship
	req := usecase.RolePermissionDeleteRequest{
		RoleID:       uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		PermissionID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
	}

	// Execute
	resp, err := rolePermissionSvc.Delete(testCtx, req)

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !resp.Success {
		t.Errorf("Expected success to be true, got %v", resp.Success)
	}
}

// TestGetPermissionsByRoleID tests the GetPermissionsByRoleID method of RolePermissionService
// It retrieves permissions by role ID
func TestGetPermissionsByRoleID(t *testing.T) {
	// Get permissions by role ID
	_, err := rolePermissionSvc.GetPermissionsByRoleID(testCtx, uuid.MustParse("00000000-0000-0000-0000-000000000001"))

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// We don't expect any permissions since we didn't create any
}

// TestGetRolesByPermissionID tests the GetRolesByPermissionID method of RolePermissionService
// It retrieves roles by permission ID
func TestGetRolesByPermissionID(t *testing.T) {
	// Get roles by permission ID
	_, err := rolePermissionSvc.GetRolesByPermissionID(testCtx, uuid.MustParse("00000000-0000-0000-0000-000000000002"))

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// We don't expect any roles since we didn't create any
}

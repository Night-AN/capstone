package service_test

import (
	"moon/internal/domain/usecase"
	"testing"

	"github.com/google/uuid"
)

// TestCreateOrganizationRole tests the Create method of OrganizationRoleService
// It creates an organization-role relationship
func TestCreateOrganizationRole(t *testing.T) {
	// Create organization-role relationship
	req := usecase.OrganizationRoleCreateRequest{
		OrganizationID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		RoleID:         uuid.MustParse("00000000-0000-0000-0000-000000000002"),
	}

	// Execute
	resp, err := organizationRoleSvc.Create(testCtx, req)

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !resp.Success {
		t.Errorf("Expected success to be true, got %v", resp.Success)
	}
}

// TestDeleteOrganizationRole tests the Delete method of OrganizationRoleService
// It deletes an organization-role relationship
func TestDeleteOrganizationRole(t *testing.T) {
	// Delete organization-role relationship
	req := usecase.OrganizationRoleDeleteRequest{
		OrganizationID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
		RoleID:         uuid.MustParse("00000000-0000-0000-0000-000000000002"),
	}

	// Execute
	resp, err := organizationRoleSvc.Delete(testCtx, req)

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !resp.Success {
		t.Errorf("Expected success to be true, got %v", resp.Success)
	}
}

// TestGetRolesByOrganizationID tests the GetRolesByOrganizationID method of OrganizationRoleService
// It retrieves roles by organization ID
func TestGetRolesByOrganizationID(t *testing.T) {
	// Get roles by organization ID
	_, err := organizationRoleSvc.GetRolesByOrganizationID(testCtx, uuid.MustParse("00000000-0000-0000-0000-000000000001"))

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// We don't expect any roles since we didn't create any
}

// TestGetOrganizationsByRoleID tests the GetOrganizationsByRoleID method of OrganizationRoleService
// It retrieves organizations by role ID
func TestGetOrganizationsByRoleID(t *testing.T) {
	// Get organizations by role ID
	_, err := organizationRoleSvc.GetOrganizationsByRoleID(testCtx, uuid.MustParse("00000000-0000-0000-0000-000000000002"))

	// Verify the response
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// We don't expect any organizations since we didn't create any
}

package service

import (
	"context"
	"fmt"
	"time"

	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"

	"github.com/google/uuid"
)

// RoleService defines the interface for role-related operations
// It provides methods for creating, retrieving, updating, and deleting roles,
// as well as managing role-permission relationships
// All methods accept a context and request object, and return a response object
// The context is used for cancellation, timeouts, and passing request-scoped values
// The request and response objects are defined in the usecase package
// The service layer orchestrates business logic and coordinates between repositories
// It does not directly interact with the database, but delegates to repositories
// It also handles validation, error handling, and business rules
// The service layer is the primary entry point for application logic
// and is used by handlers/controllers to process requests
// The service layer should be stateless and thread-safe
// It should not maintain any internal state between method calls
// All dependencies should be injected via the constructor
// The service layer should be tested in isolation from repositories
// using mocks or fakes for repository dependencies
type RoleService interface {
	// CreateRole creates a new role with the specified details
	// It returns a response containing the created role's information
	// If creation fails, it returns an error response
	CreateRole(ctx context.Context, req usecase.RoleCreateRequest) usecase.RoleCreateResponse

	// GetRole retrieves a role by its ID
	// It returns a response containing the role's information
	// If the role is not found, it returns an error response
	GetRole(ctx context.Context, req usecase.RoleGetRequest) usecase.RoleGetResponse

	// UpdateRole updates an existing role with the specified details
	// It returns a response containing the updated role's information
	// If update fails, it returns an error response
	UpdateRole(ctx context.Context, req usecase.RoleUpdateRequest) usecase.RoleUpdateResponse

	// DeleteRole deletes a role by its ID
	// It returns a response indicating whether the deletion was successful
	// If deletion fails, it returns an error response
	DeleteRole(ctx context.Context, req usecase.RoleDeleteRequest) usecase.RoleDeleteResponse

	// ListRoles retrieves all roles
	// It returns a response containing the list of roles
	// If retrieval fails, it returns an error response
	ListRoles(ctx context.Context, req usecase.RoleListRequest) usecase.RoleListResponse

	// AssignPermission assigns a permission to a role
	// It returns a response indicating whether the assignment was successful
	// If assignment fails, it returns an error response
	AssignPermission(ctx context.Context, req usecase.RoleAssignPermissionRequest) usecase.RoleAssignPermissionResponse

	// RemovePermission removes a permission from a role
	// It returns a response indicating whether the removal was successful
	// If removal fails, it returns an error response
	RemovePermission(ctx context.Context, req usecase.RoleRemovePermissionRequest) usecase.RoleRemovePermissionResponse

	// GetRolePermissions retrieves all permissions assigned to a role
	// It returns a response containing the role's permissions
	// If retrieval fails, it returns an error response
	GetRolePermissions(ctx context.Context, req usecase.RolePermissionsRequest) usecase.RolePermissionsResponse

	// GetRoleUsers retrieves all users assigned to a role
	// It returns a response containing the role's users
	// If retrieval fails, it returns an error response
	GetRoleUsers(ctx context.Context, req usecase.RoleUsersRequest) usecase.RoleUsersResponse
}

// roleService implements the RoleService interface
// It provides methods for creating, retrieving, updating, and deleting roles
// as well as managing role-permission relationships
// All methods delegate to the RoleRepository for data access
// The roleService is initialized with a RoleRepository instance
// It is stateless and thread-safe
// It should be tested in isolation from the repository using mocks or fakes
type roleService struct {
	// roleRepo is the repository used for role data access
	roleRepo repository.RoleRepository
	// userRepo is the repository used for user data access
	userRepo repository.UserRepository
}

// NewRoleService creates a new instance of roleService
// It takes a RoleRepository instance and a UserRepository instance as parameters
// It returns a RoleService
func NewRoleService(roleRepo repository.RoleRepository, userRepo repository.UserRepository) RoleService {
	return &roleService{roleRepo, userRepo}
}

// CreateRole creates a new role with the specified details
// It returns a response containing the created role's information
// If creation fails, it returns an error response
func (rs *roleService) CreateRole(ctx context.Context, req usecase.RoleCreateRequest) usecase.RoleCreateResponse {
	// Create a new role aggregate
	role := aggregate.Role{
		RoleID:        uuid.New(),
		RoleName:      req.RoleName,
		Description:   req.Description,
		RoleCode:      req.RoleCode,
		RoleFlag:      req.RoleFlag,
		SensitiveFlag: req.SensitiveFlag,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	// Save the role to the database
	err := rs.roleRepo.SaveRole(ctx, role)
	if err != nil {
		// If saving fails, return an empty response
		return usecase.RoleCreateResponse{}
	}

	// Return the created role
	return usecase.RoleCreateResponse{
		RoleID:   role.RoleID,
		RoleName: role.RoleName,
		RoleCode: role.RoleCode,
	}
}

// GetRole retrieves a role by its ID
// It returns a response containing the role's information
// If the role is not found, it returns an error response
func (rs *roleService) GetRole(ctx context.Context, req usecase.RoleGetRequest) usecase.RoleGetResponse {
	// Find the role by ID
	role, err := rs.roleRepo.FindRoleByID(ctx, req.RoleID)
	if err != nil {
		// If role is not found, return an empty response
		return usecase.RoleGetResponse{}
	}

	// Return the role
	return usecase.RoleGetResponse{
		RoleID:        role.RoleID,
		RoleName:      role.RoleName,
		Description:   role.Description,
		RoleCode:      role.RoleCode,
		RoleFlag:      role.RoleFlag,
		SensitiveFlag: role.SensitiveFlag,
		CreatedAt:     role.CreatedAt,
	}
}

// UpdateRole updates an existing role with the specified details
// It returns a response containing the updated role's information
// If update fails, it returns an error response
func (rs *roleService) UpdateRole(ctx context.Context, req usecase.RoleUpdateRequest) usecase.RoleUpdateResponse {
	// Create an updated role aggregate
	role := aggregate.Role{
		RoleID:        req.RoleID,
		RoleName:      req.RoleName,
		Description:   req.Description,
		RoleCode:      req.RoleCode,
		RoleFlag:      req.RoleFlag,
		SensitiveFlag: req.SensitiveFlag,
		CreatedAt:     time.Now(), // Using current time for simplicity
		UpdatedAt:     time.Now(),
	}

	// Save the updated role to the database
	err := rs.roleRepo.SaveRole(ctx, role)
	if err != nil {
		// If saving fails, return an empty response
		return usecase.RoleUpdateResponse{}
	}

	// Return the updated role
	return usecase.RoleUpdateResponse{
		RoleID:   role.RoleID,
		RoleName: role.RoleName,
		RoleCode: role.RoleCode,
	}
}

// DeleteRole deletes a role by its ID
// It returns a response indicating whether the deletion was successful
// If deletion fails, it returns an error response
func (rs *roleService) DeleteRole(ctx context.Context, req usecase.RoleDeleteRequest) usecase.RoleDeleteResponse {
	// Delete the role from the database
	err := rs.roleRepo.DeleteRole(ctx, req.RoleID)
	if err != nil {
		// If deletion fails, return an empty response
		return usecase.RoleDeleteResponse{}
	}

	// Return a success response
	return usecase.RoleDeleteResponse{
		Success: true,
	}
}

// AssignPermission assigns a permission to a role
// It returns a response indicating whether the assignment was successful
// If assignment fails, it returns an error response
func (rs *roleService) AssignPermission(ctx context.Context, req usecase.RoleAssignPermissionRequest) usecase.RoleAssignPermissionResponse {
	// Assign the permission to the role
	err := rs.roleRepo.AssignPermission(ctx, req.RoleID, req.PermissionID)
	if err != nil {
		// If assignment fails, return an empty response
		return usecase.RoleAssignPermissionResponse{}
	}

	// Return a success response
	return usecase.RoleAssignPermissionResponse{
		Success: true,
	}
}

// RemovePermission removes a permission from a role
// It returns a response indicating whether the removal was successful
// If removal fails, it returns an error response
func (rs *roleService) RemovePermission(ctx context.Context, req usecase.RoleRemovePermissionRequest) usecase.RoleRemovePermissionResponse {
	// Remove the permission from the role
	err := rs.roleRepo.RemovePermission(ctx, req.RoleID, req.PermissionID)
	if err != nil {
		// If removal fails, return an empty response
		return usecase.RoleRemovePermissionResponse{}
	}

	// Return a success response
	return usecase.RoleRemovePermissionResponse{
		Success: true,
	}
}

// GetRolePermissions retrieves all permissions assigned to a role
// It returns a response containing the role's permissions
// If retrieval fails, it returns an error response
func (rs *roleService) GetRolePermissions(ctx context.Context, req usecase.RolePermissionsRequest) usecase.RolePermissionsResponse {
	// Get the role's permissions
	permissions, err := rs.roleRepo.GetRolePermissions(ctx, req.RoleID)
	if err != nil {
		// If retrieval fails, return an empty response
		return usecase.RolePermissionsResponse{}
	}

	// Convert permissions to response format
	permissionResponses := make([]usecase.PermissionResponse, len(permissions))
	for i, permission := range permissions {
		permissionResponses[i] = usecase.PermissionResponse{
			PermissionID:   permission.PermissionID,
			PermissionName: permission.PermissionName,
			PermissionCode: permission.PermissionCode,
			SensitiveFlag:  permission.SensitiveFlag,
		}
	}

	// Return the role's permissions
	return usecase.RolePermissionsResponse{
		RoleID:      req.RoleID,
		Permissions: permissionResponses,
	}
}

// ListRoles retrieves all roles
// It returns a response containing the list of roles
// If retrieval fails, it returns an error response
func (rs *roleService) ListRoles(ctx context.Context, req usecase.RoleListRequest) usecase.RoleListResponse {
	// Get all roles
	roles, err := rs.roleRepo.FindAllRoles(ctx)
	if err != nil {
		// If retrieval fails, return an empty response
		return usecase.RoleListResponse{}
	}

	// Convert roles to response format
	roleItems := make([]usecase.RoleListItem, len(roles))
	for i, role := range roles {
		roleItems[i] = usecase.RoleListItem{
			RoleID:        role.RoleID,
			RoleName:      role.RoleName,
			RoleCode:      role.RoleCode,
			RoleFlag:      role.RoleFlag,
			SensitiveFlag: role.SensitiveFlag,
			CreatedAt:     role.CreatedAt,
		}
	}

	// Return the list of roles
	return usecase.RoleListResponse{
		Roles: roleItems,
	}
}

// GetRoleUsers retrieves all users assigned to a role
// It returns a response containing the role's users
// If retrieval fails, it returns an error response
func (rs *roleService) GetRoleUsers(ctx context.Context, req usecase.RoleUsersRequest) usecase.RoleUsersResponse {
	// Get the role's users
	users, err := rs.userRepo.FindUsersByRoleID(ctx, req.RoleID)
	if err != nil {
		// If retrieval fails, log the error and return an empty response
		fmt.Printf("Error retrieving role users: %v\n", err)
		return usecase.RoleUsersResponse{
			RoleID: req.RoleID,
			Users:  []usecase.UserGetResponse{},
		}
	}

	// Convert users to response format
	userResponses := make([]usecase.UserGetResponse, len(users))
	for i, user := range users {
		userResponses[i] = usecase.UserGetResponse{
			UserID:   user.UserID,
			Nickname: user.Nickname,
			FullName: user.FullName,
			Email:    user.Email,
		}
	}

	// Return the role's users
	return usecase.RoleUsersResponse{
		RoleID: req.RoleID,
		Users:  userResponses,
	}
}

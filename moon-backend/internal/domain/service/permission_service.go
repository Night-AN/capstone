package service

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"
	"time"

	"github.com/google/uuid"
)

// PermissionService defines the interface for permission-related operations
// It provides methods for creating, retrieving, updating, and deleting permissions
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
type PermissionService interface {
	// CreatePermission creates a new permission with the specified details
	// It returns a response containing the created permission's information
	// If creation fails, it returns an error response
	CreatePermission(ctx *context.Context, req usecase.PermissionCreateRequest) usecase.PermissionCreateResponse

	// GetPermission retrieves a permission by its ID
	// It returns a response containing the permission's information
	// If the permission is not found, it returns an error response
	GetPermission(ctx *context.Context, req usecase.PermissionGetRequest) usecase.PermissionGetResponse

	// UpdatePermission updates an existing permission with the specified details
	// It returns a response containing the updated permission's information
	// If update fails, it returns an error response
	UpdatePermission(ctx *context.Context, req usecase.PermissionUpdateRequest) usecase.PermissionUpdateResponse

	// DeletePermission deletes a permission by its ID
	// It returns a response indicating whether the deletion was successful
	// If deletion fails, it returns an error response
	DeletePermission(ctx *context.Context, req usecase.PermissionDeleteRequest) usecase.PermissionDeleteResponse

	// ListPermissions retrieves all permissions matching the specified criteria
	// It returns a response containing the list of permissions
	// If retrieval fails, it returns an error response
	ListPermissions(ctx *context.Context, req usecase.PermissionListRequest) usecase.PermissionListResponse
}

// permissionService implements the PermissionService interface
// It provides methods for creating, retrieving, updating, and deleting permissions
// All methods delegate to the PermissionRepository for data access
// The permissionService is initialized with a PermissionRepository instance
// It is stateless and thread-safe
// It should be tested in isolation from the repository using mocks or fakes
type permissionService struct {
	// permissionRepo is the repository used for permission data access
	permissionRepo repository.PermissionRepository
}

// NewPermissionService creates a new instance of permissionService
// It takes a PermissionRepository instance as a parameter
// It returns a PermissionService
func NewPermissionService(permissionRepo repository.PermissionRepository) PermissionService {
	return &permissionService{permissionRepo}
}

// CreatePermission creates a new permission with the specified details
// It returns a response containing the created permission's information
// If creation fails, it returns an error response
func (ps *permissionService) CreatePermission(ctx *context.Context, req usecase.PermissionCreateRequest) usecase.PermissionCreateResponse {
	// Create a new permission aggregate
	permission := aggregate.Permission{
		PermissionID:   uuid.New(),
		PermissionName: req.PermissionName,
		Description:    req.Description,
		PermissionCode: req.PermissionCode,
		SensitiveFlag:  req.SensitiveFlag,
		CreatedAt:      time.Now(),
	}

	// Save the permission to the database
	err := ps.permissionRepo.SavePermission(ctx, permission)
	if err != nil {
		// If saving fails, return an empty response
		return usecase.PermissionCreateResponse{}
	}

	// Return the created permission
	return usecase.PermissionCreateResponse{
		PermissionID:   permission.PermissionID,
		PermissionName: permission.PermissionName,
		PermissionCode: permission.PermissionCode,
	}
}

// GetPermission retrieves a permission by its ID
// It returns a response containing the permission's information
// If the permission is not found, it returns an error response
func (ps *permissionService) GetPermission(ctx *context.Context, req usecase.PermissionGetRequest) usecase.PermissionGetResponse {
	// Find the permission by ID
	permission, err := ps.permissionRepo.FindPermissionByID(ctx, req.PermissionID)
	if err != nil {
		// If permission is not found, return an empty response
		return usecase.PermissionGetResponse{}
	}

	// Return the permission
	return usecase.PermissionGetResponse{
		PermissionID:   permission.PermissionID,
		PermissionName: permission.PermissionName,
		Description:    permission.Description,
		PermissionCode: permission.PermissionCode,
		SensitiveFlag:  permission.SensitiveFlag,
		CreatedAt:      permission.CreatedAt,
	}
}

// UpdatePermission updates an existing permission with the specified details
// It returns a response containing the updated permission's information
// If update fails, it returns an error response
func (ps *permissionService) UpdatePermission(ctx *context.Context, req usecase.PermissionUpdateRequest) usecase.PermissionUpdateResponse {
	// Check if the permission is sensitive
	existingPermission, err := ps.permissionRepo.FindPermissionByID(ctx, req.PermissionID)
	if err != nil {
		// If permission is not found, return an empty response
		return usecase.PermissionUpdateResponse{}
	}
	
	// If permission is sensitive, reject the update
	if existingPermission.SensitiveFlag {
		// Return an empty response to indicate failure
		return usecase.PermissionUpdateResponse{}
	}

	// Create an updated permission aggregate
	permission := aggregate.Permission{
		PermissionID:   req.PermissionID,
		PermissionName: req.PermissionName,
		Description:    req.Description,
		PermissionCode: req.PermissionCode,
		SensitiveFlag:  req.SensitiveFlag,
		CreatedAt:      time.Now(), // Using current time for simplicity
	}

	// Save the updated permission to the database
	err = ps.permissionRepo.SavePermission(ctx, permission)
	if err != nil {
		// If saving fails, return an empty response
		return usecase.PermissionUpdateResponse{}
	}

	// Return the updated permission
	return usecase.PermissionUpdateResponse{
		PermissionID:   permission.PermissionID,
		PermissionName: permission.PermissionName,
		PermissionCode: permission.PermissionCode,
	}
}

// DeletePermission deletes a permission by its ID
// It returns a response indicating whether the deletion was successful
// If deletion fails, it returns an error response
func (ps *permissionService) DeletePermission(ctx *context.Context, req usecase.PermissionDeleteRequest) usecase.PermissionDeleteResponse {
	// Check if the permission is sensitive
	existingPermission, err := ps.permissionRepo.FindPermissionByID(ctx, req.PermissionID)
	if err != nil {
		// If permission is not found, return an empty response
		return usecase.PermissionDeleteResponse{}
	}
	
	// If permission is sensitive, reject the deletion
	if existingPermission.SensitiveFlag {
		// Return an empty response to indicate failure
		return usecase.PermissionDeleteResponse{}
	}

	// Delete the permission from the database
	err = ps.permissionRepo.DeletePermission(ctx, req.PermissionID)
	if err != nil {
		// If deletion fails, return an empty response
		return usecase.PermissionDeleteResponse{}
	}

	// Return a success response
	return usecase.PermissionDeleteResponse{
		Success: true,
	}
}

// ListPermissions retrieves all permissions matching the specified criteria
// It returns a response containing the list of permissions
// If retrieval fails, it returns an error response
func (ps *permissionService) ListPermissions(ctx *context.Context, req usecase.PermissionListRequest) usecase.PermissionListResponse {
	// List permissions from the database
	permissions, err := ps.permissionRepo.ListPermissions(ctx, req.Limit, req.Offset)
	if err != nil {
		// If listing fails, return an empty response
		return usecase.PermissionListResponse{}
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

	// Return the list of permissions
	return usecase.PermissionListResponse{
		Permissions: permissionResponses,
	}
}

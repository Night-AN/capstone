package repository

import (
	"context"
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

// PermissionRepository defines the interface for permission data access operations
// It provides methods for creating, retrieving, updating, and deleting permissions
// All methods accept a context and return the requested data or an error
// The context is used for cancellation, timeouts, and passing request-scoped values
// The repository layer is responsible for all database interactions
// It should handle database-specific operations such as transactions,
// query construction, and result mapping
// The repository layer should be tested in isolation from the service layer
// using an actual database or a database-like store
// The repository layer should be stateless and thread-safe
// It should not maintain any internal state between method calls
// All dependencies should be injected via the constructor
// The repository layer should not contain any business logic
// It should only be concerned with data access and persistence
type PermissionRepository interface {
	// SavePermission saves a permission to the database
	// If the permission already exists, it updates the existing record
	// If the permission does not exist, it creates a new record
	// It returns an error if the operation fails
	SavePermission(ctx context.Context, permission aggregate.Permission) error

	// FindPermissionByID retrieves a permission by its ID
	// It returns the permission and an error if the operation fails
	// If the permission is not found, it returns an error
	FindPermissionByID(ctx context.Context, permissionID uuid.UUID) (aggregate.Permission, error)

	// FindPermissionByName retrieves permissions by their name
	// It returns a list of permissions and an error if the operation fails
	// If no permissions are found, it returns an empty list and no error
	FindPermissionByName(ctx context.Context, permissionName string) ([]aggregate.Permission, error)

	// FindPermissionByCode retrieves permissions by their code
	// It returns a list of permissions and an error if the operation fails
	// If no permissions are found, it returns an empty list and no error
	FindPermissionByCode(ctx context.Context, permissionCode string) ([]aggregate.Permission, error)

	// DeletePermission deletes a permission by its ID
	// It returns an error if the operation fails
	DeletePermission(ctx context.Context, permissionID uuid.UUID) error

	// ListPermissions retrieves all permissions matching the specified criteria
	// It returns a list of permissions and an error if the operation fails
	// If no permissions are found, it returns an empty list and no error
	ListPermissions(ctx context.Context, limit, offset int) ([]aggregate.Permission, error)
}

package repository

import (
	"context"
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

// RoleRepository defines the interface for role data access operations
// It provides methods for creating, retrieving, updating, and deleting roles
// as well as managing role-permission relationships
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
type RoleRepository interface {
	// SaveRole saves a role to the database
	// If the role already exists, it updates the existing record
	// If the role does not exist, it creates a new record
	// It returns an error if the operation fails
	SaveRole(ctx context.Context, role aggregate.Role) error

	// FindRoleByID retrieves a role by its ID
	// It returns the role and an error if the operation fails
	// If the role is not found, it returns an error
	FindRoleByID(ctx context.Context, roleID uuid.UUID) (aggregate.Role, error)

	// FindRoleByName retrieves roles by their name
	// It returns a list of roles and an error if the operation fails
	// If no roles are found, it returns an empty list and no error
	FindRoleByName(ctx context.Context, roleName string) ([]aggregate.Role, error)

	// FindRoleByCode retrieves roles by their code
	// It returns a list of roles and an error if the operation fails
	// If no roles are found, it returns an empty list and no error
	FindRoleByCode(ctx context.Context, roleCode string) ([]aggregate.Role, error)

	// FindAllRoles retrieves all roles
	// It returns a list of roles and an error if the operation fails
	// If no roles are found, it returns an empty list and no error
	FindAllRoles(ctx context.Context) ([]aggregate.Role, error)

	// DeleteRole deletes a role by its ID
	// It returns an error if the operation fails
	DeleteRole(ctx context.Context, roleID uuid.UUID) error

	// AssignPermission assigns a permission to a role
	// It returns an error if the operation fails
	AssignPermission(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) error

	// RemovePermission removes a permission from a role
	// It returns an error if the operation fails
	RemovePermission(ctx context.Context, roleID uuid.UUID, permissionID uuid.UUID) error

	// GetRolePermissions retrieves all permissions assigned to a role
	// It returns a list of permissions and an error if the operation fails
	// If no permissions are found, it returns an empty list and no error
	GetRolePermissions(ctx context.Context, roleID uuid.UUID) ([]aggregate.Permission, error)
}

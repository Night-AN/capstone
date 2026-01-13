package repository

import (
	"context"
	"errors"
	"moon/internal/domain/access_manage_control/aggregate"

	"github.com/google/uuid"
)

var (
	// UserNotFound is returned when a user cannot be found in the repository.
	UserNotFound = errors.New("user not found")
)

// UserRepository defines the contract for user data persistence operations.
// It abstracts the underlying storage mechanism for user aggregates.
type UserRepository interface {

	// SaveUser persists a user aggregate to the repository.
	// Creates a new user if it doesn't exist, or updates an existing one.
	// ctx: The context for cancellation and timeout control.
	// user: The user aggregate to be saved.
	// Returns an error if the operation fails.
	SaveUser(ctx *context.Context, user aggregate.User) error

	// ChangeUserStatus updates the status of an existing user.
	// ctx: The context for cancellation and timeout control.
	// user_id: The unique identifier of the user.
	// user_status: The new status value to set.
	// Returns UserNotFound if the user doesn't exist, or an error if the update fails.
	ChangeUserStatus(ctx *context.Context, user_id uuid.UUID, user_status string) error

	// FindUserByID retrieves a user aggregate by its UUID.
	// ctx: The context for cancellation and timeout control.
	// user_id: The unique identifier of the user.
	// Returns the user aggregate if found, or UserNotFound if not found.
	FindUserByID(ctx *context.Context, user_id uuid.UUID) (aggregate.User, error)

	// FindUserByEmail retrieves a user aggregate by its email address.
	// ctx: The context for cancellation and timeout control.
	// email: The email address of the user.
	// Returns the user aggregate if found, or UserNotFound if not found.
	FindUserByEmail(ctx *context.Context, email string) (aggregate.User, error)
}

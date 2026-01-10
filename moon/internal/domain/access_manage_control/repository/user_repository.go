package repository

import (
	"context"
	"errors"
	"moon/internal/domain/access_manage_control/aggregate"

	"github.com/google/uuid"
)

var (
	UserNotFound = errors.New("user not found")
)

type UserRepository interface {
	SaveUser(ctx *context.Context, user aggregate.User) error
	ChangeUserStatus(ctx *context.Context, user_id uuid.UUID, user_status string) error

	FindUserByID(ctx *context.Context, user_id uuid.UUID) (aggregate.User, error)
	FindUserByEmail(ctx *context.Context, email string) (aggregate.User, error)
}

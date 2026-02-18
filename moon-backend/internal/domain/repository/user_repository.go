package repository

import (
	"context"
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

type UserRepository interface {
	SaveUser(ctx *context.Context, user aggregate.User) error
	ChangeUserStatus(ctx *context.Context, user_id uuid.UUID, user_status string) error
	FindUserByID(ctx *context.Context, user_id uuid.UUID) (aggregate.User, error)
	FindUserByEmail(ctx *context.Context, email string) (aggregate.User, error)
	ListUsers(ctx *context.Context) ([]aggregate.User, error)
}

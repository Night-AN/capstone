package repository

import (
	"context"
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user aggregate.User) error
	DeleteUser(ctx context.Context, user_id uuid.UUID) error
	FindUserByID(ctx context.Context, user_id uuid.UUID) (aggregate.User, error)
	FindUserByEmail(ctx context.Context, email string) (aggregate.User, error)
	ListUsers(ctx context.Context) ([]aggregate.User, error)
	FindUsersByOrganizationID(ctx context.Context, organization_id uuid.UUID) ([]aggregate.User, error)
	AssignRoleToUser(ctx context.Context, user_id uuid.UUID, role_id uuid.UUID) error
	RemoveRoleFromUser(ctx context.Context, user_id uuid.UUID, role_id uuid.UUID) error
	FindRolesByUserID(ctx context.Context, user_id uuid.UUID) ([]aggregate.Role, error)
	FindUsersByRoleID(ctx context.Context, role_id uuid.UUID) ([]aggregate.User, error)
}

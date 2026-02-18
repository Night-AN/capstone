package repository

import (
	"context"
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

type ResourceRepository interface {
	SaveResource(ctx *context.Context, resource aggregate.Resource) error
	UpdateResource(ctx *context.Context, resource aggregate.Resource) error
	DeleteResource(ctx *context.Context, resource_id uuid.UUID) error
	FindResourceByID(ctx *context.Context, resource_id uuid.UUID) (aggregate.Resource, error)
	ListAllResources(ctx *context.Context) ([]aggregate.Resource, error)
}

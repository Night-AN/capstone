package usecase

import (
	"github.com/google/uuid"
)

type ResourceDeleteRequest struct {
	ResourceID uuid.UUID `json:"resource_id"`
}

type ResourceDeleteResponse struct {
	Success bool `json:"success"`
}

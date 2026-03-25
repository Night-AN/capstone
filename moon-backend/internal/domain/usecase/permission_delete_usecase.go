package usecase

import (
	"github.com/google/uuid"
)

type PermissionDeleteRequest struct {
	PermissionID uuid.UUID `json:"permission_id"`
}

type PermissionDeleteResponse struct {
	Success bool `json:"success"`
}

package usecase

import (
	"github.com/google/uuid"
)

type RoleDeleteRequest struct {
	RoleID uuid.UUID `json:"role_id"`
}

type RoleDeleteResponse struct {
	Success bool `json:"success"`
}

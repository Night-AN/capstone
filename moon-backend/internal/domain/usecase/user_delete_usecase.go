package usecase

import (
	"github.com/google/uuid"
)

type UserDeleteRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

type UserDeleteResponse struct {
	Success bool `json:"success"`
}

package usecase

import (
	"github.com/google/uuid"
)

type UserUpdateRequest struct {
	UserID         uuid.UUID  `json:"user_id"`
	Nickname       string     `json:"nickname"`
	FullName       string     `json:"full_name"`
	Email          string     `json:"email"`
	OrganizationID *uuid.UUID `json:"organization_id"`
}

type UserUpdateResponse struct {
	Success bool `json:"success"`
}

type UserChangePasswordRequest struct {
	UserID      uuid.UUID `json:"user_id"`
	OldPassword string    `json:"old_password"`
	Password    string    `json:"password"`
}

type UserChangePasswordResponse struct {
	Success bool `json:"success"`
}

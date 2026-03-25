package usecase

import (
	"github.com/google/uuid"
)

type UserRegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterResponse struct {
	Success bool `json:"success"`
}

type UserCreateRequest struct {
	Nickname       string    `json:"nickname"`
	FullName       string    `json:"full_name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	OrganizationID uuid.UUID `json:"organization_id"`
}

type UserCreateResponse struct {
	Success bool `json:"success"`
}

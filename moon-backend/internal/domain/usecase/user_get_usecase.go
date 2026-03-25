package usecase

import (
	"time"

	"github.com/google/uuid"
)

type UserItem struct {
	UserID         uuid.UUID `json:"user_id"`
	Nickname       string    `json:"nickname"`
	FullName       string    `json:"full_name"`
	Email          string    `json:"email"`
	OrganizationID uuid.UUID `json:"organization_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Success bool `json:"success"`
}

type UserGetRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

type UserGetResponse = UserItem

type UserListRequest struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Count    int64 `json:"count"`
	Total    int64 `json:"total"`
}

type UserListResponse struct {
	Users      []UserItem `json:"users"`
	Total      int64      `json:"total"`
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
	TotalPages int        `json:"total_pages"`
}

type UserGetRoleRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

type UserGetRoleResponse struct {
	UserID uuid.UUID       `json:"user_id"`
	Roles  []*RoleTreeNode `json:"roles"`
}

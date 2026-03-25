package usecase

import (
	"time"

	"github.com/google/uuid"
)

type RoleItem struct {
	RoleID          uuid.UUID `json:"role_id"`
	RoleName        string    `json:"role_name"`
	RoleDescription *string   `json:"role_description"`
	RoleCode        string    `json:"role_code"`
	RoleFlag        string    `json:"role_flag"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type RoleTreeNode struct {
	RoleID       uuid.UUID `json:"role_id"`
	RoleName     string    `json:"role_name"`
	RoleCode     string    `json:"role_code"`
	RoleCodePart string    `json:"role_code_part"`

	PermissionTree []*PermissionTreeNode    `json:"permission_tree"`
	Children       map[string]*RoleTreeNode `json:"children"`
}

type RoleGetRequest struct {
	RoleID uuid.UUID `json:"role_id"`
}

type RoleGetResponse = RoleItem

type RoleGetPermissionResponse struct {
	RoleID      uuid.UUID             `json:"role_id"`
	Permissions []*PermissionTreeNode `json:"permissions"`
}

type RoleListRequest struct {
	Page     int   `json:"page"`
	PageSize int   `json:"page_size"`
	Count    int64 `json:"count"`
	Total    int64 `json:"total"`
}

type RoleListResponse = []*RoleItem

type RoleListTreeRequest struct {
	RoleID *uuid.UUID `json:"role_id"`
}

type RoleListTreeResponse struct {
	RoleTree *RoleTreeNode `json:"role_tree"`
}

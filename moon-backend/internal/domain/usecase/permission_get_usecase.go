package usecase

import (
	"time"

	"github.com/google/uuid"
)

type PermissionItem struct {
	PermissionID          uuid.UUID `json:"permission_id"`
	PermissionName        string    `json:"permission_name"`
	PermissionCode        string    `json:"permission_code"`
	PermissionDescription *string   `json:"permission_description"`
	PermissionFlag        string    `json:"permission_flag"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type PermissionGetRequest struct {
	PermissionID uuid.UUID `json:"permission_id"`
}

type PermissionGetResponse = PermissionItem

type PermissionListTreeRequest struct {
	RoleID *uuid.UUID `json:"role_id"`
}

type PermissionListTreeResponse struct {
	PermissionTree *PermissionTreeNode `json:"permission_tree"`
}

type PermissionTreeNode struct {
	PermissionID       uuid.UUID                      `json:"permission_id"`
	PermissionName     string                         `json:"permission_name"`
	PermissionCode     string                         `json:"permission_code"`
	PermissionCodePart string                         `json:"permission_code_part"`
	Children           map[string]*PermissionTreeNode `json:"children"`
}

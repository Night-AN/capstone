package usecase

import (
	"github.com/google/uuid"
)

type PermissionUpdateRequest struct {
	PermissionID          uuid.UUID `json:"permission_id"`
	PermissionName        string    `json:"permission_name"`
	PermissionCode        string    `json:"permission_code"`
	PermissionDescription *string   `json:"permission_description"`
	PermissionFlag        string    `json:"permission_flag"`
}

type PermissionUpdateResponse struct {
	Success bool `json:"success"`
}

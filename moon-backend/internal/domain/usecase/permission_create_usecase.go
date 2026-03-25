package usecase

type PermissionCreateRequest struct {
	PermissionName        string  `json:"permission_name"`
	PermissionCode        string  `json:"permission_code"`
	PermissionDescription *string `json:"permission_description"`
	PermissionFlag        string  `json:"permission_flag"`
}

type PermissionCreateResponse struct {
	Success bool `json:"success"`
}

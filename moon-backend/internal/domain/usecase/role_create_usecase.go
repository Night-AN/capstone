package usecase

type RoleCreateRequest struct {
	RoleName        string  `json:"role_name"`
	RoleCode        string  `json:"role_code"`
	RoleDescription *string `json:"role_description"`
	RoleFlag        string  `json:"role_flag"`
}

type RoleCreateResponse struct {
	Success bool `json:"success"`
}

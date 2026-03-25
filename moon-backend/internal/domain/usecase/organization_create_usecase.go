package usecase

type OrganizationCreateRequest struct {
	OrganizationName        string `json:"organization_name"`
	OrganizationCode        string `json:"organization_code"`
	OrganizationDescription string `json:"organization_description"`
	OrganizationFlag        string `json:"organization_flag"`
}

type OrganizationCreateResponse struct {
	Success bool `json:"success"`
}

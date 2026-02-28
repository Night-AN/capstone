package usecase

type ModelConfigDeleteRequest struct {
	ConfigID string `json:"config_id"`
}

type ModelConfigDeleteResponse struct {
	Success bool `json:"success"`
}

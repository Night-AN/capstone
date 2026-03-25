package usecase

type IntangibleAssetTypeCreateRequest struct {
	IntangibleAssetTypeName string `json:"intangible_asset_type_name"`
	IntangibleAssetTypeCode string `json:"intangible_asset_type_code"`
	IntangibleAssetTypeFlag string `json:"intangible_asset_type_flag"`
}

type IntangibleAssetTypeCreateResponse struct {
	Success bool `json:"success"`
}

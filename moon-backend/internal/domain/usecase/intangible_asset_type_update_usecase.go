package usecase

// IntangibleAssetTypeUpdateRequest represents the request for updating an intangible asset type
type IntangibleAssetTypeUpdateRequest struct {
	IntangibleAssetTypeName string `json:"intangible_asset_type_name"`
	IntangibleAssetTypeCode string `json:"intangible_asset_type_code"`
	IntangibleAssetTypeFlag string `json:"intangible_asset_type_flag"`
}

// IntangibleAssetTypeUpdateResponse represents the response for updating an intangible asset type
type IntangibleAssetTypeUpdateResponse struct {
	Success bool `json:"success"`
}

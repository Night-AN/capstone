package usecase

import (
	"time"

	"github.com/google/uuid"
)

type IntangibleAssetTypeItem struct {
	IntangibleAssetTypeID   uuid.UUID  `json:"intangible_asset_type_id"`
	IntangibleAssetTypeName string     `json:"intangible_asset_type_name"`
	IntangibleAssetTypeCode string     `json:"intangible_asset_type_code"`
	IntangibleAssetTypeFlag string     `json:"intangible_asset_type_flag"`
	CreatedAt               time.Time  `json:"created_at"`
	UpdatedAt               *time.Time `json:"updated_at"`
}

type IntangibleAssetTypeTreeNode struct {
	IntangibleAssetTypeName     string                                  `json:"intangible_asset_type_name"`
	IntangibleAssetTypeCode     string                                  `json:"intangible_asset_type_code"`
	IntangibleAssetTypeCodePart string                                  `json:"intangible_asset_type_code_part"`
	Children                    map[string]*IntangibleAssetTypeTreeNode `json:"children"`
}

type IntangibleAssetTypeGetRequest struct {
	IntangibleAssetTypeID uuid.UUID `json:"intangible_asset_type_id"`
}

type IntangibleAssetTypeGetResponse = IntangibleAssetTypeItem

type IntangibleAssetTypeListRequest struct {
	OrganizationID *uuid.UUID `json:"organization_id"`
}

type IntangibleAssetTypeListResponse struct {
	IntangibleAssetTypes []*IntangibleAssetTypeItem `json:"intangible_asset_types"`
}

type IntangibleAssetTypeTreeRequest struct {
	IntangibleAssetTypeID uuid.UUID `json:"intangible_asset_type_id"`
}

type IntangibleAssetTypeTreeResponse struct {
	IntangibleAssetTypeTree []*IntangibleAssetTypeTreeNode `json:"intangible_asset_types"`
}

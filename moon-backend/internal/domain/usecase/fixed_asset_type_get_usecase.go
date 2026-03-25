package usecase

import (
	"time"

	"github.com/google/uuid"
)

type FixedAssetTypeItem struct {
	FixedAssetTypeID   uuid.UUID `json:"fixed_asset_type_id"`
	FixedAssetTypeName string    `json:"fixed_asset_type_name"`
	FixedAssetTypeCode string    `json:"fixed_asset_type_code"`
	FixedAssetTypeFlag string    `json:"fixed_asset_type_flag"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type FixedAssetTypeTreeNode struct {
	FixedAssetTypeName     string                             `json:"fixed_asset_type_name"`
	FixedAssetTypeCode     string                             `json:"fixed_asset_type_code"`
	FixedAssetTypeCodePart string                             `json:"fixed_asset_type_code_part"`
	Children               map[string]*FixedAssetTypeTreeNode `json:"children"`
}

type FixedAssetTypeGetRequest struct {
	FixedAssetTypeID uuid.UUID `json:"fixed_asset_type_id"`
}

type FixedAssetTypeGetResponse = FixedAssetTypeItem

type FixedAssetTypeListRequest struct {
	OrganizationID *uuid.UUID `json:"organization_id"`
}

type FixedAssetTypeListResponse struct {
	FixedAssetTypes []*FixedAssetTypeItem `json:"fixed_asset_types"`
}

type FixedAssetTypeTreeRequest struct {
	FixedAssetTypeID uuid.UUID `json:"fixed_asset_type_id"`
}

type FixedAssetTypeTreeResponse struct {
	FixedAssetTypes []*FixedAssetTypeTreeNode `json:"fixed_asset_types"`
}

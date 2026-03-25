package usecase

import (
	"moon/internal/domain/aggregate"
	"time"

	"github.com/google/uuid"
)

type FixedAssetTypeCreateRequest struct {
	FixedAssetTypeName string `json:"fixed_asset_type_name"`
	FixedAssetTypeCode string `json:"fixed_asset_type_code"`
	FixedAssetTypeFlag string `json:"fixed_asset_type_flag"`
}

type FixedAssetTypeCreateResponse struct {
	Success bool `json:"success"`
}

func ConvertFixedAssetTypeCreateRequestToAggregate(req FixedAssetTypeCreateRequest) aggregate.FixedAssetType {
	return aggregate.FixedAssetType{
		FixedAssetTypeID:   uuid.New(),
		FixedAssetTypeName: req.FixedAssetTypeName,
		FixedAssetTypeCode: req.FixedAssetTypeCode,
		FixedAssetTypeFlag: req.FixedAssetTypeFlag,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
}

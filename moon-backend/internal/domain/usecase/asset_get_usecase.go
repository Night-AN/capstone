package usecase

import (
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

// AssetGetRequest defines the request structure for getting an asset
type AssetGetRequest struct {
	AssetID uuid.UUID `json:"asset_id"`
}

// AssetGetResponse defines the response structure for getting an asset
type AssetGetResponse struct {
	AssetID          uuid.UUID `json:"asset_id"`
	AssetName        string    `json:"asset_name"`
	AssetCode        string    `json:"asset_code"`
	AssetDescription string    `json:"asset_description"`
	AssetType        string    `json:"asset_type"`
	Manufacturer     string    `json:"manufacturer"`
	Model            string    `json:"model"`
	SerialNumber     string    `json:"serial_number"`
	IPAddress        string    `json:"ip_address"`
	MACAddress       string    `json:"mac_address"`
	Location         string    `json:"location"`
	Department       string    `json:"department"`
	Owner            string    `json:"owner"`
	Status           string    `json:"status"`
	PurchaseDate     string    `json:"purchase_date"`
	WarrantyEndDate  string    `json:"warranty_end_date"`
	CreatedAt        string    `json:"created_at"`
	UpdatedAt        string    `json:"updated_at"`
}

// ConvertAssetAggregateToAssetGetResponse converts an asset aggregate to an asset get response
func ConvertAssetAggregateToAssetGetResponse(asset aggregate.Asset) AssetGetResponse {
	response := AssetGetResponse{
		AssetID:          asset.AssetID,
		AssetName:        asset.AssetName,
		AssetCode:        asset.AssetCode,
		AssetDescription: asset.Description,
		AssetType:        asset.AssetType,
		Manufacturer:     asset.Manufacturer,
		Model:            asset.Model,
		SerialNumber:     asset.SerialNumber,
		IPAddress:        asset.IPAddress,
		MACAddress:       asset.MACAddress,
		Location:         asset.Location,
		Department:       asset.Department,
		Owner:            asset.Owner,
		Status:           asset.Status,
		PurchaseDate:     asset.PurchaseDate,
		WarrantyEndDate:  asset.WarrantyEndDate,
		CreatedAt:        asset.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:        asset.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return response
}

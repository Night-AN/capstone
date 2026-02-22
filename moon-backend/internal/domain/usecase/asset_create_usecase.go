package usecase

import (
	"moon/internal/domain/aggregate"
	"time"

	"github.com/google/uuid"
)

// AssetCreateRequest defines the request structure for creating an asset
type AssetCreateRequest struct {
	AssetName        string `json:"asset_name"`
	AssetCode        string `json:"asset_code"`
	AssetDescription string `json:"asset_description"`
	AssetType        string `json:"asset_type"`
	AssetClass       string `json:"asset_class"`
	Manufacturer     string `json:"manufacturer"`
	Model            string `json:"model"`
	SerialNumber     string `json:"serial_number"`
	IPAddress        string `json:"ip_address"`
	MACAddress       string `json:"mac_address"`
	Location         string `json:"location"`
	Department       string `json:"department"`
	Owner            string `json:"owner"`
	ContactInfo      string `json:"contact_info"`
	Status           string `json:"status"`
	PurchaseDate     string `json:"purchase_date"`
	WarrantyEndDate  string `json:"warranty_end_date"`
	Value            string `json:"value"`
	Notes            string `json:"notes"`
}

// AssetCreateResponse defines the response structure for creating an asset
type AssetCreateResponse struct {
	AssetID   uuid.UUID `json:"asset_id"`
	AssetName string    `json:"asset_name"`
	AssetCode string    `json:"asset_code"`
}

// ConvertAssetCreateRequestToAssetAggregate converts an asset create request to an asset aggregate
func ConvertAssetCreateRequestToAssetAggregate(req AssetCreateRequest) aggregate.Asset {
	return aggregate.Asset{
		AssetID:          uuid.New(),
		AssetName:        req.AssetName,
		AssetCode:        req.AssetCode,
		Description:      req.AssetDescription,
		AssetType:        req.AssetType,
		Manufacturer:     req.Manufacturer,
		Model:            req.Model,
		SerialNumber:     req.SerialNumber,
		IPAddress:        req.IPAddress,
		MACAddress:       req.MACAddress,
		Location:         req.Location,
		Department:       req.Department,
		Owner:            req.Owner,
		Status:           req.Status,
		PurchaseDate:     req.PurchaseDate,
		WarrantyEndDate:  req.WarrantyEndDate,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}
}

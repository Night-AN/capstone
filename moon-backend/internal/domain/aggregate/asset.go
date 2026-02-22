package aggregate

import (
	"time"

	"github.com/google/uuid"
)

// Asset represents a network asset, such as a server, workstation, or network device
type Asset struct {
	// AssetID is the unique identifier for the asset
	AssetID uuid.UUID `gorm:"column:asset_id"`

	// AssetName is the name of the asset
	AssetName string `gorm:"column:asset_name"`

	// AssetCode is the unique code for the asset
	AssetCode string `gorm:"column:asset_code"`

	// AssetType is the type of the asset, such as server, workstation, network_device, etc.
	AssetType string `gorm:"column:asset_type"`

	// OrganizationID is the ID of the organization that the asset belongs to
	OrganizationID uuid.UUID `gorm:"column:organization_id"`

	// Manufacturer is the manufacturer of the asset
	Manufacturer string `gorm:"column:manufacturer"`

	// Model is the model of the asset
	Model string `gorm:"column:model"`

	// SerialNumber is the serial number of the asset
	SerialNumber string `gorm:"column:serial_number"`

	// IPAddress is the IP address of the asset
	IPAddress string `gorm:"column:ip_address"`

	// MACAddress is the MAC address of the asset
	MACAddress string `gorm:"column:mac_address"`

	// Location is the physical location of the asset
	Location string `gorm:"column:location"`

	// Department is the department that the asset belongs to
	Department string `gorm:"column:department"`

	// Owner is the person responsible for the asset
	Owner string `gorm:"column:owner"`

	// ContactInfo is the contact information for the asset owner
	ContactInfo string `gorm:"column:contact_info"`

	// Status is the status of the asset, such as active, inactive, maintenance, etc.
	Status string `gorm:"column:status"`

	// PurchaseDate is the date when the asset was purchased
	PurchaseDate string `gorm:"column:purchase_date"`

	// WarrantyEndDate is the date when the asset's warranty ends
	WarrantyEndDate string `gorm:"column:warranty_end_date"`

	// Value is the value of the asset
	Value string `gorm:"column:value"`

	// Notes is any additional notes about the asset
	Notes string `gorm:"column:notes"`

	// Description is a description of the asset
	Description string `gorm:"column:asset_description"`

	// CreatedAt is the timestamp when the asset was created
	CreatedAt time.Time `gorm:"column:created_at"`

	// UpdatedAt is the timestamp when the asset was last updated
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName specifies the table name for the Asset struct
func (Asset) TableName() string {
	return "biz.asset"
}

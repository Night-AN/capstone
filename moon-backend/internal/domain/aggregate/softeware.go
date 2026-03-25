package aggregate

import "github.com/google/uuid"

type Software struct {
	// SoftwareID is the unique identifier for the software
	SoftwareID uuid.UUID `gorm:"column:software_id"`

	// SoftwareName is the name of the software
	SoftwareName string `gorm:"column:software_name"`

	// SoftwareType is the type of the software, such as application, system, etc.
	SoftwareType string `gorm:"column:software_type"`

	// Version is the version of the software
	Version string `gorm:"column:version"`

	// LicenseType is the type of license for the software, such as open_source, proprietary, etc.
	LicenseType string `gorm:"column:license_type"`

	// Organization Reference
	// Usage: Access control, resource allocation, data segregation.
	Organization []Organization `gorm:"many2many:biz.software_organization;"`
}

func (Software) TableName() string {
	return "biz.software"
}

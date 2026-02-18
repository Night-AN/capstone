package aggregate

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// BoolString is a custom type that handles conversion between string and bool
// It is used to handle cases where the database stores boolean values as strings
// such as "true", "false", or empty string

type BoolString bool

// Scan implements the sql.Scanner interface
func (b *BoolString) Scan(value interface{}) error {
	if value == nil {
		*b = BoolString(false)
		return nil
	}

	switch v := value.(type) {
	case string:
		if v == "" || v == "false" || v == "0" {
			*b = BoolString(false)
		} else if v == "true" || v == "1" {
			*b = BoolString(true)
		} else {
			return fmt.Errorf("invalid boolean string: %s", v)
		}
	case bool:
		*b = BoolString(v)
	case []byte:
		if string(v) == "" || string(v) == "false" || string(v) == "0" {
			*b = BoolString(false)
		} else if string(v) == "true" || string(v) == "1" {
			*b = BoolString(true)
		} else {
			return fmt.Errorf("invalid boolean string: %s", string(v))
		}
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}

	return nil
}

// Value implements the driver.Valuer interface
func (b BoolString) Value() (driver.Value, error) {
	return bool(b), nil
}

// Resource represents a protected resource entity (menu, API endpoint, button, data, etc.)
// that can be assigned permissions. Supports hierarchical structure for menu navigation
// and logical grouping.
type Resource struct {

	// ResourceID is the unique identifier for the resource, using UUID to ensure global uniqueness.
	ResourceID uuid.UUID `gorm:"column:resource_id"`

	// ResourceName is the display name of the resource, used for UI presentation and permission assignment.
	// Examples: "User Management", "Export API", "Delete Button"
	ResourceName string `gorm:"column:resource_name"`

	// ResourceCode is the unique identifier for the resource, used in permission keys.
	// Constraints: Must be URL-safe, lowercase recommended, use hyphens instead of underscores.
	// Examples: "user-management", "export-api", "delete-btn"
	// Usage: Appears in PermissionCode as {resourceCode}:{type}:{action}
	ResourceCode string `gorm:"column:resource_code"`

	// ResourceDescription provides detailed explanation of the resource's purpose and security context.
	// Can be nil if self-explanatory.
	// Usage: Permission documentation, audit reports, admin UI help text
	ResourceDescription *string `gorm:"column:resource_description"`

	// ResourceFlag indicates whether the resource is sensitive and cannot be modified or deleted.
	// When true, related permissions cannot be updated or deleted.
	// Usage: Protecting critical system resources
	ResourceFlag BoolString `gorm:"column:resource_flag"`

	// ResourceType indicates the category of the resource, determining its usage context.
	// Valid values: "menu", "api", "button", "field", "data", "report"
	// Usage: Drives UI rendering, permission grouping, and access control logic
	ResourceType string `gorm:"column:resource_type"`

	// ResourcePath stores the unique identifier path for hierarchical resources.
	// For menus: The frontend route path (e.g., "/system/user-management")
	// For APIs: The backend endpoint pattern (e.g., "/api/v1/users/*")
	// For buttons: The component identifier (e.g., "user-list:delete-btn")
	// Usage: Frontend routing, API pattern matching, UI component binding
	ResourcePath *string `gorm:"column:resource_path"`

	// RequestMethod indicates the HTTP method for API resources.
	// Valid values: "GET", "POST", "PUT", "DELETE", "PATCH"
	// Usage: API routing and access control
	RequestMethod *string `gorm:"column:request_method"`

	// SensitiveFlag marks permissions that grant access to sensitive data or critical operations.
	// When true, additional audit logging and approval workflows may be triggered.
	// Examples: Permissions accessing PII, financial data, or administrative actions should have this flag set to true
	// Usage: Drives enhanced logging, requires additional approval, and flags in compliance reports
	SensitiveFlag bool `gorm:"column:sensitive_flag"`

	// CreatedAt records the timestamp when the resource was first registered in the system.
	// Usage: Auditing, resource lifecycle management, compliance reporting.
	CreatedAt time.Time `gorm:"column:created_at"`

	// UpdatedAt records the timestamp when the resource was last updated in the system.
	// Usage: Auditing, resource lifecycle management, compliance reporting.
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

// TableName specifies the table name for the Resource struct
func (Resource) TableName() string {
	return "systems.resource"
}

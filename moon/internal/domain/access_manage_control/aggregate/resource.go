package aggregate

import (
	"time"

	"github.com/google/uuid"
)

// Resource represents a protected resource entity (menu, API endpoint, button, data, etc.)
// that can be assigned permissions. Supports hierarchical structure for menu navigation
// and logical grouping.
type Resource struct {

	// ResourceID is the unique identifier for the resource, using UUID to ensure global uniqueness.
	ResourceID uuid.UUID

	// ResourceName is the display name of the resource, used for UI presentation and permission assignment.
	// Examples: "User Management", "Export API", "Delete Button"
	ResourceName string

	// ResourceCode is the unique identifier for the resource, used in permission keys.
	// Constraints: Must be URL-safe, lowercase recommended, use hyphens instead of underscores.
	// Examples: "user-management", "export-api", "delete-btn"
	// Usage: Appears in PermissionCode as {resourceCode}:{type}:{action}
	ResourceCode string

	// ResourceType indicates the category of the resource, determining its usage context.
	// Valid values: "menu", "api", "button", "field", "data", "report"
	// Usage: Drives UI rendering, permission grouping, and access control logic
	ResourceType string

	// ParentResourceID is the UUID of the parent resource for hierarchical relationships.
	// Supporting scenarios like menu/sub-menu, module/sub-module structure.
	// For root-level resources, this should be nil.
	// Usage: Building resource trees, cascading permission inheritance
	ParentResourceID *uuid.UUID

	// ResourcePath stores the unique identifier path for hierarchical resources.
	// For menus: The frontend route path (e.g., "/system/user-management")
	// For APIs: The backend endpoint pattern (e.g., "/api/v1/users/*")
	// For buttons: The component identifier (e.g., "user-list:delete-btn")
	// Usage: Frontend routing, API pattern matching, UI component binding
	ResourcePath *string

	// DisplayOrder determines the sorting order in UI navigation and resource lists.
	// Smaller values appear first; default is 0.
	// Usage: Menu ordering, resource list display prioritization
	DisplayOrder int

	// IconClass stores the CSS class name or icon identifier for menu resources.
	// Only applicable when ResourceType is "menu".
	// Examples: "icon-user", "fa-solid fa-gear"
	// Usage: Frontend menu rendering, visual identification
	IconClass *string

	// Description provides detailed explanation of the resource's purpose and security context.
	// Can be nil if self-explanatory.
	// Usage: Permission documentation, audit reports, admin UI help text
	Description *string

	// IsEnabled indicates whether the resource is active and accessible in the system.
	// When false, related permissions become ineffective and UI elements are hidden.
	// Usage: Feature toggling, maintenance mode, gradual rollout
	IsEnabled bool

	// CreatedAt records the timestamp when the resource was first registered in the system.
	// Usage: Auditing, resource lifecycle management, compliance reporting.
	CreatedAt time.Time
}

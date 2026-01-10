package aggregate

import (
	"time"

	"github.com/google/uuid"
)

// Organization represents an organizational entity that supports hierarchical tree structure management.
// The OrganizationHierarchyPath field enables efficient hierarchical queries and permission inheritance
// without requiring recursive operations.
type Organization struct {

	// OrganizationID is the unique identifier for the organization, using UUID to ensure global uniqueness.
	OrganizationID uuid.UUID

	// OrganizationName is the display name of the organization, used for UI presentation and logging.
	OrganizationName string

	// OrganizationCode is the unique code for the organization, used to identify the node in the hierarchy path.
	// Constraint: Must be a URL/path-safe string;must using lowercase letters, numbers, and hyphens.
	// Examples: head-office="corp", technology-dept="tech", backend-team="backend"
	OrganizationCode string

	// OrganizationHierarchyPath stores the complete path code chain from the root organization to the current organization.
	// Format: Colon-separated OrganizationCodes without spaces.
	// Examples:
	//   Root organization (Head Office): "corp"
	//   Child organization (Tech Dept): "corp:tech"
	//   Grandchild organization (Backend Team): "corp:tech:backend"
	//
	// Core Use Cases:
	// 1. Quickly determine ancestor/descendant relationships between organizations (via simple string matching)
	// 2. Support efficient organization tree retrieval with LIKE-based queries
	// 3. Implement permission inheritance (parent orgs automatically get access to child org resources)
	// 4. Multi-tenant data isolation and visibility scope control
	// 5. Avoid recursive queries and improve performance
	//
	// Maintenance Rules:
	// - On creation: Set empty string for root; concatenate parent path for child organizations
	// - On move: Must cascade update this field for all descendant organizations
	// - On deletion: Must delete or transfer all child organizations first
	OrganizationHierarchyPath string

	// CreatedAt records the creation timestamp for auditing, sorting, and statistics.
	CreatedAt time.Time
}

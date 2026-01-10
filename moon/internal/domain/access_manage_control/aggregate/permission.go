package aggregate

import (
	"time"

	"github.com/google/uuid"
)

// Permission represents a granular access control entity that defines a specific operation or access right.
// Permissions follow the standard code format: {resource}:{type}:{action}:{subaction?}:{scope?}
// Permissions are collected into roles and assigned to users for fine-grained authorization.
type Permission struct {
	// PermissionID is the unique identifier for the permission, using UUID to ensure global uniqueness
	// and prevent collision across distributed systems.
	PermissionID uuid.UUID

	// PermissionName is the human-readable display name for UI, logging, and documentation.
	// Examples: "Create User", "Export Sensitive Data", "View System Menu"
	PermissionName string

	// Description provides detailed explanation of what this permission grants and its security implications.
	// Can be nil if no detailed description is needed.
	// Usage: Shown in permission management UI, audit reports, and access review workflows
	Description *string

	// Key is the programmatic identifier used in code for permission checks, following the standard format.
	// Standard Format: {resource}:{type}:{action}:{subaction?}:{scope?}
	//   - resource: Resource name (e.g., user, role, data, menu)
	//   - type: Permission type (e.g., api, button, menu, field)
	//   - action: Action name (e.g., create, read, update, delete, export)
	//   - subaction: Optional sub-action (e.g., excel, pdf, batch)
	//   - scope: Optional scope qualifier (e.g., own, all, tenant, admin)
	//
	// Examples:
	//   "user:api:create" - API permission to create users
	//   "user:api:update:batch" - API permission to batch update users
	//   "user:api:delete:admin" - Admin API permission to delete users
	//   "menu:view:user-management" - Menu permission to view user management
	//   "data:export:excel:sensitive" - Permission to export sensitive data to Excel
	//   "role:button:assign:tenant" - Button permission to assign roles within tenant scope
	//
	// Constraints:
	//   - Must be unique, URL-safe, lowercase recommended
	//   - Colon-separated hierarchy, avoid special characters
	//   - Wildcard support: user:*:view matches all user-related view
	PermissionCode string

	// SensitiveFlag marks permissions that grant access to sensitive data or critical operations.
	// When true, additional audit logging and approval workflows may be triggered.
	// Examples: Permissions accessing PII, financial data, or administrative actions should have this flag set to true
	// Usage: Drives enhanced logging, requires additional approval, and flags in compliance reports
	SensitiveFlag bool

	// CreatedAt records the timestamp when the permission was first defined in the system.
	// Used for auditing, permission lifecycle management, and compliance reporting.
	CreatedAt time.Time
}

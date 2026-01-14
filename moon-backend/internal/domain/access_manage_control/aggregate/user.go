package aggregate

import (
	"time"

	"github.com/google/uuid"
)

// User represents a system user entity that contains authentication and profile information.
// It serves as the core identity for access control and auditing purposes.
type User struct {

	// UserID is the unique identifier for the user, using UUID to ensure global uniqueness
	// and prevent ID collision in distributed systems.
	UserID uuid.UUID

	// Username is the unique login name used by the user to authenticate into the system.
	// Constraints: Typically unique, URL-safe, supports letters, numbers and specific symbols, recommended length 3-50 characters.
	// Usage: Used for authentication, audit logging and display purposes.
	Username string

	// HashedPassword is the password processed with strong hashing algorithms like bcrypt/argon2.
	// Never store plaintext passwords - this is a fundamental security requirement.
	// Usage: Compare hash values during authentication.
	HashedPassword string

	// Email is the user's email address, typically used for account recovery, notifications and important communications.
	// Constraints: Must conform to standard email format, usually unique or requires verification in the system.
	// Usage: Password reset, system notifications, two-factor authentication, etc.
	Email string

	// PhoneNumber is the user's mobile number, used for SMS verification and two-factor authentication.
	// Constraints: Recommended to store in international format (E.164), e.g., "+8613800138000".
	// Usage: SMS OTP, account security verification, emergency contact.
	PhoneNumber string

	// AccountStatus indicates the current status of the user account.
	// Common values: "Active", "Inactive", "Suspended", "Deleted", "Locked".
	// Usage: Controls account access, compliance requirements, lifecycle management.
	AccountStatus string

	// CreatedAt records the timestamp when the user account was created.
	// Usage: Auditing, account lifecycle analysis, compliance reporting.
	CreatedAt time.Time
}

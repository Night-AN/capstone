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
	UserID uuid.UUID `gorm:"column:user_id;type:uuid"`

	// Nickname is the display name chosen by the user for public identification.
	// Constraints: Typically non-unique, supports unicode characters, recommended length 2-50 characters.
	// Usage: Displayed in UI, social interactions, and public-facing contexts.
	Nickname string `gorm:"column:nickname;type:text"`

	// FullName is the user's complete legal or real name for official identification.
	// Constraints: Supports unicode characters, recommended length 2-100 characters.
	// Usage: Legal compliance, official documentation, identity verification.
	FullName string `gorm:"column:full_name;type:text"`

	// Email is the user's email address, typically used for account recovery, notifications and important communications.
	// Constraints: Must conform to standard email format, usually unique or requires verification in the system.
	// Usage: Password reset, system notifications, two-factor authentication, etc.
	Email string `gorm:"column:email;type:text"`

	// PasswordHash is the password processed with strong hashing algorithms like bcrypt/argon2.
	// Never store plaintext passwords - this is a fundamental security requirement.
	// Usage: Compare hash values during authentication.
	PasswordHash string `gorm:"column:password_hash;type:text"`

	// CreatedAt records the timestamp when the user account was created.
	// Usage: Auditing, account lifecycle analysis, compliance reporting.
	CreatedAt time.Time `gorm:"column:created_at;type:timestamptz"`

	// UpdatedAt records the timestamp when the user account was last modified.
	// Usage: Change tracking, cache invalidation, concurrency control, data synchronization.
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamptz"`
}

func (User) TableName() string {
	return "systems.users"
}

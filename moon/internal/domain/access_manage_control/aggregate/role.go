package aggregate

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	RoleID      uuid.UUID
	RoleName    string
	Description *string
	RoleCode    string
	SystemFlag  bool

	Permissions []Permission
	CreatedAt   time.Time
}

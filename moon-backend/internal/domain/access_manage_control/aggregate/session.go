package aggregate

import "github.com/google/uuid"

type Session struct {
	SessionID    uuid.UUID
	UserID       uuid.UUID
	AccessToken  string
	RefreshToken string
}

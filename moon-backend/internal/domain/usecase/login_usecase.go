package usecase

import (
	"moon/internal/domain/aggregate"
	"time"

	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	UserID    uuid.UUID `json:"user_id"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
}

func ConvertUserAggregateToLoginResponse(user aggregate.User) LoginResponse {
	return LoginResponse{
		user.UserID,
		user.Nickname,
		user.CreatedAt,
	}

}

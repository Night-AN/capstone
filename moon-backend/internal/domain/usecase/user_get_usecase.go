package usecase

import (
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

type UserGetRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

type UserGetResponse struct {
	UserID   uuid.UUID `json:"user_id"`
	Nickname string    `json:"nickname"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
}

func ConvertUserAggregateToUserGetResponse(user aggregate.User) UserGetResponse {
	return UserGetResponse{
		UserID:   user.UserID,
		Nickname: user.Nickname,
		FullName: user.FullName,
		Email:    user.Email,
	}
}

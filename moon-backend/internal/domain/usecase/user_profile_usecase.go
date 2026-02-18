package usecase

import (
	"moon/internal/domain/aggregate"
	"time"

	"github.com/google/uuid"
)

type UserProfileRequest struct {
	UserID uuid.UUID
}

type UserProfileResponse struct {
	UserID uuid.UUID

	Nickname string

	FullName string

	Email string

	CreatedAt time.Time

	UpdatedAt time.Time
}

func ConvertUserAggregateToUserProfileResponse(user aggregate.User) UserProfileResponse {
	return UserProfileResponse{
		user.UserID,
		user.Nickname,
		user.FullName,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	}

}

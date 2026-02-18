package usecase

import (
	"moon/internal/domain/aggregate"
	"time"

	"github.com/google/uuid"
)

type UserCreateRequest struct {
	Nickname string `json:"nickname"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateResponse struct {
	UserID   uuid.UUID `json:"user_id"`
	Nickname string    `json:"nickname"`
	Email    string    `json:"email"`
}

func ConvertUserCreateRequestToUserAggregate(req UserCreateRequest) aggregate.User {
	return aggregate.User{
		UserID:    uuid.New(),
		Nickname:  req.Nickname,
		FullName:  req.FullName,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

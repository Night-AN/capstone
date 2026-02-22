package usecase

import (
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

type UserUpdateRequest struct {
	UserID         uuid.UUID  `json:"user_id"`
	Nickname       string     `json:"nickname"`
	FullName       string     `json:"full_name"`
	Email          string     `json:"email"`
	OrganizationID *uuid.UUID `json:"organization_id"`
}

type UserUpdateResponse struct {
	UserID   uuid.UUID `json:"user_id"`
	Nickname string    `json:"nickname"`
	Email    string    `json:"email"`
}

func ConvertUserUpdateRequestToUserAggregate(req UserUpdateRequest) aggregate.User {
	return aggregate.User{
		UserID:         req.UserID,
		Nickname:       req.Nickname,
		FullName:       req.FullName,
		Email:          req.Email,
		OrganizationID: req.OrganizationID,
	}
}

package usecase

import (
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

type UserListRequest struct {
}

type UserListResponse struct {
	Users []UserListItem `json:"users"`
}

type UserListItem struct {
	UserID   uuid.UUID `json:"user_id"`
	Nickname string    `json:"nickname"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
}

func ConvertUserAggregatesToUserListResponse(users []aggregate.User) UserListResponse {
	response := UserListResponse{
		Users: make([]UserListItem, len(users)),
	}

	for i, user := range users {
		response.Users[i] = UserListItem{
			UserID:   user.UserID,
			Nickname: user.Nickname,
			FullName: user.FullName,
			Email:    user.Email,
		}
	}

	return response
}

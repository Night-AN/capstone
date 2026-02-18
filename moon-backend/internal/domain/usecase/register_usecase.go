package usecase

import "moon/internal/domain/aggregate"

type RegisterRequest struct {
	Nickname string `json:"nickname"`

	FullName string `json:"full_name"`

	Email string `json:"email"`

	PasswordHash string `json:"password"`
}

type RegisterResponse struct {
	Status bool `json:"status"`
}

func ConvertRegisterRequestToUserAggregate(req RegisterRequest) aggregate.User {
	return aggregate.User{
		Nickname:     req.Nickname,
		FullName:     req.FullName,
		Email:        req.Email,
		PasswordHash: req.PasswordHash,
	}
}

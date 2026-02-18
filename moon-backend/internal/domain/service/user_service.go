package service

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/errors"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"

	"github.com/google/uuid"
)

type UserService interface {
	Login(ctx *context.Context, email string, password string) (usecase.LoginResponse, errors.DomainError)
	Logout(ctx *context.Context, user_id uuid.UUID) (bool, errors.DomainError)
	Register(ctx *context.Context, user usecase.RegisterRequest) (usecase.RegisterResponse, errors.DomainError)

	QueryUserProfile(ctx *context.Context, user_id uuid.UUID) (usecase.UserProfileResponse, errors.DomainError)
	UpdateUserProfile(ctx *context.Context, user aggregate.User) (bool, errors.DomainError)

	CreateUser(ctx *context.Context, req usecase.UserCreateRequest) (usecase.UserCreateResponse, errors.DomainError)
	UpdateUser(ctx *context.Context, req usecase.UserUpdateRequest) (usecase.UserUpdateResponse, errors.DomainError)
	DeleteUser(ctx *context.Context, req usecase.UserDeleteRequest) (usecase.UserDeleteResponse, errors.DomainError)
	GetUserByID(ctx *context.Context, req usecase.UserGetRequest) (usecase.UserGetResponse, errors.DomainError)
	ListUsers(ctx *context.Context, req usecase.UserListRequest) (usecase.UserListResponse, errors.DomainError)
}

func NewUserService(user_repo repository.UserRepository) UserService {
	userSession := make(map[uuid.UUID]aggregate.User, 10)
	return &userService{user_repo, userSession}
}

type userService struct {
	UserRepository repository.UserRepository
	UserSession    map[uuid.UUID]aggregate.User
}

func (us *userService) Login(ctx *context.Context, email string, password string) (usecase.LoginResponse, errors.DomainError) {
	user, err := us.UserRepository.FindUserByEmail(ctx, email)
	if err != nil {
		delete(us.UserSession, user.UserID)
		return usecase.LoginResponse{}, errors.NewDomainWithError("401", "Login Err", err)

	}
	us.UserSession[user.UserID] = user
	return usecase.ConvertUserAggregateToLoginResponse(user), errors.DomainError{}
}

func (us *userService) Logout(ctx *context.Context, user_id uuid.UUID) (bool, errors.DomainError) {
	delete(us.UserSession, user_id)
	return true, errors.DomainError{}
}

func (us *userService) Register(ctx *context.Context, user usecase.RegisterRequest) (usecase.RegisterResponse, errors.DomainError) {
	req := usecase.ConvertRegisterRequestToUserAggregate(user)
	err := us.UserRepository.SaveUser(ctx, req)
	if err != nil {
		return usecase.RegisterResponse{Status: false}, errors.NewDomainError("401", "register err")
	}
	return usecase.RegisterResponse{Status: true}, errors.DomainError{}
}

func (us *userService) QueryUserProfile(ctx *context.Context, user_id uuid.UUID) (usecase.UserProfileResponse, errors.DomainError) {
	user, err := us.UserRepository.FindUserByID(ctx, user_id)
	if err != nil {
		return usecase.UserProfileResponse{}, errors.NewDomainWithError("401", "Query Profile Err", err)
	}
	usecase.ConvertUserAggregateToUserProfileResponse(user)
	return usecase.UserProfileResponse{}, errors.DomainError{}
}

func (us *userService) UpdateUserProfile(ctx *context.Context, user aggregate.User) (bool, errors.DomainError) {
	return true, errors.DomainError{}
}

func (us *userService) CreateUser(ctx *context.Context, req usecase.UserCreateRequest) (usecase.UserCreateResponse, errors.DomainError) {
	user := usecase.ConvertUserCreateRequestToUserAggregate(req)
	err := us.UserRepository.SaveUser(ctx, user)
	if err != nil {
		return usecase.UserCreateResponse{}, errors.NewDomainWithError("401", "Create User Err", err)
	}
	return usecase.UserCreateResponse{
		UserID:   user.UserID,
		Nickname: user.Nickname,
		Email:    user.Email,
	}, errors.DomainError{}
}

func (us *userService) UpdateUser(ctx *context.Context, req usecase.UserUpdateRequest) (usecase.UserUpdateResponse, errors.DomainError) {
	user := usecase.ConvertUserUpdateRequestToUserAggregate(req)
	err := us.UserRepository.SaveUser(ctx, user)
	if err != nil {
		return usecase.UserUpdateResponse{}, errors.NewDomainWithError("401", "Update User Err", err)
	}
	return usecase.UserUpdateResponse{
		UserID:   user.UserID,
		Nickname: user.Nickname,
		Email:    user.Email,
	}, errors.DomainError{}
}

func (us *userService) DeleteUser(ctx *context.Context, req usecase.UserDeleteRequest) (usecase.UserDeleteResponse, errors.DomainError) {
	// 使用ChangeUserStatus方法将用户状态改为"deleted"来实现删除功能
	err := us.UserRepository.ChangeUserStatus(ctx, req.UserID, "deleted")
	if err != nil {
		return usecase.UserDeleteResponse{Success: false}, errors.NewDomainWithError("401", "Delete User Err", err)
	}
	return usecase.UserDeleteResponse{Success: true}, errors.DomainError{}
}

func (us *userService) ListUsers(ctx *context.Context, req usecase.UserListRequest) (usecase.UserListResponse, errors.DomainError) {
	// 调用UserRepository.ListUsers方法获取所有用户
	users, err := us.UserRepository.ListUsers(ctx)
	if err != nil {
		return usecase.UserListResponse{}, errors.NewDomainWithError("401", "List Users Err", err)
	}

	// 将aggregate.User转换为usecase.UserListItem
	userItems := make([]usecase.UserListItem, len(users))
	for i, user := range users {
		userItems[i] = usecase.UserListItem{
			UserID:   user.UserID,
			Nickname: user.Nickname,
			FullName: user.FullName,
			Email:    user.Email,
		}
	}

	return usecase.UserListResponse{Users: userItems}, errors.DomainError{}
}

func (us *userService) GetUserByID(ctx *context.Context, req usecase.UserGetRequest) (usecase.UserGetResponse, errors.DomainError) {
	user, err := us.UserRepository.FindUserByID(ctx, req.UserID)
	if err != nil {
		return usecase.UserGetResponse{}, errors.NewDomainWithError("401", "Get User Err", err)
	}
	return usecase.ConvertUserAggregateToUserGetResponse(user), errors.DomainError{}
}

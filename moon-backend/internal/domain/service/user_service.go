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
	GetUsersByOrganizationID(ctx *context.Context, organization_id uuid.UUID) ([]usecase.UserListItem, errors.DomainError)

	AssignRoleToUser(ctx *context.Context, user_id uuid.UUID, role_id uuid.UUID) (bool, errors.DomainError)
	RemoveRoleFromUser(ctx *context.Context, user_id uuid.UUID, role_id uuid.UUID) (bool, errors.DomainError)
	GetUserRoles(ctx *context.Context, user_id uuid.UUID) ([]aggregate.Role, errors.DomainError)
	CheckUserHasRole(ctx *context.Context, user_id uuid.UUID, role_code string) (bool, errors.DomainError)
	CheckUserHasPermission(ctx *context.Context, user_id uuid.UUID, permission_code string) (bool, errors.DomainError)
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
	// 使用DeleteUser方法从数据库中删除用户记录
	err := us.UserRepository.DeleteUser(ctx, req.UserID)
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

func (us *userService) AssignRoleToUser(ctx *context.Context, user_id uuid.UUID, role_id uuid.UUID) (bool, errors.DomainError) {
	// 检查用户是否存在
	_, err := us.UserRepository.FindUserByID(ctx, user_id)
	if err != nil {
		return false, errors.NewDomainWithError("404", "User not found", err)
	}

	// 分配角色给用户
	err = us.UserRepository.AssignRoleToUser(ctx, user_id, role_id)
	if err != nil {
		return false, errors.NewDomainWithError("500", "Assign role failed", err)
	}

	return true, errors.DomainError{}
}

func (us *userService) RemoveRoleFromUser(ctx *context.Context, user_id uuid.UUID, role_id uuid.UUID) (bool, errors.DomainError) {
	// 检查用户是否存在
	_, err := us.UserRepository.FindUserByID(ctx, user_id)
	if err != nil {
		return false, errors.NewDomainWithError("404", "User not found", err)
	}

	// 从用户移除角色
	err = us.UserRepository.RemoveRoleFromUser(ctx, user_id, role_id)
	if err != nil {
		return false, errors.NewDomainWithError("500", "Remove role failed", err)
	}

	return true, errors.DomainError{}
}

func (us *userService) GetUserRoles(ctx *context.Context, user_id uuid.UUID) ([]aggregate.Role, errors.DomainError) {
	// 检查用户是否存在
	_, err := us.UserRepository.FindUserByID(ctx, user_id)
	if err != nil {
		return nil, errors.NewDomainWithError("404", "User not found", err)
	}

	// 获取用户的角色
	roles, err := us.UserRepository.FindRolesByUserID(ctx, user_id)
	if err != nil {
		return nil, errors.NewDomainWithError("500", "Get user roles failed", err)
	}

	return roles, errors.DomainError{}
}

func (us *userService) CheckUserHasRole(ctx *context.Context, user_id uuid.UUID, role_code string) (bool, errors.DomainError) {
	// 检查用户是否存在
	_, err := us.UserRepository.FindUserByID(ctx, user_id)
	if err != nil {
		return false, errors.NewDomainWithError("404", "User not found", err)
	}

	// 获取用户的角色
	roles, err := us.UserRepository.FindRolesByUserID(ctx, user_id)
	if err != nil {
		return false, errors.NewDomainWithError("500", "Get user roles failed", err)
	}

	// 检查用户是否有指定的角色
	for _, role := range roles {
		if role.RoleCode == role_code {
			return true, errors.DomainError{}
		}
	}

	return false, errors.DomainError{}
}

func (us *userService) CheckUserHasPermission(ctx *context.Context, user_id uuid.UUID, permission_code string) (bool, errors.DomainError) {
	// 检查用户是否存在
	_, err := us.UserRepository.FindUserByID(ctx, user_id)
	if err != nil {
		return false, errors.NewDomainWithError("404", "User not found", err)
	}

	// 获取用户的角色
	roles, err := us.UserRepository.FindRolesByUserID(ctx, user_id)
	if err != nil {
		return false, errors.NewDomainWithError("500", "Get user roles failed", err)
	}

	// 检查用户是否是SuperAdmin，如果是则拥有所有权限
	for _, role := range roles {
		if role.RoleCode == "sa" {
			return true, errors.DomainError{}
		}
	}

	// 这里简化实现，实际应该检查角色是否有指定的权限
	// 由于我们还没有实现角色-权限的检查逻辑，这里暂时返回false
	return false, errors.DomainError{}
}

func (us *userService) GetUsersByOrganizationID(ctx *context.Context, organization_id uuid.UUID) ([]usecase.UserListItem, errors.DomainError) {
	// 获取组织的用户列表
	users, err := us.UserRepository.FindUsersByOrganizationID(ctx, organization_id)
	if err != nil {
		return nil, errors.NewDomainWithError("500", "Get organization users failed", err)
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

	return userItems, errors.DomainError{}
}

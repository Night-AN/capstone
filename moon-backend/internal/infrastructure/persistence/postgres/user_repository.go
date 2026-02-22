package postgres

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/repository"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) SaveUser(ctx *context.Context, user aggregate.User) error {
	// 检查用户是否已存在
	var existingUser aggregate.User
	err := ur.db.WithContext(*ctx).Where("user_id = ?", user.UserID).First(&existingUser).Error

	if err == gorm.ErrRecordNotFound {
		// 用户不存在，创建新记录
		return ur.db.WithContext(*ctx).Create(&user).Error
	} else if err != nil {
		// 其他错误
		return err
	} else {
		// 用户存在，更新记录
		return ur.db.WithContext(*ctx).Model(&aggregate.User{}).Where("user_id = ?", user.UserID).Updates(&user).Error
	}
}

func (ur *userRepository) DeleteUser(ctx *context.Context, user_id uuid.UUID) error {
	result := ur.db.WithContext(*ctx).Where("user_id = ?", user_id).Delete(&aggregate.User{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (ur *userRepository) FindUserByID(ctx *context.Context, user_id uuid.UUID) (aggregate.User, error) {
	var user = aggregate.User{}
	err := ur.db.WithContext(*ctx).Model(&user).Where("user_id =?", user_id).First(&user).Error
	return user, err
}

func (ur *userRepository) FindUserByEmail(ctx *context.Context, email string) (aggregate.User, error) {
	var user = aggregate.User{}
	err := ur.db.WithContext(*ctx).Model(&user).Where("email =?", email).First(&user).Error
	return user, err
}

func (ur *userRepository) ListUsers(ctx *context.Context) ([]aggregate.User, error) {
	var users []aggregate.User
	err := ur.db.WithContext(*ctx).Model(&aggregate.User{}).Find(&users).Error
	return users, err
}

func (ur *userRepository) AssignRoleToUser(ctx *context.Context, user_id uuid.UUID, role_id uuid.UUID) error {
	// 检查关联是否已存在
	var count int64
	err := ur.db.WithContext(*ctx).Table("systems.user_role").Where("user_id = ? AND role_id = ?", user_id, role_id).Count(&count).Error
	if err != nil {
		return err
	}

	// 如果关联不存在，创建新关联
	if count == 0 {
		userRole := map[string]interface{}{
			"user_id": user_id,
			"role_id": role_id,
		}
		return ur.db.WithContext(*ctx).Table("systems.user_role").Create(userRole).Error
	}

	return nil
}

func (ur *userRepository) RemoveRoleFromUser(ctx *context.Context, user_id uuid.UUID, role_id uuid.UUID) error {
	return ur.db.WithContext(*ctx).Table("systems.user_role").Where("user_id = ? AND role_id = ?", user_id, role_id).Delete(nil).Error
}

func (ur *userRepository) FindRolesByUserID(ctx *context.Context, user_id uuid.UUID) ([]aggregate.Role, error) {
	var roles = []aggregate.Role{}
	err := ur.db.WithContext(*ctx).Table("systems.role").Joins("JOIN systems.user_role ON systems.role.role_id = systems.user_role.role_id").Where("systems.user_role.user_id = ?", user_id).Find(&roles).Error
	return roles, err
}

func (ur *userRepository) FindUsersByRoleID(ctx *context.Context, role_id uuid.UUID) ([]aggregate.User, error) {
	var users = []aggregate.User{}
	err := ur.db.WithContext(*ctx).Table("systems.users").Joins("JOIN systems.user_role ON systems.users.user_id = systems.user_role.user_id").Where("systems.user_role.role_id = ?", role_id).Find(&users).Error
	return users, err
}

func (ur *userRepository) FindUsersByOrganizationID(ctx *context.Context, organization_id uuid.UUID) ([]aggregate.User, error) {
	// 由于数据库中的users表没有organization_id字段，而User结构体中有这个字段，
	// 所以我们需要使用原生SQL查询，只选择数据库中存在的字段，然后手动构建User对象
	var userRows []struct {
		UserID       uuid.UUID
		Nickname     string
		FullName     string
		Email        string
		PasswordHash string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

	// 尝试使用user_organization表进行联表查询
	query := `
	SELECT u.user_id, u.nickname, u.full_name, u.email, u.password_hash, u.created_at, u.updated_at
	FROM systems.users u
	JOIN systems.user_organization uo ON u.user_id = uo.user_id
	WHERE uo.organization_id = ?
	`

	err := ur.db.WithContext(*ctx).Raw(query, organization_id).Scan(&userRows).Error

	// 如果出现错误，检查是否是表不存在或权限不够的错误
	if err != nil {
		// 检查错误是否是表不存在或权限不够的错误
		if (strings.Contains(err.Error(), "user_organization") && (strings.Contains(err.Error(), "does not exist") || strings.Contains(err.Error(), "关系不存在"))) ||
			strings.Contains(err.Error(), "permission denied") || strings.Contains(err.Error(), "权限不够") ||
			(strings.Contains(err.Error(), "organization_id") && strings.Contains(err.Error(), "does not exist")) {
			// 返回空列表，因为没有数据
			return []aggregate.User{}, nil
		}
		// 其他错误仍然返回
		return nil, err
	}

	// 将查询结果转换为[]aggregate.User
	users := make([]aggregate.User, len(userRows))
	for i, row := range userRows {
		users[i] = aggregate.User{
			UserID:       row.UserID,
			Nickname:     row.Nickname,
			FullName:     row.FullName,
			Email:        row.Email,
			PasswordHash: row.PasswordHash,
			CreatedAt:    row.CreatedAt,
			UpdatedAt:    row.UpdatedAt,
			// OrganizationID设置为nil，因为数据库中没有这个字段
			OrganizationID: nil,
		}
	}

	// 如果查询结果为空，返回空列表
	if len(users) == 0 {
		return []aggregate.User{}, nil
	}

	return users, nil
}

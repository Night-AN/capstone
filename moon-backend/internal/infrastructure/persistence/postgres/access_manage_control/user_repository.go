package accessmanagecontrol

import (
	"context"
	"moon/internal/domain/access_manage_control/aggregate"
	"moon/internal/domain/access_manage_control/repository"

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
	return ur.db.WithContext(*ctx).Save(user).Error
}

func (ur *userRepository) ChangeUserStatus(ctx *context.Context, user_id uuid.UUID, user_status string) error {
	result := ur.db.WithContext(*ctx).Model(&aggregate.User{}).
		Where("id = ?").Update("account_status", user_status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return repository.UserNotFound
	}

	return nil
}

func (ur *userRepository) FindUserByID(ctx *context.Context, user_id uuid.UUID) (aggregate.User, error) {
	var user = aggregate.User{}
	err := ur.db.WithContext(*ctx).Model(&user).Find("user_id =?", user_id).First(&user).Error
	return user, err
}

func (ur *userRepository) FindUserByEmail(ctx *context.Context, email string) (aggregate.User, error) {
	var user = aggregate.User{}
	err := ur.db.WithContext(*ctx).Model(&user).Find("email =?", email).First(&user).Error
	return user, err
}

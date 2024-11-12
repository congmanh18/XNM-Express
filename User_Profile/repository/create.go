package repository

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/entity"
)

// CreateUser implements UserRepository.
func (u *userRepository) CreateUser(ctx context.Context, user *entity.User) error {
	return u.DB.Executor.WithContext(ctx).Create(user).Debug().Error
}

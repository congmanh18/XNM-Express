package repository

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/entity"
)

// DeleteUser implements UserRepository.
func (u *userRepository) DeleteUser(ctx context.Context, id string) error {
	return u.DB.Executor.WithContext(ctx).Delete(&entity.User{}, id).Debug().Error
}

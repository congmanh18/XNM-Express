package repository

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/entity"
)

// GetUserByUsername implements UserRepository.
func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	panic("unimplemented")
}

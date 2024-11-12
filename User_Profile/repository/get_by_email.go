package repository

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/entity"
)

// GetUserByEmail implements UserRepository.
func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	panic("unimplemented")
}

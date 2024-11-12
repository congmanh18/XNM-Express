package repository

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/entity"
)

// GetUserByID implements UserRepository.
func (u *userRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	panic("unimplemented")
}

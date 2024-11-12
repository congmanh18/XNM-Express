package repository

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/entity"
)

func (u *userRepository) GetUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	panic("unimplemented")
}

package usecase

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/entity"
	"github.com/congmanh18/XNM-Express/User_Profile/repository"
	"github.com/congmanh18/lucas-core/security"
	ucerr "github.com/congmanh18/lucas-core/usecase"
)

type CreateUserUsecase interface {
	Execute(ctx context.Context, user *entity.User) *ucerr.UCError
}

type createUserUsecase struct {
	userRepo repository.UserRepository
}

func NewCreateUserUsecase(userRepo repository.UserRepository) CreateUserUsecase {
	return &createUserUsecase{userRepo: userRepo}
}

func (c *createUserUsecase) Execute(ctx context.Context, user *entity.User) *ucerr.UCError {
	hashPassword, err := security.HashPassword(*user.PasswordHash)
	if err != nil {
		return ucerr.New(
			ErrHashPassword.Message,
			ErrHashPassword.Code,
			err,
		)
	}

	user.PasswordHash = &hashPassword
	if err := c.userRepo.CreateUser(ctx, user); err != nil {
		return ucerr.New(
			ErrCreateUser.Message,
			ErrCreateUser.Code,
			err,
		)
	}
	return nil
}

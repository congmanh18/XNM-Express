package usecase

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/entity"
	"github.com/congmanh18/XNM-Express/User_Profile/repository"
	ucerr "github.com/congmanh18/lucas-core/usecase"
)

type UpdateUserUsecase interface {
	Execute(ctx context.Context, user *entity.User) *ucerr.UCError
}

type updateUserUsecase struct {
	userRepo repository.UserRepository
}

func NewUpdateUserUsecase(userRepo repository.UserRepository) UpdateUserUsecase {
	return &updateUserUsecase{userRepo: userRepo}
}

func (u *updateUserUsecase) Execute(ctx context.Context, user *entity.User) *ucerr.UCError {
	if err := u.userRepo.UpdateUser(ctx, user); err != nil {
		return ucerr.New(
			ErrUpdateUser.Message,
			ErrUpdateUser.Code,
			err,
		)
	}
	return nil
}

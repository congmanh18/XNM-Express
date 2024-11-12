package usecase

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/repository"
	ucerr "github.com/congmanh18/lucas-core/usecase"
)

type DeleteUserUsecase interface {
	Execute(ctx context.Context, id string) *ucerr.UCError
}

type deleteUserUsecase struct {
	userRepo repository.UserRepository
}

func NewDeleteUserUsecase(userRepo repository.UserRepository) DeleteUserUsecase {
	return &deleteUserUsecase{userRepo: userRepo}
}

func (d *deleteUserUsecase) Execute(ctx context.Context, id string) *ucerr.UCError {
	if err := d.userRepo.DeleteUser(ctx, id); err != nil {
		return ucerr.New(
			ErrDeleteUser.Message,
			ErrDeleteUser.Code,
			err,
		)
	}
	return nil
}

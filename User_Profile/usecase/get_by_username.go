package usecase

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/entity"
	"github.com/congmanh18/XNM-Express/User_Profile/repository"
	ucerr "github.com/congmanh18/lucas-core/usecase"
)

type GetUserByUsernameUsecase interface {
	Execute(ctx context.Context, username string) (*entity.User, *ucerr.UCError)
}

type getUserByUsernameUsecase struct {
	userRepo repository.UserRepository
}

func NewGetUserByUsernameUsecase(userRepo repository.UserRepository) GetUserByUsernameUsecase {
	return &getUserByUsernameUsecase{userRepo: userRepo}
}

func (g *getUserByUsernameUsecase) Execute(ctx context.Context, username string) (*entity.User, *ucerr.UCError) {
	user, err := g.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, ucerr.New(
			ErrGetUserByUsername.Message,
			ErrGetUserByUsername.Code,
			err,
		)
	}
	return user, nil
}

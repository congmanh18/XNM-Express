package usecase

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/entity"
	"github.com/congmanh18/XNM-Express/User_Profile/repository"
	ucerr "github.com/congmanh18/lucas-core/usecase"
)

type GetUserByEmailUsecase interface {
	Execute(ctx context.Context, email string) (*entity.User, *ucerr.UCError)
}

type getUserByEmailUsecase struct {
	userRepo repository.UserRepository
}

func NewGetUserByEmailUsecase(userRepo repository.UserRepository) GetUserByEmailUsecase {
	return &getUserByEmailUsecase{userRepo: userRepo}
}

func (g *getUserByEmailUsecase) Execute(ctx context.Context, email string) (*entity.User, *ucerr.UCError) {
	user, err := g.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, ucerr.New(
			ErrGetUserByEmail.Message,
			ErrGetUserByEmail.Code,
			err,
		)
	}
	return user, nil
}

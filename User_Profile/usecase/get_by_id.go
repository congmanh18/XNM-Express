package usecase

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/entity"
	"github.com/congmanh18/XNM-Express/User_Profile/repository"
	ucerr "github.com/congmanh18/lucas-core/usecase"
)

type GetUserByIDUsecase interface {
	Execute(ctx context.Context, id string) (*entity.User, *ucerr.UCError)
}

type getUserByIDUsecase struct {
	userRepo repository.UserRepository
}

func NewGetUserByIDUsecase(userRepo repository.UserRepository) GetUserByIDUsecase {
	return &getUserByIDUsecase{userRepo: userRepo}
}

func (g *getUserByIDUsecase) Execute(ctx context.Context, id string) (*entity.User, *ucerr.UCError) {
	user, err := g.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, ucerr.New(
			ErrGetUserByID.Message,
			ErrGetUserByID.Code,
			err,
		)
	}
	return user, nil
}

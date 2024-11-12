package usecase

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/entity"
	"github.com/congmanh18/XNM-Express/User_Profile/repository"
	ucerr "github.com/congmanh18/lucas-core/usecase"
)

type GetUserByPhoneUsecase interface {
	Execute(ctx context.Context, phone string) (*entity.User, *ucerr.UCError)
}

type getUserByPhoneUsecase struct {
	userRepo repository.UserRepository
}

func NewGetUserByPhoneUsecase(userRepo repository.UserRepository) GetUserByPhoneUsecase {
	return &getUserByPhoneUsecase{userRepo: userRepo}
}

func (g *getUserByPhoneUsecase) Execute(ctx context.Context, phone string) (*entity.User, *ucerr.UCError) {
	user, err := g.userRepo.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, ucerr.New(
			ErrGetUserByPhone.Message,
			ErrGetUserByPhone.Code,
			err,
		)
	}
	return user, nil
}

package repository

import (
	"context"

	"github.com/congmanh18/XNM-Express/User_Profile/entity"
	"github.com/congmanh18/lucas-core/db/postgresql"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByPhone(ctx context.Context, phone string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userRepository struct {
	DB *postgresql.Database
}

func NewUserRepository(db *postgresql.Database) UserRepository {
	return &userRepository{DB: db}
}

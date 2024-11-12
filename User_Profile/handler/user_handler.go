package handler

import (
	"github.com/congmanh18/XNM-Express/User_Profile/usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	HandleCreateUser(c echo.Context) error
	HandleUpdateUser(c echo.Context) error
	HandleDeleteUser(c echo.Context) error
	HandleGetUserByID(c echo.Context) error
	HandleGetUserByUsername(c echo.Context) error
	HandleGetUserByEmail(c echo.Context) error
}

type userHandler struct {
	createUserUsecase        usecase.CreateUserUsecase
	updateUserUsecase        usecase.UpdateUserUsecase
	deleteUserUsecase        usecase.DeleteUserUsecase
	getUserByIDUsecase       usecase.GetUserByIDUsecase
	getUserByUsernameUsecase usecase.GetUserByUsernameUsecase
	getUserByEmailUsecase    usecase.GetUserByEmailUsecase
}

// HandleDeleteUser implements UserHandler.
func (h *userHandler) HandleDeleteUser(c echo.Context) error {
	panic("unimplemented")
}

// HandleGetUserByEmail implements UserHandler.
func (h *userHandler) HandleGetUserByEmail(c echo.Context) error {
	panic("unimplemented")
}

// HandleGetUserByID implements UserHandler.
func (h *userHandler) HandleGetUserByID(c echo.Context) error {
	panic("unimplemented")
}

// HandleGetUserByUsername implements UserHandler.
func (h *userHandler) HandleGetUserByUsername(c echo.Context) error {
	panic("unimplemented")
}

// HandleUpdateUser implements UserHandler.
func (h *userHandler) HandleUpdateUser(c echo.Context) error {
	panic("unimplemented")
}

type UserHandlerInject struct {
	CreateUserUsecase        usecase.CreateUserUsecase
	UpdateUserUsecase        usecase.UpdateUserUsecase
	DeleteUserUsecase        usecase.DeleteUserUsecase
	GetUserByIDUsecase       usecase.GetUserByIDUsecase
	GetUserByUsernameUsecase usecase.GetUserByUsernameUsecase
	GetUserByEmailUsecase    usecase.GetUserByEmailUsecase
}

func NewUserHandler(inject UserHandlerInject) UserHandler {
	return &userHandler{
		createUserUsecase:        inject.CreateUserUsecase,
		updateUserUsecase:        inject.UpdateUserUsecase,
		deleteUserUsecase:        inject.DeleteUserUsecase,
		getUserByIDUsecase:       inject.GetUserByIDUsecase,
		getUserByUsernameUsecase: inject.GetUserByUsernameUsecase,
		getUserByEmailUsecase:    inject.GetUserByEmailUsecase,
	}
}

package server

import (
	"github.com/congmanh18/XNM-Express/User_Profile/handler"
	"github.com/congmanh18/XNM-Express/User_Profile/model"
	"github.com/congmanh18/XNM-Express/User_Profile/provider"
	"github.com/congmanh18/XNM-Express/User_Profile/repository"
	"github.com/congmanh18/XNM-Express/User_Profile/usecase"
	"github.com/congmanh18/lucas-core/transport/http/route"
)

func InitializeRoutes(userHandler handler.UserHandler) []route.GroupRoute {
	return Routes(userHandler)
}

func InitializeAppProvider(serviceConf model.ServiceConfig) *provider.AppProvider {
	return provider.NewAppProvider(serviceConf)
}

func InitializeUserHandler(appProvider *provider.AppProvider) handler.UserHandler {
	userRepo := repository.NewUserRepository(appProvider.Postgres)
	return handler.NewUserHandler(handler.UserHandlerInject{
		CreateUserUsecase: usecase.NewCreateUserUsecase(userRepo),
	})
}

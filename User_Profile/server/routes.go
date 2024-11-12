package server

import (
	"github.com/congmanh18/XNM-Express/User_Profile/handler"
	"github.com/congmanh18/lucas-core/transport/http/method"
	"github.com/congmanh18/lucas-core/transport/http/route"
)

func Routes(
	user handler.UserHandler,
) []route.GroupRoute {
	return []route.GroupRoute{
		{
			Prefix: "/users",
			Routes: []route.Route{
				{
					Path:    "/register",
					Method:  method.POST,
					Handler: user.HandleCreateUser,
				},
			},
		},
	}
}

package handler

import (
	"github.com/congmanh18/XNM-Express/User_Profile/mapper"
	"github.com/congmanh18/XNM-Express/User_Profile/model/req"
	"github.com/congmanh18/lucas-core/transport/http/response"
	"github.com/labstack/echo/v4"
)

// HandleCreateUser implements UserHandler.
func (h *userHandler) HandleCreateUser(c echo.Context) error {
	createReq := req.CreateUserReq{}
	if err := c.Bind(&createReq); err != nil {
		return response.Error(
			c,
			ErrInvalidRequest.Code,
			ErrInvalidRequest.Message+": "+err.Error(),
		)
	}
	if err := validate.Struct(createReq); err != nil {
		return response.Error(
			c,
			ErrValidation.Code,
			ErrValidation.Message+": "+err.Error(),
		)
	}

	err := h.createUserUsecase.Execute(
		c.Request().Context(),
		mapper.TransformCreateReqToEntity(createReq),
	)
	if err != nil {
		return response.Error(
			c,
			err.Code(),
			err.Error()+": "+err.Debug().Error(),
		)
	}
	return response.OK(
		c,
		OKCreateUser.Code,
		OKCreateUser.Message,
		nil,
	)
}

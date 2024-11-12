package mapper

import (
	"github.com/congmanh18/XNM-Express/User_Profile/entity"
	"github.com/congmanh18/XNM-Express/User_Profile/model/req"
	"github.com/congmanh18/lucas-core/pointer"
	"github.com/congmanh18/lucas-core/record"
	"github.com/google/uuid"
)

func TransformCreateReqToEntity(req req.CreateUserReq) *entity.User {
	return &entity.User{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.Password,
	}
}

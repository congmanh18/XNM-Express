package req

type CreateUserReq struct {
	Username *string `json:"username,omitempty" validate:"required"`
	Password *string `json:"password,omitempty" validate:"required"`
	Email    *string `json:"email,omitempty" validate:"required,email"`
	Phone    *string `json:"phone,omitempty" validate:"required"`
	Role     *string `json:"role,omitempty" validate:"required"`
}

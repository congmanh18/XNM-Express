package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type Error struct {
	Code    int
	Message string
}

var (
	ErrInvalidRequest = &Error{Code: http.StatusBadRequest, Message: "Invalid request"}
	ErrValidation     = &Error{Code: http.StatusUnprocessableEntity, Message: "Validation error"}
	ErrUnauthorized   = &Error{Code: http.StatusUnauthorized, Message: "Unauthorized"}
	ErrForbidden      = &Error{Code: http.StatusForbidden, Message: "Forbidden"}
	ErrNotFound       = &Error{Code: http.StatusNotFound, Message: "Not found"}
	ErrInternal       = &Error{Code: http.StatusInternalServerError, Message: "Internal server error"}
)

// =================================================
type OK struct {
	Code    int
	Message string
}

var (
	OKCreateUser        = &OK{Code: http.StatusOK, Message: "Create user success"}
	OKUpdateUser        = &OK{Code: http.StatusOK, Message: "Update user success"}
	OKDeleteUser        = &OK{Code: http.StatusOK, Message: "Delete user success"}
	OKGetUserByID       = &OK{Code: http.StatusOK, Message: "Get user by id success"}
	OKGetUserByUsername = &OK{Code: http.StatusOK, Message: "Get user by username success"}
	OKGetUserByEmail    = &OK{Code: http.StatusOK, Message: "Get user by email success"}
)

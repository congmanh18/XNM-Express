package usecase

import (
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

var (
	ErrCreateUser        = &Error{Code: http.StatusInternalServerError, Message: "Error creating user"}
	ErrHashPassword      = &Error{Code: http.StatusInternalServerError, Message: "Error hashing password"}
	ErrUpdateUser        = &Error{Code: http.StatusInternalServerError, Message: "Error updating user"}
	ErrDeleteUser        = &Error{Code: http.StatusInternalServerError, Message: "Error deleting user"}
	ErrGetUserByID       = &Error{Code: http.StatusInternalServerError, Message: "Error getting user by id"}
	ErrGetUserByEmail    = &Error{Code: http.StatusInternalServerError, Message: "Error getting user by email"}
	ErrGetUserByPhone    = &Error{Code: http.StatusInternalServerError, Message: "Error getting user by phone"}
	ErrGetUserByUsername = &Error{Code: http.StatusInternalServerError, Message: "Error getting user by username"}
	ErrNotFound          = &Error{Code: http.StatusNotFound, Message: "User not found"}
	ErrConflict          = &Error{Code: http.StatusConflict, Message: "User already exists"}
	ErrBadRequest        = &Error{Code: http.StatusBadRequest, Message: "Bad request"}
	ErrUnauthorized      = &Error{Code: http.StatusUnauthorized, Message: "Unauthorized"}
	ErrForbidden         = &Error{Code: http.StatusForbidden, Message: "Forbidden"}
	ErrInternal          = &Error{Code: http.StatusInternalServerError, Message: "Internal server error"}
)

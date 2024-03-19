package apperrors

import (
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

type NotExistsError struct {
	Message string
}

func (error *NotExistsError) Error() string {
	return error.Message
}

type AlreadyExistsError struct {
	Message string
}

func (error *AlreadyExistsError) Error() string {
	return error.Message
}

type InvalidDataErr struct {
	Message string
}

func (error *InvalidDataErr) Error() string {
	return error.Message
}

type TokenError struct {
	Message string
}

func (error *TokenError) Error() string {
	return error.Message
}

type HashError struct {
	Message string
}

func (error *HashError) Error() string {
	return error.Message
}

type DBoperationErr struct {
	Message string
}

func (error *DBoperationErr) Error() string {
	return error.Message
}

type AuthError struct {
	Message string
}

func (error *AuthError) Error() string {
	return error.Message
}

type AppError interface {
	Error() string
}

func MatchError(appErr AppError) *ResponseError {
	switch ae := appErr.(type) {
	case *DBoperationErr, *TokenError, *HashError:
		return &ResponseError{
			Message: ae.Error(),
			Status:  http.StatusInternalServerError,
		}
	case *InvalidDataErr, *NotExistsError:
		return &ResponseError{
			Message: ae.Error(),
			Status:  http.StatusBadRequest,
		}
	case *AuthError:
		return &ResponseError{
			Message: ae.Message,
			Status:  http.StatusUnauthorized,
		}
	case *AlreadyExistsError:
		return &ResponseError{
			Message: ae.Message,
			Status:  http.StatusConflict,
		}
	}
	return nil
}

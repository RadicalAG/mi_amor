package api_error

import (
	"errors"
	httpNet "net/http"
	"radical/red_letter/internal/internal_error"
)

type APIError struct {
	Status  int
	Message string
}

func NewApiError(status int, msg string) APIError {
	return APIError{
		Status:  status,
		Message: msg,
	}
}

func FromError(err error) APIError {
	var svcError internal_error.InternalError
	var apiError APIError
	if errors.As(err, &svcError) {
		apiError.Message = svcError.AppError().Error()
		svcErr := svcError.SvcError()
		switch svcErr {
		case internal_error.ErrBadRequest:
			apiError.Status = httpNet.StatusBadRequest
		case internal_error.ErrNotFound:
			apiError.Status = httpNet.StatusNotFound
		case internal_error.ErrInternalServer:
			apiError.Status = httpNet.StatusInternalServerError
		}
	} else {
		return NewApiError(httpNet.StatusInternalServerError, "internal server error")
	}
	return apiError
}

func (n APIError) Error() string {
	return n.Message
}

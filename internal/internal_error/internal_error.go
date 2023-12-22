package internal_error

import (
	"errors"
	"fmt"
)

var (
	ErrBadRequest     = errors.New("bad request")
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("not found")
)

type InternalError struct {
	appError error
	svcError error
}

func NewError(svcError, appError error) error {
	return InternalError{
		svcError: svcError,
		appError: appError,
	}
}
func (e InternalError) AppError() error {
	return e.appError
}
func (e InternalError) SvcError() error {
	return e.svcError
}

func (e InternalError) Error() string {
	return errors.Join(e.svcError, e.appError).Error()
}

func CannotBeEmptyError(name string) error {
	return NewError(ErrBadRequest, fmt.Errorf("%v cannot be empty", name))
}

func InvalidError(name string) error {
	return NewError(ErrBadRequest, fmt.Errorf("invalid %v value", name))
}

func InternalServerError(msg string) error {
	if msg == "" {
		msg = "internal server error"
	}
	return NewError(ErrInternalServer, fmt.Errorf(msg))
}

func NotFoundError(name string) error {
	return NewError(ErrNotFound, fmt.Errorf("%v not found", name))
}

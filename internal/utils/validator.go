package utils

import (
	"net/mail"
	"strings"
)

type validator struct {
}

func NewValidator() *validator {
	return &validator{}
}

type Validator interface {
	IsBlank(s string) bool
	IsValidEmail(email string) bool
}

func (v *validator) IsBlank(s string) bool {
	return strings.Trim(s, " ") == ""
}

func (v *validator) IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

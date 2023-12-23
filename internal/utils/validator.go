package utils

import (
	"net/mail"
	"radical/red_letter/internal/internal_error"
	"strings"
	"unicode"
)

type validator struct {
}

func NewValidator() *validator {
	return &validator{}
}

type Validator interface {
	IsBlank(s string) bool
	IsValidEmail(email string) bool
	IsValidPassword(password string) (bool, error)
}

func (v *validator) IsBlank(s string) bool {
	return strings.Trim(s, " ") == ""
}

func (v *validator) IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// func (v *validator) IsValidPassword(password string) (bool, error) {

// 	if len(password) < 8 {
// 		return false, internal_error.BadRequestError("Password must be at least 8 characters long")
// 	}

// 	for _, char := range password {
// 		switch {
// 		case !unicode.	IsUpper(char):
// 			return false, internal_error.BadRequestError("Password must contain at least one uppercase letter")
// 		case !unicode.IsLower(char):
// 			return false, internal_error.BadRequestError("Password must contain at least one lowercase letter")
// 		case !unicode.IsNumber(char):
// 			return false, internal_error.BadRequestError("Password must contain at least one numeric digit")
// 		case !unicode.IsPunct(char) && !unicode.IsSymbol(char):
// 			return false, internal_error.BadRequestError("Password must contain at least one symbol or punctuation mark")
// 		default:
// 			return false, internal_error.BadRequestError("Invalid character found in password")
// 		}
// 	}

// 	return true, nil
// }

func (v *validator) IsValidPassword(password string) (bool, error) {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false, internal_error.BadRequestError("Invalid character found in password")
		}
	}

	if !upp {
		return false, internal_error.BadRequestError("Password must contain at least one uppercase letter")
	}
	if !low {
		return false, internal_error.BadRequestError("Password must contain at least one lowercase letter")
	}
	if !num {
		return false, internal_error.BadRequestError("Password must contain at least one numeric digit")
	}
	if !sym {
		return false, internal_error.BadRequestError("Password must contain at least one symbol or punctuation mark")
	}
	if tot < 8 {
		return false, internal_error.BadRequestError("Password must be at least 8 characters long")
	}

	return true, nil
}

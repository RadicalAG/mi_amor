package service

import (
	"context"
	"log"
	"radical/red_letter/internal/internal_error"
	"radical/red_letter/internal/model"
	"radical/red_letter/internal/repository"
	"radical/red_letter/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	repo      repository.UserRepository
	validator utils.Validator
}

func NewAuthService(repo repository.UserRepository, validator utils.Validator) *authService {
	return &authService{
		repo:      repo,
		validator: validator,
	}
}

type AuthService interface {
	RegisterUser(ctx context.Context, name, email, password string) error
}

func (a *authService) RegisterUser(ctx context.Context, name, email, password string) error {
	// validate requests
	err := a.validateRequest(ctx, name, email, password)
	if err != nil {
		return err
	}

	// password hashing
	hashedPassword, err := hashPassword(password)
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		return internal_error.InternalServerError("error creating user")
	}

	// create user
	err = a.repo.CreateUser(ctx, &model.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	})
	if err != nil {
		return err
	}
	return nil

}

func (a *authService) validateRequest(ctx context.Context, name, email, password string) error {
	// empty value validation
	if a.validator.IsBlank(name) {
		return internal_error.CannotBeEmptyError("name")
	}
	if a.validator.IsBlank(email) {
		return internal_error.CannotBeEmptyError("email")
	}
	if a.validator.IsBlank(password) {
		return internal_error.CannotBeEmptyError("password")
	}

	// email validation
	if !a.validator.IsValidEmail(email) {
		return internal_error.InvalidError("email")
	}

	existingUser, _ := a.repo.GetUserByEmail(ctx, email)
	if existingUser != nil {
		log.Printf("email %v is already registered\n", email)
		return internal_error.BadRequestError("email is already registered")
	}

	// password validation
	_, err := a.validator.IsValidPassword(password)
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		return err
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// func (a *authService) CheckPassword(providedPassword string) error {
// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

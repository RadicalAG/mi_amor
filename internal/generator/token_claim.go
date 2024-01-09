package generator

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenGenerator struct {
	jwtKey []byte
}

type TokenClaim struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewTokenGenerator(secretKey string) *tokenGenerator {
	return &tokenGenerator{
		jwtKey: []byte(secretKey),
	}
}

type TokenGenerator interface {
	GenerateJWT(email string, id string) (tokenString string, err error)
	ValidateAndDecodeToken(signedToken string) (claims *TokenClaim, err error)
}

func (tc *tokenGenerator) GenerateJWT(email string, id string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &TokenClaim{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(tc.jwtKey)
	return
}

func (tc *tokenGenerator) ValidateAndDecodeToken(signedToken string) (claims *TokenClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&TokenClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tc.jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

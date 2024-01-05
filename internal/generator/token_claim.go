package generator

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type tokenClaim struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	jwtKey []byte
	jwt.StandardClaims
}

func NewTokenClaim(secretKey string) *tokenClaim {
	return &tokenClaim{
		jwtKey: []byte(secretKey),
	}
}

type TokenClaim interface {
	GenerateJWT(email string, id string) (tokenString string, err error)
	ValidateAndDecodeToken(signedToken string) (claims *tokenClaim, err error)
}

func (tc *tokenClaim) GenerateJWT(email string, id string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &tokenClaim{
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

func (tc *tokenClaim) ValidateAndDecodeToken(signedToken string) (claims *tokenClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&tokenClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(tc.jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*tokenClaim)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

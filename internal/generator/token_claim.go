package generator

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

type tokenClaim struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewTokenClaim() *tokenClaim {
	return &tokenClaim{}
}

type TokenClaim interface {
	GenerateJWT(email string, id string) (tokenString string, err error)
	ValidateToken(signedToken string) (err error)
}

func (tg *tokenClaim) GenerateJWT(email string, id string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &tokenClaim{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func (tg *tokenClaim) ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&tokenClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*tokenClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}

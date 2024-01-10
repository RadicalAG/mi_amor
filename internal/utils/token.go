package utils

import (
	"radical/red_letter/internal/api_error"
	"radical/red_letter/internal/generator"
	"radical/red_letter/internal/internal_error"

	"github.com/gin-gonic/gin"
)

func GetClaimsFromToken(c *gin.Context) (*generator.TokenClaim, error) {
	res, exists := c.Get("tokenClaims")
	if !exists {
		return nil, api_error.FromError(internal_error.BadRequestError("unauthorized"))
	}
	claim, ok := res.(*generator.TokenClaim)
	if !ok {
		return nil, api_error.FromError(internal_error.InternalServerError(""))
	}

	return claim, nil
}

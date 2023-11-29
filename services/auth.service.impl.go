package services

import (
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/ntnghiatn/rest-go-gin-api/models"
)

type AuthServiceImpl struct {
	ctx context.Context
}

// CreateToken implements AuthServive.
func (a *AuthServiceImpl) CreateToken(*models.User) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{})

}

func NewAuthService(ctx context.Context) AuthServive {
	return &AuthServiceImpl{
		ctx: ctx,
	}
}

package services

import (
	"github.com/golang-jwt/jwt"
	"github.com/ntnghiatn/rest-go-gin-api/models"
)

type AuthServive interface {
	CreateToken(*models.User) *jwt.Token
}

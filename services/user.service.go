package services

import "github.com/ntnghiatn/rest-go-gin-api/models"

type UserServive interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}

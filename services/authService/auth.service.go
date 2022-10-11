package authService

import "api.com/api/models"

type AuthService interface {
	RegisterUser(*models.User) (string,error)
	LoginUser(*models.UserLogin) (string,error)
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
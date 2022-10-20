package userService

import "api.com/api/models"

type UserService interface {
	GetListedProductsByUserId(*string) ([]*models.Product, error)
	GetUnlistedProductsByUserId(*string) ([]*models.Product, error)
	GetProductsByUserId(*string) ([]*models.Product, error)
	GetUserById(*string) (*models.User, error)

}
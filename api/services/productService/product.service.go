package productService

import "api.com/api/models"

type ProductService interface {
	CreateProduct(*models.Product,*string) (string,error)
	GetAllProducts() ([]*models.Product, error)
	GetProductById(*string) (*models.Product, error)
}
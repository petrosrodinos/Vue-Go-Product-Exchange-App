package routes

import (
	"api.com/api/middlewares"

	"api.com/api/controllers/mediaController"
	"api.com/api/controllers/productController"
	"github.com/gin-gonic/gin"
)

type ProductRoutes struct {
	productcontroller productController.ProductController
	// mediacontroller mediaController.mediaController
}

func NewProductRoutes(productcontroller productController.ProductController) ProductRoutes{
	return ProductRoutes{
		productcontroller: productcontroller,
		// mediacontroller:mediacontroller
	}
}

func (uc *ProductRoutes) RegisterProductRoutes(rg *gin.RouterGroup) {
	productroutes := rg.Group("/products")
	productroutes.POST("",middlewares.AuthorizeJWT(), uc.productcontroller.CreateProduct)
	productroutes.GET("", uc.productcontroller.GetAllProducts)
	productroutes.GET("/:id", uc.productcontroller.GetProductById)
	productroutes.POST("/file/formdata", mediaController.FileUpload())
    productroutes.POST("/file/remote", mediaController.RemoteUpload())
}
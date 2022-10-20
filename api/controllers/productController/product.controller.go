package productController

import (
	"net/http"

	"api.com/api/models"
	"api.com/api/services/productService"
	"github.com/gin-gonic/gin"
)


type ProductController struct {
	ProductService productService.ProductService
}

func New(productservice productService.ProductService) ProductController{
	return ProductController{
		ProductService: productservice,
	}
}

func (uc *ProductController) CreateProduct(ctx *gin.Context){
	var user models.Product
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message": err.Error()})
		return
	}
	userId,exist := ctx.Get("userId")
	if !exist{
		ctx.JSON(http.StatusForbidden,gin.H{"message": "anauthorized"})
		return
	}
	finalUserId := userId.(string)
	_,err := uc.ProductService.CreateProduct(&user,&finalUserId)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "Product created successfully"})
}

func (uc *ProductController) GetAllProducts(ctx *gin.Context){
	products,err := uc.ProductService.GetAllProducts()
	if err != nil{
		ctx.JSON(http.StatusNotFound,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "Products found successfully", "products": products})

}

func (uc *ProductController) GetProductById(ctx *gin.Context){
	productId := ctx.Param("id")
	product,err := uc.ProductService.GetProductById(&productId)
	if err != nil{
		ctx.JSON(http.StatusNotFound,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "Product found successfully", "product": product})

}



package userController

import (
	"net/http"

	"api.com/api/services/userService"
	"github.com/gin-gonic/gin"
)


type UserController struct {
	UserService userService.UserService
}

func New(userservice userService.UserService) UserController{
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) GetProductsByUserId(ctx *gin.Context){
	userId := ctx.Param("id")
	products,err := uc.UserService.GetProductsByUserId(&userId)
	if err != nil{
		ctx.JSON(http.StatusNotFound,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "Products found successfully", "products": products})

}

func (uc *UserController) GetUserById(ctx *gin.Context){
	userId := ctx.Param("id")
	//633015f375e5bf67fd4c0f3e
	user,err := uc.UserService.GetUserById(&userId)
	if err != nil{
		ctx.JSON(http.StatusNotFound,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "User found successfully", "user": user})

}


func (uc *UserController) GetListedProductsByUserId(ctx *gin.Context){
	userId := ctx.Param("id")
	products,err := uc.UserService.GetListedProductsByUserId(&userId)
	if err != nil{
		ctx.JSON(http.StatusNotFound,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "Products found successfully", "products": products})

}

func (uc *UserController) GetUnlistedProductsByUserId(ctx *gin.Context){
	userId := ctx.Param("id")
	products,err := uc.UserService.GetUnlistedProductsByUserId(&userId)
	if err != nil{
		ctx.JSON(http.StatusNotFound,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "Products found successfully", "products": products})

}




package routes

import (
	"api.com/api/controllers/userController"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	usercontroller userController.UserController
}

func NewUserRoutes(usercontroller userController.UserController) UserRoutes{
	return UserRoutes{
		usercontroller: usercontroller,
	}
}

func (uc *UserRoutes) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroutes := rg.Group("/user")
	userroutes.GET("/:id", uc.usercontroller.GetUserById)
	userroutes.GET("/:id/products", uc.usercontroller.GetProductsByUserId)
	userroutes.GET("/:id/products/listed", uc.usercontroller.GetListedProductsByUserId)
	userroutes.GET("/:id/products/unlisted", uc.usercontroller.GetUnlistedProductsByUserId)

}
package routes

import (
	"api.com/api/controllers/authController"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	usercontroller authController.AuthController
}

func NewAuthRoutes(usercontroller authController.AuthController) AuthRoutes{
	return AuthRoutes{
		usercontroller: usercontroller,
	}
}

func (uc *AuthRoutes) RegisterAuthRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/auth")
	userroute.POST("/register", uc.usercontroller.RegisterUser)
	userroute.POST("/login", uc.usercontroller.LoginUser)
	userroute.GET("/get/:name", uc.usercontroller.GetUser)
	userroute.GET("/getall", uc.usercontroller.GetAll)
	userroute.PATCH("/update", uc.usercontroller.UpdateUser)
	userroute.DELETE("/delete/:name", uc.usercontroller.DeleteUser)
}
package authController

import (
	"net/http"

	"api.com/api/models"
	"api.com/api/services/authService"
	"api.com/api/utils"
	"github.com/gin-gonic/gin"
)


type AuthController struct {
	AuthService authService.AuthService
}

func New(authservice authService.AuthService) AuthController{
	return AuthController{
		AuthService: authservice,
	}
}

func (uc *AuthController) RegisterUser(ctx *gin.Context){
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message": err.Error()})
		return
	}
	userId,err := uc.AuthService.RegisterUser(&user)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message": err.Error()})
		return
	}
	jwtToken := utils.JWTAuthService().GenerateToken(userId)


	ctx.JSON(http.StatusOK,gin.H{"message": "User created successfully","token": jwtToken,"id":userId})

}

func (uc *AuthController) LoginUser(ctx *gin.Context){
	var user models.UserLogin
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message": err.Error()})
		return
	}
	userId,err := uc.AuthService.LoginUser(&user)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message": err.Error()})
		return
	}
	jwtToken := utils.JWTAuthService().GenerateToken(userId)


	ctx.JSON(http.StatusOK,gin.H{"message": "Login successfull","token": jwtToken,"id":userId})

}

func (uc *AuthController) GetUser(ctx *gin.Context){
	username := ctx.Param("name")
	user,err := uc.AuthService.GetUser(&username)
	if err != nil{
		ctx.JSON(http.StatusNotFound,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "User found successfully", "user": user})

}

func (uc *AuthController) GetAll(ctx *gin.Context){
	users,err := uc.AuthService.GetAll()
	if err != nil{
		ctx.JSON(http.StatusNotFound,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "Users found successfully", "users": users})
}

func (uc *AuthController) UpdateUser(ctx *gin.Context) {
	 var user models.User
	 if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message": err.Error()})
		return
	}
	err := uc.AuthService.UpdateUser(&user)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "User updated successfully"})
}

func (uc *AuthController) DeleteUser(ctx *gin.Context){
	username := ctx.Param("name")
	err := uc.AuthService.DeleteUser(&username)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message": "User deleted successfully"})
}

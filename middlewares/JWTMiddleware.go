package middlewares

import (
	"net/http"

	"api.com/api/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if(authHeader == ""){
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := utils.JWTAuthService().ValidateToken(tokenString)
		if(err != nil){
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("userId", claims["userId"]);
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return

		}
		
	}
}
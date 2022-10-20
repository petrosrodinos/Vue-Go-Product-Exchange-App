package utils

import (
	"errors"
	"time"

	config "api.com/api/configs"

	"github.com/dgrijalva/jwt-go"
)

//jwt service
type JWTService interface {
	GenerateToken(string) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	UserId string   `json:"userId"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: config.EnvJWTKey(),
		issure:    "product-exchange",
	}
}


func (service *jwtServices) GenerateToken(userId string) string {
	claims := &authCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unauthorized")
		}
	
		return []byte(service.secretKey), nil
	})

	if(err != nil){
		return nil, err
	}

	return token,nil

}

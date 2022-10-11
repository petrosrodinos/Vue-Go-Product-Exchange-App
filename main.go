package main

import (
	"context"
	"fmt"
	"log"
	"time"

	configs "api.com/api/configs"
	"api.com/api/controllers/authController"
	"api.com/api/controllers/productController"
	"api.com/api/controllers/tradeController"
	"api.com/api/controllers/userController"
	"api.com/api/routes"
	"api.com/api/services/authService"
	"api.com/api/services/productService"
	"api.com/api/services/tradeService"
	"api.com/api/services/userService"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine
	authservice authService.AuthService
	authcontroller authController.AuthController
	authroutes routes.AuthRoutes
	usercollection *mongo.Collection
	productcollection *mongo.Collection
	productservice productService.ProductService
	productcontroller productController.ProductController
	productroutes routes.ProductRoutes

	tradecollection *mongo.Collection
	tradeservice tradeService.TradeService
	tradecontroller tradeController.TradeController
	traderoutes routes.TradeRoutes

	userservice userService.UserService
	usercontroller userController.UserController
	userroutes routes.UserRoutes

	ctx context.Context
	mongoclient *mongo.Client
	err error
)

func init(){
	ctx = context.TODO()
	// mongoclient, err = mongo.NewClient(options.Client().ApplyURI())
	mongoconn := options.Client().ApplyURI(configs.EnvMongoUrl())
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil{
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	fmt.Println(configs.EnvMongoUrl())

	initAuth()
	initProducts()
	initUser()
	initTrades();

	server = gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization","Access-Control-Allow-Headers"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	
	server.Use(cors.New(config))
}

func initAuth(){
	usercollection = mongoclient.Database("product-exchange").Collection("users")
	authservice = authService.NewAuthService(usercollection, ctx)
	authcontroller = authController.New(authservice)
	authroutes = routes.NewAuthRoutes(authcontroller);

}

func initProducts(){
	productcollection = mongoclient.Database("product-exchange").Collection("products")
	productservice = productService.NewProductService(productcollection, ctx)
	productcontroller = productController.New(productservice)
	productroutes = routes.NewProductRoutes(productcontroller);
}

func initUser(){
	userservice = userService.NewUserService(productcollection,usercollection, ctx)
	usercontroller = userController.New(userservice)
	userroutes = routes.NewUserRoutes(usercontroller);
}

func initTrades(){
	tradecollection = mongoclient.Database("product-exchange").Collection("trades")
	tradeservice = tradeService.NewTradeService(tradecollection,productcollection, ctx)
	tradecontroller = tradeController.New(tradeservice)
	traderoutes = routes.NewTradeRoutes(tradecontroller);
}

func main() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/api/v1/")
	authroutes.RegisterAuthRoutes(basepath)	
	productroutes.RegisterProductRoutes(basepath)
	userroutes.RegisterUserRoutes(basepath)
	traderoutes.RegisterTradeRoutes(basepath)
	log.Fatal(server.Run(":9090"))
}
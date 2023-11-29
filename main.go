package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ntnghiatn/rest-go-gin-api/controllers"
	"github.com/ntnghiatn/rest-go-gin-api/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userService    services.UserServive
	userController controllers.UserController
	authController controllers.AuthController
	ctx            context.Context
	userCollection *mongo.Collection
	mongoClient    *mongo.Client
	err            error
)

// Khởi tạo các đối tượng được khai báo
// 1. mongoClient <- (ctx và client options); ctx <- context.TODO(); clientOptions <- Khởi tạo từ hàm Client trong mongo/options
// 2. err:  được hứng lấy from các khởi tạo khác.
// 3.
func init() {
	ctx := context.TODO()

	clientOps := options.Client().ApplyURI("mongodb+srv://nghiango:Nghia385685@cluster0.gr4nd.mongodb.net/userdb")

	mongoClient, err = mongo.Connect(ctx, clientOps)
	if err != nil {
		log.Fatal(err)
	}

	// check connect mongo by call func Ping of client..
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	//
	userCollection = mongoClient.Database("userdb").Collection("users")
	userService = services.NewUserService(userCollection, ctx)
	userController = controllers.New(userService)
	authController = controllers.NewAuthController(services.NewAuthService(ctx))

	// Khởi tạo server
	server = gin.New()
}

func main() {
	defer func() {
		mongoClient.Disconnect(ctx)
		log.Fatal("mongo disconnection")
	}()
	apiV1 := server.Group("/api/v1")
	userController.RegisterUserRoutes(apiV1)
	authController.RegisterAuthRoutes(apiV1)

	log.Fatal(server.Run(":9090"))
}

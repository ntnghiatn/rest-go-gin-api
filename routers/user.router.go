package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ntnghiatn/rest-go-gin-api/controllers"
	"github.com/ntnghiatn/rest-go-gin-api/middlewares"
	"github.com/ntnghiatn/rest-go-gin-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

// var nikaHandler controllers.NikayaController

func RegisterUserRoutes(db *mongo.Database, ctx context.Context, rg *gin.RouterGroup) {
	userHandler := controllers.New(services.NewUserService(db.Collection("users"), ctx))
	// fmt.Println("---------------------------------------nikaHandler:", nikaHandler)
	// }
	route := rg.Group("/users").Use(middlewares.Auth())
	route.GET("/list", userHandler.GetAll)
	route.GET("/create", userHandler.CreateUser)
}

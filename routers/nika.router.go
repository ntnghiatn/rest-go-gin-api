package routers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ntnghiatn/rest-go-gin-api/controllers"
	"github.com/ntnghiatn/rest-go-gin-api/middlewares"
	"github.com/ntnghiatn/rest-go-gin-api/services"
)

// var nikaHandler controllers.NikayaController

func RegisterNikayaRoutes(ctx context.Context, rg *gin.RouterGroup) {
	nikaHandler := controllers.NewNikayaHandler(services.NewNikayaService(ctx))
	// fmt.Println("---------------------------------------nikaHandler:", nikaHandler)
	// }
	route := rg.Group("/nikaya").Use(middlewares.Auth())
	route.GET("/trungbokinh", nikaHandler.GenTrungBoKinh())
	route.GET("/truongbokinh", nikaHandler.GenTruongBoKinh())
}

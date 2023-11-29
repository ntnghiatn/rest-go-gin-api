package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ntnghiatn/rest-go-gin-api/middlewares"
	"github.com/ntnghiatn/rest-go-gin-api/models"
	"github.com/ntnghiatn/rest-go-gin-api/services"
)

type UserController struct {
	UserService services.UserServive
}

func New(userService services.UserServive) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user").Use(middlewares.Auth())
	// sau khi .Use(middlewares...) thì kiểu trả về sẽ không phải là
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/list", uc.GetAll)
	userRoute.GET("/get/:name", uc.GetUser)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/delete/:name", uc.DeleteUser)
}

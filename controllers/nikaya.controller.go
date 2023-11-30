package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ntnghiatn/rest-go-gin-api/middlewares"
	"github.com/ntnghiatn/rest-go-gin-api/services"
)

type NikayaController struct {
	NikayaServive services.NikayaServive
}

func NewNikayaHandler(nikayaService services.NikayaServive) NikayaController {
	return NikayaController{
		NikayaServive: nikayaService,
	}
}

func (nc *NikayaController) GenTrungBoKinh(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"lesson": nc.NikayaServive.TrungBoKinhRand()})
}

func (nc *NikayaController) RegisterNikayaRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/nikaya").Use(middlewares.Auth())
	route.GET("/trungbokinh", nc.GenTrungBoKinh)
}

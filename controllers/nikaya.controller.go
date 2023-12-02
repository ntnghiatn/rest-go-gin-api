package controllers

import (
	"github.com/gin-gonic/gin"
	// "github.com/ntnghiatn/rest-go-gin-api/middlewares"
	"github.com/ntnghiatn/rest-go-gin-api/services"
)

type NikayaController struct {
	NikayaServive services.NikayaServive
}

// init Controller (init được thực hiện ở router)
func NewNikayaHandler(nikayaService services.NikayaServive) NikayaController {
	return NikayaController{
		NikayaServive: nikayaService,
	}
}

// implement
func (nc *NikayaController) GenTrungBoKinh() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"lesson": nc.NikayaServive.TrungBoKinhRand()})
	}
}

// implement
func (nc *NikayaController) GenTruongBoKinh() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"lesson": nc.NikayaServive.TruongBoKinhRand()})
	}
}

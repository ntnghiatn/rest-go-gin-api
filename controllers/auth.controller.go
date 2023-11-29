package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ntnghiatn/rest-go-gin-api/services"
	"github.com/ntnghiatn/rest-go-gin-api/utils"
)

type AuthController struct {
	AuthService services.AuthServive
}

func NewAuthController(authService services.AuthServive) AuthController {
	return AuthController{
		AuthService: authService,
	}
}

func (ac *AuthController) CreateToken(ctx *gin.Context) {

	// mySigningKey := []byte("AllYourBase")
	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// signingkey := utils.LoadRSAPrivateKeyFromDisk(dir + "/private.pem")

	// // Create the Claims
	// claims := &jwt.StandardClaims{
	// 	ExpiresAt: time.Now().Add(time.Duration(30) * time.Second).Unix(),
	// 	Issuer:    "test",
	// 	Subject:   "ciquan",
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	// ss, err := token.SignedString(signingkey)
	token, err := utils.GenToken("hello")
	fmt.Println(token, err)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, gin.H{"message": err.Error()})
	}
	ctx.JSON(200, gin.H{"token": token})
}

func (ac *AuthController) RegisterAuthRoutes(rg *gin.RouterGroup) {
	route := rg.Group("/auth")
	route.POST("/token", ac.CreateToken)

}

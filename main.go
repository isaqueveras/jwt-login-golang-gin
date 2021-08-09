package main

import (
	"net/http"

	"github.com/isaqueveras/jwt-login/src/controller"
	"github.com/isaqueveras/jwt-login/src/services"

	"github.com/gin-gonic/gin"
)

func main() {
	var loginService services.LoginService = services.StaticLoginService()
	var jwtService services.JWTService = services.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.New()

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)

		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	server.Run(":8080")
}

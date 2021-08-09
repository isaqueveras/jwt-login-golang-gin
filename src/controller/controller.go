package controller

import (
	"github.com/isaqueveras/jwt-login/src/dto"
	"github.com/isaqueveras/jwt-login/src/services"

	"github.com/gin-gonic/gin"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService services.LoginService
	jWtService   services.JWTService
}

func LoginHandler(loginService services.LoginService,
	jWtService services.JWTService) LoginController {

	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials

	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}

	isUserAuthenticated := controller.loginService.LoginUser(credential.ID, credential.Email, credential.Nome)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.ID, credential.Email, true)
	}

	return ""
}

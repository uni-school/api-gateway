package controller_auth

import (
	"github.com/labstack/echo/v4"
	service_auth "github.com/uni-school/api-gateaway/pkg/core/auth/service"
)

type AuthController struct {
	AuthService service_auth.IAuthService
}

func InitAuthController(authService service_auth.IAuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (authController *AuthController) Mount(group *echo.Group) {
	group.POST("/register", authController.RegisterUser)
}

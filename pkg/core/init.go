package core

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	"github.com/uni-school/api-gateaway/shared/config"
	"github.com/uni-school/api-gateaway/wire"
)

func InitRoutes(e *echo.Echo) {
	{
		auth := e.Group("/auth")
		authController := wire.AuthController()
		authController.Mount(auth)
	}
	{
		user := e.Group("/users")
		user.Use(echojwt.WithConfig(config.ConfigureJWT()))
		userController := wire.UserController()
		userController.Mount(user)
	}
}

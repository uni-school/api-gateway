package core

import (
	"github.com/labstack/echo/v4"

	"github.com/uni-school/api-gateaway/wire"
)

func InitRoutes(e *echo.Echo) {
	{
		auth := e.Group("/auth")
		authController := wire.AuthController()
		authController.Mount(auth)
	}
}

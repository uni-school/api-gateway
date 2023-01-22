package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	my_validator "github.com/uni-school/api-gateaway/libs/validator"
	"github.com/uni-school/api-gateaway/pkg/core"
	"github.com/uni-school/api-gateaway/shared/config"
	"github.com/uni-school/api-gateaway/shared/custom"
	"github.com/uni-school/api-gateaway/shared/util"
)

func init() {
	config.ConfigApps()
}

func main() {
	e := echo.New()
	e.Validator = &my_validator.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = custom.CustomHTTPErrorHandler

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.RequestID(),
	)

	core.InitRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", util.GetHostApp(), util.GetPortApp())))
}

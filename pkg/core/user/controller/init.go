package controller_user

import (
	"github.com/labstack/echo/v4"
	service_user "github.com/uni-school/api-gateaway/pkg/core/user/service"
	"github.com/uni-school/api-gateaway/shared/constant"
	"github.com/uni-school/api-gateaway/shared/middleware"
)

type UserController struct {
	UserService service_user.IUserService
}

func InitUserController(userService service_user.IUserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (userController *UserController) Mount(group *echo.Group) {
	group.GET("/students", userController.GetListStudent, middleware.AllowedRoles(constant.ADMIN))
}

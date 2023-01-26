//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	service_auth "github.com/uni-school/api-gateaway/pkg/core/auth/service"
	controller_auth "github.com/uni-school/api-gateaway/pkg/core/auth/controller"
	
	controller_user "github.com/uni-school/api-gateaway/pkg/core/user/controller"
	service_user "github.com/uni-school/api-gateaway/pkg/core/user/service"
	resource_user "github.com/uni-school/api-gateaway/pkg/core/user/resource"
)

func AuthController() *controller_auth.AuthController {
	wire.Build(
		controller_auth.InitAuthController,
		service_auth.InitAuthService,
		resource_user.InitUserResource,
	)
	return &controller_auth.AuthController{}
}


func UserController() *controller_user.UserController {
	wire.Build(
		controller_user.InitUserController,
		service_user.InitUserService,
		resource_user.InitUserResource,
	)
	return &controller_user.UserController{}
}
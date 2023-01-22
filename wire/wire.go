//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"

	resource_user "github.com/uni-school/api-gateaway/pkg/core/user/resource"
	service_auth "github.com/uni-school/api-gateaway/pkg/core/auth/service"
	controller_auth "github.com/uni-school/api-gateaway/pkg/core/auth/controller"
)

func AuthController() *controller_auth.AuthController {
	wire.Build(
		controller_auth.InitAuthController,
		service_auth.InitAuthService,
		resource_user.InitUserResource,
	)
	return &controller_auth.AuthController{}
}

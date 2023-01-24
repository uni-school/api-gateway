package service_auth

import (
	dto_service_auth "github.com/uni-school/api-gateaway/pkg/core/auth/service/dto"
	resource_user "github.com/uni-school/api-gateaway/pkg/core/user/resource"
)

type IAuthService interface {
	RegisterUser(payload *dto_service_auth.RegisterUserDTO) error
	LoginUser(payload *dto_service_auth.LoginUserDTO) (*dto_service_auth.UserLoggedInDTO, error)
}

type AuthService struct {
	UserResource resource_user.IUserResource
}

func InitAuthService(userResource resource_user.IUserResource) IAuthService {
	return &AuthService{
		UserResource: userResource,
	}
}

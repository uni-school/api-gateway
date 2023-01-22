package service_auth

import (
	"context"

	dto_service_auth "github.com/uni-school/api-gateaway/pkg/core/auth/service/dto"
	user_microservice "github.com/uni-school/api-gateaway/proto/user"
)

func (s *AuthService) RegisterUser(payload *dto_service_auth.RegisterUserDTO) error {
	var (
		ctx = context.Background()
	)

	_, err := s.UserResource.CreateUser(ctx, &user_microservice.CreateUserRequest{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Role:     string(payload.Role),
	})
	if err != nil {
		return err
	}

	return nil
}

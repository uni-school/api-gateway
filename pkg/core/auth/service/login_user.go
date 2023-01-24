package service_auth

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	dto_service_auth "github.com/uni-school/api-gateaway/pkg/core/auth/service/dto"
	user_microservice "github.com/uni-school/api-gateaway/proto/user"
	"github.com/uni-school/api-gateaway/shared/config"
	"github.com/uni-school/api-gateaway/shared/constant"
	"github.com/uni-school/api-gateaway/shared/types"
	"github.com/uni-school/api-gateaway/shared/util"
)

func (s *AuthService) LoginUser(payload *dto_service_auth.LoginUserDTO) (*dto_service_auth.UserLoggedInDTO, error) {
	var (
		ctx = context.Background()
	)

	userProtoGrpcGetResult, err := s.UserResource.GetUserByEmail(ctx, &user_microservice.GetUserByEmailRequest{
		Email: payload.Email,
	})
	if err != nil {
		return nil, err
	}

	match := util.CheckPasswordHash(payload.Password, userProtoGrpcGetResult.GetData().GetPassword())
	if !match {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "username and password not match")
	}

	claims := &types.JwtCustomClaims{
		ID:    userProtoGrpcGetResult.GetData().GetId(),
		Name:  userProtoGrpcGetResult.GetData().GetName(),
		Email: userProtoGrpcGetResult.GetData().GetEmail(),
		Role:  constant.UserRole(userProtoGrpcGetResult.GetData().GetRole()),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.AppConfig.JWT.SecretKey))
	if err != nil {
		return nil, err
	}

	return &dto_service_auth.UserLoggedInDTO{
		User: dto_service_auth.UserDTOForLogin{
			ID: userProtoGrpcGetResult.GetData().GetId(),
		},
		Token: t,
	}, nil
}

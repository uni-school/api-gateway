package controller_auth

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/uni-school/api-gateaway/libs/response"
	request_controller_auth "github.com/uni-school/api-gateaway/pkg/core/auth/controller/request"
	dto_service_auth "github.com/uni-school/api-gateaway/pkg/core/auth/service/dto"
	"github.com/uni-school/api-gateaway/shared/validator"
)

func (c *AuthController) RegisterUser(ctx echo.Context) error {
	reqBody := new(request_controller_auth.RegisterUserRequestBody)
	if err := validator.ValidateRequest(ctx, reqBody); err != nil {
		return err
	}

	if err := c.AuthService.RegisterUser(&dto_service_auth.RegisterUserDTO{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: reqBody.Password,
		Role:     reqBody.Role,
	}); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, response.ResponseSuccess[any]("success register user", nil))
}

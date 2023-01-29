package controller_auth

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/uni-school/api-gateaway/libs/response"
	request_controller_auth "github.com/uni-school/api-gateaway/pkg/core/auth/controller/request"
	response_controller_auth "github.com/uni-school/api-gateaway/pkg/core/auth/controller/response"
	dto_service_auth "github.com/uni-school/api-gateaway/pkg/core/auth/service/dto"
	"github.com/uni-school/api-gateaway/shared/validator"
)

func (c *AuthController) LoginUser(ctx echo.Context) error {
	reqBody := new(request_controller_auth.LoginUserRequestBody)
	if err := validator.ValidateRequest(ctx, reqBody); err != nil {
		return err
	}

	userDTOLoginResult, err := c.AuthService.LoginUser(&dto_service_auth.LoginUserDTO{
		Email: reqBody.Email,
		Password: reqBody.Password,
	})
	if err != nil {
		return err
	}

	userReponseLoginResult := &response_controller_auth.UserLoggedInResponse{
		User: response_controller_auth.UserResponseForLogin{
			ID: userDTOLoginResult.User.ID,
		},
		Token: userDTOLoginResult.Token,
	}

	return ctx.JSON(http.StatusOK, response.ResponseSuccess("success login user", userReponseLoginResult))
}

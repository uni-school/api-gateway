package request_controller_auth

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uni-school/api-gateaway/libs/util"
	"github.com/uni-school/api-gateaway/shared/constant"
	"github.com/uni-school/api-gateaway/shared/constraint"
	"github.com/uni-school/api-gateaway/shared/validator"
)

type RegisterUserRequestBody struct {
	Name     string            `json:"name" validate:"required"`
	Email    string            `json:"email" validate:"required,email"`
	Password string            `json:"password" validate:"required"`
	Role     constant.UserRole `json:"role" validate:"required"`
}

func (r *RegisterUserRequestBody) CustomValidate() error {
	if !validator.IsPassword(r.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid password: password must have minimum 8 characters which contain at least one uppercase letter, one lowercase letter, one number and one special character")
	}

	if !validator.IsAllowedUserRole(r.Role) {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprint("unallowed role: only these role are allowed", util.ConvertMapKeysToString(constraint.MapAllowedUserRole)))
	}

	return nil
}

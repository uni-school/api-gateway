package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/uni-school/api-gateaway/libs/util"
	"github.com/uni-school/api-gateaway/shared/constant"
	"github.com/uni-school/api-gateaway/shared/types"
)

func AllowedRoles(roles ...constant.UserRole) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := c.Get("user").(*jwt.Token).Claims
			user, ok := claims.(*types.JwtCustomClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnprocessableEntity, fmt.Sprintf("invalid claims type. correct type: %T", claims))
			}

			if len(roles) > 0 {
				allowed := util.SliceContain(roles, user.Role)
				if !allowed {
					return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprint("user role not allowed. allowed roles:", roles))
				}
			}

			return next(c)
		}
	}
}

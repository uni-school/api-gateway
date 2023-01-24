package types

import (
	"github.com/golang-jwt/jwt"
	"github.com/uni-school/api-gateaway/shared/constant"
)

type JwtCustomClaims struct {
	ID    string            `json:"id"`
	Name  string            `json:"name"`
	Email string            `json:"email"`
	Role  constant.UserRole `json:"role"`
	jwt.StandardClaims
}

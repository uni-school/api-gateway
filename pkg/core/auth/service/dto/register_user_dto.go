package dto_service_auth

import (
	"github.com/uni-school/api-gateaway/shared/constant"
)

type RegisterUserDTO struct {
	Name     string
	Email    string
	Password string
	Role     constant.UserRole
}

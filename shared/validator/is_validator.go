package validator

import (
	"regexp"

	"github.com/uni-school/api-gateaway/shared/constant"
	"github.com/uni-school/api-gateaway/shared/constraint"
)

func IsPassword(password string) bool {
	passwordRegex := regexp.MustCompile(`^[(a-z)(A-Z)(0-9)(!@#$%^&*.?\-)]{8,}$`)
	return passwordRegex.MatchString(password)
}

func IsAllowedUserRole(role constant.UserRole) bool {
	return constraint.MapAllowedUserRole[role]
}

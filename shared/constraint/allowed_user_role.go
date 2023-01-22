package constraint

import (
	"github.com/uni-school/api-gateaway/shared/constant"
)

var (
	MapAllowedUserRole = map[constant.UserRole]bool{
		constant.STUDENT: true,
	}
)

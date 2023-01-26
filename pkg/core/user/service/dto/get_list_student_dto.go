package dto_service_user

import (
	"github.com/uni-school/api-gateaway/shared/constant"
)

type GetListStudentDTO struct {
	ID        string
	Name      string
	Email     string
	Role      constant.UserRole
	CreatedAt string
	UpdatedAt string
}

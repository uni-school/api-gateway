package response_controller_user

import (
	"github.com/uni-school/api-gateaway/shared/constant"
)

type GetListStudentResponse struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Email     string            `json:"email"`
	Role      constant.UserRole `json:"role"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}

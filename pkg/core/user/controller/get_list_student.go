package controller_user

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/uni-school/api-gateaway/libs/response"
	response_controller_user "github.com/uni-school/api-gateaway/pkg/core/user/controller/response"
)

func (c *UserController) GetListStudent(ctx echo.Context) error {
	userDTOGetListStudentResult, err := c.UserService.GetListStudent()
	if err != nil {
		return err
	}

	userResGetListStudentResponse := make([]*response_controller_user.GetListStudentResponse, 0)
	for _, value := range userDTOGetListStudentResult {
		userResGetListStudentResponse = append(userResGetListStudentResponse, &response_controller_user.GetListStudentResponse{
			ID:        value.ID,
			Name:      value.Name,
			Email:     value.Email,
			Role:      value.Role,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}

	return ctx.JSON(http.StatusOK, response.ResponseSuccess("success get list student", userResGetListStudentResponse))
}

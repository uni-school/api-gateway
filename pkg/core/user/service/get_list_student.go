package service_user

import (
	"context"

	dto_service_user "github.com/uni-school/api-gateaway/pkg/core/user/service/dto"
	user_microservice "github.com/uni-school/api-gateaway/proto/user"
	"github.com/uni-school/api-gateaway/shared/constant"
)

func (s *UserService) GetListStudent() ([]*dto_service_user.GetListStudentDTO, error) {
	var (
		ctx = context.Background()
	)

	userProtoGrpcGetListResult, err := s.UserResource.GetListUserByRole(ctx, &user_microservice.GetListUserByRoleRequest{
		Role: string(constant.STUDENT),
	})
	if err != nil {
		return nil, err
	}

	userDTOGetListStudentResponse := make([]*dto_service_user.GetListStudentDTO, 0)
	for _, value := range userProtoGrpcGetListResult.Data {
		userDTOGetListStudentResponse = append(userDTOGetListStudentResponse, &dto_service_user.GetListStudentDTO{
			ID:        value.Id,
			Name:      value.Name,
			Email:     value.Email,
			Role:      constant.UserRole(value.Role),
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}

	return userDTOGetListStudentResponse, nil
}

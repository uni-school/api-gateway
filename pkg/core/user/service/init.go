package service_user

import (
	dto_service_user "github.com/uni-school/api-gateaway/pkg/core/user/service/dto"
	resource_user "github.com/uni-school/api-gateaway/pkg/core/user/resource"
)

type IUserService interface {
	GetListStudent() ([]*dto_service_user.GetListStudentDTO, error)
}

type UserService struct {
	UserResource resource_user.IUserResource
}

func InitUserService(userResource resource_user.IUserResource) IUserService {
	return &UserService{
		UserResource: userResource,
	}
}

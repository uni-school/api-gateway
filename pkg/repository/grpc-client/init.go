package grpc_client_repository

import (
	pb_user_microservice "github.com/uni-school/api-gateaway/proto/user"
	user_microservice_grpc_client_repository "github.com/uni-school/api-gateaway/pkg/repository/grpc-client/user-microservice"
)

type GRPCClientRepository struct {
	UserMicroservice pb_user_microservice.UserServiceClient
}

func InitGRPCClientRepository() *GRPCClientRepository {
	var (
		userMicroservice = user_microservice_grpc_client_repository.InitUserMicroservice()
	)

	return &GRPCClientRepository{
		UserMicroservice: userMicroservice,
	}
}

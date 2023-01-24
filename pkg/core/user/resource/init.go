package resource_user

import (
	"context"

	grpc_client_repository "github.com/uni-school/api-gateaway/pkg/repository/grpc-client"
	pb_user_microservice "github.com/uni-school/api-gateaway/proto/user"
	"google.golang.org/grpc"
)

type IUserResource interface {
	CreateUser(ctx context.Context, in *pb_user_microservice.CreateUserRequest, opts ...grpc.CallOption) (*pb_user_microservice.CreateUserResponse, error)
	GetUserByEmail(ctx context.Context, in *pb_user_microservice.GetUserByEmailRequest, opts ...grpc.CallOption) (*pb_user_microservice.GetUserByEmailResponse, error)
}

type UserResource struct {
	GRPCClient *grpc_client_repository.GRPCClientRepository
}

func InitUserResource() IUserResource {
	var (
		GRPCClient = grpc_client_repository.InitGRPCClientRepository()
	)

	return &UserResource{
		GRPCClient: GRPCClient,
	}
}

package resource_user

import (
	"context"

	pb_user_microservice "github.com/uni-school/api-gateaway/proto/user"
	"google.golang.org/grpc"
)

func (r *UserResource) CreateUser(ctx context.Context, in *pb_user_microservice.CreateUserRequest, opts ...grpc.CallOption) (*pb_user_microservice.CreateUserResponse, error) {
	return r.GRPCClient.UserMicroservice.CreateUser(ctx, in, opts...)
}

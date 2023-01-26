package resource_user

import (
	"context"

	pb_user_microservice "github.com/uni-school/api-gateaway/proto/user"
	"google.golang.org/grpc"
)

func (r *UserResource) GetListUserByRole(ctx context.Context, in *pb_user_microservice.GetListUserByRoleRequest, opts ...grpc.CallOption) (*pb_user_microservice.GetListUserByRoleResponse, error) {
	return r.GRPCClient.UserMicroservice.GetListUserByRole(ctx, in, opts...)
}

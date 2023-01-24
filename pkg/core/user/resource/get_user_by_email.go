package resource_user

import (
	"context"

	pb_user_microservice "github.com/uni-school/api-gateaway/proto/user"
	"google.golang.org/grpc"
)

func (r *UserResource) GetUserByEmail(ctx context.Context, in *pb_user_microservice.GetUserByEmailRequest, opts ...grpc.CallOption) (*pb_user_microservice.GetUserByEmailResponse, error) {
	return r.GRPCClient.UserMicroservice.GetUserByEmail(ctx, in, opts...)
}

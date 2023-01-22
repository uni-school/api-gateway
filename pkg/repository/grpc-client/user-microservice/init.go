package user_microservice_grpc_client_repository

import (
	"github.com/sirupsen/logrus"
	pb_user_microservice "github.com/uni-school/api-gateaway/proto/user"
	"github.com/uni-school/api-gateaway/shared/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitUserMicroservice() pb_user_microservice.UserServiceClient {
	conn, err := grpc.Dial(config.AppConfig.GRPCServer.UserMicroservice.BaseURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Errorln("cannot connect user microservice", err)
	}

	client := pb_user_microservice.NewUserServiceClient(conn)

	return client
}

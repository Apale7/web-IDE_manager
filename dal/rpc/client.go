package rpc

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	containerManager "web-IDE_manager/proto/container_server"
	user_center "web-IDE_manager/proto/user-center"
)

var (
	userCenterClient       user_center.UserCenterClient
	containerManagerClient containerManager.ManagerClient
)

func init() {
	userCenterClient = user_center.NewUserCenterClient(getConn("localhost:8888"))
	containerManagerClient = containerManager.NewManagerClient(getConn("193.112.177.167:8666"))
}

func getConn(addr string) *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		logrus.Fatalf("%+v", err)
	}
	return conn
}

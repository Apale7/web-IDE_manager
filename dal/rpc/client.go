package rpc

import (
	"web-IDE_manager/proto/docker_manager"
	user_center "web-IDE_manager/proto/user-center"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	userCenterClient    user_center.UserCenterClient
	dockerManagerClient docker_manager.DockerManagerClient
)

func init() {
	userCenterClient = user_center.NewUserCenterClient(getConn("111.230.172.240:9999"))
	// userCenterClient = user_center.NewUserCenterClient(getConn("localhost:9999"))
	dockerManagerClient = docker_manager.NewDockerManagerClient(getConn("111.230.172.240:8888"))
	// dockerManagerClient = docker_manager.NewDockerManagerClient(getConn("localhost:8888"))
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

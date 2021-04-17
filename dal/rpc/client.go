package rpc

import (
	"fmt"
	"os"
	"web-IDE_manager/proto/docker_manager"
	user_center "web-IDE_manager/proto/user-center"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	userCenterClient    user_center.UserCenterClient
	dockerManagerClient docker_manager.DockerManagerClient
)

var (
	userCenterAddr    string
	dockerManagerAddr string
)

func init() {
	initAddrs() //初始化各服务的地址
	userCenterClient = user_center.NewUserCenterClient(getConn(":9999"))
	// userCenterClient = user_center.NewUserCenterClient(getConn("localhost:9999"))
	dockerManagerClient = docker_manager.NewDockerManagerClient(getConn(":8888"))
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

func initAddrs() {
	userCenterAddr = os.Getenv("user_center_addr")
	if userCenterAddr == "" {
		userCenterAddr = ":9999"
	}

	dockerManagerAddr = os.Getenv("docker_manager_addr")
	if dockerManagerAddr == "" {
		dockerManagerAddr = ":8888"
	}

	fmt.Println(userCenterAddr)
	fmt.Println(dockerManagerAddr)
}

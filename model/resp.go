package model

import (
	"web-IDE_manager/proto/docker_manager"
	user_center "web-IDE_manager/proto/user-center"
)

type ContainerResp struct {
	docker_manager.Container
	Status string `json:"status" form:"status"`
}

type LoginResp struct {
	user_center.LoginResponse
	Auth []string `json:"auth"`
}

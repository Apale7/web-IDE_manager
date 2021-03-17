package model

import "web-IDE_manager/proto/docker_manager"

type ContainerResp struct {
	docker_manager.Container
	Status string `json:"status" form:"status"`
}

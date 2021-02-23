package rpc

import (
	"context"
	containerManager "web-IDE_manager/proto/container_server"
)

func GetContainers(ctx context.Context) (containers []*containerManager.ContainerAttr, err error) {
	resp, err := containerManagerClient.ListContainers(ctx, &containerManager.ListContainers_Request{})
	if err != nil {
		return
	}
	return resp.Containers, nil
}

func GetImages(ctx context.Context) (images []*containerManager.ImageAttr, err error) {
	resp, err := containerManagerClient.ListImages(ctx, &containerManager.ListImages_Request{})
	if err != nil {
		return
	}
	return resp.Images, nil
}

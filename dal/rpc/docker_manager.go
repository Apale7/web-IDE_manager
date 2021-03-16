package rpc

import (
	"context"
	"fmt"
	"web-IDE_manager/proto/docker_manager"
)

func GetContainers(ctx context.Context, userID uint, containerID string) (containers []*docker_manager.Container, err error) {
	req := &docker_manager.GetContainerRequest{UserId: uint32(userID), ContainerId: containerID}
	resp, err := dockerManagerClient.GetContainer(ctx, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	return resp.Containers, nil
}

func GetImages(ctx context.Context, userID uint, imageID string) (images []*docker_manager.Image, err error) {
	req := &docker_manager.GetImageRequest{UserId: uint32(userID), ImageId: imageID}
	resp, err := dockerManagerClient.GetImage(ctx, req)
	if err != nil {
		return
	}
	return resp.Images, nil
}

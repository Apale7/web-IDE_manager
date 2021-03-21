package rpc

/*
	所有的权限检测和参数校验由路由中间件完成，rpc不关注有没有权限
*/
import (
	"context"
	"fmt"
	"web-IDE_manager/model"
	"web-IDE_manager/proto/docker_manager"

	"google.golang.org/protobuf/types/known/emptypb"
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

func DeleteContainer(ctx context.Context, userID uint, containerID string) (err error) {
	req := &docker_manager.DeleteContainerRequest{
		UserId:      uint32(userID),
		ContainerId: containerID,
	}
	_, err = dockerManagerClient.DeleteContainer(ctx, req)
	if err != nil {
		return
	}

	return
}

func CreateContainer(ctx context.Context, reqParams model.CreateContainerReqBody) (err error) {
	req := &docker_manager.CreateContainerRequest{
		Username:      reqParams.Username,
		ContainerName: reqParams.ContainerName,
		UserId:        uint32(reqParams.UserID),
		ImageId:       reqParams.ImageID,
	}
	_, err = dockerManagerClient.CreateContainer(ctx, req)
	if err != nil {
		return
	}
	return
}

func PruneContainer(ctx context.Context) (err error) {
	_, err = dockerManagerClient.PruneContainers(ctx, &emptypb.Empty{})
	if err != nil {
		return
	}

	return
}

func GetImages(ctx context.Context, userID uint, imageID string, isAdmin bool) (images []*docker_manager.Image, err error) {
	req := &docker_manager.GetImageRequest{UserId: uint32(userID), ImageId: imageID, IsAdmin: isAdmin}
	resp, err := dockerManagerClient.GetImage(ctx, req)
	if err != nil {
		return
	}
	return resp.Images, nil
}

func CreateImage(ctx context.Context, params model.CreateImageReqBody) (err error) {
	req := &docker_manager.CreateImageRequest{
		Type:          params.Type,
		UserId:        params.UserID,
		Dockerfile:    []byte(params.Dockerfile),
		RepositoryUrl: params.RespositryURL,
		Tag:           params.Tag,
		Username:      params.Username,
		Password:      params.Password,
		ImageUrl:      params.ImageUrl,
	}
	_, err = dockerManagerClient.CreateImage(ctx, req)
	if err != nil {
		return
	}
	return
}

func DeleteImage(ctx context.Context, userID uint, imageID string) (err error) {
	req := &docker_manager.DeleteImageRequest{
		UserId:  uint32(userID),
		ImageId: imageID,
	}
	_, err = dockerManagerClient.DeleteImage(ctx, req)
	if err != nil {
		return
	}
	return
}

func PruneImage(ctx context.Context) (err error) {
	_, err = dockerManagerClient.PruneImages(ctx, &emptypb.Empty{})
	if err != nil {
		return
	}

	return
}

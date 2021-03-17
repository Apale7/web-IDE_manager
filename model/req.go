package model

type ContainerbaseInfo struct {
	UserID      int    `json:"user_id" form:"user_id"`
	ContainerID string `json:"container_id" form:"container_id"`
}

type ImageBaseInfo struct {
	UserID  int    `json:"user_id" form:"user_id"`
	ImageID string `json:"image_id" form:"image_id"`
}

type GetContainerReqBody struct {
	ContainerbaseInfo
}

type DeleteContainerReqBody struct {
	ContainerbaseInfo
}

type CreateContainerReqBody struct {
	ImageBaseInfo
	ContainerName string    `json:"container_name" form:"container_name"`
}
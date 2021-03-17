package model

import user_center "web-IDE_manager/proto/user-center"

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
	ContainerName string `json:"container_name" form:"container_name"`
	Username      string `json:"username" form:"username"`
}

type GetGroupReqBody struct {
	user_center.Group
	MemBerID uint32 `json:"member_id" form:"member_id"`
}

type GetImageReqBody struct {
	ImageBaseInfo
	// 是管理员，则返回管理员拥有的所有image；非管理员，则返回所有有权限使用的image
	IsAdmin bool `json:"is_admin" form:"is_admin"`
}

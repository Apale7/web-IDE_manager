package model

import "web-IDE_manager/proto/docker_manager"

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
	GroupID   uint32 `json:"group_id" form:"group_id"`
	OwnerID   uint32 `json:"owner_id" form:"owner_id"`
	MemBerID  uint32 `json:"member_id" form:"member_id"`
	GroupName string `json:"group_name" form:"group_name"`
	HaveMe    bool   `json:"have_me" form:"have_me"`
}

type GetImageReqBody struct {
	ImageBaseInfo
	// 是管理员，则返回管理员拥有的所有image；非管理员，则返回所有有权限使用的image
	IsAdmin bool `json:"is_admin" form:"is_admin"`
}

type CreateImageReqBody struct {
	Type docker_manager.CreateImageType

	UserID uint32 `json:"user_id" form:"user_id"`
	// by dockerfile
	Dockerfile string `json:"dockerfile" form:"dockerfile"`

	// by pull
	RespositryURL string `json:"respositry_url" form:"respositry_url"`
	Tag           string `json:"tag" form:"tag"`
	Username      string `json:"username" form:"username"`
	Password      string `json:"password" form:"password"`
	// by upload
	ImageUrl string `json:"image_url" form:"image_url"`
}

type DeleteImageReqBody struct {
	UserID  uint32 `json:"user_id" form:"user_id"`
	ImageID string `json:"image_id" form:"image_id"`
}
type LoginReq struct {
	Base
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RefreshReq struct {
	Base
	RefreshToken string `json:"refresh_token"`
}

type CreateGroupReqBody struct {
	UserID    uint32 `json:"user_id" form:"user_id"`
	GroupName string `json:"group_name" form:"group_name"`
}

type JoinGroupReqBody struct {
	UserID  uint32 `json:"user_id" form:"user_id"`
	GroupID uint32 `json:"group_id" form:"group_id"`
}

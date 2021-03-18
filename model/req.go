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
	ContainerName string `json:"container_name" form:"container_name"`
	Username      string `json:"username" form:"username"`
}

type GetGroupReqBody struct {
	GroupID   uint32 `json:"group_id" form:"group_id"`
	OwnerID   uint32 `json:"owner_id" form:"owner_id"`
	MemBerID  uint32 `json:"member_id" form:"member_id"`
	GroupName string `json:"group_name" form:"group_name"`
}

type GetImageReqBody struct {
	ImageBaseInfo
	// 是管理员，则返回管理员拥有的所有image；非管理员，则返回所有有权限使用的image
	IsAdmin bool `json:"is_admin" form:"is_admin"`
}

package model

type GetContainerReqBody struct {
	UserID      int    `json:"user_id" form:"user_id"`
	ContainerID string `json:"container_id" form:"container_id"`
}

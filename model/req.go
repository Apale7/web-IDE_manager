package model

type GetContainerReqBody struct {
	UserID      uint32 `json:"user_id"`
	ContainerID string `json:"container_id"`
}

package main

import (
	"web-IDE_manager/handler"

	"github.com/gin-gonic/gin"
)

func collectRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		container := api.Group("/container")
		{
			container.GET("/get_containers", handler.GetContainers)
		}
		image := api.Group("/image")
		{
			image.GET("/get_images", handler.GetImages)
		}
		group := api.Group("/group")
		{
			group.GET("/get_groups", handler.GetGroup)
		}
	}
}

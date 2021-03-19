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
			container.GET("/get", handler.GetContainers)
			container.POST("/create", handler.CreateContainer)
			container.POST("/delete", handler.DeleteContainer)
		}
		image := api.Group("/image")
		{
			image.GET("/get", handler.GetImages)
			image.POST("/create", handler.CreateImage)
			image.POST("/delete", handler.DeleteImage)
		}
		group := api.Group("/group")
		{
			group.GET("/get", handler.GetGroup)
			// group.POST("/create", handler.CreateGroup)
		}
		user := api.Group("/user")
		{
			user.POST("/login", handler.Login)
		}
	}
}

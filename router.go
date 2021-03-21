package main

import (
	"web-IDE_manager/handler"
	"web-IDE_manager/middleware"

	"github.com/gin-gonic/gin"
)

func collectRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		container := api.Group("/container", middleware.JWTAuthMiddleware)
		{
			container.GET("/get", handler.GetContainers)
			container.POST("/create", handler.CreateContainer)
			container.POST("/delete", handler.DeleteContainer)
		}
		image := api.Group("/image", middleware.JWTAuthMiddleware)
		{
			image.GET("/get", handler.GetImages)
			image.POST("/create", handler.CreateImage)
			image.POST("/delete", handler.DeleteImage)
		}
		group := api.Group("/group", middleware.JWTAuthMiddleware)
		{
			group.GET("/get", handler.GetGroup)
			group.POST("/create", handler.CreateGroup)
			group.POST("/join", handler.JoinGroup)
			group.GET("/get_group_members", handler.GetGroupMembers)
		}
		user := api.Group("/user")
		{
			user.POST("/login", handler.Login)
			user.POST("/refresh", handler.Refresh)
		}
	}
}

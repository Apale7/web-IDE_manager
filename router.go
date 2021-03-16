package main

import (
	"web-IDE_manager/handler"

	"github.com/gin-gonic/gin"
)

func collectRoutes(r *gin.Engine) {
	r.GET("/api/get_containers", handler.GetContainers)
	// r.GET("/api/get_images", handler.GetImages)
}

package handler

import (
	"context"
	"web-IDE_manager/dal/rpc"

	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context) {
	ctx := context.Background()
	containers, err := rpc.GetImages(ctx)
	if err != nil {
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, gin.H{"code": 0, "images": containers})
}

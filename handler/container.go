package handler

import (
	"context"
	"web-IDE_manager/dal/rpc"

	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
)

func GetContainers(c *gin.Context) {
	ctx := context.Background()
	containers, err := rpc.GetContainers(ctx)
	if err != nil {
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, gin.H{"code": 0, "containers": containers})
}

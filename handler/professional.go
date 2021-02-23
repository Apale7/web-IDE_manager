package handler

import (
	"context"
	"web-IDE_manager/dal/rpc"

	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
)

//GetFile 获取一个文件的内容
func GetFile(c *gin.Context) {

}

//GetDir 获取一个目录
func GetDir(c *gin.Context) {

}

func GetContainers(c *gin.Context) {
	ctx := context.Background()
	containers, err := rpc.GetContainers(ctx)
	if err != nil {
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, gin.H{"code": 0, "containers": containers})
}

func GetImages(c *gin.Context) {
	ctx := context.Background()
	containers, err := rpc.GetImages(ctx)
	if err != nil {
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, gin.H{"code": 0, "images": containers})
}

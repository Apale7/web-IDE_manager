package handler

import (
	"context"
	"web-IDE_manager/dal/rpc"
	"web-IDE_manager/model"

	"github.com/Apale7/common/constdef"
	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetContainers(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.GetContainerReqBody
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("Bind error, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}
	logrus.Infof("req: %+v", reqBody)
	containers, err := rpc.GetContainers(ctx, uint(reqBody.UserID), reqBody.ContainerID)
	if err != nil {
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, gin.H{"code": 0, "containers": containers})
}

func DeleteContainer(c *gin.Context) {
	
}
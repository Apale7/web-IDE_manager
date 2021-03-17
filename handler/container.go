package handler

import (
	"context"
	"errors"
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

	containerResp := make([]*model.ContainerResp, 0, len(containers))

	for _, ctn := range containers {
		res := &model.ContainerResp{
			Container: *ctn,
			Status:    ctn.Status.String(),
		}
		containerResp = append(containerResp, res)
	}

	utils.RetData(c, gin.H{"code": 0, "containers": containerResp})
}

func CreateContainer(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.CreateContainerReqBody
	if err := c.ShouldBind(&reqBody); err != nil || checkCreateContainer(reqBody) {
		logrus.Warnf("Bind error, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}

	if reqBody.Username == "" {
		resp, err := rpc.GetUserInfo(ctx, uint32(reqBody.UserID), "")
		if err != nil {
			logrus.Warnf("GetUserInfo error: %+v", err)
			utils.RetErr(c, errors.New("错误的user_id"))
			return
		}

		reqBody.Username = resp.Username
	}

	logrus.Infof("req: %+v", reqBody)

	rpc.CreateContainer(ctx, reqBody)
}

func DeleteContainer(c *gin.Context) {

}

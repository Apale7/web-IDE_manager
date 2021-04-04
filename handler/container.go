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
	if err := c.ShouldBind(&reqBody); err != nil || !checkCreateContainer(reqBody) {
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

	err := rpc.CreateContainer(ctx, reqBody)
	if err != nil {
		utils.RetErr(c, errors.New("创建容器失败"))
		logrus.Warnf("CreateContainer error, err: %v", err)
		return
	}
	utils.RetSuccess(c)
}

func DeleteContainer(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.DeleteContainerReqBody
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("Bind error, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}

	logrus.Infof("req: %+v", reqBody)

	err := rpc.DeleteContainer(ctx, uint(reqBody.UserID), reqBody.ContainerID)

	if err != nil {
		utils.RetErr(c, errors.New("删除容器失败"))
		logrus.Warnf("DeleteContainer error, err: %v", err)
		return
	}
	utils.RetSuccess(c)
}

func StartContainer(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.StartContainerReqBody
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("Bind error, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}
	logrus.Infof("req: %+v", reqBody)

	err := rpc.StartContainer(ctx, uint(reqBody.UserID), reqBody.ContainerID)
	if err != nil {
		utils.RetErr(c, errors.New("启动容器失败"))
		logrus.Warnf("StartContainer error, err: %v", err)
		return
	}
	utils.RetSuccess(c)
}

func StopContainer(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.StopContainerReqBody
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("Bind error, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}
	logrus.Infof("req: %+v", reqBody)

	err := rpc.StopContainer(ctx, uint(reqBody.UserID), reqBody.ContainerID)
	if err != nil {
		utils.RetErr(c, errors.New("关闭容器失败"))
		logrus.Warnf("StopContainer error, err: %v", err)
		return
	}
	utils.RetSuccess(c)
}

func RestartContainer(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.RestartContainerReqBody
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("Bind error, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}
	logrus.Infof("req: %+v", reqBody)

	err := rpc.RestartContainer(ctx, uint(reqBody.UserID), reqBody.ContainerID)
	if err != nil {
		utils.RetErr(c, errors.New("重启容器失败"))
		logrus.Warnf("RestartContainer error, err: %v", err)
		return
	}
	utils.RetSuccess(c)
}

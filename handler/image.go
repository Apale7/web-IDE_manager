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

func GetImages(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.GetImageReqBody
	if err := c.Bind(&reqBody); err != nil {
		logrus.Warnf("Bind error, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}
	logrus.Infof("get reqBody: %+v", reqBody)

	images, err := rpc.GetImages(ctx, uint(reqBody.UserID), reqBody.ImageID, reqBody.IsAdmin)
	if err != nil {
		logrus.Warnf("GetImages error, err: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, gin.H{"images": images})
	logrus.Infof("successful return %d image info", len(images))
}

func CreateImage(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.CreateImageReqBody
	if err := c.Bind(&reqBody); err != nil {
		logrus.Warnf("Bind error, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}
	logrus.Infof("get reqBody: %+v", reqBody)

	err := rpc.CreateImage(ctx, uint(reqBody.UserID), reqBody.Dockerfile)
	if err != nil {
		logrus.Warnf("CreateImage error, err: %v", err)
		utils.RetErr(c, errors.New("创建镜像失败"))
		return
	}
	utils.RetSuccess(c)
}

func DeleteImage(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.DeleteImageReqBody
	if err := c.Bind(&reqBody); err != nil {
		logrus.Warnf("Bind error, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}
	logrus.Infof("get reqBody: %+v", reqBody)

	err := rpc.DeleteImage(ctx, uint(reqBody.UserID), reqBody.ImageID)
	if err != nil {
		utils.RetErr(c, errors.New("删除镜像失败"))
		logrus.Warnf("DeleteImage error, err: %v", err)
		return
	}
	utils.RetSuccess(c)
}

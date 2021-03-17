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
	logrus.Infof("successful return %d image imfo", len(images))
}

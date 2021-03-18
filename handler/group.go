package handler

import (
	"context"
	"web-IDE_manager/dal/rpc"
	"web-IDE_manager/model"
	user_center "web-IDE_manager/proto/user-center"

	"github.com/Apale7/common/constdef"
	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetGroup(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.GetGroupReqBody
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("Bind error, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}

	logrus.Infof("get reqBody: %+v", reqBody)

	groupInfo := &user_center.Group{
		Id:        reqBody.GroupID,
		OwnerId:   reqBody.OwnerID,
		GroupName: reqBody.GroupName,
	}

	groups, err := rpc.GetGroup(ctx, groupInfo, reqBody.MemBerID)
	if err != nil {
		logrus.Warnf("GetGroup error, err: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, gin.H{"groups": groups})
	logrus.Infof("success return %d group info", len(groups))
}

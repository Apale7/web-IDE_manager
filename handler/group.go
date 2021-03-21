package handler

import (
	"context"
	"errors"
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

	groups, err := rpc.GetGroup(ctx, groupInfo, reqBody.MemBerID, reqBody.HaveMe)
	if err != nil {
		logrus.Warnf("GetGroup error, err: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, gin.H{"groups": groups})
	logrus.Infof("success return %d group info", len(groups))
}

func CreateGroup(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.CreateGroupReqBody
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("Bind error, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}

	logrus.Infof("get reqBody: %+v", reqBody)

	err := rpc.CreateGroup(ctx, uint(reqBody.UserID), reqBody.GroupName)
	if err != nil {
		logrus.Warnf("CreateGroup error, err: %v", err)
		utils.RetErr(c, errors.New("创建团队失败"))
		return
	}

	utils.RetSuccess(c)
}

func GetGroupMembers(c *gin.Context) {
	ctx := context.Background()
	groupID := utils.GetParamsInt64(c, "group_id")
	if groupID <= 0 {
		logrus.Warnf("错误的group_id")
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}

	members, err := rpc.GetGroupMembers(ctx, uint(groupID))
	if err != nil {
		logrus.Warnf("GetGroupMembers errors, err: %v", err)
		utils.RetErr(c, errors.New("获取团队成员失败"))
		return
	}
	for i := range members {
		members[i].PhoneNumber = members[i].PhoneNumber[:3] + "****" + members[i].PhoneNumber[7:]
	}
	utils.RetData(c, gin.H{"members": members})
}

func JoinGroup(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.JoinGroupReqBody
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("Bind error, err: %v", err)
		utils.RetErr(c, constdef.ErrInvalidParams)
		return
	}

	logrus.Infof("get reqBody: %+v", reqBody)

	err := rpc.JoinGroup(ctx, uint(reqBody.UserID), uint(reqBody.GroupID))

	if err != nil {
		logrus.Warnf("JoinGroup error, err: %v", err)
		utils.RetErr(c, errors.New("加入团队失败"))
		return
	}

	utils.RetSuccess(c)
}

package handler

import (
	"context"
	"web-IDE_manager/dal/rpc"
	"web-IDE_manager/model"

	"github.com/Apale7/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Login(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.LoginReq
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	logrus.Infof("reqBody: %+v", reqBody)
	resp, err := rpc.Login(ctx, reqBody.Username, reqBody.Password)
	if err != nil {
		logrus.Warnf("login failed, err: %v", err)
		utils.RetErr(c, err)
		return
	}
	var loginRes model.LoginResp
	loginRes.LoginResponse = *resp
	loginRes.Auth = []string{"super"}
	utils.RetData(c, loginRes)
}

func Register(c *gin.Context) {

}

func Refresh(c *gin.Context) {
	ctx := context.Background()
	var reqBody model.RefreshReq
	if err := c.ShouldBind(&reqBody); err != nil {
		logrus.Warnf("invalid params: %v", err)
		utils.RetErr(c, err)
		return
	}
	logrus.Infof("reqBody: %+v", reqBody)

	resp, err := rpc.Refresh(ctx, reqBody.RefreshToken)
	if err != nil {
		logrus.Warnf("refresh failed, err: %v", err)
		utils.RetErr(c, err)
		return
	}
	utils.RetData(c, resp)
}

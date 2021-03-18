package rpc

import (
	"context"
	user_center "web-IDE_manager/proto/user-center"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Register 注册
func Register(ctx context.Context, user *user_center.User, extra *user_center.UserExtra) (err error) {
	req := &user_center.RegisterRequest{
		User:      user,
		UserExtra: extra,
	}
	resp, err := userCenterClient.Register(ctx, req)
	if err != nil {
		logrus.Warnln(err)
		return
	}
	logrus.Infof("注册成功, id: %d\n", resp.Id)
	return
}

func Login(ctx context.Context, username, password string) (resp *user_center.LoginResponse, err error) {
	req := &user_center.LoginRequest{
		Username: username,
		Password: password,
	}
	resp, err = userCenterClient.Login(ctx, req)
	if err != nil {
		return
	}

	return
}

func CheckToken(ctx context.Context, token string) (userInfo *user_center.UserInfo, err error) {
	req := &user_center.CheckTokenRequest{Token: token}
	resp, err := userCenterClient.CheckToken(ctx, req)
	if err != nil {
		return
	}
	return resp.User, nil
}

func Refresh(ctx context.Context, refreshToken string) (resp *user_center.RefreshResponse, err error) {
	req := &user_center.RefreshRequest{RefreshToken: refreshToken}
	resp, err = userCenterClient.Refresh(ctx, req)

	if err != nil {
		return
	}

	return
}

func CreateGroup(ctx context.Context, ownerID uint, groupName string) (err error) {
	req := &user_center.CreateGroupRequest{
		OwnerId:   uint32(ownerID),
		GroupName: groupName,
	}
	_, err = userCenterClient.CreateGroup(ctx, req)
	if err != nil {
		return
	}

	return
}

func JoinGroup(ctx context.Context, userID, groupID uint) (err error) {
	req := &user_center.JoinGroupRequest{
		UserId:  uint32(userID),
		GroupId: uint32(groupID),
	}
	_, err = userCenterClient.JoinGroup(ctx, req)
	if err != nil {
		return
	}
	return
}

func ExitGroup(ctx context.Context, userID, groupID uint) (err error) {
	req := &user_center.ExitGroupRequest{
		UserId:  uint32(userID),
		GroupId: uint32(groupID),
	}
	_, err = userCenterClient.ExitGroup(ctx, req)
	if err != nil {
		return
	}
	return
}

func GetGroup(ctx context.Context, groupInfo *user_center.Group, memberID uint32) (groups []*user_center.Group, err error) {
	req := &user_center.GetGroupRequest{
		GroupInfo: groupInfo,
		MemberId:  memberID,
	}
	resp, err := userCenterClient.GetGroup(ctx, req)
	if err != nil {
		return
	}
	return resp.Groups, nil
}

func GetGroupMembers(ctx context.Context, groupID uint) (members []*user_center.UserExtra, err error) {
	req := &user_center.GetGroupMembersRequest{
		GroupId: uint32(groupID),
	}
	resp, err := userCenterClient.GetGroupMembers(ctx, req)
	if err != nil {
		return
	}
	return resp.Members, nil
}

func GetUserInfo(ctx context.Context, userID uint32, username string) (resp *user_center.GetUserInfoResponse, err error) {
	req := &user_center.GetUserInfoRequest{
		UserId:   userID,
		Username: username,
	}
	resp, err = userCenterClient.GetUserInfo(ctx, req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

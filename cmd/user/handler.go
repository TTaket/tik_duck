package main

import (
	"context"

	userservice "github.com/808-not-found/tik_duck/cmd/user/userService"
	user "github.com/808-not-found/tik_duck/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// User注册.
func (s *UserServiceImpl) UserRegister(
	ctx context.Context,
	req *user.UserRegisterRequest,
) (resp *user.UserRegisterResponse, err error) {
	// 1: i32 StatusCode //状态码，0-成功，其他值失败
	// 2: optional string StatusMsg //返回状态描述
	// 3: i64 UserId //用户id
	// 4: string Token //用户鉴权token
	// 生成回应结构体
	resp = new(user.UserRegisterResponse)
	// 校验参数
	err = req.IsValid()
	if err != nil {
		return resp, err
	}
	// 实现逻辑
	// StatusCode, StatusMsg, UserId, Token, err := userservice.UserRegisterService(ctx, req)
	StatusCode, StatusMsg, _, Token, err := userservice.UserRegisterService(ctx, req)
	if err != nil {
		return resp, err
	}
	resp.StatusCode = StatusCode
	resp.StatusMsg = &StatusMsg
	// resp.UserId = UserId //ToDo: 等待类型更新
	resp.Token = Token
	// 返回结构体
	return resp, nil
}

// 视频流
// UserGetFeed implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserGetFeed(
	ctx context.Context,
	req *user.FeedRequest,
) (resp *user.FeedResponse, err error) {
	resp = new(user.FeedResponse)

	// statusCode, statusMsg, videoList, nextTime, err := userservice.UserGetFeedService(ctx, req)
	statusCode, statusMsg, videoList, _, err := userservice.UserGetFeedService(ctx, req)
	// if err = req.IsValid(); err != nil {
	// 	resp.StatusCode = 1001
	// 	return resp, nil
	// }
	if err != nil {
		return resp, err
	}
	resp.StatusCode = statusCode
	resp.StatusMsg = &statusMsg
	resp.VideoList = videoList

	// resp.NextTime = &nextTime
	return resp, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(
	ctx context.Context,
	req *user.UserLoginRequest,
) (resp *user.UserLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(
	ctx context.Context,
	req *user.UserRequest,
) (resp *user.UserResponse, err error) {
	// TODO: Your code here...
	return
}

// UserPublishList implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserPublishList(
	ctx context.Context,
	req *user.PublishListRequest,
) (resp *user.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// UserPublishAction implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserPublishAction(
	ctx context.Context,
	req *user.PublishActionRequest,
) (resp *user.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// UserTest implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserTest(
	ctx context.Context,
	req *user.Testinfo,
) (resp *user.Testinfo, err error) {
	// TODO: Your code here...
	resp = &user.Testinfo{Testinfo: req.Testinfo}
	return
}
// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"github.com/liankui/e-commerce/apps/user/rpc/internal/logic"
	"github.com/liankui/e-commerce/apps/user/rpc/internal/svc"
	"github.com/liankui/e-commerce/apps/user/rpc/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// 登录
func (s *UserServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

// 获取用户信息
func (s *UserServer) UserInfo(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

// 添加收获地址
func (s *UserServer) AddUserReceiveAddress(ctx context.Context, in *user.UserReceiveAddressAddReq) (*user.UserReceiveAddressAddRes, error) {
	l := logic.NewAddUserReceiveAddressLogic(ctx, s.svcCtx)
	return l.AddUserReceiveAddress(in)
}

// 编辑收获地址
func (s *UserServer) EditUserReceiveAddress(ctx context.Context, in *user.UserReceiveAddressEditReq) (*user.UserReceiveAddressEditRes, error) {
	l := logic.NewEditUserReceiveAddressLogic(ctx, s.svcCtx)
	return l.EditUserReceiveAddress(in)
}

// 删除收获地址
func (s *UserServer) DelUserReceiveAddress(ctx context.Context, in *user.UserReceiveAddressDelReq) (*user.UserReceiveAddressDelRes, error) {
	l := logic.NewDelUserReceiveAddressLogic(ctx, s.svcCtx)
	return l.DelUserReceiveAddress(in)
}

// 获取收获地址列表
func (s *UserServer) GetUserReceiveAddressList(ctx context.Context, in *user.UserReceiveAddressListReq) (*user.UserReceiveAddressListRes, error) {
	l := logic.NewGetUserReceiveAddressListLogic(ctx, s.svcCtx)
	return l.GetUserReceiveAddressList(in)
}

//  添加收藏
func (s *UserServer) AddUserCollection(ctx context.Context, in *user.UserCollectionAddReq) (*user.UserCollectionAddRes, error) {
	l := logic.NewAddUserCollectionLogic(ctx, s.svcCtx)
	return l.AddUserCollection(in)
}

//  删除收藏
func (s *UserServer) DelUserCollection(ctx context.Context, in *user.UserCollectionDelReq) (*user.UserCollectionDelRes, error) {
	l := logic.NewDelUserCollectionLogic(ctx, s.svcCtx)
	return l.DelUserCollection(in)
}

//  收藏列表
func (s *UserServer) GetUserCollectionList(ctx context.Context, in *user.UserCollectionListReq) (*user.UserCollectionListRes, error) {
	l := logic.NewGetUserCollectionListLogic(ctx, s.svcCtx)
	return l.GetUserCollectionList(in)
}

// 根据主键id,查询收获地址
func (s *UserServer) GetUserReceiveAddressInfo(ctx context.Context, in *user.UserReceiveAddressInfoReq) (*user.UserReceiveAddress, error) {
	l := logic.NewGetUserReceiveAddressInfoLogic(ctx, s.svcCtx)
	return l.GetUserReceiveAddressInfo(in)
}

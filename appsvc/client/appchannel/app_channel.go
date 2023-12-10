// Code generated by goctl. DO NOT EDIT.
// Source: appsvc.proto

package appchannel

import (
	"context"

	"github.com/aheadIV/NightVoyager/appsvc/types/appsvc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	App                = appsvc.App
	AppListReq         = appsvc.AppListReq
	AppListResp        = appsvc.AppListResp
	Channel            = appsvc.Channel
	ChannelListReq     = appsvc.ChannelListReq
	ChannelListResp    = appsvc.ChannelListResp
	CreateAppReq       = appsvc.CreateAppReq
	CreateAppResp      = appsvc.CreateAppResp
	CreateChannelReq   = appsvc.CreateChannelReq
	CreateChannelResp  = appsvc.CreateChannelResp
	GetAppByIdReq      = appsvc.GetAppByIdReq
	GetAppByIdResp     = appsvc.GetAppByIdResp
	GetChannelByIdReq  = appsvc.GetChannelByIdReq
	GetChannelByIdResp = appsvc.GetChannelByIdResp
	Req                = appsvc.Req
	Resp               = appsvc.Resp

	AppChannel interface {
		CreateChannel(ctx context.Context, in *CreateChannelReq, opts ...grpc.CallOption) (*CreateChannelResp, error)
		GetChannelById(ctx context.Context, in *GetChannelByIdReq, opts ...grpc.CallOption) (*GetChannelByIdResp, error)
		ChannelList(ctx context.Context, in *ChannelListReq, opts ...grpc.CallOption) (*ChannelListResp, error)
	}

	defaultAppChannel struct {
		cli zrpc.Client
	}
)

func NewAppChannel(cli zrpc.Client) AppChannel {
	return &defaultAppChannel{
		cli: cli,
	}
}

func (m *defaultAppChannel) CreateChannel(ctx context.Context, in *CreateChannelReq, opts ...grpc.CallOption) (*CreateChannelResp, error) {
	client := appsvc.NewAppChannelClient(m.cli.Conn())
	return client.CreateChannel(ctx, in, opts...)
}

func (m *defaultAppChannel) GetChannelById(ctx context.Context, in *GetChannelByIdReq, opts ...grpc.CallOption) (*GetChannelByIdResp, error) {
	client := appsvc.NewAppChannelClient(m.cli.Conn())
	return client.GetChannelById(ctx, in, opts...)
}

func (m *defaultAppChannel) ChannelList(ctx context.Context, in *ChannelListReq, opts ...grpc.CallOption) (*ChannelListResp, error) {
	client := appsvc.NewAppChannelClient(m.cli.Conn())
	return client.ChannelList(ctx, in, opts...)
}

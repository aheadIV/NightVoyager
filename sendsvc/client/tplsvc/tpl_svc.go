// Code generated by goctl. DO NOT EDIT.
// Source: sendsvc.proto

package tplsvc

import (
	"context"

	"github.com/aheadIV/NightVoyager/sendsvc/types/sendsvc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateRecReq   = sendsvc.CreateRecReq
	CreateRecResp  = sendsvc.CreateRecResp
	CreateSigReq   = sendsvc.CreateSigReq
	CreateSigResp  = sendsvc.CreateSigResp
	CreateTplReq   = sendsvc.CreateTplReq
	CreateTplResp  = sendsvc.CreateTplResp
	DeletTplResp   = sendsvc.DeletTplResp
	DeleteSigReq   = sendsvc.DeleteSigReq
	DeleteSigResp  = sendsvc.DeleteSigResp
	DeleteTplReq   = sendsvc.DeleteTplReq
	GetRecByIdReq  = sendsvc.GetRecByIdReq
	GetRecByIdResp = sendsvc.GetRecByIdResp
	GetSigByIdReq  = sendsvc.GetSigByIdReq
	GetSigByIdResp = sendsvc.GetSigByIdResp
	GetTplByIdReq  = sendsvc.GetTplByIdReq
	GetTplByIdResp = sendsvc.GetTplByIdResp
	Rec            = sendsvc.Rec
	SigListReq     = sendsvc.SigListReq
	SigListResp    = sendsvc.SigListResp
	Sign           = sendsvc.Sign
	Tpl            = sendsvc.Tpl
	TplListReq     = sendsvc.TplListReq
	TplListResp    = sendsvc.TplListResp
	UpdateSigReq   = sendsvc.UpdateSigReq
	UpdateSigResp  = sendsvc.UpdateSigResp
	UpdateTplReq   = sendsvc.UpdateTplReq
	UpdateTplResp  = sendsvc.UpdateTplResp

	TplSvc interface {
		GetTplById(ctx context.Context, in *GetTplByIdReq, opts ...grpc.CallOption) (*GetTplByIdResp, error)
		CreateTpl(ctx context.Context, in *CreateTplReq, opts ...grpc.CallOption) (*CreateTplResp, error)
		UpdateTpl(ctx context.Context, in *UpdateTplReq, opts ...grpc.CallOption) (*UpdateTplResp, error)
		DeleteTpl(ctx context.Context, in *DeleteTplReq, opts ...grpc.CallOption) (*DeletTplResp, error)
		TplList(ctx context.Context, in *TplListReq, opts ...grpc.CallOption) (*TplListResp, error)
	}

	defaultTplSvc struct {
		cli zrpc.Client
	}
)

func NewTplSvc(cli zrpc.Client) TplSvc {
	return &defaultTplSvc{
		cli: cli,
	}
}

func (m *defaultTplSvc) GetTplById(ctx context.Context, in *GetTplByIdReq, opts ...grpc.CallOption) (*GetTplByIdResp, error) {
	client := sendsvc.NewTplSvcClient(m.cli.Conn())
	return client.GetTplById(ctx, in, opts...)
}

func (m *defaultTplSvc) CreateTpl(ctx context.Context, in *CreateTplReq, opts ...grpc.CallOption) (*CreateTplResp, error) {
	client := sendsvc.NewTplSvcClient(m.cli.Conn())
	return client.CreateTpl(ctx, in, opts...)
}

func (m *defaultTplSvc) UpdateTpl(ctx context.Context, in *UpdateTplReq, opts ...grpc.CallOption) (*UpdateTplResp, error) {
	client := sendsvc.NewTplSvcClient(m.cli.Conn())
	return client.UpdateTpl(ctx, in, opts...)
}

func (m *defaultTplSvc) DeleteTpl(ctx context.Context, in *DeleteTplReq, opts ...grpc.CallOption) (*DeletTplResp, error) {
	client := sendsvc.NewTplSvcClient(m.cli.Conn())
	return client.DeleteTpl(ctx, in, opts...)
}

func (m *defaultTplSvc) TplList(ctx context.Context, in *TplListReq, opts ...grpc.CallOption) (*TplListResp, error) {
	client := sendsvc.NewTplSvcClient(m.cli.Conn())
	return client.TplList(ctx, in, opts...)
}

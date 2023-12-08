// Code generated by goctl. DO NOT EDIT.
// Source: tplsvc.proto

package server

import (
	"context"

	"github.com/aheadIV/NightVoyager/tplsvc/internal/logic/tplsignsvc"
	"github.com/aheadIV/NightVoyager/tplsvc/internal/svc"
	"github.com/aheadIV/NightVoyager/tplsvc/types/tplsvc"
)

type TplSignSvcServer struct {
	svcCtx *svc.ServiceContext
	tplsvc.UnimplementedTplSignSvcServer
}

func NewTplSignSvcServer(svcCtx *svc.ServiceContext) *TplSignSvcServer {
	return &TplSignSvcServer{
		svcCtx: svcCtx,
	}
}

func (s *TplSignSvcServer) GetSigById(ctx context.Context, in *tplsvc.GetSigByIdReq) (*tplsvc.GetSigByIdResp, error) {
	l := tplsignsvclogic.NewGetSigByIdLogic(ctx, s.svcCtx)
	return l.GetSigById(in)
}

func (s *TplSignSvcServer) CreateSig(ctx context.Context, in *tplsvc.CreateSigReq) (*tplsvc.CreateSigResp, error) {
	l := tplsignsvclogic.NewCreateSigLogic(ctx, s.svcCtx)
	return l.CreateSig(in)
}

func (s *TplSignSvcServer) UpdateSig(ctx context.Context, in *tplsvc.UpdateSigReq) (*tplsvc.UpdateSigResp, error) {
	l := tplsignsvclogic.NewUpdateSigLogic(ctx, s.svcCtx)
	return l.UpdateSig(in)
}

func (s *TplSignSvcServer) DeleteSig(ctx context.Context, in *tplsvc.DeleteSigReq) (*tplsvc.DeleteSigResp, error) {
	l := tplsignsvclogic.NewDeleteSigLogic(ctx, s.svcCtx)
	return l.DeleteSig(in)
}

func (s *TplSignSvcServer) SigList(ctx context.Context, in *tplsvc.SigListReq) (*tplsvc.SigListResp, error) {
	l := tplsignsvclogic.NewSigListLogic(ctx, s.svcCtx)
	return l.SigList(in)
}

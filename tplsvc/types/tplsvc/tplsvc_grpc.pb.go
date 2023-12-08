// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.0
// source: tplsvc.proto

package tplsvc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TplSvc_GetTplById_FullMethodName = "/tplsvc.TplSvc/GetTplById"
	TplSvc_CreateTpl_FullMethodName  = "/tplsvc.TplSvc/CreateTpl"
	TplSvc_UpdateTpl_FullMethodName  = "/tplsvc.TplSvc/UpdateTpl"
	TplSvc_DeleteTpl_FullMethodName  = "/tplsvc.TplSvc/DeleteTpl"
	TplSvc_TplList_FullMethodName    = "/tplsvc.TplSvc/TplList"
)

// TplSvcClient is the client API for TplSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TplSvcClient interface {
	GetTplById(ctx context.Context, in *GetTplByIdReq, opts ...grpc.CallOption) (*GetTplByIdResp, error)
	CreateTpl(ctx context.Context, in *CreateTplReq, opts ...grpc.CallOption) (*CreateTplResp, error)
	UpdateTpl(ctx context.Context, in *UpdateTplReq, opts ...grpc.CallOption) (*UpdateTplResp, error)
	DeleteTpl(ctx context.Context, in *DeleteTplReq, opts ...grpc.CallOption) (*DeletTplResp, error)
	TplList(ctx context.Context, in *TplListReq, opts ...grpc.CallOption) (*TplListResp, error)
}

type tplSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewTplSvcClient(cc grpc.ClientConnInterface) TplSvcClient {
	return &tplSvcClient{cc}
}

func (c *tplSvcClient) GetTplById(ctx context.Context, in *GetTplByIdReq, opts ...grpc.CallOption) (*GetTplByIdResp, error) {
	out := new(GetTplByIdResp)
	err := c.cc.Invoke(ctx, TplSvc_GetTplById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplSvcClient) CreateTpl(ctx context.Context, in *CreateTplReq, opts ...grpc.CallOption) (*CreateTplResp, error) {
	out := new(CreateTplResp)
	err := c.cc.Invoke(ctx, TplSvc_CreateTpl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplSvcClient) UpdateTpl(ctx context.Context, in *UpdateTplReq, opts ...grpc.CallOption) (*UpdateTplResp, error) {
	out := new(UpdateTplResp)
	err := c.cc.Invoke(ctx, TplSvc_UpdateTpl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplSvcClient) DeleteTpl(ctx context.Context, in *DeleteTplReq, opts ...grpc.CallOption) (*DeletTplResp, error) {
	out := new(DeletTplResp)
	err := c.cc.Invoke(ctx, TplSvc_DeleteTpl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplSvcClient) TplList(ctx context.Context, in *TplListReq, opts ...grpc.CallOption) (*TplListResp, error) {
	out := new(TplListResp)
	err := c.cc.Invoke(ctx, TplSvc_TplList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TplSvcServer is the server API for TplSvc service.
// All implementations must embed UnimplementedTplSvcServer
// for forward compatibility
type TplSvcServer interface {
	GetTplById(context.Context, *GetTplByIdReq) (*GetTplByIdResp, error)
	CreateTpl(context.Context, *CreateTplReq) (*CreateTplResp, error)
	UpdateTpl(context.Context, *UpdateTplReq) (*UpdateTplResp, error)
	DeleteTpl(context.Context, *DeleteTplReq) (*DeletTplResp, error)
	TplList(context.Context, *TplListReq) (*TplListResp, error)
	mustEmbedUnimplementedTplSvcServer()
}

// UnimplementedTplSvcServer must be embedded to have forward compatible implementations.
type UnimplementedTplSvcServer struct {
}

func (UnimplementedTplSvcServer) GetTplById(context.Context, *GetTplByIdReq) (*GetTplByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTplById not implemented")
}
func (UnimplementedTplSvcServer) CreateTpl(context.Context, *CreateTplReq) (*CreateTplResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTpl not implemented")
}
func (UnimplementedTplSvcServer) UpdateTpl(context.Context, *UpdateTplReq) (*UpdateTplResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTpl not implemented")
}
func (UnimplementedTplSvcServer) DeleteTpl(context.Context, *DeleteTplReq) (*DeletTplResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTpl not implemented")
}
func (UnimplementedTplSvcServer) TplList(context.Context, *TplListReq) (*TplListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TplList not implemented")
}
func (UnimplementedTplSvcServer) mustEmbedUnimplementedTplSvcServer() {}

// UnsafeTplSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TplSvcServer will
// result in compilation errors.
type UnsafeTplSvcServer interface {
	mustEmbedUnimplementedTplSvcServer()
}

func RegisterTplSvcServer(s grpc.ServiceRegistrar, srv TplSvcServer) {
	s.RegisterService(&TplSvc_ServiceDesc, srv)
}

func _TplSvc_GetTplById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTplByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TplSvcServer).GetTplById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TplSvc_GetTplById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TplSvcServer).GetTplById(ctx, req.(*GetTplByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TplSvc_CreateTpl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTplReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TplSvcServer).CreateTpl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TplSvc_CreateTpl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TplSvcServer).CreateTpl(ctx, req.(*CreateTplReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TplSvc_UpdateTpl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTplReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TplSvcServer).UpdateTpl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TplSvc_UpdateTpl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TplSvcServer).UpdateTpl(ctx, req.(*UpdateTplReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TplSvc_DeleteTpl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTplReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TplSvcServer).DeleteTpl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TplSvc_DeleteTpl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TplSvcServer).DeleteTpl(ctx, req.(*DeleteTplReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TplSvc_TplList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TplListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TplSvcServer).TplList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TplSvc_TplList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TplSvcServer).TplList(ctx, req.(*TplListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TplSvc_ServiceDesc is the grpc.ServiceDesc for TplSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TplSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tplsvc.TplSvc",
	HandlerType: (*TplSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTplById",
			Handler:    _TplSvc_GetTplById_Handler,
		},
		{
			MethodName: "CreateTpl",
			Handler:    _TplSvc_CreateTpl_Handler,
		},
		{
			MethodName: "UpdateTpl",
			Handler:    _TplSvc_UpdateTpl_Handler,
		},
		{
			MethodName: "DeleteTpl",
			Handler:    _TplSvc_DeleteTpl_Handler,
		},
		{
			MethodName: "TplList",
			Handler:    _TplSvc_TplList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tplsvc.proto",
}

const (
	TplSignSvc_GetSigById_FullMethodName = "/tplsvc.TplSignSvc/GetSigById"
	TplSignSvc_CreateSig_FullMethodName  = "/tplsvc.TplSignSvc/CreateSig"
	TplSignSvc_UpdateSig_FullMethodName  = "/tplsvc.TplSignSvc/UpdateSig"
	TplSignSvc_DeleteSig_FullMethodName  = "/tplsvc.TplSignSvc/DeleteSig"
	TplSignSvc_SigList_FullMethodName    = "/tplsvc.TplSignSvc/SigList"
)

// TplSignSvcClient is the client API for TplSignSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TplSignSvcClient interface {
	GetSigById(ctx context.Context, in *GetSigByIdReq, opts ...grpc.CallOption) (*GetSigByIdResp, error)
	CreateSig(ctx context.Context, in *CreateSigReq, opts ...grpc.CallOption) (*CreateSigResp, error)
	UpdateSig(ctx context.Context, in *UpdateSigReq, opts ...grpc.CallOption) (*UpdateSigResp, error)
	DeleteSig(ctx context.Context, in *DeleteSigReq, opts ...grpc.CallOption) (*DeleteSigResp, error)
	SigList(ctx context.Context, in *SigListReq, opts ...grpc.CallOption) (*SigListResp, error)
}

type tplSignSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewTplSignSvcClient(cc grpc.ClientConnInterface) TplSignSvcClient {
	return &tplSignSvcClient{cc}
}

func (c *tplSignSvcClient) GetSigById(ctx context.Context, in *GetSigByIdReq, opts ...grpc.CallOption) (*GetSigByIdResp, error) {
	out := new(GetSigByIdResp)
	err := c.cc.Invoke(ctx, TplSignSvc_GetSigById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplSignSvcClient) CreateSig(ctx context.Context, in *CreateSigReq, opts ...grpc.CallOption) (*CreateSigResp, error) {
	out := new(CreateSigResp)
	err := c.cc.Invoke(ctx, TplSignSvc_CreateSig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplSignSvcClient) UpdateSig(ctx context.Context, in *UpdateSigReq, opts ...grpc.CallOption) (*UpdateSigResp, error) {
	out := new(UpdateSigResp)
	err := c.cc.Invoke(ctx, TplSignSvc_UpdateSig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplSignSvcClient) DeleteSig(ctx context.Context, in *DeleteSigReq, opts ...grpc.CallOption) (*DeleteSigResp, error) {
	out := new(DeleteSigResp)
	err := c.cc.Invoke(ctx, TplSignSvc_DeleteSig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplSignSvcClient) SigList(ctx context.Context, in *SigListReq, opts ...grpc.CallOption) (*SigListResp, error) {
	out := new(SigListResp)
	err := c.cc.Invoke(ctx, TplSignSvc_SigList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TplSignSvcServer is the server API for TplSignSvc service.
// All implementations must embed UnimplementedTplSignSvcServer
// for forward compatibility
type TplSignSvcServer interface {
	GetSigById(context.Context, *GetSigByIdReq) (*GetSigByIdResp, error)
	CreateSig(context.Context, *CreateSigReq) (*CreateSigResp, error)
	UpdateSig(context.Context, *UpdateSigReq) (*UpdateSigResp, error)
	DeleteSig(context.Context, *DeleteSigReq) (*DeleteSigResp, error)
	SigList(context.Context, *SigListReq) (*SigListResp, error)
	mustEmbedUnimplementedTplSignSvcServer()
}

// UnimplementedTplSignSvcServer must be embedded to have forward compatible implementations.
type UnimplementedTplSignSvcServer struct {
}

func (UnimplementedTplSignSvcServer) GetSigById(context.Context, *GetSigByIdReq) (*GetSigByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSigById not implemented")
}
func (UnimplementedTplSignSvcServer) CreateSig(context.Context, *CreateSigReq) (*CreateSigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSig not implemented")
}
func (UnimplementedTplSignSvcServer) UpdateSig(context.Context, *UpdateSigReq) (*UpdateSigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSig not implemented")
}
func (UnimplementedTplSignSvcServer) DeleteSig(context.Context, *DeleteSigReq) (*DeleteSigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSig not implemented")
}
func (UnimplementedTplSignSvcServer) SigList(context.Context, *SigListReq) (*SigListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SigList not implemented")
}
func (UnimplementedTplSignSvcServer) mustEmbedUnimplementedTplSignSvcServer() {}

// UnsafeTplSignSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TplSignSvcServer will
// result in compilation errors.
type UnsafeTplSignSvcServer interface {
	mustEmbedUnimplementedTplSignSvcServer()
}

func RegisterTplSignSvcServer(s grpc.ServiceRegistrar, srv TplSignSvcServer) {
	s.RegisterService(&TplSignSvc_ServiceDesc, srv)
}

func _TplSignSvc_GetSigById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSigByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TplSignSvcServer).GetSigById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TplSignSvc_GetSigById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TplSignSvcServer).GetSigById(ctx, req.(*GetSigByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TplSignSvc_CreateSig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TplSignSvcServer).CreateSig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TplSignSvc_CreateSig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TplSignSvcServer).CreateSig(ctx, req.(*CreateSigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TplSignSvc_UpdateSig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TplSignSvcServer).UpdateSig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TplSignSvc_UpdateSig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TplSignSvcServer).UpdateSig(ctx, req.(*UpdateSigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TplSignSvc_DeleteSig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TplSignSvcServer).DeleteSig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TplSignSvc_DeleteSig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TplSignSvcServer).DeleteSig(ctx, req.(*DeleteSigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TplSignSvc_SigList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SigListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TplSignSvcServer).SigList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TplSignSvc_SigList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TplSignSvcServer).SigList(ctx, req.(*SigListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TplSignSvc_ServiceDesc is the grpc.ServiceDesc for TplSignSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TplSignSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tplsvc.TplSignSvc",
	HandlerType: (*TplSignSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSigById",
			Handler:    _TplSignSvc_GetSigById_Handler,
		},
		{
			MethodName: "CreateSig",
			Handler:    _TplSignSvc_CreateSig_Handler,
		},
		{
			MethodName: "UpdateSig",
			Handler:    _TplSignSvc_UpdateSig_Handler,
		},
		{
			MethodName: "DeleteSig",
			Handler:    _TplSignSvc_DeleteSig_Handler,
		},
		{
			MethodName: "SigList",
			Handler:    _TplSignSvc_SigList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tplsvc.proto",
}

const (
	TplSendSvc_CreateRec_FullMethodName  = "/tplsvc.TplSendSvc/CreateRec"
	TplSendSvc_GetRecById_FullMethodName = "/tplsvc.TplSendSvc/GetRecById"
)

// TplSendSvcClient is the client API for TplSendSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TplSendSvcClient interface {
	CreateRec(ctx context.Context, in *CreateRecReq, opts ...grpc.CallOption) (*CreateRecResp, error)
	GetRecById(ctx context.Context, in *GetRecByIdReq, opts ...grpc.CallOption) (*GetRecByIdResp, error)
}

type tplSendSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewTplSendSvcClient(cc grpc.ClientConnInterface) TplSendSvcClient {
	return &tplSendSvcClient{cc}
}

func (c *tplSendSvcClient) CreateRec(ctx context.Context, in *CreateRecReq, opts ...grpc.CallOption) (*CreateRecResp, error) {
	out := new(CreateRecResp)
	err := c.cc.Invoke(ctx, TplSendSvc_CreateRec_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tplSendSvcClient) GetRecById(ctx context.Context, in *GetRecByIdReq, opts ...grpc.CallOption) (*GetRecByIdResp, error) {
	out := new(GetRecByIdResp)
	err := c.cc.Invoke(ctx, TplSendSvc_GetRecById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TplSendSvcServer is the server API for TplSendSvc service.
// All implementations must embed UnimplementedTplSendSvcServer
// for forward compatibility
type TplSendSvcServer interface {
	CreateRec(context.Context, *CreateRecReq) (*CreateRecResp, error)
	GetRecById(context.Context, *GetRecByIdReq) (*GetRecByIdResp, error)
	mustEmbedUnimplementedTplSendSvcServer()
}

// UnimplementedTplSendSvcServer must be embedded to have forward compatible implementations.
type UnimplementedTplSendSvcServer struct {
}

func (UnimplementedTplSendSvcServer) CreateRec(context.Context, *CreateRecReq) (*CreateRecResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRec not implemented")
}
func (UnimplementedTplSendSvcServer) GetRecById(context.Context, *GetRecByIdReq) (*GetRecByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecById not implemented")
}
func (UnimplementedTplSendSvcServer) mustEmbedUnimplementedTplSendSvcServer() {}

// UnsafeTplSendSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TplSendSvcServer will
// result in compilation errors.
type UnsafeTplSendSvcServer interface {
	mustEmbedUnimplementedTplSendSvcServer()
}

func RegisterTplSendSvcServer(s grpc.ServiceRegistrar, srv TplSendSvcServer) {
	s.RegisterService(&TplSendSvc_ServiceDesc, srv)
}

func _TplSendSvc_CreateRec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TplSendSvcServer).CreateRec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TplSendSvc_CreateRec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TplSendSvcServer).CreateRec(ctx, req.(*CreateRecReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TplSendSvc_GetRecById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TplSendSvcServer).GetRecById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TplSendSvc_GetRecById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TplSendSvcServer).GetRecById(ctx, req.(*GetRecByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TplSendSvc_ServiceDesc is the grpc.ServiceDesc for TplSendSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TplSendSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tplsvc.TplSendSvc",
	HandlerType: (*TplSendSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRec",
			Handler:    _TplSendSvc_CreateRec_Handler,
		},
		{
			MethodName: "GetRecById",
			Handler:    _TplSendSvc_GetRecById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tplsvc.proto",
}

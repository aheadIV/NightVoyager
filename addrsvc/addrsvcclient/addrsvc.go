// Code generated by goctl. DO NOT EDIT.
// Source: addrsvc.proto

package addrsvcclient

import (
	"context"

	"github.com/aheadIV/NightVoyager/addrsvc/addrsvc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = addrsvc.Request
	Response = addrsvc.Response

	Addrsvc interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultAddrsvc struct {
		cli zrpc.Client
	}
)

func NewAddrsvc(cli zrpc.Client) Addrsvc {
	return &defaultAddrsvc{
		cli: cli,
	}
}

func (m *defaultAddrsvc) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := addrsvc.NewAddrsvcClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

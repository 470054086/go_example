package service

import (
	"context"
	"rpc_test.com/protocol"
	"rpc_test.com/transport"
)

type ServeFunc func(network string, addr string) error
type ServeTransportFunc func(tr transport.Transport)
type HandleRequestFunc func(ctx context.Context, request *protocol.Message, response *protocol.Message, tr transport.Transport)

// 服务端的中间件
type Wrapper interface {
	WrapServe(s *SGServer, serveFunc ServeFunc) ServeFunc
	WrapServeTransport(s *SGServer, transportFunc ServeTransportFunc) ServeTransportFunc
	WrapHandleRequest(s *SGServer, requestFunc HandleRequestFunc) HandleRequestFunc
}
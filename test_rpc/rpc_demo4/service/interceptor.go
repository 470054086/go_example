package service

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"rpc_test.com/protocol"
	registry "rpc_test.com/register"
	"rpc_test.com/transport"
	"sync/atomic"
	"syscall"
)

type DefaultServerWrapper struct {
}

func (w *DefaultServerWrapper) WrapServe(s *SGServer, serveFunc ServeFunc) ServeFunc {
	return func(network string, addr string) error {
		//注册shutdownHook
		go func(s *SGServer) {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, syscall.SIGTERM)
			sig := <-ch
			if sig.String() == "terminated" {
				for _, hook := range s.Option.ShutDownHooks {
					hook(s)
				}
				os.Exit(0)
			}
		}(s)
		// 保存所有的服务
		serviceInfo, _ := json.Marshal(s.Services())
		// 生成服务对象
		provider := registry.Provider{
			ProviderKey: network + "@" + addr,
			Network:     network,
			Addr:        addr,
			Meta:        map[string]string{"services": string(serviceInfo)},
		}
		r := s.Option.Registry
		rOpt := s.Option.RegisterOption
		// 注册到服务中间
		r.Register(rOpt, provider)
		log.Printf("registered provider %v for app %s", provider, rOpt)

		return serveFunc(network, addr)
	}
}

func (w *DefaultServerWrapper) WrapServeTransport(s *SGServer, transportFunc ServeTransportFunc) ServeTransportFunc {
	return transportFunc
}

func (w *DefaultServerWrapper) WrapHandleRequest(s *SGServer, requestFunc HandleRequestFunc) HandleRequestFunc {
	return func(ctx context.Context, request *protocol.Message, response *protocol.Message, tr transport.Transport) {
		atomic.AddInt64(&s.requestInProcess, 1)
		requestFunc(ctx, request, response, tr)
		atomic.AddInt64(&s.requestInProcess, -1)
	}
}
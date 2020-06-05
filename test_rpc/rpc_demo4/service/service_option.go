package service

import (
	"rpc_test.com/codec"
	"rpc_test.com/protocol"
	registry "rpc_test.com/register"
	"rpc_test.com/transport"
	"time"
)

type Option struct {
	AppKey         string
	Registry       registry.Registry
	RegisterOption registry.RegisterOption
	Wrappers       []Wrapper
	ShutDownWait   time.Duration
	ShutDownHooks  []ShutDownHook

	ProtocolType  protocol.ProtocolType
	SerializeType codec.SerializeType
	CompressType  protocol.CompressType
	TransportType transport.TransportType
}

var DefaultOption = Option{
	ShutDownWait:  time.Second * 12,
	ProtocolType:  protocol.Default,
	SerializeType: codec.MessagePack,
	CompressType:  protocol.CompressTypeNone,
	TransportType: transport.TCPTransport,
}

type ShutDownHook func(s *SGServer)
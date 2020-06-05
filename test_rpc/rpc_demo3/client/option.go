package client

import (
	"rpc_test.com/codec"
	"rpc_test.com/protocol"
	"rpc_test.com/transport"
	"time"
)

type Option struct {
	ProtocolType  protocol.ProtocolType
	SerializeType codec.SerializeType
	CompressType  protocol.CompressType
	TransportType transport.TransportType

	RequestTimeout time.Duration
}

var DefaultOption = Option{
	ProtocolType:  protocol.Default,
	SerializeType: codec.MessagePack,
	CompressType:  protocol.CompressTypeNone,
	TransportType: transport.TCPTransport,
}
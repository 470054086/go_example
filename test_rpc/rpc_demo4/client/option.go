package client

import (
	"rpc_test.com/codec"
	"rpc_test.com/protocol"
	registry "rpc_test.com/register"
	"rpc_test.com/selector"
	"rpc_test.com/transport"
	"time"
)

type Option struct {
	ProtocolType  protocol.ProtocolType
	SerializeType codec.SerializeType
	CompressType  protocol.CompressType
	TransportType transport.TransportType
	DialTimeout    time.Duration
	RequestTimeout time.Duration
}

var DefaultOption = Option{
	ProtocolType:  protocol.Default,
	SerializeType: codec.MessagePack,
	CompressType:  protocol.CompressTypeNone,
	TransportType: transport.TCPTransport,
}

type FailMode byte

const (
	FailFast FailMode = iota
	FailOver
	FailRetry
	FailSafe
)

type SGOption struct {
	AppKey       string
	RemoteAppkey string
	FailMode     FailMode
	Retries      int
	Registry     registry.Registry
	Selector     selector.Selector
	SelectOption selector.SelectOption
	Wrappers     []Wrapper

	Option

	Meta map[string]string
}

func AddWrapper(o *SGOption, w ...Wrapper) *SGOption {
	o.Wrappers = append(o.Wrappers, w...)
	return o
}

var DefaultSGOption = SGOption{
	AppKey:   "",
	FailMode: FailFast,
	Retries:  0,
	Selector: selector.NewRandomSelector(),

	Option: DefaultOption,

	Meta: make(map[string]string),
}
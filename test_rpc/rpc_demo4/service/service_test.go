package service

import (
	"rpc_test.com/codec"
	"testing"
)



func TestSimpleServer_Register(t *testing.T) {
	var op Option
	op.SerializeType = codec.MessagePack
	server := NewSimpleServer(op)
	var r TestReflect
	server.Register(r, nil)
}

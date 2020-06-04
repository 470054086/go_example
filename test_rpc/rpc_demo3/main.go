package main

import (
	"rpc_test.com/codec"
	"rpc_test.com/service"
)

type TestReflect struct {
}

func (t *TestReflect) AddTest() {
}
func (t *TestReflect) AddDecl() {

}
func (t *TestReflect) AddTest2() {

}
func (t *TestReflect) addTest2() {

}


func main()  {
	var op service.Option
	op.SerializeType = codec.MessagePack
	server := service.NewSimpleServer(op)
	var r TestReflect
	server.Register(r, nil)
}

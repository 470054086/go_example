package main

import (
	"context"
	"log"
	"math/rand"
	"rpc_test.com/client"
	"rpc_test.com/codec"
	registry "rpc_test.com/register"
	"rpc_test.com/register/memory"
	"rpc_test.com/server"
	"rpc_test.com/service"
	"strconv"
	"time"
)

const callTimes = 1

var s service.RPCServer

func main() {
	StartServer()
	time.Sleep(1e9)
	start := time.Now()
	for i := 0; i < callTimes; i++ {
		MakeCall(codec.MessagePack)
	}
	cost := time.Now().Sub(start)
	log.Printf("cost:%s", cost)
	//StopServer()
}

func StopServer() {
	s.Close()
}

var Registry = memory.NewInMemoryRegistry()

func StartServer() {
	go func() {
		serverOpt := service.DefaultOption
		serverOpt.RegisterOption.AppKey = "my-app"
		serverOpt.Registry = Registry
		s := service.NewRPCServer(serverOpt)
		// 注册服务
		err := s.Register(server.Arith{}, make(map[string]string))
		if err != nil {
			log.Println("err!!!" + err.Error())
		}
		port := 8880
		s.Serve("tcp", ":"+strconv.Itoa(port))
	}()
	go func() {
		serverOpt := service.DefaultOption
		serverOpt.RegisterOption.AppKey = "my-app"
		serverOpt.Registry = Registry
		s := service.NewRPCServer(serverOpt)
		// 注册服务
		err := s.Register(server.Arith{}, make(map[string]string))
		if err != nil {
			log.Println("err!!!" + err.Error())
		}
		port := 8881
		s.Serve("tcp", ":"+strconv.Itoa(port))
	}()

	go func() {
		serverOpt := service.DefaultOption
		serverOpt.RegisterOption.AppKey = "my-app"
		serverOpt.Registry = Registry
		s := service.NewRPCServer(serverOpt)
		// 注册服务
		err := s.Register(server.Arith{}, make(map[string]string))
		if err != nil {
			log.Println("err!!!" + err.Error())
		}
		port := 8882
		s.Serve("tcp", ":"+strconv.Itoa(port))
	}()
}

func MakeCall(t codec.SerializeType) {
	op := &client.DefaultSGOption
	op.AppKey = "my-app"
	op.SerializeType = t
	op.RequestTimeout = time.Millisecond * 100
	op.DialTimeout = time.Millisecond * 100
	op.FailMode = client.FailRetry
	op.Retries = 3
	r := memory.G_Registry
	r.Register(registry.RegisterOption{})
	op.Registry = r
	c := client.NewSGClient(*op)

	args := server.Args{A: rand.Intn(200), B: rand.Intn(100)}
	reply := &server.Reply{}
	ctx := context.Background()
	err := c.Call(ctx, "Arith.Add", args, reply)
	if err != nil {
		log.Println("err!!!" + err.Error())
	} else if reply.C != args.A+args.B {
		log.Printf("%d + %d != %d", args.A, args.B, reply.C)
	}

}



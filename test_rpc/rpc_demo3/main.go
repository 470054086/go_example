package main

import (
	"context"
	"errors"
	"fmt"
	"rpc_test.com/client"
	"rpc_test.com/service"
	"time"
)

func main() {
	s := service.NewSimpleServer(service.DefaultOption)
	err := s.Register(Arith{}, make(map[string]string))
	if err != nil {
		panic(err)
	}
	// 启动服务
	go func() {
		err = s.Server("tcp", ":8888")
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(1e9)

	c, err := client.NewRpcClient("tcp", ":8888", client.DefaultOption)
	args := Args{A: 200, B: 100}
	reply := &Reply{}
	err = c.Call(context.TODO(), "Arith.Add", args, reply)
	fmt.Println(reply.C)



	//wg.Add(100)
	//for i := 0; i < 100; i++ {
	//	go func() {
	//		c, err := client.NewRpcClient("tcp", ":8888", client.DefaultOption)
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//		args := Args{A: rand.Intn(200), B: rand.Intn(100)}
	//		reply := &Reply{}
	//		err = c.Call(context.TODO(), "Arith.Add", args, reply)
	//		if err != nil {
	//			panic(err)
	//		}
	//		if reply.C != args.A+args.B {
	//			log.Fatal(reply.C)
	//		} else {
	//			fmt.Println(reply.C)
	//		}
	//
	//		err = c.Call(context.TODO(), "Arith.Minus", args, reply)
	//		if err != nil {
	//			panic(err)
	//		}
	//		if reply.C != args.A-args.B {
	//			log.Fatal(reply.C)
	//		} else {
	//			fmt.Println(reply.C)
	//		}
	//
	//		err = c.Call(context.TODO(), "Arith.Mul", args, reply)
	//		if err != nil {
	//			panic(err)
	//		}
	//		if reply.C != args.A*args.B {
	//			log.Fatal(reply.C)
	//		} else {
	//			fmt.Println(reply.C)
	//		}
	//
	//		err = c.Call(context.TODO(), "Arith.Divide", args, reply)
	//		if err != nil {
	//			log.Println(err)
	//
	//		}
	//		if err != nil && err.Error() == "divided by 0" {
	//			log.Println(err)
	//		} else if reply.C != args.A/args.B {
	//			log.Fatal(reply.C)
	//		} else {
	//			fmt.Println(reply.C)
	//		}
	//		wg.Done()
	//	}()
	//}
	//wg.Wait()
}

type Arith struct{}

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

//arg可以是指针类型，也可以是指针类型
func (a Arith) Add(ctx context.Context, arg *Args, reply *Reply) error {
	reply.C = arg.A + arg.B
	return nil
}

func (a Arith) Minus(ctx context.Context, arg Args, reply *Reply) error {
	reply.C = arg.A - arg.B
	return nil
}

func (a Arith) Mul(ctx context.Context, arg Args, reply *Reply) error {
	reply.C = arg.A * arg.B
	return nil
}

func (a Arith) Divide(ctx context.Context, arg *Args, reply *Reply) error {
	if arg.B == 0 {
		return errors.New("divided by 0")
	}
	reply.C = arg.A / arg.B
	return nil
}

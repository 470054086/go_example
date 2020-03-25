package main

import (
	"gocache"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)
func main()  {
	gocache.NewCache(func(options *gocache.Options) {
		options.ExpireKeyIntervalDuration = time.Second * 10
	})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigs:

	}


	//启动一个tcp/ip服务
	listen, err := net.Listen("tcp", "8082")
	if err != nil {
		panic(err)
	}
	for {
		// 监听服务端的请求
		accept, _ := listen.Accept()
		//开启协程 返回服务端请求
		// 不开启协程的话 会造成阻塞 只能一个请求一个请求的处理
		go func(conn net.Conn) {
			var buf []byte
			buf = make([]byte,2048)
			_, _= conn.Read(buf)
			sbuf := string(buf)
			if sbuf =="get" {
				// get key
			} else if sbuf == "put" {
				//xxxx
			}
		}(accept)
	}



}
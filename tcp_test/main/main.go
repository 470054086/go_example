package main

import (
	"fmt"
	"io"
	"net"
	"tcp_test/common"
	"time"
)

func main() {
	listen, _ := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   []byte("127.0.0.1"),
		Port: 8090,
		Zone: "",
	})
	listen, err := net.ListenTCP("tcp", "127.0.0.1:8090")
	listen.SetLinger()
	listen.SetReadBuffer(1024)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("服务端链接成功")
	for {
		fmt.Println("开始等待数据")
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("客户端链接成功%s", accept.RemoteAddr().String())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("开始处理数据")
		go handlerRequest(accept)
	}
}
func handlerRequest(n net.Conn) {
	i := 0
	util := common.SocketUtil{Coon: n}
	for {
		read, err := util.ReadPkg()
		fmt.Println(string(read))
		if err != nil || err == io.EOF {
			fmt.Println("循环结束")
			break
		}
		i++
		fmt.Printf("读取次数%d", i)
		fmt.Println()
	}
	fmt.Println("读取完成")
}

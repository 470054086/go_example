package main

import (
	"fmt"
	"log"
	"net"
	"tcp_test/common"
)

func main()  {
	// 启动tcp ip
	listen, err := net.Listen("tcp", ":8899")
	if err!= nil {
		panic(err)
	}
	defer listen.Close()
	log.Println("listen ok")
	for {
		if conn, err := listen.Accept(); err != nil {
			log.Println("accept error:", err)
			break
		}else{
			go handler(conn)
		}
	}
}

func handler(c net.Conn)  {
	defer c.Close()
	fd := common.SocketUtil{Coon:c}
	for {
		data, err := fd.ReadPkg() //读取数据
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Println(string(data))
	}
}

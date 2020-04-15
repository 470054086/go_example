package main

import (
	"log"
	"net"
	"time"
)

func main()  {
	// 启动tcp ip
	listen, err := net.Listen("tcp", ":8080")
	if err!= nil {
		panic(err)
	}
	defer listen.Close()
	log.Println("listen ok")
	i:=0;
	for {
		if _, err := listen.Accept(); err != nil {
			log.Println("accept error:", err)
			break
		}
		i++
		go func(i int) {
			time.Sleep(time.Second*10)
			i++
			log.Printf("%d: accept a new connection\n", i)
		}(i)

	}
}

func handler(c net.Conn )  {
	defer c.Close()
	for  {
		// read from
		// write form
		c.Write([]byte("xiaobai"))
	}
}

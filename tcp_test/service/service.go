package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"tcp_test/chatroom"
	"tcp_test/common"
	"tcp_test/constant"
)

var chat *chatroom.ChatRoom

func main() {
	// 启动tcp ip
	listen, err := net.Listen("tcp", ":9988")
	if err != nil {
		panic(err)
	}
	chat = chatroom.NewChat()
	defer listen.Close()
	for {
		if conn, err := listen.Accept(); err != nil {
			log.Println("accept error:", err)
			break
		} else {
			go handler(conn)
		}
	}
}

func handler(c net.Conn) {
	fd := common.NewSocketUtil(c)
	for {
		data, err := fd.PkgReader() //读取数据
		if err != nil {
			fmt.Println(err)
			break
		}
		someHandler(c, data)
	}
}
func someHandler(conn net.Conn, data []byte) error {
	// 进行数据解析
	var message constant.SendMessage
	err := json.Unmarshal(data, &message)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// 进行字符串解析操作
	switch message.MType {
	case constant.Connection:
		// 如果是连接的话
		chat.Add(conn, message)
	case constant.Send:
		// 如果是单播的话
		chat.Send(conn, message)

	}
	return nil
}

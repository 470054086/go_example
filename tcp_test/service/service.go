package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"tcp_test/chatroom"
	"tcp_test/constant"
	"tcp_test/protocol"
)

var chat *chatroom.ChatRoom

func main() {
	// 启动tcp ip
	listen, err := net.Listen("tcp", ":9999")
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
			// 启动一个协程来处理事务
			go handler(conn)
		}
	}
}

func handler(c net.Conn) {
	fd := protocol.NewSocketUtil(c)
	for {
		var data []byte
		_, err := fd.Read(data) //读取数据
		data = fd.GetBytes()
		if err != nil {
			fmt.Println(err)
			break
		}
		someHandler(fd, data)
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
	case constant.Leave:
		// 如果是离开当前直播间的话
		chat.Leave(conn,message)
	case constant.Close:
		// 如果为关闭连接的话
		chat.Leave(conn,message)
		conn.Close()
	}
	return nil
}

package chatroom

import (
	"fmt"
	"net"
	"sync"
	"tcp_test/constant"
)

// 定义聊天室的结构体
type chatRoom struct {
	Client     []net.Conn
	UserClient map[int64]net.Conn
	UserConn   map[*net.Conn]int64
	sy         sync.RWMutex
}

// 生成对象
func NewChat() *chatRoom {
	return &chatRoom{
		UserClient: make(map[int64]net.Conn),
		UserConn:   make(map[*net.Conn]int64),
		sy:         sync.RWMutex{},
	}
}

// 加入
func (c *chatRoom) Add(conn net.Conn, message constant.SendMessage) error {
	c.sy.Lock()
	if _, ok := c.UserClient[message.MUser]; ok {
		return constant.JoinChatExists
	}
	c.Client = append(c.Client, conn)
	c.UserClient[message.MUser] = conn
	c.UserConn[&conn] = message.MUser
	c.sy.Unlock()
	return nil
}

// 发送
func (c *chatRoom) Send(conn net.Conn, message constant.SendMessage) error {
	if message.SType == constant.SingleSendType {
		return c.singleSend(conn, message)
	} else if message.SType == constant.BroadcastSendType {
		return c.broadcastSend(conn, message)
	}
	return constant.SendTypeError
}

// 离开
func (c *chatRoom) Leave(conn net.Conn, message constant.SendMessage) (int, error) {
	index := 0
	for k, val := range c.Client {
		if val == conn {
			index = k
			break
		}
	}
	if index == 0 {
		return 0, constant.DeleteNotExists
	}
	// 删除数据
	c.Client = append(c.Client[:index], c.Client[(index+1):]...)
	delete(c.UserConn, &conn)
	delete(c.UserClient, message.MUser)
	return index, nil
}

// 单播
func (c *chatRoom) singleSend(conn net.Conn, message constant.SendMessage) error {
	if message.Receiver == 0 {
		return constant.SingleChatRevicerError
	}
	co, ok := c.UserClient[message.Receiver]
	if !ok {
		return constant.SingleChatRevicerAcceptError
	}
	fmt.Printf("发送给用户%d的信息为%s", message.Receiver, message.MData)
	_, err := co.Write(message.MData)
	if err != nil {
		return err
	}
	return nil
}

func (c *chatRoom) broadcastSend(conn net.Conn, message constant.SendMessage) error {
	for _, co := range c.Client {
		go func(co net.Conn) {
			fmt.Printf("发送给用户%d的信息为%s", c.UserConn[&co], message.MData)
			co.Write(message.MData)
		}(co)
	}
	return nil
}

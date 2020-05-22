package chatroom

import (
	"fmt"
	"net"
	"sync"
	"tcp_test/constant"
)

//{"mtype":1,"s_type":2,"m_user":1,"m_data":"","receiver":0}

// 定义聊天室的结构体
type ChatRoom struct {
	Client     []*net.Conn
	UserClient map[int64]*net.Conn
	UserConn   map[*net.Conn]int64
	sy         sync.RWMutex
}

// 生成对象
func NewChat() *ChatRoom {
	return &ChatRoom{
		UserClient: make(map[int64]*net.Conn),
		UserConn:   make(map[*net.Conn]int64),
		sy:         sync.RWMutex{},
	}
}

// 加入
func (c *ChatRoom) Add(conn net.Conn, message constant.SendMessage) error {
	c.sy.Lock()
	if _, ok := c.UserClient[message.MUser]; ok {
		return constant.JoinChatExists
	}
	c.Client = append(c.Client, &conn)
	c.UserClient[message.MUser] = &conn
	c.UserConn[&conn] = message.MUser
	c.sy.Unlock()
	fmt.Println(c.Client)
	fmt.Println(c.UserClient)
	fmt.Println(c.UserConn)
	return nil
}

// 发送
func (c *ChatRoom) Send(conn net.Conn, message constant.SendMessage) error {
	if message.SType == constant.SingleSendType {
		return c.singleSend(message)
	} else if message.SType == constant.BroadcastSendType {
		return c.broadcastSend(message)
	}
	return constant.SendTypeError
}

// 离开
func (c *ChatRoom) Leave(conn net.Conn, message constant.SendMessage) (int, error) {
	index := 0
	for k, val := range c.Client {
		if val == &conn {
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
func (c *ChatRoom) singleSend(message constant.SendMessage) error {
	if message.Receiver == 0 {
		return constant.SingleChatRevicerError
	}
	co, ok := c.UserClient[message.Receiver]
	if !ok {
		return constant.SingleChatRevicerAcceptError
	}
	fmt.Printf("发送给用户%d的信息为%s", message.Receiver, message.MData)
	_, err := (*co).Write([]byte(message.MData))
	if err != nil {
		return err
	}
	return nil
}

// 广播
func (c *ChatRoom) broadcastSend(message constant.SendMessage) error {
	for _, co := range c.Client {
		go func(co *net.Conn) {
			fmt.Printf("发送给用户%d的信息为%s", c.UserConn[co], message.MData)

			(*co).Write([]byte(message.MData))
		}(co)
	}
	return nil
}

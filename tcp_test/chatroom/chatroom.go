package chatroom

import (
	"net"
	"sync"
)
// 定义聊天室的结构体
type chatRoom struct {
	Client     []net.Conn
	UserClient map[int64]net.Conn
	sy         sync.RWMutex
}

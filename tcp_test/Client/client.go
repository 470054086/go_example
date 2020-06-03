package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"tcp_test/protocol"
	"time"
)

/*
// 加入操作
{"mtype":1,"s_type":2,"m_user":1,"m_data":"","receiver":0}
{"mtype":1,"s_type":2,"m_user":2,"m_data":"","receiver":0}
{"mtype":1,"s_type":2,"m_user":3,"m_data":"","receiver":0}

// 单播操作
{"mtype":2,"s_type":1,"m_user":1,"m_data":"xiaobaijun","receiver":2}

// 多播操作
{"mtype":2,"s_type":2,"m_user":2,"m_data":"xiaozhong","receiver":0}
*/

func main() {
	conn, err := net.DialTimeout("tcp", ":9999", time.Second*2)
	if err != nil {
		fmt.Println(err)
		return
	}
	//  关闭的channel
	closeChan := make(chan bool)
	defer conn.Close() // 关闭连接 在函数进行return的时候
	clntFd := protocol.NewSocketUtil(conn)
	// 读取来自服务器的回传
	go func(conn *protocol.SocketUtil) {
		for {
			var pkgReader []byte
			body, err2 := conn.Read(pkgReader)
			if body == 0 && err2 != nil {
				closeChan <- true
				conn.Close()
				break
			}
			// 未收到消息 则退出循环
			pkgReader = conn.GetBytes()
			if err2 != nil {
				fmt.Println(pkgReader)
			}
			fmt.Println("收到的信息为", string(pkgReader))
		}
	}(clntFd)
	// 写操作
	go func(conn *protocol.SocketUtil) {
		// 读取键盘的输入
		reader := bufio.NewReader(os.Stdin)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}
			// 对数据 进行解码
			//var message constant.SendMessage
			//err = json.Unmarshal([]byte(readString), &message)
			n, err := conn.Write([]byte(readString))
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Send %d byte data : %s", n, readString)
		}
	}(clntFd)
	// 信号量事件
	ch := make(chan os.Signal,1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	// 如果需要关闭的话 则关闭当前链接
	for {
		select {
		case <-closeChan:
			fmt.Println("我是被动关闭的")
			return
		case <-ch:
			fmt.Println("我是手动结束的")
			return

		}
	}

}

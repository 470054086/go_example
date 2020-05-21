package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"tcp_test/common"
)

/*
// 加入操作
{"mtype":1,"s_type":2,"m_user":1,"m_data":"","receiver":0}
{"mtype":1,"s_type":2,"m_user":2,"m_data":"","receiver":0}
{"mtype":1,"s_type":2,"m_user":3,"m_data":"","receiver":0}

// 单播操作
{"mtype":2,"s_type":1,"m_user":1,"m_data":"xiaobaijun","receiver":3}

// 多播操作
{"mtype":2,"s_type":2,"m_user":2,"m_data":"xiaozhong","receiver":0}
*/

func main() {
	conn, err := net.Dial("tcp", ":9988")
	if err != nil {
		fmt.Println(err)
		return
	}
	clntFd := common.NewSocketUtil(conn)
	// 读取键盘的输入
	reader := bufio.NewReader(os.Stdin)

	// 读取来自服务器的回传
	go func(conn *common.SocketUtil) {
		for {
			pkgReader, err2 := conn.PkgReader()
			if err2 != nil {
				fmt.Println(pkgReader)
			}
			fmt.Println("收到的信息为", string(pkgReader))
		}
	}(clntFd)

	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		// 对数据 进行解码
		//var message constant.SendMessage
		//err = json.Unmarshal([]byte(readString), &message)
		n, err := clntFd.PkgWrite([]byte(readString))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Send %d byte data : %s", n, readString)
	}

}

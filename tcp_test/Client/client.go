package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"tcp_test/common"
)

func main() {
	conn, err := net.Dial("tcp", ":8899")
	if err != nil {
		fmt.Println(err)
		return
	}
	clntFd := common.NewSocketUtil(conn)
	// 读取键盘的输入
	reader := bufio.NewReader(os.Stdin)
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

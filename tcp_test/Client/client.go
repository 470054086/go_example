package main

import (
	"encoding/json"
	"fmt"
	"net"
	"tcp_test/common"
	"tcp_test/constant"
)

func main() {
	conn, err := net.Dial("tcp", ":8899")
	if err != nil {
		fmt.Println(err)
		return
	}
	clntFd := common.NewSocketUtil(conn)
	p := map[string]interface{}{
		"name":    "xiaobai",
		"age":     21,
		"company": "intely",
	}
	m := constant.SendMessage{
		MType: 1,
		MUser: 1,
		MData: p,
	}
	marshal, _ := json.Marshal(m)
	n, err := clntFd.PkgWrite(marshal)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Send %d byte data : %s", n, marshal)

}

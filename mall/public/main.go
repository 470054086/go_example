package main

import (
	"flag"
	"log"
	"mall/bootstrap"
	"strconv"
)

var port int

//go env -w GOPROXY=https://goproxy.cn,direct
// 使用gin当做http请求

func main() {
	flag.IntVar(&port,"port",8080,"http port ")
	flag.Parse()
	app := bootstrap.NewApp()
	portString := ":"+strconv.Itoa(port)
	err := app.Run(portString)
	if err != nil {
		log.Fatal(err)
	}
}



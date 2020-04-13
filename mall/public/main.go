package main

import (
	"fmt"
	"github.com/pkg/errors"
	"mall/app/commom"
)

var port int

//go env -w GOPROXY=https://goproxy.cn,direct
// 使用gin当做http请求

func main() {
	//flag.IntVar(&port,"port",80,"http port ")
	//flag.Parse()
	//app := bootstrap.NewApp()
	//portString := ":"+strconv.Itoa(port)
	//err := app.Run(portString)
	//if err != nil {
	//	log.Fatal(err)
	//}
	defer func() {
		e := recover()
		err := errors.Cause(e.(error))
		switch err.(type) {
		case commom.ParamsError:
			fmt.Printf("%+v", err.(commom.ParamsError).Err)
		default:
			fmt.Println(333)
		}
	}()
	Build2()

}

func Build() error {
	panic(commom.NewParamsError(errors.New("error"), "error"))
}

func Build2() error {
	err := Build()
	return errors.Wrap(err, "build2 error")
}

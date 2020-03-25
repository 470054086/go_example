package main

import (
	"fmt"
	"simplehttp/simple"
)
type User struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

func main()  {
	e := simple.Default()
	e.Use(func(ctx *simple.Context) {
		fmt.Println("我就是丑陋的中间件")
	})
	e.Get("/user", func(ctx *simple.Context) {
		fmt.Println(11111)
	})
	e.Get("/users", func(ctx *simple.Context) {
		fmt.Println(22222)
	})
	e.Get("/user/post", func(ctx *simple.Context) {
		var r User
		ctx.BindJson(&r)
		fmt.Println(r)
	})
	
	err := e.Run(":8083")
	if err != nil {
		panic(err)
	}
}

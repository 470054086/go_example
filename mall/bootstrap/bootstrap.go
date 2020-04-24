package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"mall/Providers/SendMobile"
	"mall/dao"
	"mall/router"

)

type App struct {
	*gin.Engine
}
func NewApp() *App {
	e := &App{ gin.New()}
	// 启动路由
	e.startRouter()
	// 启动数据库
	e.startDb()
	// 启动发送短信服务
	e.startMobile()
	return e
}

// 设置路由相关
func (r *App) startRouter()  {
	// 启动路由
	router.Router(r.Engine)
}
// 启动数据库
func (r *App) startDb()  {
	dao.NewDao("root:root@tcp(localhost:3306)/go_test?charset=utf8&parseTime=True&loc=Local")
}
// 启动短信服务

func (r *App) startMobile()  {
	xue := SendMobile.NewSendXue(1000)
	ctx, cancelFunc := context.WithCancel(context.Background())
	xue.Star(ctx,cancelFunc)
}
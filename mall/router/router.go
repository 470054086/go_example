package router

import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"mall/app/controller"
	"mall/app/middleware"
	"mall/helpers"
)

// 设置路由相关函数
func Router(r *gin.Engine) {
	var hello = &controller.Hello{}
	var upload = &controller.Upload{}
	pprof.Register(r)

	r.Use(middleware.RecoverMiddle())
	r.Use(ginprom.PromMiddleware(nil))
	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))
	r.GET("/index", wapper(hello.Add))
	r.POST("/list", wapper(hello.Lists))
	r.POST("/upload", wapper(upload.Upload))
}

type fwapper func(c *gin.Context) *helpers.Response

/**
包装一下返回值
*/
func wapper(f fwapper) func(c *gin.Context) {
	return func(c *gin.Context) {
		response := f(c)
		c.JSON(response.Code, response)
	}
}

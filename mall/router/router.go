package router

import (
	"github.com/gin-gonic/gin"
	"mall/app/controller"
	"mall/app/middleware"
	"mall/helpers"
)
// 设置路由相关函数
func Router(r *gin.Engine)  {
	var hello = &controller.Hello{}
	var upload =&controller.Upload{}
	r.Use(middleware.RecoverMiddle())
	r.POST("/index",wapper(hello.Add))
	r.POST("/upload",wapper(upload.Upload))
}

type fwapper func(c *gin.Context) *helpers.Response
/**
	包装一下返回值
 */
func wapper(f fwapper) (func(c *gin.Context)) {
	return func(c *gin.Context) {
		response := f(c)
		c.JSON(response.Code,response)
	}
}
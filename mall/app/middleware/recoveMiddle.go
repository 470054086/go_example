package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"mall/app/commom"
	"net/http"
)

// 监听错误回调函数
func RecoverMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(c *gin.Context) {
			err := recover()
			if err != nil {
				switch err := errors.Cause(err.(error)).(type) {
				case commom.ParamsError:
					fmt.Printf("%+v\n", err.Err)
					data := make(map[string]interface{})
					c.Header("Content-Type", "applicaton/json")
					var code int
					if err.Code() == 0 {
						code = 403
					} else {
						code = err.Code()
					}
					c.JSON(http.StatusForbidden, gin.H{
						"code": code,
						"msg":  err.Error(),
						"data": data,
					})
				default:
					logrus.Error(err)
				}
			}
		}(c)
		c.Next()
	}

}

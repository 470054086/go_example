package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"mall/app/commom"
)

// 监听错误回调函数
func RecoverMiddle() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			err := recover()
			switch e:=errors.Cause(err.(error)).(type) {
			case commom.ParamsError:
				fmt.Printf("%+v",e)
			default:
				logrus.Error("unknow error")
			}
		}()
		context.Next()
	}

}

package middleware

import (
	"github.com/gin-gonic/gin"
)

// 监听错误回调函数
func RecoverMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(c *gin.Context) {
			err := recover()
			if err != nil {
				d:=  gin.H{
					"code": 200,
					"msg":  1111,
				}
				c.JSON(200,d)
				c.Abort()
			}

			//if err != nil {
			//	switch err:= errors.Cause(err.(error)).(type) {
			//	case commom.ParamsError:
			//		fmt.Printf("%+v\n",err.Err)
			//		data := make(map[string]interface{})
			//		c.Header("Content-Type","applicaton/json")
			//		c.JSON(200,gin.H{
			//			"code":200,
			//			"msg":err.Error(),
			//			"data": data,
			//		})
			//	default:
			//		logrus.Error("unknow error")
			//	}
			//}
		}(c)
		c.Next()
	}

}

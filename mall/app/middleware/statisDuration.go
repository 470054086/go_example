package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func StatisDuration() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		start := time.Now().UnixNano()
		c.Next()
		end := time.Now().UnixNano()
		diff := (end - start) / 1000 / 1000
		logrus.Infof("url为%s,请求的是时长为%dms", c.Request.Host+c.Request.URL.String(), diff)
	}
}

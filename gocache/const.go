package gocache

import (
	"github.com/pkg/errors"
	"time"
)

// 错误类型
var (
	KeyVaildError = errors.New("参数传递错误")
)

const (
	DefaultExpireKeyChanNumber = 100
)


// 用于传输过期时间的数据结构
type KeyExpireChan struct {
	key string
	expire time.Duration
}

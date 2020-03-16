package gocache

import (
	"github.com/pkg/errors"
	"time"
)

// 错误类型
var (
	KeyVaildError = errors.New("参数传递错误")
)

// 缓存的类型
type dataType int
const (
	keyValueType dataType=iota
	fifoArrayType
)

const (
	// 默认channel大小
	DefaultExpireKeyChanNumber = 100
	// 默认过期key扫描大小
	DefaultExpireKeyIntervalTime = 2 * time.Second
)



// key value 数据结构
type keyValue struct {
	value string
	// 过期时间
	expire int64
}

// 用于传输过期时间的数据结构
type KeyExpireChan struct {
	key string
	expire int64
	types dataType
}

// 构建传输的channel
func newKeyExpireChan(key string,expire int64,types dataType) *KeyExpireChan  {
	return &KeyExpireChan{
		key:    key,
		expire: expire,
		types:types,
	}
}
// 每个key的使用率
type CacheLru struct {
	key string // 缓存命中率
	hits int // 命中率
}
// todo 使用中间件统计即可
// 统计数量
type CacheTotal struct {
	 totalMemory int //总共的大小
	 hits int // 请求多少次
	 putNum int //put次数
	 getNum int //get次数

}

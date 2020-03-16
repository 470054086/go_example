package gocache

import "time"

type Op func(options *Options)
type Options struct {
	// key统一的过期时间 如果为0 则不过期
	KeyExpireTimeDuration time.Duration
	// 扫面过期key的间隔时间
	ExpireKeyIntervalDuration time.Duration
	// 过期时间chan的缓存大小
	ExpireKeyChanNumber  int
	// 统一key前缀名称
	PrefixKey string
}
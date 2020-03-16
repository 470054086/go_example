package gocache

import "time"

type keyValue struct {
	value string
	// 过期时间
	expire time.Duration
}

// 添加key-value
func (k *keyValue) put(key string,value interface{}) bool {
	
}

// 获取key-value
func (k *keyValue) get(key string)  string  {

}

// 删除key-value
func (k *keyValue) delete(key string) bool  {

}
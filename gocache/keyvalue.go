package gocache

import (
	"sync"
	"time"
)

type keyValueCache struct {
	key           map[string]*keyValue
	keyExpireChan chan *KeyExpireChan
	lock          sync.Mutex
}

func NewKeyValue(c chan *KeyExpireChan) keyValueCache {
	k := keyValueCache{
		key:           make(map[string]*keyValue),
		keyExpireChan: c,
		lock:          sync.Mutex{},
	}
	return k
}

// 添加key-value
func (k keyValueCache) Put(key string, value string) bool {
	k.lock.Lock()
	// 构建数据格式
	v := &keyValue{value: value, expire: 0}
	k.key[key] = v
	k.lock.Unlock()
	//构建过期key
	expireChan := newKeyExpireChan(key, 0, keyValueType)
	G.addExpireKeyChan(expireChan)
	return true
}

// 添加key-value 秒级别的
func (k keyValueCache) PutExpire(key string, value string, expire time.Duration) bool {
	k.lock.Lock()
	// 构建数据格式
	// 从当前时间开始的纳秒
	expireTime := time.Now().Unix() + int64(expire)
	v := &keyValue{value: value, expire: expireTime}
	// 写入过期时间
	k.key[key] = v
	//构建过期key
	expireChan := newKeyExpireChan(key, expireTime, keyValueType)
	G.addExpireKeyChan(expireChan)
	k.lock.Unlock()
	return true
}

// 获取key-value
func (k keyValueCache) Get(key string) (s string, b bool) {
	k.lock.Lock()
	if value, ok := k.key[key]; ok {
		// 永不过期
		if value.expire == 0 {
			k.lock.Unlock()
			return value.value, true
		} else {
			nowTime := time.Now().Unix()
			// 1. 如果还未过期
			if value.expire > nowTime {
				k.lock.Unlock()
				return value.value, true
			} else {
				// 如果过期 自动删除key  redis上面的key也需要删除
				k.Delete(key)
				G.delete(key)
				k.lock.Unlock()
				return "", false
			}
		}
	} else {
		k.lock.Unlock()
		return "", ok
	}
}

// 删除key-value
func (k keyValueCache) Delete(key string) bool {
	delete(k.key, key)
	return true
}

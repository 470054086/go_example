package gocache

import (
	"sync"
	"time"
)

//定义一个缓存的示例
type Gocache struct {
	// 存储过期key的一个map 这里应该使用有序集合更好
	allKey map[string] *KeyExpireChan
	//  key-value 数据结构
	keyValueCache
	// 数组结构队列
	fifo map[string]*fifoArray

	// LRu算法
	lru map[string] *CacheLru

	// total统计
	total map[string]*CacheTotal

	// 用于添加过期时间的channel
	keyExpireChan chan *KeyExpireChan
	// 配置文件
	option   *Options
	// 锁
	lock sync.RWMutex
}

// 单例模式
var G *Gocache

func NewCache(option ...Op) *Gocache {
	op := new(Options)
	// 配置参数
	for _, opt := range option {
		opt(op)
	}

	// 默认大小
	if op.ExpireKeyChanNumber == 0 {
		op.ExpireKeyChanNumber = DefaultExpireKeyChanNumber
	}
	if op.ExpireKeyIntervalDuration == time.Duration(0) {
		op.ExpireKeyIntervalDuration = DefaultExpireKeyIntervalTime
	}

	// 实例化
	g := &Gocache{
		allKey:        make(map[string]*KeyExpireChan),
		fifo:          make(map[string]*fifoArray),
		keyExpireChan: make(chan *KeyExpireChan, op.ExpireKeyChanNumber),
		lock:sync.RWMutex{},
		option:op,
	}
	// 初始化key-value结构
	g.keyValueCache = NewKeyValue(g.keyExpireChan)

	// 启动扫描过期时间
	go g.clearExpireKey()
	// 启动添加过期时间key
	go g.loopAddExpireKey()
	// 启动Lru算法 删除最不长使用的key
	go g.clearLruKey()
	G = g
	return g
}

// 清理过期时间
func (g *Gocache) clearExpireKey() {
	// todo 暂时使用最简单的map结构 循环清理
	// todo 下个版本使用有序集合 清理时间更好
	interval := g.option.ExpireKeyIntervalDuration
	// 生成timmer定时器
	ticker := time.NewTicker(time.Duration(interval))
	defer ticker.Stop()
	for range ticker.C {
		// todo 如果key很大的话 这里会运行多个吗
		// todo 不会存在的C 执行完了 会发送类似done信号  下个C chan才会进来
		nowTime := time.Now().Unix()
		for key,val := range g.allKey {
			if (val.expire != 0) &&  val.expire <= nowTime {
				g.deleteCaseType(val)
				delete(g.allKey,key)
			}
		}
	}
}
// 根据type删除对应类型下面的key
func (g *Gocache) deleteCaseType(val *KeyExpireChan)  {
	switch val.types {
	case keyValueType:
		g.keyValueCache.Delete(val.key)
	case fifoArrayType:
		// todo 暂时未开发
	}
}
// lur 算法
func (g *Gocache) clearLruKey()  {
}

// channel 
func (g *Gocache) addExpireKeyChan(v *KeyExpireChan) {
	g.keyExpireChan <- v
}

//  监听 循环监听过期key
func (g *Gocache) loopAddExpireKey() {
	for {
		value := <-g.keyExpireChan;
		g.allKey[value.key] = value
	}
}

// 删除过期key
func (g *Gocache) delete(key string)  bool{
	delete(g.allKey,key)
	return true
}

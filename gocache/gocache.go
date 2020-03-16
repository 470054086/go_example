package gocache

import (
	"fmt"
	"time"
)

//定义一个缓存的示例
type Gocache struct {
	// 存储过期key的一个map 这里应该使用有序集合更好
	expire map[string]time.Duration
	//  定义个map 用于存储所有的key
	key map[string]string
	// 数组结构队列
	fifo map[string]*fifoArray
	// KeyValue结构
	keyVale map[string]*keyValue
	// 用于添加过期时间的key
	keyExpireChan chan *KeyExpireChan
	option Options
	// 继承他们的方法
	keyValue
	fifoArray
}

func NewCache(option ...Op)  *Gocache {
	op := new(Options)
	// 配置参数
	for _,opt := range option {
		opt(op)
	}
	// 默认大小
	if op.ExpireKeyChanNumber == 0 {
		op.ExpireKeyChanNumber = DefaultExpireKeyChanNumber
	}
	// 实例化
	g := &Gocache{
		expire:  make(map[string]time.Duration),
		key:     make(map[string]string),
		fifo:    make(map[string]*fifoArray),
		keyVale: make(map[string]*keyValue),
		keyExpireChan:make(chan *KeyExpireChan,op.ExpireKeyChanNumber),
	}

	// 启动扫描过期时间
	go g.clearExpireKey()
	// 启动添加过期时间key
	go g.addExpireKey()
	return g;
}

// 清理过期时间
func (g *Gocache) clearExpireKey()  {

}

// 添加过期时间
func (g *Gocache) addExpireKey()  {
	for value := range g.keyExpireChan {
		fmt.Println(value)
	}
}



package cache

import (
	"time"
	"sync"
	"math"
)

const  (
	DefaultTime time.Duration = 0
)

type Cache struct {
	*cache
}


//执行缓存的写入操作
func (c *cache) Set(key string,value interface{},expir time.Duration)  {
	if expir == DefaultTime {
		expir = c.defaultExpiration
	}
	expirUnix := time.Now().Add(expir).UnixNano();
	c.mu.Lock()
	c.items[key] = Item{
		Object:value,
		Expiration:expirUnix,
	}
	c.mu.Unlock();
}


//获取缓存中的数据
func (c *cache) Get(key string) (interface{},bool) {
	c.mu.RLock();
	item,ok := c.items[key];
	if !ok {
		c.mu.Unlock()
		return nil,false
	}
	time := time.Now().UnixNano()
	if item.Expiration < time {
		c.mu.Unlock()
		delete(c.items,key)
		return nil,false
	}
	return item.Object,true
}

func (c *cache) GetExpire(key string)(int64,bool)  {
	c.mu.RLock();
	item,ok := c.items[key];
	if !ok {
		c.mu.Unlock()
		return 0,false
	}
	time := time.Now().UnixNano()
	if item.Expiration < time {
		c.mu.Unlock()
		delete(c.items,key)
		return 0,false
	}
	res := int64(math.Ceil(float64(item.Expiration-time)/1000))
	return res,true
}

func (c *cache) Del(key string)  {
	c.mu.Lock()
	delete(c.items,key)
	c.mu.Unlock()
}

type cache struct {
	items             map[string]Item
	mu                sync.RWMutex
	defaultExpiration time.Duration
}
//执行清理的操作
func (c *cache)clearItem()  {
	//进行锁的操作
	times:= time.Now().UnixNano()
	//这里为什么加锁就挂了 ???
	//c.mu.Lock()
	for key,value := range c.items {
		if value.Expiration>0 &&  value.Expiration < times {
			delete(c.items,key)
		}
	}
	//c.mu.Unlock()
}


type Item struct {
	Object interface{}
	Expiration int64
}


//生成一个cache
func newCache(defaultExporation time.Duration,item map[string]Item) *cache {
	if defaultExporation == 0 {
		defaultExporation = -1
	}
	c := &cache{
		items:item,
		defaultExpiration:defaultExporation,
	}
	return c;
}

func cleanTimer(cleanTime time.Duration,c *cache)  {
	ticker := time.NewTicker(cleanTime)
	for {
		select {
		case <-ticker.C:
			c.clearItem()
		default:
		}
	}
}

//生成一个Cache
func New(defaultExporation,cleanUpTime time.Duration) *Cache {
	//生成一个items
	item := make(map[string]Item)
	cache := newCache(defaultExporation,item)
	C := &Cache{cache}
	//如果清除的时间大于1的话 就执行清除操作
	if cleanUpTime > 0 {
		go cleanTimer(cleanUpTime,cache)
	}
	return C

}


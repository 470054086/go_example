package gocache

import (
	"time"
)

// 队列服务
type fifoArray struct {
	// 数组队列
	value []interface{}
	// 过期时间
	expire time.Duration
	// 链表指定长度
	len int
}

// 入队
//func (f *fifoArray) lpush(key string,value interface{}) bool  {
//
//}
//// 出队
//func (f *fifoArray) lpop(key string) interface{} {
//
//}
//
//// 入栈
//func (f *fifoArray) rpush(key string,value interface{}) bool  {
//
//}
//// 出栈
//func (f *fifoArray) rpop(key string) interface{}  {
//
//}
//// 队列长度
//func (f *fifoArray) llen() int  {
//	return f.len
//}
//// 返回指定位置的值
//func (f * fifoArray) lrange(key string,index ...int) []interface{}  {
//
//}





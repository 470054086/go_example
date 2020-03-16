package test

import (
	"gocache"
	"testing"
	"time"
)
var C *gocache.Gocache
func init()  {
	c := gocache.NewCache(func(options *gocache.Options) {
		options.ExpireKeyIntervalDuration = time.Second * 10
	})
	C = c
}
func TestGoCachePutAndGet( t *testing.T)  {
	name:="xiaobai"
	C.Put("name",name)
	get, _ := C.Get("name")
	if get != name {
		t.Error("TestGoCachePutAndGet error")
	}
}
func TestGoCachePutExpireAndGet( t *testing.T)  {
	name:="xiaobai"
	C.PutExpire("name",name,4)
	time.Sleep(time.Second * 5 )
	get, _ := C.Get("name")
	if get != "" {
		t.Error("TestGoCachePutExpireAndGet error")
	}
}





package main

import (
	"gocache"
	"time"
)
func main()  {
	c := gocache.NewCache(func(options *gocache.Options) {
		options.ExpireKeyIntervalDuration = time.Second * 10
	})
	c.PutExpire("name","xiaobai",4)
	for {
		time.Sleep(time.Second *3 )
	}
}
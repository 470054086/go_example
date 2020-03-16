package main

import (
	"gocache"
	"os"
	"os/signal"
	"syscall"
	"time"
)
func main()  {
	gocache.NewCache(func(options *gocache.Options) {
		options.ExpireKeyIntervalDuration = time.Second * 10
	})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigs:

	}
}
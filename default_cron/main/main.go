package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"xiaobai.com/go_example/default_cron/Scheduler"
)

func main() {
	c := Scheduler.NewCron()
	c.AddFunc("print_r", "*/2 * * * * * *", func(entry *Scheduler.Entry) error {
		fmt.Printf("我是第一个任务%s", entry.Name())
		return nil
	})
	c.AddFunc("print_r/2", "*/4 * * * * * *", func(entry *Scheduler.Entry) error {
		fmt.Printf("我是第二个任务%s", entry.Name())
		return nil
	})
	c.Start()
	c.AddFunc("print_3/3", "*/3 * * * * * *", func(entry *Scheduler.Entry) error {
		fmt.Printf("我是第三个任务%s", entry.Name())
		return nil
	})
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	time.Sleep(time.Minute * 2 )
	select {
	case <-ch:
		return
	}
}

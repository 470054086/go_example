package Scheduler

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"sort"
	"sync"
	"time"
)

type Cron struct {
	entris   []*Entry
	add      chan *Entry
	remove   EntryId
	sy       sync.Mutex
	running  bool //是否运行
	schedule Scheduler
	NextId   int
}
type EntryId int

type Entry struct {
	id        EntryId
	name      string
	parse     string
	mustParse *cronexpr.Expression
	job       job
	nextTime  time.Time
}

func (e *Entry) Name() string {
	return e.name
}

type job func(entry *Entry) error

type opt func(cron *Cron)

// 创建计划任务
func NewCron(opts ...opt) *Cron {
	c := &Cron{
		entris:   nil,
		add:      make(chan *Entry),
		remove:   0,
		sy:       sync.Mutex{},
		running:  false,
		schedule: &SimpleScheduler{},
	}
	for _, op := range opts {
		op(c)
	}
	return c
}

// 添加计划任务
func (c *Cron) AddFunc(name, parse string, f job) EntryId {
	c.sy.Lock()
	defer c.sy.Unlock()
	//构造一个entry请求
	c.NextId++
	entry := &Entry{
		id:    EntryId(c.NextId),
		name:  name,
		parse: parse,
		job:   f,
	}
	if c.running {
		c.add <- entry
	} else {
		c.entris = append(c.entris, entry)
	}
	return EntryId(c.NextId)
}

func (c *Cron) Start() {
	c.sy.Lock()
	defer c.sy.Unlock()
	if c.running {
		return
	}
	go c.start()
}

func (c *Cron) start() {
	now := time.Now()
	// 获取所有任务的下一次执行时间
	for _, entry := range c.entris {
		parse := cronexpr.MustParse(entry.parse)
		entry.nextTime = parse.Next(now)
		entry.mustParse = parse
	}

	for {
		//  进行一次排序
		sort.Slice(c.entris, func(i, j int) bool {
			return c.entris[i].nextTime.Before(c.entris[j].nextTime)
		})
		timer := time.NewTimer(c.entris[0].nextTime.Sub(now))
		timeDiff := &time.Ticker{C: nil}
		for {
			select {
			case <-timer.C:
				now := time.Now()
				// 计算下次执行任务的时间
				for _, entry := range c.entris {
					// 因此已经进行了排序 第一个不成功的话 就说明下面都还没到时间
					if entry.nextTime.After(now) {
						break
					}
					fmt.Printf("执行任务的名称为%s,时间为%s",
						entry.name,
						time.Now().Format("2006-01-02 15:04:05"))
					// 执行命令
					go c.schedule.Run(entry)
					// 更新下次执行时间
					now2 := time.Now()
					entry.nextTime = entry.mustParse.Next(now2)
					fmt.Printf("下次执行任务名称为%s,名称为%s\n",
						entry.name,
						entry.nextTime.Format("2006-01-02 15:04:05"))
					timer = time.NewTimer(entry.nextTime.Sub(now2))
				}
				// 计算下次最近的需要执行多久
				sort.Slice(c.entris, func(i, j int) bool {
					return c.entris[i].nextTime.Before(c.entris[j].nextTime)
				})
				now3 := time.Now()
				timeDiff = time.NewTicker(c.entris[0].nextTime.Sub(now3))
				// 添加
			case e := <-c.add:
				c.entris = append(c.entris, e)
			case <-timeDiff.C:
				fmt.Println("当前执行休眠")
			}
		}
	}

}

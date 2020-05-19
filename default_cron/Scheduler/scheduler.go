package Scheduler

type Scheduler interface {
	Run(entry *Entry)
}

type SimpleScheduler struct {
}

// 执行任务
func (s *SimpleScheduler) Run(entry *Entry) {
	entry.job(entry)
}

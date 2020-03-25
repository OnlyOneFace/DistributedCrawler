package scheduler

import (
	"DistributedCrawler/engine"
)

/*
* @
* @Author:
* @Date: 2020/3/23 18:54
 */
// 任务队列的实例
type SimpleScheduler struct {
	workerChan chan *engine.Request // 请求队列
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(cap int) {
	s.workerChan = make(chan *engine.Request, cap) // 请求队列
}

func (s *SimpleScheduler) Submit(r *engine.Request) {
	go func() {
		if r != nil {
			s.workerChan <- r
		}
	}()
}

func (s *SimpleScheduler) TaskChan() chan *engine.Request {
	return s.workerChan
}

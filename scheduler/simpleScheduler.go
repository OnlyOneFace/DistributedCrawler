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

func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan *engine.Request) {
	s.workerChan = in
}

func (s *SimpleScheduler) Submit(r *engine.Request) {
	go func() {
		if r != nil {
			s.workerChan <- r
		}
	}()
}

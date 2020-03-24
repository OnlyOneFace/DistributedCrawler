package scheduler

import (
	"DistributedCrawler/engine"
)

/*
* @
* @Author:
* @Date: 2020/3/23 18:54
 */

type SimpleScheduler struct {
	workerChan chan *engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan *engine.Request) {
	s.workerChan = in
}

func (s *SimpleScheduler) Submit(r *engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

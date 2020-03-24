package engine

import (
	"fmt"
)

/*
* @
* @Author:
* @Date: 2020/3/23 18:40
 */

type CurEngine struct {
	Scheduler   AllScheduler // 调度器接口
	WorkerCount int
}

func (ce *CurEngine) Run(seeds ...*Request) {
	in := make(chan *Request)
	ce.Scheduler.ConfigureMasterWorkerChan(in)
	out := make(chan *ParseResult)
	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(in, out)
	}
	for _, r := range seeds {
		ce.Scheduler.Submit(r)
	}
	itemcount := 0
	for result := range out {
		for _, item := range result.Items {
			fmt.Printf("Got item #%d %v", itemcount, item)
			itemcount++
		}
		for _, r := range result.Requests {
			if r != nil {
				ce.Scheduler.Submit(r)
			}

		}
	}
}

func createWorker(in chan *Request, out chan *ParseResult) {
	go func() {
		for r := range in {
			if r != nil {
				result, err := work(r)
				if err != nil {
					continue
				}
				out <- result
			}
		}
	}()
}

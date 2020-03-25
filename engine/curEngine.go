package engine

import (
	"DistributedCrawler/fetcher"
	"log"
)

/*
* @
* @Author:
* @Date: 2020/3/23 18:40
 */
// 引擎分发(单机并发版)
type CurEngine struct {
	Scheduler     AllScheduler // 调度器接口
	WorkerCount   int
	Out           chan *Response // 响应队列
	ItemChan      chan interface{}
	IsDuplication map[string]bool
}

func (c *CurEngine) Run(seeds ...*Request) {
	c.Out = make(chan *Response, 10)          // 配置响应队列
	c.Scheduler.ConfigureMasterWorkerChan(10) // 配置任务队列
	// 起工作协程
	for i := 0; i < c.WorkerCount; i++ {
		go createWorker2(c.Scheduler.TaskChan(), c.Out)
	}
	// 给任务队列发送初始请求任务
	for _, r := range seeds {
		if _, ok := c.IsDuplication[r.Url]; ok {
			continue
		} else {
			c.IsDuplication[r.Url] = true
		}
		c.Scheduler.Submit(r)
	}
	// 处理输出队列
	for {
		result := <-c.Out
		go func() {
			c.ItemChan <- result.Items
		}()
		if result.Requests == nil {
			continue
		}
		if _, ok := c.IsDuplication[result.Requests.Url]; ok { // 去重
			continue
		} else {
			c.IsDuplication[result.Requests.Url] = true
		}
		c.Scheduler.Submit(result.Requests)
	}
}

type workerFunc func(chan *Request, chan *Response)

func createWorker3() workerFunc {
	return func(requests chan *Request, responses chan *Response) {
		for r := range requests {
			work(r, responses)
		}
	}
}

func createWorker2(requests chan *Request, responses chan *Response) {
	for r := range requests {
		work(r, responses)
	}
}

func work(r *Request, out chan *Response) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher:error fetching url %s:%v", r.Url, err)
		return
	}
	r.ParseFunc(body, out)
}

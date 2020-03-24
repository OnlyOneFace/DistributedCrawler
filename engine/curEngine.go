package engine

import (
	"fmt"
	"sync"
)

/*
* @
* @Author:
* @Date: 2020/3/23 18:40
 */
// 引擎分发
type CurEngine struct {
	Scheduler   AllScheduler // 调度器接口
	WorkerCount int
	Out         chan *Response // 响应队列
	wg          sync.WaitGroup // 控制工作池退出
}

func (c *CurEngine) Run(seeds ...*Request) {
	var in = make(chan *Request, 10)          // 请求队列
	c.Out = make(chan *Response, 10)          // 配置响应队列
	c.Scheduler.ConfigureMasterWorkerChan(in) // 配置任务队列
	// 起工作协程
	for i := 0; i < c.WorkerCount; i++ {
		go createWorker(in, c.Out)
	}
	// 给任务队列发送初始请求任务
	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}
	// 处理输出队列
	itemsCount := 0
	for result := range c.Out {
		fmt.Printf("Got item %d %v\n", itemsCount, result.Items) // 结果输出
		// 将响应带的请求发送到任务队列
		c.Scheduler.Submit(result.Requests)
		itemsCount++
	}
}

func createWorker(in chan *Request, out chan *Response) {
	for r := range in {
		work(r, out)
	}
}

package main

import (
	"DistributedCrawler/engine"
	"DistributedCrawler/persist"
	"DistributedCrawler/scheduler"
	"DistributedCrawler/zhenai/parser"
)

/*
* @
* @Author:
* @Date: 2020/3/23 12:58
 */
const seed = "http://www.zhenai.com/zhenghun"

func main() {
	e := &engine.CurEngine{
		Scheduler:     &scheduler.SimpleScheduler{},
		WorkerCount:   100,
		ItemChan:      persist.ItemSaver(),
		IsDuplication: make(map[string]bool),
	}
	e.Run(&engine.Request{
		Url:       seed,
		ParseFunc: parser.ParseCityList,
	})
}

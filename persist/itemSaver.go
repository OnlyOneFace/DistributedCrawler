package persist

import (
	"DistributedCrawler/rpcSupport"
)

/*
* @
* @Author:
* @Date: 2020/3/25 10:28
 */
// 单机并发版
func ItemSaver(host string) chan interface{} {
	out := make(chan interface{}, 10)
	client, err := rpcSupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			item := <-out
			var result string
			if err = client.Call("ItemSaverService.Save", item, &result); err != nil {
				panic(err)
			}
		}
	}()
	return out
}

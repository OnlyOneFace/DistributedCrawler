package persist

import (
	"DistributedCrawler/rpcSupport"
	"fmt"
)

/*
* @
* @Author:
* @Date: 2020/3/25 10:28
 */
// 单机并发版->分布式版
func ItemSaver(host string) chan interface{} {
	out := make(chan interface{}, 10)
	client, err := rpcSupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	var count int
	go func() {
		for {
			item := <-out
			var result string
			if err = client.Call("ItemSaverService.Save", fmt.Sprintf("%v %v", count, item), &result); err != nil {
				panic(err)
			}
			fmt.Println(result)
			count++
		}
	}()
	return out
}

package persist

import "log"

/*
* @
* @Author:
* @Date: 2020/3/25 10:28
 */
func ItemSaver() chan interface{} {
	out := make(chan interface{}, 10)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item %d %v\n", itemCount, item)
			itemCount++
		}
	}()
	return out
}

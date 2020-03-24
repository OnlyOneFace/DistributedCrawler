package engine

import (
	"DistributedCrawler/fetcher"
	"log"
)

/*
* @
* @Author:
* @Date: 2020/3/23 15:01
 */
// 请求结构体
type Request struct {
	Url       string
	ParseFunc func([]byte, chan *Response)
}

// 响应结构体
type Response struct {
	Requests *Request
	Items    interface{}
}

// 空的处理情况
func NilParser([]byte) *Response {
	return nil
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

package engine

import (
	"DistributedCrawler/fetcher"
	"log"
)

/*
* @
* @Author:
* @Date: 2020/3/23 15:11
 */
type SimpleEngine struct{}

func (se *SimpleEngine) Run(seeds ...*Request) {
	var requests = make([]*Request, 0, len(seeds))
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := work(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func work(r *Request) (*ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher:error fetching url %s:%v", r.Url, err)
		return nil, nil
	}
	return r.ParseFunc(body), nil
}

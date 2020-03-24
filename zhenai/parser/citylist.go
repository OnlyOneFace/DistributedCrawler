package parser

import (
	"DistributedCrawler/engine"
	"regexp"
)

/*
* @
* @Author:
* @Date: 2020/3/23 15:00
 */
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte, out chan *engine.Response) {
	matches := regexp.MustCompile(cityListRe).FindAllSubmatch(contents, -1)
	for _, m := range matches {
		out <- &engine.Response{
			Requests: &engine.Request{
				Url:       string(m[1]),
				ParseFunc: ParseCity,
			},
			Items: string(m[2]),
		}
	}
}

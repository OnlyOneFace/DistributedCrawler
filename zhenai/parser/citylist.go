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

func ParseCityList(contents []byte) *engine.ParseResult {
	matches := regexp.MustCompile(cityListRe).FindAllSubmatch(contents, -1)
	result := &engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, &engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
	}
	return result
}

package parser

import (
	"DistributedCrawler/engine"
	"regexp"
)

/*
* @
* @Author:
* @Date: 2020/3/23 15:55
 */
const cityRe = `<a href="(http://[^.]+.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) *engine.ParseResult {
	matches := regexp.MustCompile(cityRe).FindAllSubmatch(contents, -1)
	result := &engine.ParseResult{}
	for _, m := range matches {
		name := "User " + string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, &engine.Request{
			// Url:       string(m[1]),
			Url:       "",
			ParseFunc: engine.NilParser,
			// ParseFunc: func(c []byte) *engine.ParseResult {
			// 	return ParseProfile(c, name)
			// },
		})
	}
	return result
}

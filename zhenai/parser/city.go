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

func ParseCity(contents []byte, out chan *engine.Response) {
	matches := regexp.MustCompile(cityRe).FindAllSubmatch(contents, -1)
	for _, m := range matches {
		name := "User " + string(m[2])
		out <- &engine.Response{
			Items: name,
		}
	}
}

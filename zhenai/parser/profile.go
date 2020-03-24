package parser

import (
	"DistributedCrawler/engine"
	"DistributedCrawler/model"
	"regexp"
	"strconv"
)

/*
* @
* @Author:
* @Date: 2020/3/23 16:10
 */
var (
	ageRe    = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)岁</div>`)
	heightRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)cm</div>`)
	weightRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)kg</div>`)
	incomeRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">月收入:([^<]+)</div>`)
	HokouRe  = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn pink">籍贯:([^<]+)</div>`)
	xinzuoRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^(]+).+\)</div>`)
)

func ParseProfile(contents []byte, name string) *engine.ParseResult {
	profile := &model.Profile{}
	// 姓名
	profile.Name = name
	// 年龄
	match := extractString(contents, ageRe)
	profile.Age, _ = strconv.Atoi(match)
	// 身高
	match = extractString(contents, heightRe)
	profile.Height, _ = strconv.Atoi(match)
	// 体重
	match = extractString(contents, weightRe)
	profile.Weight, _ = strconv.Atoi(match)
	// 月收入
	profile.Income = extractString(contents, incomeRe)
	// 籍贯
	profile.Hokou = extractString(contents, HokouRe)
	// 星座
	profile.Xinzuo = extractString(contents, xinzuoRe)
	return &engine.ParseResult{Items: []interface{}{profile}}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	matches := re.FindSubmatch(contents)
	if len(matches) >= 2 {
		return string(matches[1])
	}
	return ""
}

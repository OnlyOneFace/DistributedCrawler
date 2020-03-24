package engine

/*
* @
* @Author:
* @Date: 2020/3/23 15:01
 */

type Request struct {
	Url       string
	ParseFunc func([]byte) *ParseResult
}

type ParseResult struct {
	Requests []*Request
	Items    []interface{}
}

func NilParser([]byte) *ParseResult {
	return &ParseResult{}
}

package parser

import (
	"io/ioutil"
	"testing"
)

/*
* @
* @Author:
* @Date: 2020/3/23 15:27
 */
func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)
	if len(result.Requests) != 470 {
		t.Errorf("result should have %d requests;but have %d", 470, len(result.Requests))
	}
	if len(result.Items) != 470 {
		t.Errorf("result should have %d requests;but have %d", 470, len(result.Items))
	}
	var (
		expectedUrls = []string{
			"http://www.zhenai.com/zhenghun/aba",
		}
		expectedCitys = []string{
			"阿坝",
		}
	)
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s;but was %s", i, url, result.Requests[i].Url)
		}
	}
	for i, city := range expectedCitys {
		if result.Items[i] != city {
			t.Errorf("expected city #%d: %s;but was %s", i, city, result.Items[i])
		}
	}
}

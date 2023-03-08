package parser

import (
	"goprojects/crawler/engine"
	"regexp"
)

// example : "<a href="http://localhost:8080/mock/www.zhenai.com/zhenghun/aba" class = "">阿坝</a>"
const cityURLRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]+>([^<]+)</a>`

//// example : "<a href="http://www.zhenai.com/zhenghun/aba" data-v-602e7f5e>阿坝</a>
//const cityURLRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data[^>]+>([^<]+)</a>`

func ParseCityList(body []byte) engine.ParseResult {
	r := regexp.MustCompile(cityURLRe)
	matches := r.FindAllSubmatch(body, -1)

	parsedRequests := []engine.Request{}
	parsedItems := []engine.Item{}
	limit := 20
	for _, match := range matches {
		parsedRequests = append(parsedRequests, engine.Request{Url: string(match[1]), ParseFunc: ParseCity})
		parsedItems = append(parsedItems, engine.Item{Payload: match[2]})

		limit--
		if limit <= 0 {
			break
		}

	}
	return engine.ParseResult{Requests: parsedRequests, Items: parsedItems}
}

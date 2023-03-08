package parser

import (
	"fmt"
	"goprojects/crawler/engine"
	"regexp"
)

// user link
// example: <a href="http://localhost:8080/mock/album.zhenai.com/u/8256018539338750764">寂寞成影萌宝</a>
// example: <a href="http://localhost:8080/mock/www.zhenai.com/zhenghun/aba/4">4</a>
const userURLRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)">([^<]+)</a>`
const cityPageURLRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[a-zA-Z]+/[0-9]+)">([0-9]+)</a>`

// todo: crawler multiple page of a city

func ParseCity(body []byte) engine.ParseResult {
	r := regexp.MustCompile(userURLRe)
	matches := r.FindAllSubmatch(body, -1)

	parseRequests := []engine.Request{}
	parseItems := []engine.Item{}
	for _, match := range matches {
		url := string(match[1])
		parseRequests = append(parseRequests, engine.Request{
			Url: url,
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, url)
			},
		})
		parseItems = append(parseItems, engine.Item{Payload: fmt.Sprintf("Profile %s", match[2])})
	}

	r = regexp.MustCompile(cityPageURLRe)
	matches = r.FindAllSubmatch(body, -1)

	for _, match := range matches {
		parseRequests = append(parseRequests, engine.Request{
			Url:       string(match[1]),
			ParseFunc: ParseCity,
		})
		parseItems = append(parseItems, engine.Item{Payload: fmt.Sprintf("City %s Page %s", match[1], match[2])})
	}

	return engine.ParseResult{Requests: parseRequests, Items: parseItems}
}

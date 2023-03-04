package engine

import (
	"goprojects/crawler/fetcher"
	"log"
)

type SimpleEngine struct{}

func (e *SimpleEngine) Run(seeds ...Request) {
	queue := []Request{}
	queue = append(queue, seeds...)

	for len(queue) > 0 {
		request := queue[0]
		queue = queue[1:]

		result, err := worker(request)
		if err != nil {
			log.Printf("Error: worker error %v", err)
			continue
		}

		queue = append(queue, result.Requests...)
		log.Printf("Info: url %s success, items %s\n", request.Url, result.Items)
	}
	log.Printf("Info: queue finished. Exit. \n")
}

func worker(request Request) (ParseResult, error) {
	body, err := fetcher.FetchURL(request.Url)
	if err != nil {
		return ParseResult{}, err
	}
	result := request.ParseFunc(body)
	return result, nil
}

func PlaceholderParseFunc(body []byte) ParseResult {
	return ParseResult{}
}

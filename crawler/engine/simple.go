package engine

import (
	"log"
)

type SimpleEngine struct{}

func (e *SimpleEngine) Run(seeds ...Request) {
	queue := []Request{}
	queue = append(queue, seeds...)

	for len(queue) > 0 {
		request := queue[0]
		queue = queue[1:]

		result, err := Worker(request)
		if err != nil {
			log.Printf("Error: worker error %v", err)
			continue
		}

		queue = append(queue, result.Requests...)
		log.Printf("Info: url %s success, items %s\n", request.Url, result.Items)
	}
	log.Printf("Info: queue finished. Exit. \n")
}

func PlaceholderParseFunc(body []byte) ParseResult {
	return ParseResult{}
}

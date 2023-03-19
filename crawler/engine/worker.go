package engine

import "goprojects/crawler/fetcher"

func Worker(request Request) (ParseResult, error) {
	body, err := fetcher.FetchURL(request.Url)
	if err != nil {
		return ParseResult{}, err
	}
	result := request.ParseFunc(body)
	return result, nil
}

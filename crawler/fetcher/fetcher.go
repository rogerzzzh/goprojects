package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func FetchURL(url string) ([]byte, error) {
	<-rateLimiter
	log.Printf("Fetcher Info: Fetching %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Fetcher Error: Fetching URL %s with error %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Fetcher Error: URL %s response code %v error", url, resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Fetcher Error: Reading URL %s response error %v", url, err)
	}
	return body, nil
}

package persist

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
)

func getClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses:              []string{"https://localhost:9200"},
		Username:               "elastic",
		Password:               "jmM5KaDXWj09nntaRzLj",
		CertificateFingerprint: "72d9dc7c48057d40a4703df45c6ca72989d5cae0a1413798a9edc80ca016827d",
	}
	return elasticsearch.NewClient(cfg)
}

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	counter := 0
	go func() {
		for item := range out {
			log.Printf("Saver Info: Got #%d item to save, %s", counter, item)
			go save(item)
		}
	}()
	return out
}

func save(item interface{}) (string, error) {
	_, err := getClient()
	if err != nil {
		panic(err)
	}

	es, err := getClient()
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(item)
	if err != nil {
		panic(err)
	}

	req := esapi.IndexRequest{
		Index:   "test_index",
		Body:    bytes.NewReader(data),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	//res, err := es.Index("data", bytes.NewReader(data)).Do(context.Background())

	if res.IsError() {
		log.Fatalf("Saver Info: save failed, got response %s", res)
		return "", fmt.Errorf("Save Item to ES failed")
	} else {
		log.Printf("Saver Info: save successfully, got response %s", res)
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("[%s] %s; version=%s, id=%s", res.Status(), r["result"], r["_version"], r["_id"])
		}
		return r["_id"].(string), nil
	}
}

package persist

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"goprojects/crawler/engine"
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

func ItemSaver(index string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	es, err := getClient()
	if err != nil {
		return nil, err
	}
	counter := 0
	go func() {
		for item := range out {
			counter++
			log.Printf("Saver Info: Got #%d item to save, %s", counter, item)
			go save(es, index, item)
		}
	}()
	return out, nil
}

func save(es *elasticsearch.Client, index string, item engine.Item) error {
	data, err := json.Marshal(item)
	log.Printf("Saver Info: data = %s", data)
	if err != nil {
		panic(err)
	}

	if item.Id == "" {
		return errors.New("Saver error: no id provided in item. Skipped.")
	}

	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: item.Id,
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("Saver Info: save failed, got response %s", res)
		return errors.New("Save Item to ES failed")
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("Saver Info: save successfully, [%s] %s; version=%s, id=%s", res.Status(), r["result"], r["_version"], r["_id"])
		}
		return nil
	}
}

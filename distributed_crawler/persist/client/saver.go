package client

import (
	"goprojects/crawler/engine"
	"goprojects/distributed_crawler/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	if err != nil {
		return nil, err
	}
	counter := 0
	go func() {
		for item := range out {
			counter++
			log.Printf("Saver Info: Got #%d item to Save, %s", counter, item)
			result := ""
			client.Call("ItemSaverService.Save", item, &result)
		}
	}()
	return out, nil
}

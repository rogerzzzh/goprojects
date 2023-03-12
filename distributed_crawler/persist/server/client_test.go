package main

import (
	"goprojects/crawler/engine"
	"goprojects/crawler/zhenai"
	"goprojects/distributed_crawler/rpcsupport"
	"testing"
	"time"
)

func TestItemSave(t *testing.T) {
	// start itemsaver servie
	const host = ":1234"
	const index = "test_data"
	go serveRpc(host, index)

	// sleep to wait the service to boot
	time.Sleep(time.Second * 10)

	// create itemsaver client
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		t.Errorf("err = %s", err)
	}

	// call save
	profile := engine.Item{
		Id:  "abc",
		Url: "http://www.baidu.com",
		Payload: zhenai.UserProfile{
			Age:        15,
			Gender:     "Male",
			Name:       "Mundo",
			Height:     220,
			Income:     "3000",
			Marriage:   "",
			Education:  "",
			Occupation: "",
			Weight:     200,
		},
	}

	result := ""
	err = client.Call("ItemSaverService.Save", profile, &result)
	if err != nil || result != "ok" {
		t.Errorf("result = %s, err = %s", result, err)
	}
}

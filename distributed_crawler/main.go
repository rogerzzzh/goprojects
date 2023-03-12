package main

import (
	"goprojects/crawler/engine"
	"goprojects/crawler/scheduler"
	"goprojects/crawler/zhenai/parser"
	"goprojects/distributed_crawler/persist/client"
)

//const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun"

const url = "http://localhost:8080/mock/www.zhenai.com/zhenghun/aba"

//// example : "<a href="http://www.zhenai.com/zhenghun/aba" data-v-602e7f5e>阿坝</a>
//const url = "http://www.zhenai.com/zhenghun"
//const cityURLRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data[^>]+>([^<]+)</a>`

func main() {
	//e := engine.SimpleEngine{}
	//e := engine.ConcurrentEngine{Scheduler: &scheduler.SimpleScheduler{}, WorkerCount: 10}

	itemSaveChan, err := client.ItemSaver(":1234")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 10, ItemSaver: itemSaveChan}
	//e.Run(engine.Request{
	//	Url:       url,
	//	ParseFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url:       url,
		ParseFunc: parser.ParseCity,
	})
}

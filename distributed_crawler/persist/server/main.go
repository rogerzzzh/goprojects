package main

import (
	"goprojects/crawler/persist"
	"goprojects/distributed_crawler/config"
	persist2 "goprojects/distributed_crawler/persist"
	"goprojects/distributed_crawler/rpcsupport"
)

func main() {
	err := serveRpc(config.ItemSaverPort, config.ElasticIndex)
	if err != nil {
		panic(err)
	}
}

func serveRpc(host string, index string) error {
	es, err := persist.GetClient()
	if err != nil {
		return err
	}

	err = rpcsupport.ServeRpc(host, &persist2.ItemSaverService{Es: es, Index: index})
	return err
}

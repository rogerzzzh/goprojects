package main

import (
	"goprojects/crawler/persist"
	persist2 "goprojects/distributed_crawler/persist"
	"goprojects/distributed_crawler/rpcsupport"
)

func main() {
	const host = ":1234"
	const index = "test_data"
	err := serveRpc(host, index)
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

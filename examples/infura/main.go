package main

import (
	"encoding/json"
	"github.com/dkropachev/ethscan/pkg/memtxstore"
	"github.com/dkropachev/ethscan/pkg/subscriber"
	"time"
)

func main() {
	sub, err := subscriber.NewStoreSubscriber("https://mainnet.infura.io/v3/<API-KEY>", memtxstore.New())
	if err != nil {
		panic(err)
	}
	if err := sub.Start(); err != nil {
		panic(err)
	}
	defer sub.Stop()
	sub.Subscribe("0x1234567890")

	for range time.NewTicker(time.Second * 5).C {
		txs, err := sub.GetTransactions("0x1234567890")
		if err != nil {
			panic(err)
		}
		for _, tx := range txs {
			txTxt, err := json.Marshal(tx)
			if err != nil {
				panic(err)
			}
			println(string(txTxt))
		}
	}
}

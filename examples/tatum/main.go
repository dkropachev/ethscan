package main

import (
	"encoding/json"
	"github.com/dkropachev/ethscan/pkg/memtxstore"
	subscriber2 "github.com/dkropachev/ethscan/pkg/subscriber"
	"net/http"
	"time"
)

type TatumHTTPClient struct {
	apiKey string
}

func (t *TatumHTTPClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-Api-Key", t.apiKey)
	return http.DefaultClient.Do(req)
}

func main() {
	sub, err := subscriber2.NewStoreSubscriber(
		"https://<NODE-NAME>.rpc.tatum.io/<API-KEY>",
		memtxstore.New(),
		subscriber2.WithHTTPClient(&TatumHTTPClient{
			apiKey: "t-xxxxxxxxxxxxxxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxxxxxxx",
		}))
	if err != nil {
		panic(err)
	}
	if err := sub.Start(); err != nil {
		panic(err)
	}
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

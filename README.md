# Ethernet [JSON-RPC](https://ethereum.org/en/developers/docs/apis/json-rpc/) Subscriber

This is a simple subscriber for the Ethernet JSON-RPC Ethereum protocol. 
It listens for changes on Ethernet chain and track transactions for requested addresses.

## Installation

```bash
go install github.com/dkropachev/ethscan@latest
```

## Examples

### [Infura](https://www.infura.io/)

```bash
ethscan --endpoint https://mainnet.infura.io/v3/<API-KEY> --wallets 0xc940323bdacd868c319e9039ea5fddd35745e62d --start-block 19762452
```

### [Alchemy](https://www.alchemy.com/)
```bash
ethscan --endpoint https://eth-mainnet.g.alchemy.com/v2/<API-KEY> --wallets 0xc940323bdacd868c319e9039ea5fddd35745e62d --start-block 19762452
```

### [Tatum](https://tatum.io/)
```bash
ethscan --endpoint https://<NODE-NAME>.rpc.tatum.io/ --header "X-Api-Key: <API-KEY>" --wallets 0xc940323bdacd868c319e9039ea5fddd35745e62d --start-block 19762452
```

## Programmatic API

### In-Memory Subscriber

The library provides a subscriber that tracks transactions for requested addresses and stores them in memory


Usage of the library is simple. 
You can create a new instance of the `subscriber` pointing it to an Ethernet API endpoint:

```go
    sub, err := subscriber.NewMemSubscriber("https://mainnet.infura.io/v3/<API-KEY>")
	if err != nil {
        panic(err)
    }
```

Subscribe to wanted addresses:

```go
    sub.Subscribe("0x1234567890abcdef1234567890abcdef12345678")
```

Start the `subscriber` by calling the `Start` method:

```go
    if err := sub.Start(); err != nil {
        panic(err)
    }
    defer sub.Stop()
```

And read transactions:
```go
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
```


### Chan Subscriber

The library provides a subscriber that tracks transactions for requested addresses and sends them into a channel


Usage of the library is simple.
You can create a new instance of the `subscriber` pointing it to an Ethernet API endpoint:

```go
	sub, err := subscriber.NewChanSubscriber(endpoint, opts...)
	if err != nil {
		panic(err)
	}
```

Subscribe to wanted addresses:

```go
	sub.Subscribe("0x1234567890abcdef1234567890abcdef12345678")
```

Start the `subscriber` by calling the `Start` method:

```go
	if err := sub.Start(); err != nil {
		panic(err)
	}
	defer sub.Stop()
```

And read transactions from the channel:
```go	
    for tx := range sub.GetTransactionChan() {
        if tx == nil {
            break
        }
        txTxt, err := json.Marshal(tx)
        if err != nil {
            panic(err)
        }
        println(string(txTxt))
        txTxt, err := json.Marshal(tx)
        if err != nil {
            panic(err)
        }
        println(string(txTxt))
    }
```

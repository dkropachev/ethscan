package memtxstore_test

import (
	"github.com/dkropachev/ethscan/pkg/memtxstore"
	"github.com/dkropachev/ethscan/pkg/synclist"
	"github.com/dkropachev/ethscan/pkg/types"
	"math/big"
	"math/rand/v2"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquatableParallel(t *testing.T) {
	lst := memtxstore.New()
	wg := sync.WaitGroup{}

	atomicList := synclist.EquatableList[*types.Transaction]{}

	for range 10 {
		wg.Add(2)
		go func() {
			defer wg.Done()
			for range 10000 {
				to := randomAddr()
				from := randomAddr()
				tx := types.Transaction{
					Hash:             uniqueHash(),
					BlockHash:        uniqueHash(),
					BlockNumber:      randomInt(),
					From:             from,
					To:               to,
					Gas:              randomInt(),
					GasPrice:         randomInt(),
					Input:            types.BinData{1, 2, 3, 4},
					Nonce:            randomInt(),
					TransactionIndex: randomInt(),
					Value:            randomInt(),
				}
				if tx.To == targetAddress || tx.From == targetAddress {
					atomicList.AppendIfNotExists(&types.Transaction{})
				}

				err := lst.StoreTransaction(&tx)
				if err != nil {
					t.Error(err)
				}
			}
		}()
		go func() {
			defer wg.Done()
			for range 10000 {
				_, err := lst.GetTransactions(targetAddress.String())
				if err != nil {
					t.Error(err)
				}
			}
		}()
	}
	wg.Wait()
	txs, err := lst.GetTransactions(targetAddress.String())
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, atomicList.GetAll(), txs)
}

var atomicHashNum atomic.Int64

func uniqueHash() types.EthHash {
	hashID := atomicHashNum.Add(1)
	return types.EthHash{byte(hashID), byte(hashID >> 8), byte(hashID >> 16), byte(hashID >> 24), byte(hashID >> 32), byte(hashID >> 40), byte(hashID >> 48), byte(hashID >> 56)}
}

var targetAddress = types.EthAddress{10, 20, 30, 40, 50, 60, 70, 80}

func randomAddr() types.EthAddress {
	sameAdd := rand.N(10) < 3
	if sameAdd {
		return targetAddress
	}
	addr := rand.Int64()
	return types.EthAddress{byte(addr), byte(addr >> 8), byte(addr >> 16), byte(addr >> 24), byte(addr >> 32), byte(addr >> 40), byte(addr >> 48), byte(addr >> 56)}
}

func randomInt() types.BigInt {
	return types.BigInt(*big.NewInt(rand.Int64()))
}

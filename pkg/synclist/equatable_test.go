package synclist_test

import (
	"github.com/dkropachev/ethscan/pkg/synclist"
	"github.com/dkropachev/ethscan/pkg/types"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquatableParallel(t *testing.T) {
	lst := synclist.EquatableList[*types.Transaction]{}
	wg := sync.WaitGroup{}
	tx1Cnt := atomic.Int32{}
	tx1 := &types.Transaction{
		Hash: types.EthHash{10, 20, 30},
	}
	for range 10 {
		wg.Add(4)
		go func() {
			defer wg.Done()
			for range 1000 {
				lst.Append(tx1)
				tx1Cnt.Add(1)
			}
		}()
		go func() {
			defer wg.Done()
			for range 1000 {
				if lst.AppendIfNotExists(tx1) {
					tx1Cnt.Add(1)
				}
			}
		}()
		go func() {
			defer wg.Done()
			for range 10000 {
				lst.GetAll()
			}
		}()
		go func() {
			defer wg.Done()
			for range 10000 {
				lst.Contains(tx1)
			}
		}()
	}
	wg.Wait()
	assert.Equal(t, tx1Cnt.Load(), int32(lst.Len()))
	lst.GetAll()
}

package synclist_test

import (
	"github.com/dkropachev/ethscan/pkg/synclist"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComparableParallel(t *testing.T) {
	lst := synclist.ComparableList[string]{}
	wg := sync.WaitGroup{}
	fooCnt := atomic.Int32{}
	for range 10 {
		wg.Add(4)
		go func() {
			defer wg.Done()
			for range 1000 {
				lst.Append("foo")
				fooCnt.Add(1)
			}
		}()
		go func() {
			defer wg.Done()
			for range 1000 {
				if lst.AppendIfNotExists("foo") {
					fooCnt.Add(1)
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
				lst.Contains("foo")
			}
		}()
	}
	wg.Wait()
	assert.Equal(t, fooCnt.Load(), int32(lst.Len()))
	lst.GetAll()
}

package processors

import (
	stderr "errors"
	"github.com/dkropachev/ethscan/pkg/types"
)

type txStore interface {
	StoreTransaction(tx *types.Transaction) error
}

type TxStore struct {
	inChan <-chan *types.Transaction
	store  txStore
	errors chan error
}

func NewTxStore(inChan <-chan *types.Transaction, store txStore) *TxStore {
	out := &TxStore{
		inChan: inChan,
		store:  store,
		errors: make(chan error, 10),
	}
	go out.body()
	return out
}

func (p *TxStore) body() {
	defer close(p.errors)
	for tx := range p.inChan {
		if tx == nil {
			return
		}
		err := p.store.StoreTransaction(tx)
		if err != nil {
			select {
			case p.errors <- err:
			default:
				return
			}
		}
	}
}

func (p *TxStore) LastError() error {
	var errs []error
outer:
	for {
		select {
		case err := <-p.errors:
			errs = append(errs, err)
		default:
			break outer
		}
	}
	return stderr.Join(errs...)
}

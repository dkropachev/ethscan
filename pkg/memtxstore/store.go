package memtxstore

import (
	"github.com/dkropachev/ethscan/pkg/synclist"
	"github.com/dkropachev/ethscan/pkg/types"
	"math/big"
	"slices"
	"sync"
)

type Store struct {
	addrMap map[string]*synclist.EquatableList[*types.Transaction]

	addMapMutex sync.RWMutex
}

func New() *Store {
	return &Store{
		addrMap: make(map[string]*synclist.EquatableList[*types.Transaction]),
	}
}

func (s *Store) storeTx(address string, tx *types.Transaction) error {
	var val *synclist.EquatableList[*types.Transaction]
	s.addMapMutex.RLock()

	if val = s.addrMap[address]; val == nil {
		s.addMapMutex.RUnlock()
		s.addMapMutex.Lock()
		if val = s.addrMap[address]; val == nil {
			val = &synclist.EquatableList[*types.Transaction]{}
			s.addrMap[address] = val
		}
		s.addMapMutex.Unlock()
		s.addMapMutex.RLock()
	}
	val.AppendIfNotExists(tx)
	s.addMapMutex.RUnlock()
	return nil
}

func (s *Store) StoreTransaction(tx *types.Transaction) error {
	to := tx.To.String()
	from := tx.From.String()

	if err := s.storeTx(to, tx); err != nil {
		return err
	}

	if to != from {
		if err := s.storeTx(from, tx); err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) GetTransactions(address string) ([]*types.Transaction, error) {
	s.addMapMutex.RLock()
	defer s.addMapMutex.RUnlock()

	if lst := s.addrMap[address]; lst != nil {
		return lst.GetAll(), nil
	}
	return nil, nil
}

func (s *Store) GetTransactionsAfterBlock(blkId big.Int, address string) ([]*types.Transaction, error) {
	allTxs, err := s.GetTransactions(address)
	if err != nil {
		return nil, err
	}

	return slices.DeleteFunc(allTxs, func(tx *types.Transaction) bool {
		return tx.BlockNumber.AsBigInt().Cmp(&blkId) <= 0
	}), nil
}

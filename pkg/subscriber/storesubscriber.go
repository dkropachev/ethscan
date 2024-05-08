package subscriber

import (
	stderr "errors"
	"github.com/dkropachev/ethscan/pkg/blksubscriber"
	processors2 "github.com/dkropachev/ethscan/pkg/processors"
	"github.com/dkropachev/ethscan/pkg/types"
	"math/big"

	"github.com/pkg/errors"
)

type StoreSubscriber struct {
	blkSub           *blksubscriber.Subscriber[types.BlockDetailed]
	walletFilter     *processors2.TxWalletFilter
	txStoreProcessor *processors2.TxStore
	store            txStore
}

type txStore interface {
	StoreTransaction(tx *types.Transaction) error
	GetTransactions(address string) ([]*types.Transaction, error)
	GetTransactionsAfterBlock(blkId big.Int, address string) ([]*types.Transaction, error)
}

func NewStoreSubscriber(endpoint string, store txStore, opts ...Option) (*StoreSubscriber, error) {
	blkSub, err := blksubscriber.New[types.BlockDetailed](endpoint, convOptions(opts)...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create block subscriber")
	}
	walletFilter := processors2.NewTxWalletFilter(processors2.NewBlockToTxProcessor(blkSub.GetBlockChan()).Out())
	txStoreProcessor := processors2.NewTxStore(walletFilter.Out(), store)
	return &StoreSubscriber{
		blkSub:           blkSub,
		store:            store,
		walletFilter:     walletFilter,
		txStoreProcessor: txStoreProcessor,
	}, nil
}

func (s *StoreSubscriber) Subscribe(address string) bool {
	return s.walletFilter.AddWallet(address)
}

func (s *StoreSubscriber) IsRunning() bool {
	return s.blkSub.IsRunning()
}

func (s *StoreSubscriber) LastError() error {
	return stderr.Join(
		errors.Wrap(s.txStoreProcessor.LastError(), "transaction store processor error"),
		errors.Wrap(s.blkSub.LastError(), "block subscriber error"),
	)
}

func (s *StoreSubscriber) GetTransactions(address string) ([]*types.Transaction, error) {
	return s.store.GetTransactions(address)
}

func (s *StoreSubscriber) GetTransactionsAfterBlock(blkId big.Int, address string) ([]*types.Transaction, error) {
	return s.store.GetTransactionsAfterBlock(blkId, address)
}

func (s *StoreSubscriber) GetCurrentBlock() big.Int {
	return s.blkSub.GetCurrentBlock()
}

func (s *StoreSubscriber) Stop() {
	s.blkSub.Stop()
}

func (s *StoreSubscriber) Start() error {
	return errors.Wrap(s.blkSub.Start(), "failed to run block subscriber")
}

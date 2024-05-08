package subscriber

import (
	"github.com/dkropachev/ethscan/pkg/blksubscriber"
	processors2 "github.com/dkropachev/ethscan/pkg/processors"
	"github.com/dkropachev/ethscan/pkg/types"
	"math/big"

	"github.com/pkg/errors"
)

type ChanSubscriber struct {
	blkSub       *blksubscriber.Subscriber[types.BlockDetailed]
	walletFilter *processors2.TxWalletFilter
}

func NewChanSubscriber(endpoint string, opts ...Option) (*ChanSubscriber, error) {
	blkSub, err := blksubscriber.New[types.BlockDetailed](endpoint, convOptions(opts)...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create block subscriber")
	}

	return &ChanSubscriber{
		blkSub:       blkSub,
		walletFilter: processors2.NewTxWalletFilter(processors2.NewBlockToTxProcessor(blkSub.GetBlockChan()).Out()),
	}, nil
}

func (s *ChanSubscriber) Subscribe(address string) bool {
	return s.walletFilter.AddWallet(address)
}

func (s *ChanSubscriber) GetCurrentBlock() big.Int {
	return s.blkSub.GetCurrentBlock()
}

func (s *ChanSubscriber) IsRunning() bool {
	return s.blkSub.IsRunning()
}

func (s *ChanSubscriber) LastError() error {
	return errors.Wrap(s.blkSub.LastError(), "block subscriber error")
}

func (s *ChanSubscriber) GetTransactionChan() <-chan *types.Transaction {
	return s.walletFilter.Out()
}

func (s *ChanSubscriber) Stop() {
	s.blkSub.Stop()
}

func (s *ChanSubscriber) Start() error {
	return errors.Wrap(s.blkSub.Start(), "failed to run block subscriber")
}

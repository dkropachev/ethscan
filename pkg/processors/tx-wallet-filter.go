package processors

import (
	"github.com/dkropachev/ethscan/pkg/synclist"
	"github.com/dkropachev/ethscan/pkg/types"
)

type TxWalletFilter struct {
	inChan  <-chan *types.Transaction
	outChan chan *types.Transaction
	wallets synclist.ComparableList[string]
}

func NewTxWalletFilter(inChan <-chan *types.Transaction) *TxWalletFilter {
	out := &TxWalletFilter{
		inChan:  inChan,
		outChan: make(chan *types.Transaction, 1000),
	}
	go out.body()
	return out
}

func (p *TxWalletFilter) body() {
	defer close(p.outChan)
	for tx := range p.inChan {
		if tx == nil {
			return
		}
		if p.wallets.Contains(tx.From.String(), tx.To.String()) {
			p.outChan <- tx
		}
	}
}

func (p *TxWalletFilter) AddWallet(wallet string) bool {
	return p.wallets.AppendIfNotExists(wallet)
}

func (p *TxWalletFilter) Out() <-chan *types.Transaction {
	return p.outChan
}

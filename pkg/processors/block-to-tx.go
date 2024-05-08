package processors

import (
	"github.com/dkropachev/ethscan/pkg/types"
)

type BlockToTx struct {
	blkChan <-chan *types.BlockDetailed
	outChan chan *types.Transaction
}

func NewBlockToTxProcessor(blkChan <-chan *types.BlockDetailed) *BlockToTx {
	out := &BlockToTx{
		blkChan: blkChan,
		outChan: make(chan *types.Transaction, 1000),
	}
	go out.body()
	return out
}

func (p *BlockToTx) body() {
	defer close(p.outChan)
	for blk := range p.blkChan {
		if blk == nil {
			return
		}
		for _, tx := range blk.Transactions {
			p.outChan <- tx
		}
	}
}

func (p *BlockToTx) Out() <-chan *types.Transaction {
	return p.outChan
}

package subscriber

import (
	"github.com/dkropachev/ethscan/pkg/blksubscriber"
	"math/big"
	"net/http"
	"time"
)

type Option blksubscriber.Option

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func WithHTTPClient(cl httpClient) Option {
	return Option(blksubscriber.WithHTTPClient(cl))
}

func WithPoolingPeriod(period time.Duration) Option {
	return Option(blksubscriber.WithPoolingPeriod(period))
}

func WithStartBlock(blkId *big.Int) Option {
	return Option(blksubscriber.WithStartBlock(blkId))
}

func WithEndBlock(blkId *big.Int) Option {
	return Option(blksubscriber.WithEndBlock(blkId))
}

func convOptions(in []Option) []blksubscriber.Option {
	out := make([]blksubscriber.Option, len(in))
	for i, opt := range in {
		out[i] = blksubscriber.Option(opt)
	}
	return out
}

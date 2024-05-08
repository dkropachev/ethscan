package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/dkropachev/ethscan/pkg/blksubscriber"
	subscriber2 "github.com/dkropachev/ethscan/pkg/subscriber"
	"github.com/dkropachev/ethscan/pkg/types"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

type Options struct {
	endpoint      string
	header        string
	startBlock    string
	startBlockInt *big.Int
	endBlock      string
	endBlockInt   *big.Int
	poolingPeriod time.Duration
	wallets       string
	quite         bool
	target        string
}

func (o *Options) Parse() {
	flag.StringVar(&o.header, "header", "", "curl-style header to send with the request. Example: --header \"Authorization: Bearer <TOKEN>\"")
	flag.StringVar(&o.endpoint, "endpoint", "", "ethereum JSON-RPC endpoint")
	flag.StringVar(&o.startBlock, "start-block", "", "start block number")
	flag.StringVar(&o.endBlock, "end-block", "", "end block number")
	flag.DurationVar(&o.poolingPeriod, "poolingPeriod", time.Second, "pooling period")
	flag.StringVar(&o.wallets, "wallets", "", "wallets to subscribe, separated by comma")
	flag.BoolVar(&o.quite, "quite", false, "print out only transactions, no logs or messages")
	flag.StringVar(&o.target, "target", "tx", "target objects to print, options: tx, block, block-detailed")
	flag.Parse()
}

func (o *Options) Validate() error {
	var err error

	if o.endpoint == "" {
		return errors.New("endpoint option is required")
	}
	if o.wallets == "" {
		return errors.New("wallets option is required")
	}

	o.startBlockInt, err = parseBigInt(o.startBlock, "start-block")
	if err != nil {
		return err
	}

	o.endBlockInt, err = parseBigInt(o.endBlock, "end-block")
	if err != nil {
		return err
	}

	switch o.target {
	case "tx", "block", "block-detailed":
	default:
		return errors.Errorf("unknown target: %s\n", o.target)
	}

	return nil
}

func (o *Options) Run() error {
	switch o.target {
	case "tx":
		return subscribeTransaction(o.endpoint, strings.Split(o.wallets, ","), o.quite, o.buildSubscriberOptions()...)
	case "block":
		return subscribeBlocks[types.Block](o.endpoint, o.quite, o.buildBlkSubscriberOptions()...)
	case "block-detailed":
		return subscribeBlocks[types.BlockDetailed](o.endpoint, o.quite, o.buildBlkSubscriberOptions()...)
	default:
		return errors.Errorf("unknown target: %s\n", o.target)
	}
}

func (o *Options) buildSubscriberOptions() []subscriber2.Option {
	opts := []subscriber2.Option{
		subscriber2.WithPoolingPeriod(o.poolingPeriod),
	}

	if httpClient := newSimpleHTTPClient(o.header); httpClient != nil {
		opts = append(opts, subscriber2.WithHTTPClient(httpClient))
	}

	if o.startBlock != "" {
		opts = append(opts, subscriber2.WithStartBlock(o.startBlockInt))
	}

	if o.endBlock != "" {
		opts = append(opts, subscriber2.WithEndBlock(o.endBlockInt))
	}
	return opts
}

func (o *Options) buildBlkSubscriberOptions() []blksubscriber.Option {
	opts := []blksubscriber.Option{
		blksubscriber.WithPoolingPeriod(o.poolingPeriod),
	}

	if httpClient := newSimpleHTTPClient(o.header); httpClient != nil {
		opts = append(opts, blksubscriber.WithHTTPClient(httpClient))
	}

	if o.startBlock != "" {
		opts = append(opts, blksubscriber.WithStartBlock(o.startBlockInt))
	}

	if o.endBlock != "" {
		opts = append(opts, blksubscriber.WithEndBlock(o.endBlockInt))
	}
	return opts
}

func subscribeBlocks[T types.BlockType](
	endpoint string,
	quite bool,
	opts ...blksubscriber.Option,
) error {
	sub, err := blksubscriber.New[T](endpoint, opts...)
	if err != nil {
		return errors.Wrap(err, "failed to create subscriber")
	}

	if err = sub.Start(); err != nil {
		return errors.Wrap(err, "failed to start subscriber")
	}

	defer sub.Stop()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		// Run Cleanup
		sub.Stop()
	}()

	if !quite {
		fmt.Println("Listening for blocks...")
	}

	for block := range sub.GetBlockChan() {
		if block == nil {
			return errors.Wrap(sub.LastError(), "subscriber failed with error")
		}

		blockTxt, err := json.Marshal(block)
		if err != nil {
			return errors.Wrap(err, "failed to marshal block")
		}

		println(string(blockTxt))
	}
	return nil
}

func subscribeTransaction(
	endpoint string,
	walletList []string,
	quite bool,
	opts ...subscriber2.Option,
) error {
	sub, err := subscriber2.NewChanSubscriber(endpoint, opts...)
	if err != nil {
		return errors.Wrap(err, "failed to create subscriber")
	}

	for _, wallet := range walletList {
		sub.Subscribe(wallet)
	}

	if err = sub.Start(); err != nil {
		return errors.Wrap(err, "failed to start subscriber")
	}

	defer sub.Stop()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		// Run Cleanup
		sub.Stop()
	}()

	if !quite {
		fmt.Println("Listening for transactions...")
	}

	for tx := range sub.GetTransactionChan() {
		if tx == nil {
			return errors.Wrap(sub.LastError(), "subscriber failed with error")
		}

		txTxt, err := json.Marshal(tx)
		if err != nil {
			return errors.Wrap(err, "failed to marshal tx")
		}

		println(string(txTxt))
	}
	return nil
}

func parseBigInt(val, optionName string) (*big.Int, error) {
	if val == "" {
		return big.NewInt(0), nil
	}
	var out *big.Int
	var isOk bool
	if strings.HasPrefix(val, "0x") {
		out, isOk = new(big.Int).SetString(strings.TrimPrefix(val, "0x"), 16)
	} else {
		out, isOk = new(big.Int).SetString(val, 10)
	}
	if !isOk {
		return nil, errors.Errorf("failed to Parse %s to big.Int\n", optionName)
	}
	return out, nil
}

type simpleHTTPClient struct {
	headers http.Header
}

func newSimpleHTTPClient(header string) *simpleHTTPClient {
	tmp := strings.SplitN(header, ":", 1)
	if len(tmp) != 2 {
		return nil
	}
	headers := http.Header{}
	headers.Set(tmp[0], strings.TrimSpace(tmp[1]))
	return &simpleHTTPClient{
		headers: headers,
	}
}

func (c *simpleHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}

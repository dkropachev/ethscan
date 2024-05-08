package blksubscriber

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dkropachev/ethscan/pkg/types"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
)

type (
	rcpError struct {
		Error struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
	}

	options struct {
		poolingPeriod time.Duration
		currentBlock  atomic.Pointer[big.Int]
		endBlock      atomic.Pointer[big.Int]
		client        httpClient
	}

	Option func(opts *options)

	Subscriber[T types.BlockType] struct {
		url           url.URL
		ctx           context.Context
		ctxCancel     context.CancelFunc
		running       atomic.Bool
		blockDetailed bool
		blocksChan    chan *T
		lastError     atomic.Pointer[error]
		options
	}
)

func (o *options) apply(mods ...Option) {
	for _, opt := range mods {
		opt(o)
	}
}

func WithHTTPClient(cl httpClient) Option {
	return func(opts *options) {
		opts.client = cl
	}
}

func WithPoolingPeriod(period time.Duration) Option {
	return func(opts *options) {
		opts.poolingPeriod = period
	}
}

func WithStartBlock(blkId *big.Int) Option {
	return func(opts *options) {
		opts.currentBlock.Store(blkId)
	}
}

func WithEndBlock(blkId *big.Int) Option {
	return func(opts *options) {
		opts.endBlock.Store(blkId)
	}
}

const defaultPoolingPeriod = time.Second

var typeOfBlockDetailed = reflect.TypeOf(types.BlockDetailed{})

func New[T types.BlockType](endpoint string, opts ...Option) (*Subscriber[T], error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse endpoint URL")
	}
	if u.Scheme != "ws" && u.Scheme != "http" && u.Scheme != "https" {
		return nil, errors.Errorf("unsupported URL scheme %q", u.Scheme)
	}
	ctx, cancel := context.WithCancel(context.Background())

	var blockDetailed bool
	if reflect.TypeOf(*new(T)) == typeOfBlockDetailed {
		blockDetailed = true
	}

	out := &Subscriber[T]{
		url:           *u,
		blocksChan:    make(chan *T, 1000),
		ctx:           ctx,
		ctxCancel:     cancel,
		blockDetailed: blockDetailed,
		options: options{
			client:        http.DefaultClient,
			poolingPeriod: defaultPoolingPeriod,
		},
	}
	out.options.apply(opts...)
	return out, nil
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func (s *Subscriber[T]) SetClient(re httpClient) {
	s.client = re
}

func (s *Subscriber[T]) Stop() {
	s.ctxCancel()
}

func (s *Subscriber[T]) checkIfRPCError(buff []byte) error {
	var rpcErr rcpError
	err := json.Unmarshal(buff, &rpcErr)
	if err != nil {
		return errors.Wrapf(err, "failed to unmarshal response body %q", string(buff))
	}
	if rpcErr.Error.Code != 0 && rpcErr.Error.Message != "" {
		return errors.Errorf("rpc error %d: %q", rpcErr.Error.Code, rpcErr.Error.Message)
	}
	return nil
}

func (s *Subscriber[T]) getCurrentBlockNumber() (*big.Int, error) {
	// https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_blocknumber

	req, err := http.NewRequest(
		http.MethodPost,
		s.url.String(),
		strings.NewReader("{\"jsonrpc\":\"2.0\",\"method\":\"eth_blockNumber\",\"params\":[],\"id\":1}"),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request to get current block number")
	}
	req = req.WithContext(s.ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current block number")
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	if err = s.checkIfRPCError(buff); err != nil {
		return nil, err
	}

	var respBody struct {
		ID     int    `json:"id"`
		Result string `json:"result"`
	}

	err = json.Unmarshal(buff, &respBody)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal response body %q", string(buff))
	}
	if respBody.Result == "" {
		return nil, errors.Errorf("unexpected response %q", string(buff))
	}
	result, ok := new(big.Int).SetString(strings.TrimPrefix(respBody.Result, "0x"), 16)
	if !ok {
		return nil, errors.Errorf("failed to parse block number %q", respBody.Result)
	}
	return result, nil
}

func (s *Subscriber[T]) getBlockInfo(blockNum *big.Int) (*T, error) {
	// https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_getBlockByNumber

	body := "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"" + fmt.Sprintf("0x%x", blockNum) + "\", " + fmt.Sprintf("%t", s.blockDetailed) + "],\"id\":1}"
	req, err := http.NewRequest(
		http.MethodPost,
		s.url.String(),
		strings.NewReader(body),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request to get current block number")
	}
	req = req.WithContext(s.ctx)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get current block number")
	}

	buff := bytes.NewBuffer(make([]byte, 0, 1024))

	var respBody struct {
		Result T
	}

	err = json.NewDecoder(io.TeeReader(resp.Body, &limitWriter{w: buff, limit: 1024})).Decode(&respBody)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal response body %q", buff.String())
	}

	if respBody.Result.IsEmpty() {
		if err = s.checkIfRPCError(buff.Bytes()); err != nil {
			return nil, err
		}
		return nil, errors.Errorf("unexpected response %q", buff.String())
	}

	return &respBody.Result, nil
}

var bigIntUno = big.NewInt(1)

// httpStart implements http-based block subscription
func (s *Subscriber[T]) httpStart() error {
	var currentBlock *big.Int
	var err error

	if currentBlock = s.currentBlock.Load(); currentBlock == nil {
		currentBlock, err = s.getCurrentBlockNumber()
		if err != nil {
			return errors.Wrap(err, "failed to get current block number on init")
		}
	}

	s.currentBlock.Store(currentBlock)
	s.running.Store(true)
	go s.httpSubscriberBody()
	return nil
}

func (s *Subscriber[T]) httpSubscriberBody() {
	defer func() {
		close(s.blocksChan)
		s.running.Store(false)
	}()
	timer := time.NewTicker(s.poolingPeriod)
	defer timer.Stop()

	currentBlock := s.currentBlock.Load()
	for {
		select {
		case <-s.ctx.Done():
			return
		case <-timer.C:
		}

		topKnownBlock, err := s.getCurrentBlockNumber()
		if err != nil {
			s.lastError.Store(toPtr(errors.Wrap(err, "failed to get current block number")))
			return
		}

		for ; currentBlock.Cmp(topKnownBlock) < 0; currentBlock.Add(currentBlock, bigIntUno) {
			endBlock := s.endBlock.Load()
			if endBlock.Int64() != 0 && currentBlock.Cmp(endBlock) > 0 {
				return
			}
			block, err := s.getBlockInfo(currentBlock)
			if err != nil {
				s.lastError.Store(toPtr(errors.Wrapf(err, "failed to read block %x info", currentBlock)))
				return
			}
			s.blocksChan <- block
		}

		currentBlock = topKnownBlock
		s.currentBlock.Store(currentBlock)
	}
}

func (s *Subscriber[T]) GetCurrentBlock() big.Int {
	val := s.currentBlock.Load()
	if val == nil {
		return big.Int{}
	}
	return *val
}

func (s *Subscriber[T]) GetBlockChan() <-chan *T {
	return s.blocksChan
}

// wsStart implements websocket-based block subscription
func (s *Subscriber[T]) wsStart() error {
	// https: //docs.infura.io/api/networks/ethereum/json-rpc-methods/subscription-methods/eth_subscribe

	return errors.New("not implemented")
}

func (s *Subscriber[T]) Start() error {
	switch s.url.Scheme {
	case "http", "https":
		return s.httpStart()
	case "ws":
		return s.wsStart()
	default:
		return errors.Errorf("unsupported scheme %q", s.url.Scheme)
	}
}

func (s *Subscriber[T]) LastError() error {
	if err := s.lastError.Load(); err != nil {
		return *err
	}
	return nil
}

func (s *Subscriber[T]) IsRunning() bool {
	return s.running.Load()
}

func toPtr[T any](in T) *T {
	return &in
}

type limitWriter struct {
	w     io.Writer
	limit int
}

func (w *limitWriter) Write(p []byte) (n int, err error) {
	if w.limit == 0 {
		return len(p), nil
	}
	if len(p) > w.limit {
		n, err = w.w.Write(p[:w.limit])
	} else {
		n, err = w.w.Write(p)
	}
	w.limit -= n
	return n, err
}

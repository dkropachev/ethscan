package types

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strings"

	"github.com/pkg/errors"
)

const (
	// HashLength is the expected length of the hash
	HashLength = 32
	// AddressLength is the expected length of the address
	AddressLength = 20
)

type EthAddress [AddressLength]byte

func (a EthAddress) String() string {
	var buf [len(a)*2 + 2]byte
	copy(buf[:2], "0x")
	hex.Encode(buf[2:], a[:])
	return string(buf[:])
}

func (a *EthAddress) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	data = removeHexPrefix(removeQuotes(data))
	if len(data) != AddressLength*2 {
		return errors.Errorf("invalid address length: %d", len(data))
	}
	_, err := hex.Decode((*a)[:], data)
	if err != nil {
		return errors.Wrap(err, "failed to decode hex data")
	}
	return nil
}

func (a EthAddress) MarshalJSON() ([]byte, error) {
	return binSliceToHex(a[:]), nil
}

type EthHash [HashLength]byte

func (h *EthHash) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	data = removeHexPrefix(removeQuotes(data))
	if len(data) != HashLength*2 {
		return errors.Errorf("invalid hash length: %d", len(data))
	}
	_, err := hex.Decode((*h)[:], data)
	if err != nil {
		return errors.Wrap(err, "failed to decode hex data")
	}
	return nil
}

func (h EthHash) MarshalJSON() ([]byte, error) {
	return binSliceToHex(h[:]), nil
}

type BinData []byte

func (d *BinData) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	data = removeHexPrefix(removeQuotes(data))
	out := make([]byte, len(data)/2+1)
	_, err := hex.Decode(out, data)
	if err != nil {
		return errors.Wrap(err, "failed to decode hex data")
	}
	*d = out
	return nil
}

func (d BinData) MarshalJSON() ([]byte, error) {
	return binSliceToHex(d), nil
}

func (d *BinData) Equal(o BinData) bool {
	return len(*d) == len(o) && string(*d) == string(o)
}

type Withdrawal struct {
	Address        EthAddress `json:"address"`
	Amount         BigInt     `json:"amount"`
	Index          BigInt     `json:"index"`
	ValidatorIndex BigInt     `json:"validatorIndex"`
}

// BlockBase represents a block in the Ethereum blockchain.
type BlockBase struct {
	Difficulty      BigInt         `json:"difficulty"`
	TotalDifficulty BigInt         `json:"totalDifficulty"`
	ExtraData       BinData        `json:"extraData"`
	GasLimit        BigInt         `json:"gasLimit"`
	GasUsed         BigInt         `json:"gasUsed"`
	Miner           EthAddress     `json:"miner"`
	Nonce           BigInt         `json:"nonce"`
	Number          BigInt         `json:"number"`
	Hash            EthHash        `json:"hash"`
	MixHash         EthHash        `json:"mixHash"`
	ParentHash      EthHash        `json:"parentHash"`
	ReceiptHash     EthHash        `json:"receiptsRoot"`
	RootHash        EthHash        `json:"stateRoot"`
	TxHash          EthHash        `json:"transactionsRoot"`
	UnclesHash      EthHash        `json:"sha3Uncles"`
	Size            BigInt         `json:"size"`
	Timestamp       BigInt         `json:"timestamp"`
	Transactions    []*Transaction `json:"transactions"`
	Uncles          []EthHash      `json:"uncles"`
	Withdrawals     []Withdrawal   `json:"withdrawals"`
	BaseFeePerGas   BigInt         `json:"baseFeePerGas"`
	BlobGasUsed     BigInt         `json:"blobGasUsed"`
	ExcessBlobGas   BigInt         `json:"excessBlobGas"`
}

func (b BlockBase) IsEmpty() bool {
	return b.Number.AsBigInt().Int64() == 0
}

type Block struct {
	BlockBase
	Transactions []EthHash `json:"transactions"`
}

type BlockDetailed struct {
	BlockBase
	Transactions []*Transaction `json:"transactions"`
}

type BlockType interface {
	Block | BlockDetailed

	IsEmpty() bool
}

type BigInt big.Int

func (b *BigInt) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	i, ok := new(big.Int).SetString(strings.TrimPrefix(s, "0x"), 16)
	if !ok {
		return errors.Errorf("failed to parse %q into big.Int", s)
	}
	*b = BigInt(*i)
	return nil
}

func (b BigInt) MarshalJSON() ([]byte, error) {
	return []byte(b.AsBigInt().String()), nil
}

func (b *BigInt) AsBigInt() *big.Int {
	return (*big.Int)(b)
}

type Transaction struct {
	BlockHash        EthHash    `json:"blockHash"`
	BlockNumber      BigInt     `json:"blockNumber"`
	From             EthAddress `json:"from"`
	Gas              BigInt     `json:"gas"`
	GasPrice         BigInt     `json:"gasPrice"`
	Hash             EthHash    `json:"hash"`
	Input            BinData    `json:"input"`
	Nonce            BigInt     `json:"nonce"`
	To               EthAddress `json:"to"`
	TransactionIndex BigInt     `json:"transactionIndex"`
	Value            BigInt     `json:"value"`
}

func (t *Transaction) Equal(o *Transaction) bool {
	return t.BlockHash == o.BlockHash &&
		t.BlockNumber.AsBigInt().Cmp(o.BlockNumber.AsBigInt()) == 0 &&
		t.From == o.From &&
		t.Gas.AsBigInt().Cmp(o.Gas.AsBigInt()) == 0 &&
		t.GasPrice.AsBigInt().Cmp(o.GasPrice.AsBigInt()) == 0 &&
		t.Hash == o.Hash &&
		t.Input.Equal(o.Input) &&
		t.Nonce.AsBigInt().Cmp(o.Nonce.AsBigInt()) == 0 &&
		t.To == o.To &&
		t.TransactionIndex.AsBigInt().Cmp(o.TransactionIndex.AsBigInt()) == 0 &&
		t.Value.AsBigInt().Cmp(o.Value.AsBigInt()) == 0
}

func removeQuotes(data []byte) []byte {
	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		return data[1 : len(data)-1]
	}
	return data
}

func removeHexPrefix(data []byte) []byte {
	if len(data) >= 2 && data[0] == '0' && (data[1] == 'x' || data[1] == 'X') {
		return data[2:]
	}
	return data
}

func binSliceToHex(data []byte) []byte {
	out := make([]byte, len(data)*2+4)
	out[0] = '"'
	out[1] = '0'
	out[2] = 'x'
	hex.Encode(out[3:], data)
	out[len(out)-1] = '"'
	return out
}

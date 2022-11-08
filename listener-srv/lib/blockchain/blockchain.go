package blockchain

import (
	"github.com/truekupo/cluster/lib/blockchain/chain"
)

type Blockachain interface {
	Symbol() string
	Name() string
	Enabled() bool
	Network() string
	Url() string
	Secret() string
	HeadBlockNumber() (uint64, error)
	GetBlock(uint64) (*chain.Block, error)
	GetHistoryTransactions(string) ([]chain.Transaction, error)
}

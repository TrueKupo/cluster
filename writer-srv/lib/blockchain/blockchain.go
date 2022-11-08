package blockchain

import (
	"github.com/shopspring/decimal"

	"github.com/truekupo/cluster/lib/blockchain/chain"
)

type Blockachain interface {
	Symbol() string
	Name() string
	Enabled() bool
	Network() string
	Url() string
	GetBalanceOf(string) (string, error)
	SendFromTo(string, string, string, decimal.Decimal) (*chain.Transaction, error)
	GetTxStatusByHash(string) (string, error)
}

package blockchain

import (
	"github.com/shopspring/decimal"

	solana "github.com/stafiprotocol/solana-go-sdk/client"
)

type Blockachain interface {
	Symbol() string
	Name() string
	Enabled() bool
	Network() string
	Url() string
	CreateStakeAccount(SignerPrivateBase58 string, Amount decimal.Decimal) (string, string, error)
	DelegateStake(SignerPrivateBase58 string, StakePublicBase58 string, VotePublicBase58 string) (string, error)
	DeactivateStake(SignerPrivateBase58 string, StakePublicBase58 string) (string, error)
	WithdrawStake(SignerPrivateBase58 string, StakePublicBase58 string, Amount decimal.Decimal) (string, error)
	StakeActivationStatus(StakePublicBase58 string) (string, uint64, uint64, error)
	StakeAccountInfo(StakePublicBase58 string) (*solana.StakeAccountRsp, error)
}

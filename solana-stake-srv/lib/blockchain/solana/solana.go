package solana

import (
	"context"
	"fmt"
	"sync"

	"github.com/mr-tron/base58"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	solana "github.com/stafiprotocol/solana-go-sdk/client"
	"github.com/stafiprotocol/solana-go-sdk/common"
	"github.com/stafiprotocol/solana-go-sdk/stakeprog"
	"github.com/stafiprotocol/solana-go-sdk/sysprog"
	"github.com/stafiprotocol/solana-go-sdk/types"

	"github.com/truekupo/cluster/lib/blockchain/chain"
	gc "github.com/truekupo/cluster/lib/config"
	"github.com/truekupo/cluster/lib/logger"
)

const (
	errorSkipped = -32007
)

type Solana struct {
	sync.Mutex

	chain.Chain

	endpoint string
	client   *solana.Client
}

var (
	log *logrus.Entry = nil
)

func New(conf *gc.Chain) (*Solana, error) {
	log = logger.LogModule("solana")

	s := Solana{
		Chain: chain.New(conf),
	}

	s.init_blockchain()

	s.client = solana.NewClient([]string{s.endpoint})

	log.WithField("endpoint", s.endpoint).WithField("net", s.Network()).WithField("url", s.Url()).WithField("symbol", s.Symbol()).Info("initialize blockchain")

	return &s, nil
}

func (s *Solana) init_blockchain() {
	switch s.Network() {
	case "mainnet":
		s.endpoint = solana.MainnetRPCEndpoint
	case "testnet":
		s.endpoint = "https://api.testnet.solana.com" //solana.TestnetRPCEndpoint
	case "devnet":
		s.endpoint = solana.DevnetRPCEndpoint
	default:
		s.endpoint = solana.TestnetRPCEndpoint
	}
}

func (s *Solana) StakeAccountInfo(StakePublicBase58 string) (*solana.StakeAccountRsp, error) {
	r, err := s.client.GetStakeAccountInfo(context.Background(), StakePublicBase58)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Solana) StakeActivationStatus(StakePublicBase58 string) (string, uint64, uint64, error) {
	r, err := s.client.GetStakeActivation(context.Background(), StakePublicBase58, solana.GetStakeActivationConfig{})
	if err != nil {
		return "", 0, 0, err
	}

	return string(r.State), r.Active, r.Inactive, nil
}

func (s *Solana) CreateStakeAccount(SignerPrivateBase58 string, Amount decimal.Decimal) (string, string, error) {
	// Create private key
	private_key_bytes, err := base58.Decode(SignerPrivateBase58)
	if err != nil {
		return "", "", err
	}

	// Create wallet from private key
	auth_wallet := types.AccountFromPrivateKeyBytes(private_key_bytes)

	// Create stake wallet
	stake_wallet := types.NewAccount()

	log.WithField("stake_address", stake_wallet.PublicKey.ToBase58()).Info("CreateStakeAccount")

	// Get blockhash
	res, err := s.client.GetRecentBlockhash(context.Background())
	if err != nil {
		return "", "", err
	}

	// Create Raw Transaction
	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions: []types.Instruction{
			sysprog.CreateAccount(
				auth_wallet.PublicKey,
				stake_wallet.PublicKey,
				common.StakeProgramID,
				uint64(Amount.Shift(9).IntPart()),
				solana.StakeAccountInfoLengthDefault,
			),
			stakeprog.Initialize(
				stake_wallet.PublicKey,
				stakeprog.Authorized{
					Staker:     auth_wallet.PublicKey,
					Withdrawer: auth_wallet.PublicKey,
				},
				stakeprog.Lockup{0, 0, auth_wallet.PublicKey},
			),
		},
		Signers:         []types.Account{auth_wallet, stake_wallet},
		FeePayer:        auth_wallet.PublicKey,
		RecentBlockHash: res.Blockhash,
	})
	if err != nil {
		return "", "", err
	}

	// Send Raw transaction
	txHash, err := s.client.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		return "", "", err
	}

	return txHash, stake_wallet.PublicKey.ToBase58(), nil
}

func (s *Solana) DelegateStake(SignerPrivateBase58 string, StakePublicBase58 string, VotePublicBase58 string) (string, error) {
	// Create private key
	private_key_bytes, err := base58.Decode(SignerPrivateBase58)
	if err != nil {
		return "", err
	}

	// Create wallet from private key
	auth_wallet := types.AccountFromPrivateKeyBytes(private_key_bytes)

	// Get Stake Public Key
	StakePublicKey, err := publicKeyFromBase58(StakePublicBase58)
	if err != nil {
		return "", err
	}

	// Get Vote Public Key
	VotePublicKey, err := publicKeyFromBase58(VotePublicBase58)
	if err != nil {
		return "", err
	}

	// Get blockhash
	res, err := s.client.GetRecentBlockhash(context.Background())
	if err != nil {
		return "", err
	}

	// Create transaction
	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions: []types.Instruction{
			stakeprog.DelegateStake(
				StakePublicKey,
				auth_wallet.PublicKey,
				VotePublicKey,
			),
		},
		Signers:         []types.Account{auth_wallet},
		FeePayer:        auth_wallet.PublicKey,
		RecentBlockHash: res.Blockhash,
	})
	if err != nil {
		return "", err
	}

	// Send Raw transaction
	txHash, err := s.client.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

func (s *Solana) DeactivateStake(SignerPrivateBase58 string, StakePublicBase58 string) (string, error) {
	// Create private key
	private_key_bytes, err := base58.Decode(SignerPrivateBase58)
	if err != nil {
		return "", err
	}

	// Create wallet from private key
	auth_wallet := types.AccountFromPrivateKeyBytes(private_key_bytes)

	// Get Stake Public Key
	StakePublicKey, err := publicKeyFromBase58(StakePublicBase58)
	if err != nil {
		return "", err
	}

	// Get blockhash
	res, err := s.client.GetRecentBlockhash(context.Background())
	if err != nil {
		return "", err
	}

	// Create transaction
	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions: []types.Instruction{
			stakeprog.Deactivate(
				StakePublicKey,
				auth_wallet.PublicKey,
			),
		},
		Signers:         []types.Account{auth_wallet},
		FeePayer:        auth_wallet.PublicKey,
		RecentBlockHash: res.Blockhash,
	})
	if err != nil {
		return "", err
	}

	// Send Raw transaction
	txHash, err := s.client.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

func (s *Solana) WithdrawStake(SignerPrivateBase58 string, StakePublicBase58 string, Amount decimal.Decimal) (string, error) {
	// Create private key
	private_key_bytes, err := base58.Decode(SignerPrivateBase58)
	if err != nil {
		return "", err
	}

	// Create wallet from private key
	auth_wallet := types.AccountFromPrivateKeyBytes(private_key_bytes)

	// Get Stake Public Key
	StakePublicKey, err := publicKeyFromBase58(StakePublicBase58)
	if err != nil {
		return "", err
	}

	// Get blockhash
	res, err := s.client.GetRecentBlockhash(context.Background())
	if err != nil {
		return "", err
	}

	// Create transaction
	rawTx, err := types.CreateRawTransaction(types.CreateRawTransactionParam{
		Instructions: []types.Instruction{
			stakeprog.Withdraw(
				StakePublicKey,
				auth_wallet.PublicKey,
				auth_wallet.PublicKey,
				uint64(Amount.Shift(9).IntPart()),
				auth_wallet.PublicKey,
			),
		},
		Signers:         []types.Account{auth_wallet},
		FeePayer:        auth_wallet.PublicKey,
		RecentBlockHash: res.Blockhash,
	})
	if err != nil {
		return "", err
	}

	// Send Raw transaction
	txHash, err := s.client.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

func publicKeyFromBase58(in string) (out common.PublicKey, err error) {
	val, err := base58.Decode(in)
	if err != nil {
		return out, fmt.Errorf("decode: %w", err)
	}

	if len(val) != 32 {
		return out, fmt.Errorf("invalid length, expected 32, got %d", len(val))
	}

	copy(out[:], val)
	return
}

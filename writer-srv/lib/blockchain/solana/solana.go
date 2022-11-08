package solana

import (
	"context"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	ee "github.com/truekupo/cluster/lib/errors"
	"github.com/truekupo/cluster/lib/logger"

	"github.com/truekupo/cluster/lib/blockchain/chain"
	gc "github.com/truekupo/cluster/lib/config"
)

type Solana struct {
	chain.Chain

	endpoint string
	client   *rpc.Client
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

	s.client = rpc.New(s.endpoint)

	log.WithField("net", s.Network()).WithField("url", s.Url()).WithField("symbol", s.Symbol()).Info("initialize blockchain")

	return &s, nil
}

func (s *Solana) init_blockchain() {
	switch s.Network() {
	case "mainnet":
		s.endpoint = rpc.MainNetBeta_RPC
	case "testnet":
		s.endpoint = rpc.TestNet_RPC
	default:
		s.endpoint = rpc.TestNet_RPC
	}
}

func (s *Solana) GetBalanceOf(Addr string) (string, error) {
	address, err := solana.PublicKeyFromBase58(Addr)
	if err != nil {
		return "", ee.ErrBadRequest
	}

	resp, err := s.client.GetBalance(context.Background(), address, rpc.CommitmentConfirmed)
	if err != nil {
		return "", ee.ErrNotFound
	}

	return decimal.NewFromInt(int64(resp.Value)).Shift(-9).String(), nil
}

func (s *Solana) SendFromTo(PublicAddrFrom string, PublicAddrTo string, PrivateAddrFrom string, Amount decimal.Decimal) (*chain.Transaction, error) {
	wallet, err := solana.WalletFromPrivateKeyBase58(PrivateAddrFrom)
	if err != nil {
		return nil, ee.ErrBadRequest
	}

	recent, err := s.client.GetRecentBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return nil, ee.ErrInternal
	}

	public_from, err := solana.PublicKeyFromBase58(PublicAddrFrom)
	if err != nil {
		return nil, ee.ErrBadRequest
	}

	public_to, err := solana.PublicKeyFromBase58(PublicAddrTo)
	if err != nil {
		return nil, ee.ErrBadRequest
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				uint64(Amount.Shift(9).IntPart()),
				public_from,
				public_to).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(public_from),
	)

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			return &wallet.PrivateKey
		},
	)

	sig, err := s.client.SendTransactionWithOpts(context.Background(), tx, false, rpc.CommitmentFinalized)
	if err != nil {
		return nil, ee.ErrInternal
	}

	return &chain.Transaction{
		Hash:      sig.String(),
		CreatedAt: uint64(time.Now().Unix()),
		BlockNum:  0,
		From:      PublicAddrFrom,
		To:        PublicAddrTo,
		Amount:    Amount,
		Fee:       decimal.NewFromInt(0),
		Direction: chain.TransactionDirectionOut,
		Status:    chain.TransactionStatusNew,
	}, nil
}

func (s *Solana) GetTxByHash(TxHash string) (*chain.Transaction, error) {
	/*
		sig, err := solana.SignatureFromBase58(TxHash)
		if err != nil {
			return nil, e.ErrBadRequest
		}

		out, err := s.client.GetTransaction(context.Background(), sig, &rpc.GetTransactionOpts{Encoding: solana.EncodingBase64})
		if err != nil {
			return nil, e.ErrNotFound
		}
	*/

	return nil, ee.ErrNotFound
}

func (s *Solana) GetTxStatusByHash(TxHash string) (string, error) {
	sig, err := solana.SignatureFromBase58(TxHash)
	if err != nil {
		return "", ee.ErrBadRequest
	}

	out, err := s.client.GetTransaction(context.Background(), sig,
		&rpc.GetTransactionOpts{
			Encoding: solana.EncodingBase58,
		},
	)

	if err != nil {
		return "", ee.ErrNotFound
	}

	if out.Meta == nil {
		return "", ee.ErrInternal
	}

	if out.Meta.Err != nil {
		return chain.TransactionStatusFailed, nil
	}

	return chain.TransactionStatusFinalized, nil
}

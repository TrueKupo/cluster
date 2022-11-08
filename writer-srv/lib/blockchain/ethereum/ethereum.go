package ethereum

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	ee "github.com/truekupo/cluster/lib/errors"
	"github.com/truekupo/cluster/lib/logger"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/truekupo/cluster/lib/blockchain/chain"
	gc "github.com/truekupo/cluster/lib/config"
	eu "github.com/truekupo/cluster/lib/utils/ethereum"
)

type Ethereum struct {
	chain.Chain

	hdnet  *chaincfg.Params
	client *ethclient.Client
}

var (
	log *logrus.Entry = nil
)

func New(conf *gc.Chain) (*Ethereum, error) {
	log = logger.LogModule("ethereum")

	e := &Ethereum{
		Chain: chain.New(conf),
	}

	e.init_blockchain()

	client, err := ethclient.Dial(e.Url())
	if err != nil {
		return nil, err
	}

	e.client = client

	log.WithField("net", e.Network()).WithField("url", e.Url()).WithField("symbol", e.Symbol()).Info("initialize blockchain")

	return e, nil
}

func (e *Ethereum) init_blockchain() {
	switch e.Network() {
	case "mainnet":
		e.hdnet = &chaincfg.MainNetParams
	case "testnet":
		e.hdnet = &chaincfg.TestNet3Params
	default:
		e.hdnet = &chaincfg.TestNet3Params
	}
}

func (e *Ethereum) GetBalanceOf(Addr string) (string, error) {
	account := common.HexToAddress(Addr)

	balance, err := e.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return "", ee.ErrNotFound
	}

	return decimal.NewFromBigInt(balance, -18).String(), nil
}

func (e *Ethereum) SendFromTo(PublicAddrFrom string, PublicAddrTo string, PrivateAddrFrom string, Amount decimal.Decimal) (*chain.Transaction, error) {
	//Amount = Amount.Shift(18)

	to_address := common.HexToAddress(PublicAddrTo)
	from_address := common.HexToAddress(PublicAddrFrom)

	nonce, err := e.client.PendingNonceAt(context.Background(), from_address)
	if err != nil {
		return nil, ee.ErrInternal
	}

	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, ee.ErrInternal
	}

	gasLimit := uint64(21000)
	tx := types.NewTransaction(nonce, to_address, Amount.Shift(18).BigInt(), gasLimit, gasPrice, []byte{})

	privateKey, err := eu.PrivateKeyHexToECDSA(PrivateAddrFrom)
	if err != nil {
		return nil, ee.ErrBadRequest
	}

	signer := types.HomesteadSigner{}

	// Sign the transaction
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return nil, ee.ErrInternal
	}

	sender, err := types.Sender(signer, signedTx)
	if err != nil {
		return nil, ee.ErrInternal
	}

	if sender != from_address {
		return nil, ee.ErrInternal
	}

	// Send the transaction
	err = e.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, ee.ErrInternal
	}

	return &chain.Transaction{
		Hash:      signedTx.Hash().Hex(),
		CreatedAt: uint64(time.Now().Unix()),
		BlockNum:  0,
		From:      PublicAddrFrom,
		To:        PublicAddrTo,
		Amount:    Amount,
		Fee:       decimal.NewFromInt(int64(gasLimit) * gasPrice.Int64()).Shift(-18),
		Direction: chain.TransactionDirectionOut,
		Status:    chain.TransactionStatusNew,
	}, nil
}

func (e *Ethereum) GetTxStatusByHash(TxHash string) (string, error) {
	// Get transaction
	_, pending, err := e.client.TransactionByHash(context.Background(), common.HexToHash(TxHash))
	if err != nil {
		if err == ethereum.NotFound {
			return chain.TransactionStatusNew, ee.ErrNotFound
		}
		return chain.TransactionStatusNew, ee.ErrInternal
	}
	if pending {
		return chain.TransactionStatusNew, nil
	}

	// Get receipt
	receipt, err := e.client.TransactionReceipt(context.Background(), common.HexToHash(TxHash))
	if err != nil {
		log.Error(err)
		return chain.TransactionStatusProcessing, ee.ErrInternal
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return chain.TransactionStatusFailed, nil
	}

	// Get head block number
	head_block_number, err := e.client.BlockNumber(context.Background())
	if err != nil {
		return chain.TransactionStatusNew, ee.ErrInternal
	}

	if receipt.BlockNumber == nil {
		return chain.TransactionStatusNew, nil
	}

	// Check for stable
	if receipt.BlockNumber.Uint64()+2 >= head_block_number {
		return chain.TransactionStatusProcessing, nil
	}

	return chain.TransactionStatusFinalized, nil
}

func (e *Ethereum) GetTxByHash(TxHash string) (*chain.Transaction, error) {
	/*
		tx, pending, err := e.client.TransactionByHash(context.Background(), common.HexToHash(TxHash))
		if err != nil {
			return nil, ee.ErrNotFound
		}

		msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()), big.NewInt(0))
		if err != nil {
			return nil, err
		}

		to := ""
		if msg.To() != nil {
			to = msg.To().String()
		}

		amount := decimal.NewFromBigInt(tx.Value(), 0).Shift(-18)
		fee := decimal.NewFromBigInt(tx.Cost(), 0).Sub(decimal.NewFromBigInt(tx.Value(), 0)).Shift(-18)
		status := chain.TransactionStatusNew

		if !pending {
			receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(TxHash))
			if err != nil {
				return nil, err
			}
		}

		return &chain.Transaction{
			Hash:      tx.Hash().Hex(),
			CreatedAt: tx.time,
			BlockNum:  0,
			From:      msg.From().String(),
			To:        to,
			Amount:    amount,
			Fee:       fee,
			//Direction: chain.TransactionDirectionOut,
			Status: status,
		}, nil
	*/

	return nil, ee.ErrInternal
}

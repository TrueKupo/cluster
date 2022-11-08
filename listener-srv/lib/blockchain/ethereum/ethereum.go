package ethereum

import (
	"context"
	"math/big"
	"sync"
	"time"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nanmu42/etherscan-api"

	"github.com/truekupo/cluster/lib/blockchain/chain"
	gc "github.com/truekupo/cluster/lib/config"
	"github.com/truekupo/cluster/lib/logger"
)

type Ethereum struct {
	sync.Mutex

	chain.Chain

	hdnet  *chaincfg.Params
	client *ethclient.Client

	etherscan_client *etherscan.Client
	etherscan_net    etherscan.Network

	headBlockNum uint64
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

	//
	client, err := ethclient.Dial(e.Url())
	if err != nil {
		return nil, err
	}

	e.client = client

	// Connect to etherscan.io
	e.etherscan_client = etherscan.New(e.etherscan_net, e.Chain.Secret())

	go e.process_blockchain()

	log.WithField("net", e.Network()).WithField("etherscan_net", e.etherscan_net).WithField("url", e.Url()).WithField("symbol", e.Symbol()).Info("initialize blockchain")

	return e, nil
}

func (e *Ethereum) init_blockchain() {
	e.etherscan_net = etherscan.Ropsten

	switch e.Network() {
	case "mainnet":
		e.hdnet = &chaincfg.MainNetParams
		e.etherscan_net = etherscan.Mainnet
	case "testnet":
		e.hdnet = &chaincfg.TestNet3Params
	default:
		e.hdnet = &chaincfg.TestNet3Params
	}
}

func (e *Ethereum) process_blockchain() {
	for {
		time.Sleep(5 * time.Second)

		block_num, err := e.headBlockNumber()
		if err != nil {
			continue
		}

		e.Lock()
		e.headBlockNum = block_num
		e.Unlock()
	}
}

func (e *Ethereum) GetBlock(BlockNum uint64) (*chain.Block, error) {
	block, err := e.client.BlockByNumber(context.Background(), big.NewInt(int64(BlockNum)))
	if err != nil {
		return nil, err
	}

	bb := chain.NewBlock(BlockNum, "eth", block.Time())

	chainID, err := e.client.NetworkID(context.Background())
	if err != nil {
		return nil, err
	}

	head_block, _ := e.HeadBlockNumber()

	signer := types.LatestSignerForChainID(chainID)

	for _, tx := range block.Transactions() {
		value := decimal.NewFromBigInt(tx.Value(), -18)
		cost := decimal.NewFromBigInt(tx.Cost(), -18)
		status := chain.TransactionStatusFinalized

		if head_block-3 < BlockNum {
			status = chain.TransactionStatusProcessing
		}

		base_tx := chain.Transaction{
			CreatedAt: block.Time(),
			Hash:      tx.Hash().Hex(),
			From:      "",
			To:        "",
			Status:    status,
			Amount:    value,
			Fee:       cost.Sub(value),
		}
		if tx.To() != nil {
			base_tx.To = tx.To().Hex()
		}

		sender, err := signer.Sender(tx)
		if err == nil {
			base_tx.From = sender.Hex()
		}

		bb.AddTx(base_tx)
	}

	return bb, nil
}

func (e *Ethereum) HeadBlockNumber() (uint64, error) {
	e.Lock()
	defer e.Unlock()

	return e.headBlockNum, nil
}

func (e *Ethereum) headBlockNumber() (uint64, error) {
	header, err := e.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}

	return header.Number.Uint64(), nil
}

func (e *Ethereum) GetHistoryTransactions(Address string) ([]chain.Transaction, error) {
	//log.WithField("address", Address).WithField("chain", "ETH").Info("GetHistoryTransactions")

	txs, err := e.etherscan_client.NormalTxByAddress(Address, nil, nil, 0, 0, true)
	if err != nil {
		return nil, err
	}

	head_block, _ := e.HeadBlockNumber()

	rtxs := make([]chain.Transaction, 0, len(txs))
	for _, tx := range txs {
		value := decimal.NewFromBigInt(tx.Value.Int(), -18)
		gas_price := decimal.NewFromBigInt(tx.GasPrice.Int(), -18)
		status := chain.TransactionStatusFinalized

		if head_block-3 < uint64(tx.BlockNumber) {
			status = chain.TransactionStatusProcessing
		}

		if tx.IsError != 0 {
			status = chain.TransactionStatusFailed
		}

		rtxs = append(rtxs, chain.Transaction{
			CreatedAt: uint64(tx.TimeStamp.Time().Unix()),
			BlockNum:  uint64(tx.BlockNumber),
			Hash:      tx.Hash,
			From:      tx.From,
			To:        tx.To,
			Status:    status,
			Amount:    value,
			Fee:       gas_price.Mul(decimal.NewFromInt(int64(tx.Gas))),
		})
	}

	return rtxs, nil
}

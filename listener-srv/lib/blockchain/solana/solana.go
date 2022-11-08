package solana

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/rpc/jsonrpc"

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
	client   *rpc.Client

	headBlockNum uint64
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

	go s.process_blockchain()

	log.WithField("net", s.Network()).WithField("url", s.Url()).WithField("symbol", s.Symbol()).Info("initialize blockchain")

	return &s, nil
}

func (s *Solana) init_blockchain() {
	switch s.Network() {
	case "mainnet":
		s.endpoint = rpc.MainNetBeta_RPC
	case "testnet":
		s.endpoint = rpc.TestNet_RPC
	case "devnet":
		s.endpoint = rpc.DevNet_RPC
	default:
		s.endpoint = rpc.TestNet_RPC
	}
}

func (s *Solana) process_blockchain() {
	for {
		time.Sleep(5 * time.Second)

		block_num, err := s.headBlockNumber()
		if err != nil {
			continue
		}

		s.Lock()
		s.headBlockNum = block_num
		s.Unlock()
	}
}

func (s *Solana) headBlockNumber() (uint64, error) {
	blockhash, err := s.client.GetRecentBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return 0, err
	}

	return blockhash.Context.Slot, nil
}

func (s *Solana) HeadBlockNumber() (uint64, error) {
	s.Lock()
	defer s.Unlock()

	return s.headBlockNum, nil
}

func (s *Solana) GetBlock(BlockNum uint64) (*chain.Block, error) {
	includeRewards := false
	block, err := s.client.GetBlockWithOpts(context.Background(), BlockNum,
		&rpc.GetBlockOpts{
			Encoding:           solana.EncodingBase64,
			Commitment:         rpc.CommitmentFinalized,
			TransactionDetails: rpc.TransactionDetailsFull,
			Rewards:            &includeRewards,
		},
	)
	if err != nil {
		rpcError, ok := err.(*jsonrpc.RPCError)
		if ok && rpcError.Code == errorSkipped {
			return chain.NewBlock(BlockNum, "sol", uint64(0)), nil
		}
		return nil, err
	}

	bb := chain.NewBlock(BlockNum, "sol", uint64(*block.BlockTime))

	// Parse block transactions
	for _, twm := range block.Transactions {
		btx, err := s.IsTransferTransactionWithMeta(twm, uint64(*block.BlockTime), uint64(BlockNum))
		if err != nil {
			continue
		}

		bb.AddTx(btx)
	}

	return bb, nil
}

func (s *Solana) IsTransferTransactionWithMeta(twm rpc.TransactionWithMeta, utm uint64, slot uint64) (chain.Transaction, error) {
	tx, err := twm.GetTransaction()
	if err != nil {
		return chain.Transaction{}, err
	}

	return s.IsTransferTransaction(tx, twm.Meta, utm, slot)
}

func (s *Solana) IsTransferTransactionResult(gtr rpc.GetTransactionResult, utm uint64, slot uint64) (chain.Transaction, error) {
	tx := new(solana.Transaction)
	err := tx.UnmarshalWithDecoder(bin.NewBinDecoder(gtr.Transaction.GetBinary()))
	if err != nil {
		return chain.Transaction{}, err
	}

	return s.IsTransferTransaction(tx, gtr.Meta, utm, slot)
}

func (s *Solana) IsTransferTransaction(tx *solana.Transaction, meta *rpc.TransactionMeta, utm uint64, slot uint64) (chain.Transaction, error) {
	for _, inst := range tx.Message.Instructions {
		progKey, err := tx.ResolveProgramIDIndex(inst.ProgramIDIndex)
		if err != nil {
			continue
		}

		if progKey != solana.SystemProgramID {
			continue
		}

		accounts := inst.ResolveInstructionAccounts(&tx.Message)

		instruction, err := system.DecodeInstruction(accounts, inst.Data)
		if err != nil {
			break
		}

		if instruction.TypeID.Uint32() == system.Instruction_Transfer {
			transfer_instruction, ok := instruction.Impl.(*system.Transfer)
			if !ok {
				break
			}

			tx_status := chain.TransactionStatusFinalized
			if meta.Err != nil {
				tx_status = chain.TransactionStatusFailed
			}

			base_tx := chain.Transaction{
				CreatedAt: utm,
				BlockNum:  slot,
				Hash:      tx.Signatures[0].String(),
				From:      transfer_instruction.AccountMetaSlice[0].PublicKey.String(),
				To:        transfer_instruction.AccountMetaSlice[1].PublicKey.String(),
				Status:    tx_status,
				Amount:    decimal.NewFromInt(int64(*transfer_instruction.Lamports)).Shift(-9),
				Fee:       decimal.NewFromInt(int64(meta.Fee)).Shift(-9),
			}

			return base_tx, nil
		}
	}

	return chain.Transaction{}, fmt.Errorf("!transfer transaction")
}

func (s *Solana) GetHistoryTransactions(Address string) ([]chain.Transaction, error) {
	addr, err := solana.PublicKeyFromBase58(Address)
	if err != nil {
		return nil, err
	}

	signatures, err := s.client.GetSignaturesForAddress(context.Background(), addr)
	if err != nil {
		return nil, err
	}

	trxs := []chain.Transaction{}

	for _, sig := range signatures {
		out, err := s.client.GetTransaction(context.Background(), sig.Signature, &rpc.GetTransactionOpts{
			Encoding: solana.EncodingBase64,
		})
		if err != nil {
			continue
		}

		btx, err := s.IsTransferTransactionResult(*out, uint64(*sig.BlockTime), sig.Slot)
		if err != nil {
			continue
		}

		trxs = append(trxs, btx)
	}

	log.WithField("address", Address).WithField("chain", "SOL").WithField("count", len(signatures)).Info("GetHistoryTransactions")

	return trxs, nil
}

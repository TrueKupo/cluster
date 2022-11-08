package listener

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gocraft/dbr/v2"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	"github.com/truekupo/cluster/lib/blockchain/chain"
	"github.com/truekupo/cluster/lib/cluster"
	"github.com/truekupo/cluster/listener-srv/lib/blockchain"
	"github.com/truekupo/cluster/listener-srv/lib/blockchain/ethereum"
	"github.com/truekupo/cluster/listener-srv/lib/blockchain/solana"
	"github.com/truekupo/cluster/listener-srv/lib/config"
	"github.com/truekupo/cluster/listener-srv/lib/db/middleware"
	"github.com/truekupo/cluster/listener-srv/lib/db/models"
	"github.com/truekupo/cluster/listener-srv/lib/db/repo"

	"github.com/truekupo/cluster/lib/logger"
)

type ListenerHandler struct {
	sync.Mutex

	conf     *config.Config
	dbs      *dbr.Session
	init     bool
	duration time.Duration

	b blockchain.Blockachain

	w *AddressWatcher

	client *cluster.Cluster
}

var (
	log *logrus.Entry = nil
)

func NewListenerHandler(conf *config.Config) (*ListenerHandler, error) {
	log = logger.LogModule("listener-handler")

	h := ListenerHandler{
		conf:   conf,
		dbs:    nil,
		init:   false,
		client: cluster.New(&conf.Cluster),
		w:      NewAddressWatcher(),
	}

	// init blockchain
	cc := conf.Chain

	switch cc.Symbol {
	case "ETH":
		e, err := ethereum.New(&cc)
		if err != nil {
			log.WithField("error", err).Error("ethereum:new")
		}

		h.b = e
	case "SOL":
		s, err := solana.New(&cc)
		if err != nil {
			log.WithField("error", err).Error("solana:new")
		}

		h.b = s
	default:
		panic("No active blockchain")
	}

	var err error
	// set request duration
	h.duration, err = time.ParseDuration(cc.RequestTm)
	if err != nil {
		switch cc.Symbol {
		case "ETH":
			h.duration = time.Second * 3
		case "SOL":
			h.duration = time.Millisecond * 300
		}
	}

	// connect to DB
	dsd, err := middleware.NewDSD(conf)
	if err != nil {
		log.Error("listener:handler:new ", err)
		return nil, err
	}

	h.dbs = dsd.Conn.NewSession(nil)

	// load watched addresses
	err = h.loadWatchList()
	if err != nil {
		return nil, err
	}

	// watch handler
	go h.listener_process()

	// monitoring transactions status
	go h.status_watcher_process()

	// transactions history
	go h.history_process()

	return &h, nil
}

func (h *ListenerHandler) IsInit() bool {
	h.Lock()
	defer h.Unlock()

	return h.isInit()
}

func (h *ListenerHandler) isInit() bool {
	return h.init
}

func (h *ListenerHandler) loadWatchList() error {
	as, err := repo.Address(h.dbs).FindActive()
	if err != nil {
		return err
	}

	h.Lock()
	defer h.Unlock()

	for _, a := range as {
		h.w.Add(a.Addr, a.Id, a.AccountId)
	}

	h.init = true

	return nil
}

func (h *ListenerHandler) listener_process() {
	for {
		// Get block number from db
		block_num, err := repo.Block(h.dbs).BlockNumBySymbol(h.b.Symbol())
		if err != nil {
			log.Error(err)
		}

		// Get head block number
		head_block_num, err := h.b.HeadBlockNumber()
		if err != nil {
			log.Error(err)
		}

		if err == nil {
			log.WithField("block_num", block_num).WithField("head_block_num", head_block_num).Debug("listener_process")

			if block_num == 0 {
				block_num = head_block_num - 10
			}

			for i := block_num + 1; i <= head_block_num; i++ {
				log.WithField("num", i).Debug("process new block")

				// Get Next Block
				block, err := h.b.GetBlock(i)
				if err != nil {
					log.WithField("error", err.Error()).Error("GetBlock")
					break
				}

				h.process_transactions(block.CreatedAt, block.Num, block.TxList)

				// Update block number in DB
				err = repo.Block(h.dbs).UpdateBlockNumBySymbol(h.b.Symbol(), i)
				if err != nil {
					log.WithField("error", err.Error()).Error("UpdateBlockNumBySymbol")
					break
				}
			}

			time.Sleep(h.duration)
		}

		time.Sleep(3 * time.Second)
	}
}

func (h *ListenerHandler) status_watcher_process() {
	for {
		// Get transactions with processing status
		txs, err := repo.Tx(h.dbs).FindByStatus("Processing")
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}

		for _, tx := range txs {
			status, err := h.client.FetchTransactionStatus(h.b.Symbol(), tx.Hash)
			if err != nil {
				log.WithField("error", err.Error()).Error("FetchTransactionStatus")
				continue
			}

			repo.Tx(h.dbs).UpdateStatus(tx.Id, status)
		}

		time.Sleep(3 * time.Second)
	}
}

func (h *ListenerHandler) history_process() {
	for {
		as, err := repo.Address(h.dbs).GetAddressesWithHistoryNotRequested()
		if err == nil {
			for _, a := range as {
				TxList, err := h.b.GetHistoryTransactions(a.Addr)
				if err != nil {
					log.WithField("addr", a.Addr).WithField("error", err.Error()).Error("GetHistoryTransactions")
					continue
				}

				log.WithField("addr", a.Addr).WithField("count", len(TxList)).Info("GetHistoryTransactions")

				for _, tx := range TxList {
					h.ProcessTransaction(tx.CreatedAt, tx.BlockNum, &tx)
				}

				repo.Address(h.dbs).SetHistoryRequested(a.Id)
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func (h *ListenerHandler) process_transactions(TmBlockCreatedAt uint64, BlockNum uint64, TxList []chain.Transaction) {
	log.WithField("block_num", BlockNum).WithField("tx_count", len(TxList)).Debug("process_transactions")

	for _, tx := range TxList {
		h.ProcessTransaction(TmBlockCreatedAt, BlockNum, &tx)
	}

}

func (h *ListenerHandler) ProcessTransaction(TmBlockCreatedAt uint64, BlockNum uint64, Tx *chain.Transaction) {
	log.WithField("block_num", BlockNum).WithField("tx_hash", Tx.Hash).Debug("process_transaction")

	from := strings.ToLower(Tx.From)
	to := strings.ToLower(Tx.To)

	h.Lock()
	from_address_id, from_account_id, from_err := h.w.Get(from)
	to_address_id, to_account_id, to_err := h.w.Get(to)
	h.Unlock()

	if from_err != nil && to_err != nil {
		return
	}

	if (from_err != nil && to_account_id == 0) || (to_err != nil && from_account_id == 0) {
		return
	}

	if (to_account_id == 0) && (from_account_id == 0) {
		return
	}

	direction := chain.TransactionDirectionIn
	if from_account_id != 0 {
		direction = chain.TransactionDirectionOut
	}

	if from_account_id != 0 && to_account_id != 0 {
		direction = chain.TransactionDirectionInOut
	}

	if from_err != nil {
		a, err := repo.Address(h.dbs).Insert(&models.Address{
			AccountId:      0,
			Addr:           from,
			RequestHistory: true,
		})
		if err != nil {
			log.Error(err)
			return
		}

		from_address_id = a.Id

		h.w.Add(Tx.From, from_address_id, 0)
	}

	if to_err != nil {
		a, err := repo.Address(h.dbs).Insert(&models.Address{
			AccountId:      0,
			Addr:           to,
			RequestHistory: true,
		})
		if err != nil {
			log.Error(err)
			return
		}

		to_address_id = a.Id

		h.w.Add(Tx.To, to_address_id, 0)
	}

	_, err := repo.Tx(h.dbs).InsertOrUpdate(&models.Tx{
		CreatedAt:  TmBlockCreatedAt,
		Hash:       Tx.Hash,
		BlockNum:   BlockNum,
		Direction:  direction,
		FromAddrId: from_address_id,
		ToAddrId:   to_address_id,
		Status:     Tx.Status,
		Amount:     Tx.Amount.String(),
		Fee:        Tx.Fee.String(),
	})
	if err != nil {
		log.Error(err)
	}

	log.WithField("TxHash", Tx.Hash).WithField("From", from).WithField("To", to).WithField("Amount", Tx.Amount).WithField("Fee", Tx.Fee).Info("Insert Tx")
}

func (h *ListenerHandler) AddWatchAddress(AccountUuid string, Address string) error {
	if !h.IsInit() {
		return fmt.Errorf("Listener not init")
	}

	new_account_id, err := repo.Account(h.dbs).GetOrCreateAccountId(AccountUuid)
	if err != nil {
		return err
	}

	address := strings.ToLower(Address)

	h.Lock()
	address_id, account_id, err := h.w.Get(address)
	h.Unlock()

	if err == nil {
		if account_id == new_account_id {
			return nil
		}

		err = repo.Address(h.dbs).UpdateById(&models.Address{
			Id:             address_id,
			AccountId:      new_account_id,
			Addr:           address,
			RequestHistory: false,
		})
		if err != nil {
			return err
		}

		return h.w.Add(address, address_id, new_account_id)
	}

	address_id, err = repo.Address(h.dbs).GetOrCreateAddressId(new_account_id, address, false)
	if err != nil {
		return err
	}

	h.Lock()
	defer h.Unlock()

	return h.w.Add(address, address_id, new_account_id)
}

func (h *ListenerHandler) DeleteWatchAddress() error {
	if !h.IsInit() {
		return fmt.Errorf("Listener not init")
	}

	// TODO:

	h.Lock()
	defer h.Unlock()

	return nil
}

func (h *ListenerHandler) TransactionsByAddress(Address string, From int32, Limit int32) ([]chain.Transaction, error) {
	if !h.IsInit() {
		return nil, fmt.Errorf("Listener not init")
	}

	address := strings.ToLower(Address)

	address_id, err := repo.Address(h.dbs).FindIdByAddress(address)
	if err != nil {
		return nil, err
	}

	txs, err := repo.Tx(h.dbs).FindByAddressId(address_id, From, Limit)
	if err != nil {
		return nil, err
	}

	transactions := make([]chain.Transaction, 0, len(txs))
	for _, tx := range txs {
		amount, err := decimal.NewFromString(tx.Amount)
		if err != nil {
			continue
		}

		fee, err := decimal.NewFromString(tx.Fee)
		if err != nil {
			continue
		}

		transactions = append(transactions, chain.Transaction{
			CreatedAt: tx.CreatedAt,
			Hash:      tx.Hash,
			From:      repo.Address(h.dbs).AddressById(tx.FromAddrId),
			To:        repo.Address(h.dbs).AddressById(tx.ToAddrId),
			Direction: tx.Direction,
			Status:    h.realStatus(&tx), // tx.Status,
			Amount:    amount,
			Fee:       fee,
		})
	}

	return transactions, nil
}

func (h *ListenerHandler) realStatus(tx *models.Tx) string {
	if tx.Status != chain.TransactionStatusProcessing {
		return tx.Status
	}

	head_block, _ := h.b.HeadBlockNumber()
	if head_block > tx.BlockNum+3 {
		return chain.TransactionStatusFinalized
	}

	return tx.Status
}

func (h *ListenerHandler) TransactionsByAccount(Uuid string, From int32, Limit int32) ([]chain.Transaction, error) {
	if !h.IsInit() {
		return nil, fmt.Errorf("Listener not init")
	}

	account, err := repo.Account(h.dbs).FindByUuid(Uuid)
	if err != nil {
		log.WithField("uuid", Uuid).WithField("error", err).Error("FindByUuid")
		return nil, err
	}

	address_ids, err := repo.Address(h.dbs).FindAddressIdsByAccount(account.Id)
	if err != nil {
		log.WithField("account_id", account.Id).WithField("error", err).Error("FindAddressIdsByAccount")
		return nil, err
	}

	txs, err := repo.Tx(h.dbs).FindByAddressIds(address_ids, From, Limit)
	if err != nil {
		log.WithField("address_ids", address_ids).WithField("error", err).Error("FindByAddressIds")
		return nil, err
	}

	transactions := make([]chain.Transaction, 0, len(txs))
	for _, tx := range txs {
		amount, err := decimal.NewFromString(tx.Amount)
		if err != nil {
			continue
		}

		fee, err := decimal.NewFromString(tx.Fee)
		if err != nil {
			continue
		}

		transactions = append(transactions, chain.Transaction{
			CreatedAt: tx.CreatedAt,
			Hash:      tx.Hash,
			From:      repo.Address(h.dbs).AddressById(tx.FromAddrId),
			To:        repo.Address(h.dbs).AddressById(tx.ToAddrId),
			Direction: tx.Direction,
			Status:    h.realStatus(&tx),
			Amount:    amount,
			Fee:       fee,
		})
	}

	return transactions, nil
}

func (h *ListenerHandler) AddTransaction(Tx chain.Transaction) error {
	if !h.IsInit() {
		return fmt.Errorf("Listener not init")
	}

	h.ProcessTransaction(uint64(time.Now().Unix()), 0, &Tx)

	return nil
}

func (h *ListenerHandler) GetTxByHash(TxHash string) (*chain.Transaction, error) {
	if !h.IsInit() {
		return nil, fmt.Errorf("Listener not init")
	}

	tx, err := repo.Tx(h.dbs).FindByTxHash(TxHash)
	if err != nil {
		return nil, err
	}

	amount, _ := decimal.NewFromString(tx.Amount)
	fee, _ := decimal.NewFromString(tx.Fee)

	return &chain.Transaction{
		CreatedAt: tx.CreatedAt,
		Hash:      tx.Hash,
		From:      repo.Address(h.dbs).AddressById(tx.FromAddrId),
		To:        repo.Address(h.dbs).AddressById(tx.ToAddrId),
		Direction: tx.Direction,
		Status:    h.realStatus(tx),
		Amount:    amount,
		Fee:       fee,
	}, nil
}

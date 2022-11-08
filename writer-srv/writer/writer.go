package writer

import (
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	"github.com/truekupo/cluster/lib/blockchain/chain"
	"github.com/truekupo/cluster/lib/cluster"
	"github.com/truekupo/cluster/writer-srv/lib/blockchain"
	"github.com/truekupo/cluster/writer-srv/lib/blockchain/ethereum"
	"github.com/truekupo/cluster/writer-srv/lib/blockchain/solana"
	"github.com/truekupo/cluster/writer-srv/lib/config"
	"github.com/truekupo/cluster/writer-srv/lib/db/middleware"

	e "github.com/truekupo/cluster/lib/errors"
	"github.com/truekupo/cluster/lib/logger"
)

type WriterHandler struct {
	conf *config.Config

	dsd    *middleware.DSD
	client *cluster.Cluster

	b map[string]blockchain.Blockachain
}

var (
	log *logrus.Entry = nil
)

func NewWriterHandler(conf *config.Config) (*WriterHandler, error) {
	log = logger.LogModule("writer-handler")

	a := WriterHandler{
		conf:   conf,
		dsd:    nil,
		client: cluster.New(&conf.Cluster),
		b:      map[string]blockchain.Blockachain{},
	}

	for _, cc := range conf.Settings.Chains {
		switch cc.Symbol {
		case "ETH":
			e, err := ethereum.New(&cc)
			if err != nil {
				log.WithField("error", err).Error("ethereum:new")
				continue
			}

			a.b[cc.Symbol] = e
			a.b[cc.Name] = e
		case "SOL":
			s, err := solana.New(&cc)
			if err != nil {
				log.WithField("error", err).Error("solana:new")
				continue
			}

			a.b[cc.Symbol] = s
			a.b[cc.Name] = s
		}
	}

	return &a, nil
}

func (h *WriterHandler) GetBalanceOf(Symbol string, Addr string) (string, error) {
	s, ok := h.b[Symbol]
	if !ok {
		return "", e.ErrBadRequest
	}

	return s.GetBalanceOf(Addr)
}

func (h *WriterHandler) SendFromTo(Symbol string, PublicAddrFrom string, PublicAddrTo string, PrivateAddrFrom string, Amount string) (*chain.Transaction, error) {
	s, ok := h.b[Symbol]
	if !ok {
		return nil, e.ErrBadRequest
	}

	AmountDecimal, err := decimal.NewFromString(Amount)
	if err != nil {
		return nil, e.ErrBadRequest
	}

	trx, err := s.SendFromTo(PublicAddrFrom, PublicAddrTo, PrivateAddrFrom, AmountDecimal)
	if err != nil {
		return nil, err
	}

	err = h.client.OnNewOutgoingTransaction(Symbol, trx)
	if err != nil {
		log.WithField("symbol", Symbol).WithField("From", PublicAddrFrom).WithField("To", PublicAddrTo).WithField("error", err.Error()).Error("OnNewOutgoingTransaction")
	}

	return trx, nil
}

func (h *WriterHandler) GetTxStatusByHash(Symbol string, TxHash string) (string, error) {
	s, ok := h.b[Symbol]
	if !ok {
		return "", e.ErrBadRequest
	}

	return s.GetTxStatusByHash(TxHash)
}

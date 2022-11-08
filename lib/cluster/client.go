package cluster

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pr "github.com/truekupo/cluster/common/interfaces/messages/response"
	pt "github.com/truekupo/cluster/common/interfaces/messages/types"
	lp "github.com/truekupo/cluster/common/interfaces/services/listener"
	wp "github.com/truekupo/cluster/common/interfaces/services/writer"

	"github.com/truekupo/cluster/lib/blockchain/chain"
	"github.com/truekupo/cluster/lib/config"
	"github.com/truekupo/cluster/lib/logger"
	"github.com/truekupo/cluster/lib/utils/convert"
)

type Cluster struct {
	listener_clients map[string]lp.ListenerServiceClient
	writer_client    wp.WriterServiceClient

	config *config.Cluster
}

var (
	log *logrus.Entry = nil
)

func New(conf *config.Cluster) *Cluster {
	log = logger.LogModule("cluster-client")

	client := &Cluster{
		listener_clients: map[string]lp.ListenerServiceClient{},
		config:           conf,
	}

	client.init_listeners(conf.Listeners)

	client.init_writer(conf.Writer)

	return client
}

func (c *Cluster) init_writer(w config.Writer) error {
	if w.Addr == "" {
		return fmt.Errorf("Writer not configured")
	}

	// Dial Writer
	conn, err := grpc.Dial(w.Addr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("Dial writer %s failed (%s)", w.Addr, err.Error())
	}

	// Create Client
	client := wp.NewWriterServiceClient(conn)

	c.writer_client = client

	log.WithField("address", w.Addr).Info("add Writer")

	return nil
}

func (c *Cluster) init_listeners(cfg []config.Listener) error {
	for _, l := range cfg {
		// Dial Listener
		conn, err := grpc.Dial(l.Addr, grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("Dial listener %s failed (%s)", l.Addr, err.Error())
		}

		// Create Client
		client := lp.NewListenerServiceClient(conn)

		// Validate Service
		// TODO:

		c.listener_clients[l.Symbol] = client

		log.WithField("symbol", l.Symbol).WithField("address", l.Addr).Info("add Listener")
	}

	return nil
}

func (c *Cluster) FetchTransactionStatus(Code string, TxHash string) (string, error) {
	if c.writer_client == nil {
		return "", fmt.Errorf("Client to writer not initialized")
	}

	ret, err := c.writer_client.GetTxStatusByHash(context.Background(),
		&wp.GetTxStatusByHashRequest{
			Symbol: pt.CoinCode(pt.CoinCode_value[Code]),
			TxHash: TxHash,
		},
	)
	if err != nil {
		return "", err
	}
	if ret.GetRetStatus().GetCode() != pr.StatusCode_OK {
		return "", fmt.Errorf(ret.GetRetStatus().GetError())
	}

	return convert.ToStringTxStatus(ret.GetStatus()), nil
}

func (c *Cluster) OnNewOutgoingTransaction(Code string, Trx *chain.Transaction) error {
	client, ok := c.listener_clients[Code]
	if !ok {
		return fmt.Errorf("Client to %s listener not found", Code)
	}

	ret, err := client.AddTransaction(context.Background(),
		&lp.AddTransactionRequest{
			Transaction: &pt.Transaction{
				TxHash:    Trx.Hash,
				CreatedAt: Trx.CreatedAt,
				BlockNum:  Trx.BlockNum,
				FromAddr:  Trx.From,
				ToAddr:    Trx.To,
				Amount:    Trx.Amount.String(),
				Fee:       Trx.Fee.String(),
				Direction: pt.TxDirection_OUTPUT,
				Status:    pt.TxStatus_NEW,
			},
		})
	if err != nil {
		return err
	}

	if ret.GetRetStatus().GetCode() != pr.StatusCode_OK {
		return fmt.Errorf(ret.GetRetStatus().GetError())
	}

	return nil
}

package server

import (
	"context"
	"net"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pr "github.com/truekupo/cluster/common/interfaces/messages/response"
	pt "github.com/truekupo/cluster/common/interfaces/messages/types"
	proto "github.com/truekupo/cluster/common/interfaces/services/listener"
	"github.com/truekupo/cluster/lib/blockchain/chain"
	"github.com/truekupo/cluster/lib/utils/convert"
	"github.com/truekupo/cluster/listener-srv/lib/config"
	"github.com/truekupo/cluster/listener-srv/listener"

	"github.com/truekupo/cluster/lib/logger"
)

type GrpcServer struct {
	cfg      *config.Config
	listener net.Listener
	server   *grpc.Server
	internal *listener.ListenerHandler

	proto.UnimplementedListenerServiceServer
}

var (
	log *logrus.Entry = nil
)

func NewGrpcServer(cfg *config.Config, internal *listener.ListenerHandler) (*GrpcServer, error) {
	log = logger.LogModule("grpcserver")

	listener, err := net.Listen("tcp", cfg.GrpcListenAddr)
	if err != nil {
		log.WithField("error", err).WithField("grpc_listen_addr", cfg.GrpcListenAddr).Error("Could not listen to port")
		return nil, err
	}

	instance := &GrpcServer{
		cfg:      cfg,
		server:   grpc.NewServer(),
		listener: listener,
		internal: internal,
	}
	proto.RegisterListenerServiceServer(instance.server, instance)

	log.WithField("grpc_listen_addr", cfg.GrpcListenAddr).Info("new")

	return instance, nil
}

func (gs *GrpcServer) AddAddress(ctx context.Context, in *proto.AddAddressRequest) (*proto.AddAddressResponse, error) {
	log.WithField("account", in.GetAccountUuid()).WithField("address", in.GetAddress()).Debug("AddAddress")

	err := gs.internal.AddWatchAddress(in.GetAccountUuid(), in.GetAddress())
	if err != nil {
		log.WithField("account", in.GetAccountUuid()).WithField("address", in.GetAddress()).WithField("error", err.Error()).Error("AddAddress")

		return &proto.AddAddressResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_InternalServerError,
				Error: err.Error()},
		}, nil
	}

	log.WithField("account", in.GetAccountUuid()).WithField("address", in.GetAddress()).WithField("status", true).Info("AddAddress")

	return &proto.AddAddressResponse{
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) TransactionsByAddress(ctx context.Context, in *proto.TransactionsByAddressRequest) (*proto.TransactionsByAddressResponse, error) {
	log.WithField("from", in.GetFrom()).WithField("limit", in.GetLimit()).WithField("address", in.GetAddress()).Debug("TransactionsByAddress")

	transactions, err := gs.internal.TransactionsByAddress(in.GetAddress(), in.GetFrom(), in.GetLimit())
	if err != nil {
		log.WithField("from", in.GetFrom()).WithField("limit", in.GetLimit()).WithField("address", in.GetAddress()).WithField("error", err.Error()).Error("TransactionsByAddress")

		return &proto.TransactionsByAddressResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_NotFound,
				Error: err.Error()},
		}, nil
	}

	trxs := make([]*pt.Transaction, 0, len(transactions))
	for _, trx := range transactions {
		trxs = append(trxs, &pt.Transaction{
			TxHash:    trx.Hash,
			CreatedAt: trx.CreatedAt,
			BlockNum:  trx.BlockNum,
			FromAddr:  trx.From,
			ToAddr:    trx.To,
			Amount:    trx.Amount.String(),
			Fee:       trx.Fee.String(),
			Direction: convert.ToProtoTxDirection(trx.Direction),
			Status:    convert.ToProtoTxStatus(trx.Status),
		})
	}

	log.WithField("from", in.GetFrom()).WithField("limit", in.GetLimit()).WithField("address", in.GetAddress()).WithField("count", len(trxs)).WithField("status", true).Info("TransactionsByAddress")

	return &proto.TransactionsByAddressResponse{
		Transactions: trxs,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: "",
		},
	}, nil
}

func (gs *GrpcServer) TransactionsByAccount(ctx context.Context, in *proto.TransactionsByAccountRequest) (*proto.TransactionsByAccountResponse, error) {
	log.WithField("from", in.GetFrom()).WithField("limit", in.GetLimit()).WithField("account", in.GetAccountUuid()).Debug("TransactionsByAccount")

	transactions, err := gs.internal.TransactionsByAccount(in.GetAccountUuid(), in.GetFrom(), in.GetLimit())
	if err != nil {
		log.WithField("from", in.GetFrom()).WithField("limit", in.GetLimit()).WithField("account", in.GetAccountUuid()).WithField("error", err.Error()).Error("TransactionsByAccount")

		return &proto.TransactionsByAccountResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_NotFound,
				Error: err.Error()},
		}, nil
	}

	trxs := make([]*pt.Transaction, 0, len(transactions))
	for _, trx := range transactions {
		trxs = append(trxs, &pt.Transaction{
			TxHash:    trx.Hash,
			CreatedAt: trx.CreatedAt,
			BlockNum:  trx.BlockNum,
			FromAddr:  trx.From,
			ToAddr:    trx.To,
			Amount:    trx.Amount.String(),
			Fee:       trx.Fee.String(),
			Direction: convert.ToProtoTxDirection(trx.Direction),
			Status:    convert.ToProtoTxStatus(trx.Status),
		})
	}

	log.WithField("from", in.GetFrom()).WithField("limit", in.GetLimit()).WithField("account", in.GetAccountUuid()).WithField("count", len(trxs)).WithField("status", true).Info("TransactionsByAccount")

	return &proto.TransactionsByAccountResponse{
		Transactions: trxs,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: "",
		},
	}, nil
}

func (gs *GrpcServer) AddTransaction(ctx context.Context, in *proto.AddTransactionRequest) (*proto.AddTransactionResponse, error) {
	log.WithField("tx_hash", in.GetTransaction().GetTxHash()).WithField("from", in.GetTransaction().GetFromAddr()).WithField("to", in.GetTransaction().GetToAddr()).WithField("Amount", in.GetTransaction().GetAmount()).Debug("AddTransaction")

	amount, err := decimal.NewFromString(in.GetTransaction().GetAmount())
	if err != nil {
		log.WithField("tx_hash", in.GetTransaction().GetTxHash()).WithField("error", err.Error()).Error("AddTransaction")

		return &proto.AddTransactionResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_BadRequest,
				Error: "Amount: Failed format",
			},
		}, nil
	}

	fee, err := decimal.NewFromString(in.GetTransaction().GetFee())
	if err != nil {
		log.WithField("tx_hash", in.GetTransaction().GetTxHash()).WithField("error", err.Error()).Error("AddTransaction")

		return &proto.AddTransactionResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_BadRequest,
				Error: "Fee: Failed format",
			},
		}, nil
	}

	trx := chain.Transaction{
		Hash:      in.GetTransaction().GetTxHash(),
		CreatedAt: in.GetTransaction().GetCreatedAt(),
		BlockNum:  in.GetTransaction().GetBlockNum(),
		From:      in.GetTransaction().GetFromAddr(),
		To:        in.GetTransaction().GetToAddr(),
		Amount:    amount,
		Fee:       fee,
		Direction: convert.ToStringTxDirection(in.GetTransaction().GetDirection()),
		Status:    convert.ToStringTxStatus(in.GetTransaction().GetStatus()),
	}

	err = gs.internal.AddTransaction(trx)
	if err != nil {
		log.WithField("tx_hash", in.GetTransaction().GetTxHash()).WithField("error", err.Error()).Error("AddTransaction")

		return &proto.AddTransactionResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_InternalServerError,
				Error: err.Error(),
			},
		}, nil
	}

	log.WithField("tx_hash", in.GetTransaction().GetTxHash()).WithField("status", true).Info("AddTransaction")

	return &proto.AddTransactionResponse{
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: "",
		},
	}, nil
}

func (gs *GrpcServer) GetTxByHash(ctx context.Context, in *proto.GetTxByHashRequest) (*proto.GetTxByHashResponse, error) {
	log.WithField("tx_hash", in.GetTxHash()).Debug("GetTxByHash")

	tx, err := gs.internal.GetTxByHash(in.GetTxHash())
	if err != nil {
		log.WithField("tx_hash", in.GetTxHash()).WithField("error", err.Error()).Debug("GetTxByHash")

		return &proto.GetTxByHashResponse{
			Transaction: nil,
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_NotFound,
				Error: err.Error(),
			},
		}, nil
	}

	log.WithField("tx_hash", in.GetTxHash()).WithField("status", true).Info("GetTxByHash")

	return &proto.GetTxByHashResponse{
		Transaction: &pt.Transaction{
			TxHash:    tx.Hash,
			CreatedAt: tx.CreatedAt,
			BlockNum:  tx.BlockNum,
			FromAddr:  tx.From,
			ToAddr:    tx.To,
			Amount:    tx.Amount.String(),
			Fee:       tx.Fee.String(),
			Direction: convert.ToProtoTxDirection(tx.Direction),
			Status:    convert.ToProtoTxStatus(tx.Status),
		},
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: "",
		},
	}, nil
}

func (gs *GrpcServer) Start() error {
	log.Debug("start")

	go func() {
		if err := gs.server.Serve(gs.listener); err != nil {
			log.WithField("error", err).Error("start")
			return
		}
	}()

	return nil
}

func (gs *GrpcServer) Stop() {
	log.Debug("stop")

	gs.server.Stop()
	_ = gs.listener.Close()
}

package server

import (
	"context"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pr "github.com/truekupo/cluster/common/interfaces/messages/response"
	pt "github.com/truekupo/cluster/common/interfaces/messages/types"
	proto "github.com/truekupo/cluster/common/interfaces/services/writer"
	"github.com/truekupo/cluster/lib/utils/convert"
	"github.com/truekupo/cluster/writer-srv/lib/config"
	"github.com/truekupo/cluster/writer-srv/writer"

	"github.com/truekupo/cluster/lib/logger"
)

type GrpcServer struct {
	cfg      *config.Config
	listener net.Listener
	server   *grpc.Server
	writer   *writer.WriterHandler

	proto.UnimplementedWriterServiceServer
}

var (
	log *logrus.Entry = nil
)

func NewGrpcServer(cfg *config.Config, writer *writer.WriterHandler) (*GrpcServer, error) {
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
		writer:   writer,
	}
	proto.RegisterWriterServiceServer(instance.server, instance)

	log.WithField("grpc_listen_addr", cfg.GrpcListenAddr).Info("new")

	return instance, nil
}

func (gs *GrpcServer) GetBalanceOf(ctx context.Context, in *proto.GetBalanceRequest) (*proto.GetBalanceResponse, error) {
	log.WithField("symbol", in.GetSymbol()).WithField("address", in.GetAddress()).Debug("GetBalanceOf")

	amount, err := gs.writer.GetBalanceOf(in.GetSymbol().String(), in.GetAddress())
	if err != nil {
		log.WithField("error", err).WithField("symbol", in.GetSymbol()).WithField("address", in.GetAddress()).Error("GetBalanceOf")

		return &proto.GetBalanceResponse{
			Amount: "0",
			RetStatus: &pr.Status{
				Code:  convert.ToProtoError(err),
				Error: err.Error()},
		}, nil
	}

	log.WithField("symbol", in.GetSymbol()).WithField("address", in.GetAddress()).WithField("amount", amount).Info("GetBalanceOf")

	return &proto.GetBalanceResponse{
		Amount: amount,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) SendFromTo(ctx context.Context, in *proto.SendFromToRequest) (*proto.SendFromToResponse, error) {
	log.WithField("symbol", in.GetSymbol()).WithField("from", in.GetFromAddr()).WithField("to", in.GetToAddr()).WithField("amount", in.GetAmount()).Debug("SendFromTo")

	tx, err := gs.writer.SendFromTo(in.GetSymbol().String(), in.GetFromAddr(), in.GetToAddr(), in.GetFromPrivate(), in.GetAmount())
	if err != nil {
		log.WithField("symbol", in.GetSymbol()).WithField("from", in.GetFromAddr()).WithField("to", in.GetToAddr()).WithField("amount", in.GetAmount()).WithField("error", err).Error("SendFromTo")

		return &proto.SendFromToResponse{
			Transaction: nil,
			RetStatus: &pr.Status{
				Code:  convert.ToProtoError(err),
				Error: err.Error()},
		}, nil
	}

	log.WithField("symbol", in.GetSymbol()).WithField("from", in.GetFromAddr()).WithField("to", in.GetToAddr()).WithField("amount", in.GetAmount()).WithField("fee", tx.Fee).WithField("tx_hash", tx.Hash).Info("SendFromTo")

	return &proto.SendFromToResponse{
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
			Error: ""},
	}, nil
}

func (gs *GrpcServer) GetTxStatusByHash(ctx context.Context, in *proto.GetTxStatusByHashRequest) (*proto.GetTxStatusByHashResponse, error) {
	log.WithField("symbol", in.GetSymbol()).WithField("tx_hash", in.GetTxHash()).Debug("GetTxStatusByHash")

	status, err := gs.writer.GetTxStatusByHash(in.GetSymbol().String(), in.GetTxHash())
	if err != nil {
		log.WithField("symbol", in.GetSymbol()).WithField("tx_hash", in.GetTxHash()).WithField("error", err).Error("GetTxStatusByHash")

		return &proto.GetTxStatusByHashResponse{
			RetStatus: &pr.Status{
				Code:  convert.ToProtoError(err),
				Error: err.Error()},
		}, nil
	}

	log.WithField("symbol", in.GetSymbol()).WithField("tx_hash", in.GetTxHash()).WithField("status", status).Info("GetTxStatusByHash")

	return &proto.GetTxStatusByHashResponse{
		Status: convert.ToProtoTxStatus(status),
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
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

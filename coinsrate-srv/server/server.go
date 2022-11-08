package server

import (
	"context"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/truekupo/cluster/coinsrate-srv/internal"
	"github.com/truekupo/cluster/coinsrate-srv/lib/config"
	pr "github.com/truekupo/cluster/common/interfaces/messages/response"
	pt "github.com/truekupo/cluster/common/interfaces/messages/types"
	proto "github.com/truekupo/cluster/common/interfaces/services/coinsrate"

	"github.com/truekupo/cluster/lib/logger"
)

type GrpcServer struct {
	cfg      *config.Config
	listener net.Listener
	server   *grpc.Server
	internal *coinsrate.CoinsRateHandler

	proto.UnimplementedCoinsRateServiceServer
}

var (
	log *logrus.Entry = nil
)

func NewGrpcServer(cfg *config.Config, ch *coinsrate.CoinsRateHandler) (*GrpcServer, error) {
	log = logger.LogModule("grpcserver")

	listener, err := net.Listen("tcp", cfg.GrpcListenAddr)
	if err != nil {
		log.Errorf("Could not listen to port in Start() %s: %v", cfg.GrpcListenAddr, err)
		return nil, err
	}
	instance := &GrpcServer{
		cfg:      cfg,
		server:   grpc.NewServer(),
		listener: listener,
		internal: ch,
	}
	proto.RegisterCoinsRateServiceServer(instance.server, instance)

	return instance, nil
}

func (gs *GrpcServer) LastRates(ctx context.Context, in *proto.LastRatesRequest) (*proto.LastRatesResponse, error) {
	log.WithField("currency", in.GetCurrency()).WithField("coins", in.GetCoins()).Debug("LastRates")

	coins := in.GetCoins()
	rates := make([]*proto.Rate, 0, len(coins))

	for _, c := range coins {
		unix_tm, rate, err := gs.internal.LastRate(pt.CoinCode_name[int32(c)], pt.CoinCode_name[int32(in.GetCurrency())])
		if err != nil {
			continue
		}

		rates = append(rates, &proto.Rate{
			UnixTm:   unix_tm,
			Coin:     c,
			Currency: in.GetCurrency(),
			Rate:     rate,
		})
	}

	log.WithField("currency", in.GetCurrency()).WithField("coins", in.GetCoins()).WithField("rates", rates).Info("LastRates")

	return &proto.LastRatesResponse{
		Rates: rates,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) Start() error {
	log.Debug("GrpcServer:Start")

	go func() {
		//log.Printf("gRPC server is listening on address %s", gs.config.GrpcAddress)
		if err := gs.server.Serve(gs.listener); err != nil {
			//log.Fatalf("failed to serve gRPC: %v", err)
			return
		}
	}()

	return nil
}

func (gs *GrpcServer) Stop() {
	log.Debug("GrpcServer:Stop")

	gs.server.Stop()
	_ = gs.listener.Close()
}

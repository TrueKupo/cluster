package server

import (
	"context"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/truekupo/cluster/account-srv/account"
	"github.com/truekupo/cluster/account-srv/lib/config"

	pr "github.com/truekupo/cluster/common/interfaces/messages/response"
	pt "github.com/truekupo/cluster/common/interfaces/messages/types"
	proto "github.com/truekupo/cluster/common/interfaces/services/account"

	"github.com/truekupo/cluster/lib/logger"
)

type GrpcServer struct {
	cfg      *config.Config
	listener net.Listener
	server   *grpc.Server
	account  *account.AccountHandler

	proto.UnimplementedAccountServiceServer
}

var (
	log *logrus.Entry = nil
)

func NewGrpcServer(cfg *config.Config, account *account.AccountHandler) (*GrpcServer, error) {
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
		account:  account,
	}
	proto.RegisterAccountServiceServer(instance.server, instance)

	return instance, nil
}

func (gs *GrpcServer) NewMnemonic(ctx context.Context, in *proto.NewMnemonicRequest) (*proto.NewMnemonicResponse, error) {
	log.WithField("entropy", in.GetEntropy()).WithField("lang", in.GetLang()).Debug("NewMnemonic")

	mnemonic, err := gs.account.NewMnemonic(int(in.GetEntropy()), int32(in.GetLang()))
	if err != nil {
		log.WithField("entropy", in.GetEntropy()).WithField("lang", in.GetLang()).WithField("error", err.Error()).Error("NewMnemonic")

		return &proto.NewMnemonicResponse{
			Mnemonic: mnemonic,
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_InternalServerError,
				Error: err.Error()},
		}, nil
	}

	log.WithField("entropy", in.GetEntropy()).WithField("lang", in.GetLang()).WithField("status", true).Info("NewMnemonic")

	return &proto.NewMnemonicResponse{
		Mnemonic: mnemonic,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) GetSeedFromMnemonic(ctx context.Context, in *proto.GetSeedFromMnemonicRequest) (*proto.SeedFromMnemonicResponse, error) {
	log.Debug("GetSeedFromMnemonic")

	seed, err := gs.account.GetSeed(in.GetMnemonic(), in.GetPassword())
	if err != nil {
		log.WithField("error", err.Error()).Error("GetSeedFromMnemonic")

		return &proto.SeedFromMnemonicResponse{
			Seed: []byte{},
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_InternalServerError,
				Error: err.Error()},
		}, nil
	}

	log.WithField("status", true).Info("GetSeedFromMnemonic")

	return &proto.SeedFromMnemonicResponse{
		Seed: seed,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) GetMasterKeyFromMnemonic(ctx context.Context, in *proto.GetMasterKeyFromMnemonicRequest) (*proto.MasterKeyFromMnemonicResponse, error) {
	log.Debug("GetMasterKeyFromMnemonic")

	master_key, err := gs.account.GetMasterKey(in.GetMnemonic(), in.GetPassword())
	if err != nil {
		log.WithField("error", err.Error()).Error("GetMasterKeyFromMnemonic")

		return &proto.MasterKeyFromMnemonicResponse{
			MasterKey: []byte{},
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_InternalServerError,
				Error: err.Error()},
		}, nil
	}

	log.WithField("status", true).Info("GetMasterKeyFromMnemonic")

	return &proto.MasterKeyFromMnemonicResponse{
		MasterKey: master_key,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) SeedDeriveToPublicKeyHex(ctx context.Context, in *proto.SeedDeriveToAddressHexRequest) (*proto.SeedDeriveToAddressHexResponse, error) {
	log.WithField("symbol", in.GetSymbol()).WithField("path", in.GetPath()).Debug("SeedDeriveToPublicKeyHex")

	symbol, ok := pt.CoinCode_name[int32(in.GetSymbol())]
	if !ok {
		log.WithField("symbol", in.GetSymbol()).WithField("path", in.GetPath()).WithField("error", "Unknown Coin Code").Error("SeedDeriveToPublicKeyHex")

		return &proto.SeedDeriveToAddressHexResponse{
			Data: "",
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_BadRequest,
				Error: "Unknown Coin Code"},
		}, nil
	}

	addr, err := gs.account.SeedDeriveToPublicKeyHex(symbol, in.GetSeed(), in.GetPath())
	if err != nil {
		log.WithField("symbol", in.GetSymbol()).WithField("path", in.GetPath()).WithField("error", err.Error()).Error("SeedDeriveToPublicKeyHex")

		return &proto.SeedDeriveToAddressHexResponse{
			Data: "",
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_InternalServerError,
				Error: err.Error()},
		}, nil
	}

	log.WithField("symbol", in.GetSymbol()).WithField("path", in.GetPath()).WithField("status", true).WithField("addr", addr).Info("SeedDeriveToPublicKeyHex")

	return &proto.SeedDeriveToAddressHexResponse{
		Data: addr,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) SeedDeriveToPrivateKeyHex(ctx context.Context, in *proto.SeedDeriveToAddressHexRequest) (*proto.SeedDeriveToAddressHexResponse, error) {
	log.WithField("symbol", in.GetSymbol()).WithField("path", in.GetPath()).Debug("SeedDeriveToPrivateKeyHex")

	symbol, ok := pt.CoinCode_name[int32(in.GetSymbol())]
	if !ok {
		log.WithField("symbol", in.GetSymbol()).WithField("path", in.GetPath()).WithField("error", "Unknown Coin Code").Error("SeedDeriveToPrivateKeyHex")

		return &proto.SeedDeriveToAddressHexResponse{
			Data: "",
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_BadRequest,
				Error: "Unknown Coin Code"},
		}, nil
	}

	addr, err := gs.account.SeedDeriveToPrivateKeyHex(symbol, in.GetSeed(), in.GetPath())
	if err != nil {
		log.WithField("symbol", in.GetSymbol()).WithField("path", in.GetPath()).WithField("error", err.Error()).Error("SeedDeriveToPrivateKeyHex")

		return &proto.SeedDeriveToAddressHexResponse{
			Data: "",
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_InternalServerError,
				Error: err.Error()},
		}, nil
	}

	log.WithField("symbol", in.GetSymbol()).WithField("path", in.GetPath()).WithField("status", true).Info("SeedDeriveToPrivateKeyHex")

	return &proto.SeedDeriveToAddressHexResponse{
		Data: addr,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) SeedDeriveToAccountData(ctx context.Context, in *proto.SeedDeriveToAccountRequest) (*proto.SeedDeriveToAccountResponse, error) {
	log.WithField("symbol", in.GetSymbol()).WithField("path", in.GetPath()).Debug("SeedDeriveToAccountData")

	symbol, ok := pt.CoinCode_name[int32(in.GetSymbol())]
	if !ok {
		return &proto.SeedDeriveToAccountResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_BadRequest,
				Error: "Unknown Coin Code"},
		}, nil
	}

	private_hex, public_hex, public_addr, err := gs.account.SeedDeriveToAccountData(symbol, in.GetSeed(), in.GetPath())
	if err != nil {
		log.WithField("symbol", in.GetSymbol()).WithField("path", in.GetPath()).WithField("error", err.Error()).Error("SeedDeriveToAccountData")
		return &proto.SeedDeriveToAccountResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_InternalServerError,
				Error: err.Error()},
		}, nil
	}

	log.WithField("symbol", in.GetSymbol()).WithField("path", in.GetPath()).WithField("addr", public_hex).WithField("status", true).Info("SeedDeriveToAccountData")

	return &proto.SeedDeriveToAccountResponse{
		PrivateKey: private_hex,
		PublicKey:  public_hex,
		PublicAddr: public_addr,
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
			log.Error("failed to serve gRPC: %v", err)
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

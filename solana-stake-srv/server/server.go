package server

import (
	"context"
	"net"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	tmpb "google.golang.org/protobuf/types/known/timestamppb"

	pr "github.com/truekupo/cluster/common/interfaces/messages/response"
	proto "github.com/truekupo/cluster/common/interfaces/services/solana_stake"
	"github.com/truekupo/cluster/solana-stake-srv/lib/config"
	"github.com/truekupo/cluster/solana-stake-srv/stake"

	"github.com/truekupo/cluster/lib/logger"
)

type GrpcServer struct {
	cfg      *config.Config
	listener net.Listener
	server   *grpc.Server
	internal *stake.SolanaStakeHandler

	proto.UnimplementedSolanaStakeServiceServer
}

var (
	log *logrus.Entry = nil
)

func NewGrpcServer(cfg *config.Config, internal *stake.SolanaStakeHandler) (*GrpcServer, error) {
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
	proto.RegisterSolanaStakeServiceServer(instance.server, instance)

	log.WithField("grpc_listen_addr", cfg.GrpcListenAddr).Info("new")

	return instance, nil
}

func (gs *GrpcServer) ValidatorInfo(ctx context.Context, in *proto.ValidatorInfoRequest) (*proto.ValidatorInfoResponse, error) {
	log.WithField("id", in.GetId()).Debug("ValidatorInfo")

	info, err := gs.internal.ValidatorInfoById(in.GetId())
	if err != nil {
		log.WithField("id", in.GetId()).WithField("error", err.Error()).Info("ValidatorInfo")

		return &proto.ValidatorInfoResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_NotFound,
				Error: err.Error()},
		}, nil
	}

	log.WithField("Id", info.Id).WithField("network", info.Network).WithField("account", info.Account).WithField("name", info.Name).WithField("vote_account", info.VoteAccount).Info("ValidatorInfo")

	return &proto.ValidatorInfoResponse{
		Info: &proto.Validator{
			Id:                           info.Id,
			Network:                      info.Network,
			Account:                      info.Account,
			Name:                         info.Name,
			WwwUrl:                       info.WwwUrl,
			Details:                      info.Details,
			AvatarUrl:                    info.AvatarUrl,
			CreatedAt:                    tmpb.New(info.CreatedAt),
			UpdatedAt:                    tmpb.New(info.UpdatedAt),
			TotalScore:                   info.TotalScore,
			RootDistanceScore:            info.RootDistanceScore,
			VoteDistanceScore:            info.VoteDistanceScore,
			SkippedSlotScore:             info.SkippedSlotScore,
			SoftwareVersion:              info.SoftwareVersion,
			SoftwareVersionScore:         info.SoftwareVersionScore,
			StakeConcentrationScore:      info.StakeConcentrationScore,
			DataCenterConcentrationScore: info.DataCenterConcentrationScore,
			PublishedInformationScore:    info.PublishedInformationScore,
			SecurityReportScore:          info.SecurityReportScore,
			ActiveStake:                  info.ActiveStake,
			Commission:                   info.Commission,
			Delinquent:                   info.Delinquent,
			DataCenterKey:                info.DataCenterKey,
			DataCenterHost:               info.DataCenterHost,
			AutonomousSystemNumber:       info.AutonomousSystemNumber,
			VoteAccount:                  info.VoteAccount,
			EpochCredits:                 info.EpochCredits,
			SkippedSlots:                 info.SkippedSlots,
			SkippedSlotPercent:           info.SkippedSlotPercent,
			PingTime:                     info.PingTime,
			Url:                          info.Url,
		},
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) ValidatorsList(ctx context.Context, in *proto.ValidatorsListRequest) (*proto.ValidatorsListResponse, error) {
	log.WithField("From", in.GetFrom()).WithField("Limit", in.GetLimit()).WithField("SortBy", in.GetSortField()).WithField("Order", in.GetOrder()).Debug("ValidatorsList")

	list, err := gs.internal.ValidatorsList(in.GetFrom(), in.GetLimit(), strings.ToLower(in.GetSortField().String()), in.GetOrder().String())
	if err != nil {
		log.WithField("From", in.GetFrom()).WithField("Limit", in.GetLimit()).WithField("SortBy", in.GetSortField()).WithField("Order", in.GetOrder()).WithField("error", err.Error()).Info("ValidatorsList")

		return &proto.ValidatorsListResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_NotFound,
				Error: err.Error()},
		}, nil
	}

	ll := make([]*proto.Validator, 0, len(list))
	for _, info := range list {
		ll = append(ll, &proto.Validator{
			Id:                           info.Id,
			Network:                      info.Network,
			Account:                      info.Account,
			Name:                         info.Name,
			WwwUrl:                       info.WwwUrl,
			Details:                      info.Details,
			AvatarUrl:                    info.AvatarUrl,
			CreatedAt:                    tmpb.New(info.CreatedAt),
			UpdatedAt:                    tmpb.New(info.UpdatedAt),
			TotalScore:                   info.TotalScore,
			RootDistanceScore:            info.RootDistanceScore,
			VoteDistanceScore:            info.VoteDistanceScore,
			SkippedSlotScore:             info.SkippedSlotScore,
			SoftwareVersion:              info.SoftwareVersion,
			SoftwareVersionScore:         info.SoftwareVersionScore,
			StakeConcentrationScore:      info.StakeConcentrationScore,
			DataCenterConcentrationScore: info.DataCenterConcentrationScore,
			PublishedInformationScore:    info.PublishedInformationScore,
			SecurityReportScore:          info.SecurityReportScore,
			ActiveStake:                  info.ActiveStake,
			Commission:                   info.Commission,
			Delinquent:                   info.Delinquent,
			DataCenterKey:                info.DataCenterKey,
			DataCenterHost:               info.DataCenterHost,
			AutonomousSystemNumber:       info.AutonomousSystemNumber,
			VoteAccount:                  info.VoteAccount,
			EpochCredits:                 info.EpochCredits,
			SkippedSlots:                 info.SkippedSlots,
			SkippedSlotPercent:           info.SkippedSlotPercent,
			PingTime:                     info.PingTime,
			Url:                          info.Url,
		})
	}

	log.WithField("From", in.GetFrom()).WithField("Limit", in.GetLimit()).WithField("SortBy", in.GetSortField()).WithField("Order", in.GetOrder()).WithField("count", len(ll)).Info("ValidatorsList")

	return &proto.ValidatorsListResponse{
		List: ll,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) StakeActivationStatus(ctx context.Context, in *proto.StakeActivationStatusRequest) (*proto.StakeActivationStatusResponse, error) {
	log.WithField("Address", in.GetStakeAddressBase58()).Debug("StakeActivationStatus")

	status, active, inactive, err := gs.internal.StakeActivationStatus(in.GetStakeAddressBase58())
	if err != nil {
		log.WithField("Address", in.GetStakeAddressBase58()).WithField("error", err.Error()).Error("StakeActivationStatus")

		return &proto.StakeActivationStatusResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_NotFound,
				Error: err.Error()},
		}, nil
	}

	st, ok := proto.ActivationStatus_value[strings.ToUpper(status)]
	if !ok {
		log.WithField("Address", in.GetStakeAddressBase58()).WithField("error", "Unknown status").Error("StakeActivationStatus")

		return &proto.StakeActivationStatusResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_InternalServerError,
				Error: "ActivationStatus unknown"},
		}, nil
	}

	log.WithField("Address", in.GetStakeAddressBase58()).WithField("status", status).Info("StakeActivationStatus")

	return &proto.StakeActivationStatusResponse{
		Status:        proto.ActivationStatus(st),
		ActiveEpoch:   active,
		InactiveEpoch: inactive,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) StakeAccountInfo(ctx context.Context, in *proto.StakeAccountInfoRequest) (*proto.StakeAccountInfoResponse, error) {
	log.WithField("Address", in.GetStakeAddressBase58()).Debug("StakeAccountInfo")

	r, err := gs.internal.StakeAccountInfo(in.GetStakeAddressBase58())
	if err != nil {
		log.WithField("error", err.Error()).WithField("Address", in.GetStakeAddressBase58()).Info("StakeAccountInfo")

		return &proto.StakeAccountInfoResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_NotFound,
				Error: err.Error()},
		}, nil
	}

	log.WithField("Address", in.GetStakeAddressBase58()).WithField("status", true).Info("StakeAccountInfo")

	return &proto.StakeAccountInfoResponse{
		Amount:    decimal.NewFromInt(int64(r.Lamports)).Shift(-9).String(),
		Owner:     r.Owner,
		Excutable: r.Excutable,
		RentEpoch: r.RentEpoch,
		StakeAccount: &proto.StakeAccount{
			Status: proto.ActivationStatus(r.StakeAccount.Type),
			Info: &proto.StakeAccount_Info{
				Meta: &proto.StakeAccount_Info_Meta{
					RentExemptReserve: r.StakeAccount.Info.Meta.RentExemptReserve,
					Authorized: &proto.StakeAccount_Info_Meta_Authorized{
						StakeAddressBase58:    r.StakeAccount.Info.Meta.Authorized.Staker.ToBase58(),
						WithdrawAddressBase58: r.StakeAccount.Info.Meta.Authorized.Withdrawer.ToBase58(),
						Lockup: &proto.StakeAccount_Info_Meta_Authorized_Lockup{
							UnixTimeStamp:          r.StakeAccount.Info.Meta.Authorized.Lockup.UnixTimeStamp,
							Epoch:                  r.StakeAccount.Info.Meta.Authorized.Lockup.Epoch,
							CustodianAddressBase58: r.StakeAccount.Info.Meta.Authorized.Lockup.Custodian.ToBase58(),
						},
					},
				},
				Stake: &proto.StakeAccount_Info_Stake{
					Delegation: &proto.StakeAccount_Info_Stake_Delegation{
						VoterAddressBase58: r.StakeAccount.Info.Stake.Delegation.Voter.ToBase58(),
						Stake:              r.StakeAccount.Info.Stake.Delegation.Stake,
						ActivationEpoch:    r.StakeAccount.Info.Stake.Delegation.ActivationEpoch,
						DeactivationEpoch:  r.StakeAccount.Info.Stake.Delegation.DeactivationEpoch,
						WarmupCooldownRate: r.StakeAccount.Info.Stake.Delegation.WarmupCooldownRate,
					},
					CreditsObserved: r.StakeAccount.Info.Stake.CreditsObserved,
				},
			},
		},
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) CreateStakeAccount(ctx context.Context, in *proto.CreateStakeAccountRequest) (*proto.CreateStakeAccountResponse, error) {
	log.WithField("Amount", in.GetAmount()).Debug("CreateStakeAccount")

	amount, err := decimal.NewFromString(in.GetAmount())
	if err != nil {
		log.WithField("Amount", in.GetAmount()).WithField("error", "Amount: Wrong format").Error("CreateStakeAccount")

		return &proto.CreateStakeAccountResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_BadRequest,
				Error: err.Error()},
		}, nil
	}

	TxHash, stakePublicBase58, err := gs.internal.CreateStakeAccount(in.GetSignerPrivateBase58(), amount)
	if err != nil {
		log.WithField("Amount", in.GetAmount()).WithField("error", err.Error()).Info("CreateStakeAccount")

		return &proto.CreateStakeAccountResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_InternalServerError,
				Error: err.Error()},
		}, nil
	}

	log.WithField("TxHash", TxHash).WithField("StakeAddressBase58", stakePublicBase58).WithField("Amount", in.GetAmount()).Info("CreateStakeAccount")

	return &proto.CreateStakeAccountResponse{
		TxHash:             TxHash,
		StakeAddressBase58: stakePublicBase58,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) DelegateStake(ctx context.Context, in *proto.DelegateStakeRequest) (*proto.DelegateStakeResponse, error) {
	log.WithField("StakeAddress", in.GetStakeAddressBase58()).WithField("VoteAddress", in.GetVoteAddressBase58()).Debug("DelegateStake")

	TxHash, err := gs.internal.DelegateStake(in.GetSignerPrivateBase58(), in.GetStakeAddressBase58(), in.GetVoteAddressBase58())
	if err != nil {
		log.WithField("StakeAddress", in.GetStakeAddressBase58()).WithField("VoteAddress", in.GetVoteAddressBase58()).WithField("error", err.Error()).Error("DelegateStake")

		return &proto.DelegateStakeResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_InternalServerError,
				Error: err.Error()},
		}, nil
	}

	log.WithField("StakeAddress", in.GetStakeAddressBase58()).WithField("VoteAddress", in.GetVoteAddressBase58()).WithField("TxHash", TxHash).Info("DelegateStake")

	return &proto.DelegateStakeResponse{
		TxHash: TxHash,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) DeactivateStake(ctx context.Context, in *proto.DeactivateStakeRequest) (*proto.DeactivateStakeResponse, error) {
	log.WithField("StakeAddress", in.GetStakeAddressBase58()).Debug("DeactivateStake")

	TxHash, err := gs.internal.DeactivateStake(in.GetSignerPrivateBase58(), in.GetStakeAddressBase58())
	if err != nil {
		log.WithField("StakeAddress", in.GetStakeAddressBase58()).WithField("error", err.Error()).Error("DeactivateStake")

		return &proto.DeactivateStakeResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_InternalServerError,
				Error: err.Error()},
		}, nil
	}

	log.WithField("StakeAddress", in.GetStakeAddressBase58()).WithField("TxHash", TxHash).Info("DeactivateStake")

	return &proto.DeactivateStakeResponse{
		TxHash: TxHash,
		RetStatus: &pr.Status{
			Code:  pr.StatusCode_OK,
			Error: ""},
	}, nil
}

func (gs *GrpcServer) WithdrawStake(ctx context.Context, in *proto.WithdrawStakeRequest) (*proto.WithdrawStakeResponse, error) {
	log.WithField("StakeAddress", in.GetStakeAddressBase58()).Debug("WithdrawStake")

	amount, err := decimal.NewFromString(in.GetAmount())
	if err != nil {
		log.WithField("Amount", in.GetAmount()).WithField("error", "Amount: Wrong format").Error("WithdrawStake")

		return &proto.WithdrawStakeResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_BadRequest,
				Error: err.Error()},
		}, nil
	}

	TxHash, err := gs.internal.WithdrawStake(in.GetSignerPrivateBase58(), in.GetStakeAddressBase58(), amount)
	if err != nil {
		log.WithField("StakeAddress", in.GetStakeAddressBase58()).WithField("error", err.Error()).Error("WithdrawStake")

		return &proto.WithdrawStakeResponse{
			RetStatus: &pr.Status{
				Code:  pr.StatusCode_InternalServerError,
				Error: err.Error()},
		}, nil
	}

	log.WithField("StakeAddress", in.GetStakeAddressBase58()).WithField("TxHash", TxHash).Info("WithdrawStake")

	return &proto.WithdrawStakeResponse{
		TxHash: TxHash,
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

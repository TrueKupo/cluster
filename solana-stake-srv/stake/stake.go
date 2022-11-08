package stake

import (
	"context"
	"sync"
	"time"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gocraft/dbr/v2"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	client "github.com/stafiprotocol/solana-go-sdk/client"
	"github.com/truekupo/cluster/solana-stake-srv/lib/blockchain"
	"github.com/truekupo/cluster/solana-stake-srv/lib/blockchain/solana"
	"github.com/truekupo/cluster/solana-stake-srv/lib/config"

	"github.com/truekupo/cluster/solana-stake-srv/lib/db/middleware"
	"github.com/truekupo/cluster/solana-stake-srv/lib/db/models"
	"github.com/truekupo/cluster/solana-stake-srv/lib/db/repo"

	"github.com/truekupo/cluster/lib/logger"
)

type SolanaStakeHandler struct {
	sync.Mutex

	conf *config.Config
	init bool
	dbs  *dbr.Session

	endpoint string
	b        blockchain.Blockachain
}

var (
	log *logrus.Entry = nil
)

func NewSolanaStakeHandler(conf *config.Config) (*SolanaStakeHandler, error) {
	log = logger.LogModule("solana-stake-handler")

	h := SolanaStakeHandler{
		conf: conf,
		init: true,
	}

	// init blockchain
	cc := conf.Chain

	s, err := solana.New(&cc)
	if err != nil {
		log.WithField("error", err).Error("solana:new")
	}

	h.b = s

	// connect to DB
	dsd, err := middleware.NewDSD(conf)
	if err != nil {
		log.Error("listener:handler:new ", err)
		return nil, err
	}

	h.dbs = dsd.Conn.NewSession(nil)

	// choose endpoit
	switch h.conf.Net {
	case "mainnet":
		h.endpoint = rpc.MainNetBeta_RPC
	case "testnet":
		h.endpoint = rpc.TestNet_RPC
	case "devnet":
		h.endpoint = rpc.DevNet_RPC
	default:
		h.endpoint = rpc.TestNet_RPC
	}

	// start watch validators handler
	go h.validator_process()

	return &h, nil
}

func (h *SolanaStakeHandler) validator_process() {
	solana_client := rpc.New(h.endpoint)
	validators_client := NewValidatorClient(h.conf.ValidatorsAppSecret)

	for {
		va, err := solana_client.GetVoteAccounts(context.Background(), nil)
		if err != nil {
			log.Error(err)
			time.Sleep(1 * time.Minute)
			continue
		}

		log.WithField("count", len(va.Current)).Debug("GetVoteAccounts")

		for _, v := range va.Current {
			info, err := validators_client.GetValidatorInfo(h.conf.Net, v.NodePubkey.String())
			if err != nil {
				log.Error(err)
				continue
			}

			_, err = repo.Validator(h.dbs).InsertOrUpdate(&models.Validator{
				Network:                      info.Network,
				Account:                      info.Account,
				Name:                         info.Name,
				WwwUrl:                       info.WwwUrl,
				Details:                      info.Details,
				AvatarUrl:                    info.AvatarUrl,
				CreatedAt:                    info.CreatedAt,
				UpdatedAt:                    info.UpdatedAt,
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
			if err != nil {
				log.Error(err)
			}

			log.WithField("account", info.VoteAccount).Info("add/update votes account")
		}

		time.Sleep(10 * time.Minute)
	}
}

func (h *SolanaStakeHandler) ValidatorsList(From int32, Limit int32, SortBy string, Order string) ([]*ValidatorInfo, error) {
	list, err := repo.Validator(h.dbs).FindByOrder(From, Limit, SortBy, Order)
	if err != nil {
		return nil, err
	}

	ll := make([]*ValidatorInfo, 0, len(list))
	for _, info := range list {
		ll = append(ll, &ValidatorInfo{
			Id:                           info.Id,
			Network:                      info.Network,
			Account:                      info.Account,
			Name:                         info.Name,
			WwwUrl:                       info.WwwUrl,
			Details:                      info.Details,
			AvatarUrl:                    info.AvatarUrl,
			CreatedAt:                    info.CreatedAt,
			UpdatedAt:                    info.UpdatedAt,
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

	return ll, nil
}

func (h *SolanaStakeHandler) ValidatorInfoById(Id int64) (*ValidatorInfo, error) {
	info, err := repo.Validator(h.dbs).FindById(Id)
	if err != nil {
		return nil, err
	}

	return &ValidatorInfo{
		Id:                           info.Id,
		Network:                      info.Network,
		Account:                      info.Account,
		Name:                         info.Name,
		WwwUrl:                       info.WwwUrl,
		Details:                      info.Details,
		AvatarUrl:                    info.AvatarUrl,
		CreatedAt:                    info.CreatedAt,
		UpdatedAt:                    info.UpdatedAt,
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
	}, nil
}

func (h *SolanaStakeHandler) StakeActivationStatus(StakePublicBase58 string) (string, uint64, uint64, error) {
	return h.b.StakeActivationStatus(StakePublicBase58)
}

func (h *SolanaStakeHandler) StakeAccountInfo(StakePublicBase58 string) (*client.StakeAccountRsp, error) {
	return h.b.StakeAccountInfo(StakePublicBase58)
}

func (h *SolanaStakeHandler) CreateStakeAccount(SignerPrivateBase58 string, Amount decimal.Decimal) (string, string, error) {
	txHash, stakePublicBase58, err := h.b.CreateStakeAccount(SignerPrivateBase58, Amount)
	if err != nil {
		return "", "", err
	}

	return txHash, stakePublicBase58, nil
}

func (h *SolanaStakeHandler) DelegateStake(SignerPrivateBase58 string, StakePublicBase58 string, VotePublicBase58 string) (string, error) {
	txHash, err := h.b.DelegateStake(SignerPrivateBase58, StakePublicBase58, VotePublicBase58)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

func (h *SolanaStakeHandler) DeactivateStake(SignerPrivateBase58 string, StakePublicBase58 string) (string, error) {
	txHash, err := h.b.DeactivateStake(SignerPrivateBase58, StakePublicBase58)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

func (h *SolanaStakeHandler) WithdrawStake(SignerPrivateBase58 string, StakePublicBase58 string, Amount decimal.Decimal) (string, error) {
	txHash, err := h.b.WithdrawStake(SignerPrivateBase58, StakePublicBase58, Amount)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

func (h *SolanaStakeHandler) IsInit() bool {
	h.Lock()
	defer h.Unlock()

	return h.isInit()
}

func (h *SolanaStakeHandler) isInit() bool {
	return h.init
}

package account

import (
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/truekupo/cluster/account-srv/lib/blockchain"
	"github.com/truekupo/cluster/account-srv/lib/blockchain/ethereum"
	"github.com/truekupo/cluster/account-srv/lib/blockchain/solana"
	"github.com/truekupo/cluster/account-srv/lib/config"
	"github.com/truekupo/cluster/account-srv/lib/db/middleware"

	"github.com/truekupo/cluster/lib/logger"
)

type AccountHandler struct {
	conf *config.Config
	dsd  *middleware.DSD
	b    map[string]blockchain.Blockachain
}

var (
	log *logrus.Entry = nil
)

func NewAccountHandler(conf *config.Config) (*AccountHandler, error) {
	log = logger.LogModule("account-handler")

	a := AccountHandler{
		conf: conf,
		dsd:  nil,
		b:    map[string]blockchain.Blockachain{},
	}

	log.Info(conf.Settings)

	for _, cc := range conf.Settings.Chains {
		log.Info("Add " + cc.Symbol)
		switch cc.Symbol {
		case "ETH":
			e, err := ethereum.New(&cc)
			if err != nil {
				log.Error("account-handler: ", err)
			}

			a.b[cc.Symbol] = e
			a.b[cc.Name] = e
		case "SOL":
			s, err := solana.New(&cc)
			if err != nil {
				log.Error("account-handler: ", err)
			}

			a.b[cc.Symbol] = s
			a.b[cc.Name] = s
		}
	}

	return &a, nil
}

func (h *AccountHandler) NewMnemonic(entropy_size int, lang_id int32) (string, error) {
	log.Debug("account:handler:NewMnemonic")

	return blockchain.NewMnemonic(entropy_size, lang_id)
}

func (h *AccountHandler) GetSeed(mnemonic string, password string) ([]byte, error) {
	log.Debug("account:handler:GetSeed")

	return blockchain.GetSeed(mnemonic, password)
}

func (h *AccountHandler) GetMasterKey(mnemonic string, password string) ([]byte, error) {
	log.Debug("account:handler:GetMasterKey")

	return blockchain.GetMasterKey(mnemonic, password)
}

func (h *AccountHandler) SeedDeriveToPublicKeyHex(coin string, seed []byte, path string) (string, error) {
	log.Debug("account:handler:SeedDeriveToPublicKeyHex")

	s, ok := h.b[coin]
	if !ok {
		return "", fmt.Errorf("account:handler:SeedDeriveToPublicKeyHex coint code invalid")
	}

	return s.SeedDeriveToPublicKeyHex(seed, path)
}

func (h *AccountHandler) SeedDeriveToPrivateKeyHex(coin string, seed []byte, path string) (string, error) {
	log.Debug("account:handler:SeedDeriveToPrivateKeyHex")

	s, ok := h.b[coin]
	if !ok {
		return "", fmt.Errorf("account:handler:SeedDeriveToPrivateKeyHex coint code invalid")
	}

	return s.SeedDeriveToPrivateKeyHex(seed, path)
}

func (h *AccountHandler) SeedDeriveToAccountData(coin string, seed []byte, path string) (string, string, string, error) {
	log.Debug("account:handler:SeedDeriveToAccount")

	s, ok := h.b[coin]
	if !ok {
		return "", "", "", fmt.Errorf("account:handler:SeedDeriveToAccount coint code invalid")
	}

	return s.SeedDeriveToAccountData(seed, path)
}

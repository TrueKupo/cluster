package ethereum

import (
	"github.com/sirupsen/logrus"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/miguelmota/go-ethereum-hdwallet"

	"github.com/truekupo/cluster/account-srv/lib/blockchain/basechain"
	gc "github.com/truekupo/cluster/lib/config"
	"github.com/truekupo/cluster/lib/logger"
	ethutils "github.com/truekupo/cluster/lib/utils/ethereum"
)

type Ethereum struct {
	basechain.BaseChain

	hdnet *chaincfg.Params
}

var (
	log *logrus.Entry = nil
)

func New(conf *gc.Chain) (*Ethereum, error) {
	log = logger.LogModule("ethereum:new")

	e := &Ethereum{
		BaseChain: basechain.New(conf),
	}

	e.init_blockchain()

	return e, nil
}

func (e *Ethereum) init_blockchain() {
	switch e.Network() {
	case "mainnet":
		e.hdnet = &chaincfg.MainNetParams
	case "testnet":
		e.hdnet = &chaincfg.TestNet3Params
	default:
		e.hdnet = &chaincfg.TestNet3Params
	}
}

// "m/44'/60'/0'/0" => return ECDSA public key in hex string
func (e *Ethereum) SeedDeriveToPublicKeyHex(Seed []byte, Path string) (string, error) {
	wallet, err := hdwallet.NewFromSeed(Seed)
	if err != nil {
		return "", err
	}

	dp, err := hdwallet.ParseDerivationPath(Path)
	if err != nil {
		return "", err
	}

	account, err := wallet.Derive(dp, true)
	if err != nil {
		return "", err
	}

	public_hex, err := wallet.PublicKeyHex(account)
	if err != nil {
		return "", err
	}

	return public_hex, nil
}

// "m/44'/60'/0'/0" => return ECDSA private key in hex string
func (e *Ethereum) SeedDeriveToPrivateKeyHex(Seed []byte, Path string) (string, error) {
	wallet, err := hdwallet.NewFromSeed(Seed)
	if err != nil {
		return "", err
	}

	dp, err := hdwallet.ParseDerivationPath(Path)
	if err != nil {
		return "", err
	}

	account, err := wallet.Derive(dp, true)
	if err != nil {
		return "", err
	}

	private_hex, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return "", err
	}

	return private_hex, nil
}

func (e *Ethereum) SeedDeriveToAccountData(Seed []byte, Path string) (string, string, string, error) {
	wallet, err := hdwallet.NewFromSeed(Seed)
	if err != nil {
		return "", "", "", err
	}

	dp, err := hdwallet.ParseDerivationPath(Path)
	if err != nil {
		return "", "", "", err
	}

	account, err := wallet.Derive(dp, true)
	if err != nil {
		return "", "", "", err
	}

	private_hex, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return "", "", "", err
	}

	public_hex, err := wallet.PublicKeyHex(account)
	if err != nil {
		return "", "", "", err
	}

	addr, err := ethutils.PublicKeyHexToAddress(public_hex)
	if err != nil {
		return "", "", "", err
	}

	return private_hex, public_hex, addr.Hex(), nil
}

package ethereum

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
)

// PrivateKeyHexToECDSA is a helper function for converting a hexadecimal representation of a private key into ECDSA format.
func PrivateKeyHexToECDSA(privateKeyHex string) (*ecdsa.PrivateKey, error) {
	privateKey, err := ethcrypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// PublicKeyHexToECDSA is a helper function for converting a hexadecimal representation of a public key into ECDSA format.
func PublicKeyHexToECDSA(publicKeyHex string) (*ecdsa.PublicKey, error) {
	publicKeyBytes, err := hexutil.Decode("0x04" + publicKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey, err := ethcrypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}

// PublicKeyHexToAddress is a helper function for converting a hexadecimal representation of a public key into common.Address.
func PublicKeyHexToAddress(publicKeyHex string) (*common.Address, error) {
	publicKeyBytes, err := hexutil.Decode("0x04" + publicKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey, err := ethcrypto.UnmarshalPubkey(publicKeyBytes)
	if err != nil {
		return nil, err
	}

	addr := crypto.PubkeyToAddress(*publicKey)

	return &addr, nil
}

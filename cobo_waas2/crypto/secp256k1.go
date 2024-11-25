package crypto

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

type Secp256k1Signer struct {
	PrivateKey string
}

func (signer Secp256k1Signer) Sign(hash []byte) string {
	apiSecret, _ := hex.DecodeString(signer.PrivateKey)
	key, _ := btcec.PrivKeyFromBytes(apiSecret)
	sig := ecdsa.Sign(key, hash)
	return hex.EncodeToString(sig.Serialize())
}

func (signer Secp256k1Signer) GetPublicKey() string {
	apiSecret, _ := hex.DecodeString(signer.PrivateKey)
	key, _ := btcec.PrivKeyFromBytes(apiSecret)
	return hex.EncodeToString(key.PubKey().SerializeCompressed())
}

func GenerateSecp256k1ApiKey() (apiKey, apiSecret string, err error) {
	apiSecretBytes := make([]byte, 32)
	if _, err := rand.Read(apiSecretBytes); err != nil {
		return "", "", err
	}

	sk, _ := btcec.PrivKeyFromBytes(apiSecretBytes)

	apiKey = hex.EncodeToString(sk.PubKey().SerializeCompressed())
	apiSecret = hex.EncodeToString(apiSecretBytes)
	return apiKey, apiSecret, nil
}

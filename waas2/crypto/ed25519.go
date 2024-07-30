package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
)

type Ed25519Signer struct {
	Secret string
}

func (signer Ed25519Signer) Sign(hash []byte) string {
	apiSecret, _ := hex.DecodeString(signer.Secret)
	sig := ed25519.Sign(ed25519.NewKeyFromSeed(apiSecret), hash)
	return hex.EncodeToString(sig)
}

func (signer Ed25519Signer) GetPublicKey() string {
	apiSecret, _ := hex.DecodeString(signer.Secret)
	sk := ed25519.NewKeyFromSeed(apiSecret)
	pk := sk.Public().(ed25519.PublicKey)
	return hex.EncodeToString(pk)
}

func GenerateApiKey() (apiKey, apiSecret string, err error) {
	pk, sk, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}
	apiSecret = hex.EncodeToString(sk.Seed())
	apiKey = hex.EncodeToString(pk)
	return
}

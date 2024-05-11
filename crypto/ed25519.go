package crypto

import (
	"crypto/ed25519"
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

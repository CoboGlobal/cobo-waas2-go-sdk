package crypto

type ApiSigner interface {
	Sign(hash []byte) string
	GetPublicKey() string
}

package crypto

import (
	"crypto/sha256"
	"io"
	"net/http"
	"strings"
)

func GetUnsignedDigest(req *http.Request, body io.ReadSeeker, timestamp string) ([]byte, error) {
	content, err := BuildUnsignedContent(req, body, timestamp)
	if err != nil {
		return nil, err
	}
	return Hash256x2(content), nil
}

func Hash256x2(s string) []byte {
	hash1 := sha256.Sum256([]byte(s))
	hash2 := sha256.Sum256(hash1[:])
	return hash2[:]
}

func BuildUnsignedContent(req *http.Request, body io.ReadSeeker, timestamp string) (string, error) {
	bodyStr, err := io.ReadAll(body)
	if err != nil {
		return "", err
	}

	return strings.Join([]string{
		strings.ToUpper(req.Method),
		req.URL.Path,
		timestamp,
		req.URL.Query().Encode(),
		string(bodyStr),
	}, "|"), nil
}

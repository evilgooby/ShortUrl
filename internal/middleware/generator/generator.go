package generator

import (
	"ShortUrl/internal/middleware/errorHandling"
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
)

func GenerateShortLink(initialLink string) (string, error) {
	urlHashBytes := sha256.Sum256([]byte(initialLink))
	generatedNumber := new(big.Int).SetBytes(urlHashBytes[:]).Uint64()
	shortStr, err := encodeBase58([]byte(fmt.Sprintf("%d", generatedNumber)))
	if err != nil {
		return "", errorHandling.ErrGenerateShortUrl
	}
	return shortStr[:8], nil
}

func encodeBase58(bytes []byte) (string, error) {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

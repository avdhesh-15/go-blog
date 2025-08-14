package util

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
)

func GenerateKey() (*ecdsa.PrivateKey, error) {
	pemStr := strings.ReplaceAll(os.Getenv("JWT_SECRET"), `\n`, "\n")
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block")
	}
	return x509.ParseECPrivateKey(block.Bytes)
}

package interfaces

import (
	"crypto"
	"crypto/x509"
)

// IConfig is an interface for configuration
type ICrypto interface {
	GetPrivateKey() (crypto.PrivateKey, x509.Certificate)
}

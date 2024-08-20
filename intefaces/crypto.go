package interfaces

import "crypto"

// IConfig is an interface for configuration
type ICrypto interface {
	GetPrivateKey() crypto.PrivateKey
}

package mesh

import (
	"context"
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/pnet"
)

// exported functions are shared here
type DiscoConfiguration struct {
	Enabled    bool
	Rendezvous string
}
type IdentityConfiguration struct {
	PrivateKey []byte
}
type ListenConfiguration struct {
	Port       uint16
	Address    string
	InetFamily string
}
type MeshConfiguration struct {
	EnableRelay  bool
	Insecure     bool
	PreSharedKey string
}

type NetworkConfiguration struct {
	DiscoConfig    DiscoConfiguration
	ListenConfig   ListenConfiguration
	IdentityConfig IdentityConfiguration
	MeshConfig     MeshConfiguration
}

func makeMultiAddr(config NetworkConfiguration) string {
	return fmt.Sprintf("/%s/%s/tcp/%d", config.ListenConfig.InetFamily, config.ListenConfig.Address, config.ListenConfig.Port)
}
func generatePrivateKey() crypto.PrivKey {
	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 4096, rand.Reader)
	if err != nil {
		panic(err)
	}
	return priv
}

func buildBaseOptions(config NetworkConfiguration) []libp2p.Option {
	key := generatePrivateKey()
	return []libp2p.Option{
		libp2p.ListenAddrStrings(makeMultiAddr(config)),
		libp2p.Identity(key),
	}
}
func buildSecureOptions(config NetworkConfiguration) []libp2p.Option {
	if config.MeshConfig.Insecure {
		return []libp2p.Option{libp2p.NoSecurity}
	}
	// Load or generate a pre-shared key
	psk, err := pnet.DecodeV1PSK(strings.NewReader(config.MeshConfig.PreSharedKey))
	if err != nil {
		panic(err)
	}

	return []libp2p.Option{
		libp2p.PrivateNetwork(psk),
	}
}

func buildRelayOptions(config NetworkConfiguration) []libp2p.Option {
	if !config.MeshConfig.EnableRelay {
		return []libp2p.Option{}
	}
	return []libp2p.Option{
		libp2p.EnableRelay(),
	}
}

func buildOptions(config NetworkConfiguration) []libp2p.Option {
	baseOptions := buildBaseOptions(config)
	relayOptions := buildRelayOptions(config)
	secureOptions := buildSecureOptions(config)

	allOptions := append(baseOptions, relayOptions...)
	allOptions = append(allOptions, secureOptions...)

	return allOptions
}

func makeHost(config NetworkConfiguration) (host.Host, error) {
	opts := buildOptions(config)
	return libp2p.New(opts...)
}

func StartNet(context context.Context, config NetworkConfiguration) {
	host, hostErr := makeHost(config)
	if hostErr != nil {
		panic(hostErr)
	}
	startDiscovery(host, config)
}

package mesh

import (
	golog "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p/core/host"
)

func startDiscovery(host host.Host, config NetworkConfiguration) {
	if config.DiscoConfig.Enabled {
		logger := golog.Logger("discovery")
		logger.Info("Starting discovery service")
		initMDNS(host, config.DiscoConfig.Rendezvous)
	}
}

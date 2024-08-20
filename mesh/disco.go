package mesh

import "github.com/libp2p/go-libp2p/core/host"

func startDiscovery(host host.Host, config NetworkConfiguration) {
	if config.DiscoConfig.Enabled {
		initMDNS(host, config.DiscoConfig.Rendezvous)
	}
}

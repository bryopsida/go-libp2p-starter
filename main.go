package main

import (
	"context"

	"github.com/bryopsida/go-libp2p-starter/config"
	"github.com/bryopsida/go-libp2p-starter/libp2p"
)

func main() {
	// goals

	// 1) zero-conf discoevry of peers in the same network
	// 2) pub/sub messaging between peers
	// 3) ability to use relays to connect to peers outside the network
	// 4) expose service implementations to peers (increment service)
	// 5) stable peer id after first spin up
	// 6) support using s3 buckets for bootstraping
	// 7) support pre-shared key bootstraping for secure networks

	ctx, cancel := context.WithCancel(context.Background())
	cfg := config.NewViperConfig()
	defer cancel()
	config := libp2p.NetworkConfiguration{
		DiscoConfig: libp2p.DiscoConfiguration{
			Enabled:    true,
			Rendezvous: "libp2p-starter",
		},
	}
	libp2p.StartNet(ctx, config)

}

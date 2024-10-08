package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bryopsida/go-libp2p-starter/config"
	"github.com/bryopsida/go-libp2p-starter/mesh"
	golog "github.com/ipfs/go-log/v2"
)

func init() {
	golog.SetAllLoggers(golog.LevelDebug) // Change to INFO for extra info
	golog.SetupLogging(golog.Config{
		Format: golog.ColorizedOutput,
		Stdout: true,
		Level:  golog.LevelDebug,
	})
}

func main() {
	// goals

	// 1) zero-conf discoevry of peers in the same network
	// 2) pub/sub messaging between peers
	// 3) ability to use relays to connect to peers outside the network
	// 4) expose service implementations to peers (increment service)
	// 5) stable peer id after first spin up
	// 6) support using s3 buckets for bootstraping
	// 7) support pre-shared key bootstraping for secure networks

	logger := golog.Logger("main")
	ctx, cancel := context.WithCancel(context.Background())
	cfg := config.NewViperConfig()

	defer cancel()
	config := mesh.NetworkConfiguration{
		DiscoConfig: mesh.DiscoConfiguration{
			Enabled:    true,
			Rendezvous: "libp2p-starter",
		},
		ListenConfig: mesh.ListenConfiguration{
			Port:       cfg.GetListenPort(),
			Address:    cfg.GetListenAddress(),
			InetFamily: cfg.GetInetFamily(),
		},
		MeshConfig: mesh.MeshConfiguration{
			EnableRelay: true,
			Insecure:    false,
			PreSharedKey: `/key/swarm/psk/1.0.0/
/base16/
0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef`,
		},
		IdentityConfig: mesh.IdentityConfiguration{
			PrivateKey: nil,
		},
	}
	mesh.StartNet(ctx, config)

	// Set up signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	// Wait for a signal
	sig := <-sigChan
	logger.Info("Received signal", "signal", sig)
}

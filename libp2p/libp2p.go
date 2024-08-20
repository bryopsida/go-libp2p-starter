package libp2p

// exported functions are shared here
type DiscoConfiguration struct {
	Enabled    bool
	Rendezvous string
}
type NetworkConfiguration struct {
	DiscoConfig DiscoConfiguration
}

func StartNet(config NetworkConfiguration) {

}

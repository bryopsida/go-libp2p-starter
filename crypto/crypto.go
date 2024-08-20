package crypto

type crypto struct {
	config interfaces.IConfig
}

func (*crypto) GetPrivateKey() crypto.PrivateKey {
	return nil
}

func NewCrypt(config interfaces.IConfig) interfaces.ICrypto {
	return &crypto{config}
}

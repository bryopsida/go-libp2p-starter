package config

import (
	"github.com/spf13/viper"
)

const (
	listenAddress    = "listen.address"
	listenPort       = "listen.port"
	listenInetFamily = "listen.inetfamily"
	pkcs12Path       = "keystore.path"
	pkcs12Password   = "keystore.password"
)

type viperConfig struct {
	viper *viper.Viper
}

// NewViperConfig creates a new viperConfig instance
func NewViperConfig() *viperConfig {
	config := viperConfig{viper: viper.New()}
	config.setDefaults()
	config.initialize()
	return &config
}

func (c *viperConfig) setDefaults() {
	c.viper.SetDefault(listenAddress, "0.0.0.0")
	c.viper.SetDefault(listenPort, "1234")
	c.viper.SetDefault(listenInetFamily, "ip4")
	c.viper.SetDefault(pkcs12Path, "keystore.p12")
	c.viper.SetDefault(pkcs12Password, "password")
}

func (c *viperConfig) initialize() {
	c.viper.SetConfigName("config")
	c.viper.SetConfigType("yaml")
	c.viper.AddConfigPath(".")
	c.viper.AutomaticEnv()
}

// Gets the path to the keystore
func (c *viperConfig) GetPathToPKCS12() string {
	return c.viper.GetString(pkcs12Path)
}

// Gets the password to the keystore
func (c *viperConfig) GetPKCS12Password() string {
	return c.viper.GetString(pkcs12Password)
}

// Gets the listen address
func (c *viperConfig) GetListenAddress() string {
	return c.viper.GetString(listenAddress)
}

func (c *viperConfig) GetListenPort() uint16 {
	return c.viper.GetUint16(listenPort)
}

func (c *viperConfig) GetInetFamily() string {
	return c.viper.GetString(listenInetFamily)
}

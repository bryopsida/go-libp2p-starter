package interfaces

// IConfig is an interface for configuration
type IConfig interface {
	GetListenAddress() string
	GetPathToPKCS12() string
	GetPKCS12Password() string
}

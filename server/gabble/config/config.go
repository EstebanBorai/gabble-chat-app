package config

// Config represents a Gabble Server
// configuration
type Config interface {
	GetHost() string
	GetPort() string
	GetLogLevel() int
}

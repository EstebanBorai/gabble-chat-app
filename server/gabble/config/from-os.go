package config

// OSEnvConfig satisfies Gabble server
// configuration from OS environment variables
type OSEnvConfig struct {
	port       string
	host       string
	clientHost string
	clientPort string
	logLevel   int
}

// GetHost returns the "SERVER_HOST" value from the .env file
func (conf *OSEnvConfig) GetHost() string {
	return conf.host
}

// GetPort returns the "SERVER_PORT" value from the .env file
func (conf *OSEnvConfig) GetPort() string {
	return conf.port
}

// GetClientHost returns the "SERVER_PORT" value from the .env file
func (conf *OSEnvConfig) GetClientHost() string {
	return conf.clientHost
}

// GetClientPort returns the "SERVER_PORT" value from the .env file
func (conf *OSEnvConfig) GetClientPort() string {
	return conf.clientPort
}

// GetLogLevel returns the "LOG_LEVEL" value from the .env file
func (conf *OSEnvConfig) GetLogLevel() int {
	return conf.logLevel
}

// FromOsEnv reads configuration from the OS environment variables
func FromOsEnv() (*OSEnvConfig, error) {
	var conf *OSEnvConfig = new(OSEnvConfig)

	conf.host = MustGetEnv("SERVER_HOST")
	conf.port = MustGetEnv("SERVER_PORT")
	conf.clientHost = MustGetEnv("CLIENT_HOST")
	conf.clientPort = MustGetEnv("CLIENT_PORT")
	conf.logLevel = MustGetEnvInt("LOG_LEVEL")

	return conf, nil
}

package config

import (
	"github.com/joho/godotenv"
)

// EnvConfig satisfies Gabble server
// configuration from a .env file
type EnvConfig struct {
	port       string
	host       string
	clientHost string
	clientPort string
	logLevel   int
}

// GetHost returns the "SERVER_HOST" value from the .env file
func (conf *EnvConfig) GetHost() string {
	return conf.host
}

// GetPort returns the "SERVER_PORT" value from the .env file
func (conf *EnvConfig) GetPort() string {
	return conf.port
}

// GetClientHost returns the "SERVER_PORT" value from the .env file
func (conf *EnvConfig) GetClientHost() string {
	return conf.clientHost
}

// GetClientPort returns the "SERVER_PORT" value from the .env file
func (conf *EnvConfig) GetClientPort() string {
	return conf.clientPort
}

// GetLogLevel returns the "LOG_LEVEL" value from the .env file
func (conf *EnvConfig) GetLogLevel() int {
	return conf.logLevel
}

// FromEnvFile reads configuration from a .env file located
// in the CWD and returns a EnvConfig
func FromEnvFile() (*EnvConfig, error) {
	var conf *EnvConfig = new(EnvConfig)

	err := godotenv.Load()

	if err != nil {
		// Provide the capability to create a .env file if
		// the file doesnt exists, replacing the `build-env.sh`
		// script.
		// https://github.com/joho/godotenv/issues/101
		return nil, err
	}

	conf.host = MustGetEnv("SERVER_HOST")
	conf.port = MustGetEnv("SERVER_PORT")
	conf.clientHost = MustGetEnv("CLIENT_HOST")
	conf.clientPort = MustGetEnv("CLIENT_PORT")
	conf.logLevel = MustGetEnvInt("LOG_LEVEL")

	return conf, nil
}

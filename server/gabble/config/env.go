package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// EnvConfig satisfies Gabble server
// configuration from a .env file
type EnvConfig struct {
	port string
	host string
}

// GetHost returns the "SERVER_HOST" value from the .env file
func (conf *EnvConfig) GetHost() string {
	return conf.host
}

// GetPort returns the "SERVER_PORT" value from the .env file
func (conf *EnvConfig) GetPort() string {
	return conf.port
}

// FromEnv reads configuration from a .env file located
// in the CWD and returns a EnvConfig
func FromEnv() (*EnvConfig, error) {
	var conf *EnvConfig = new(EnvConfig)

	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	conf.host = mustGetEnv("SERVER_HOST")
	conf.port = mustGetEnv("SERVER_PORT")

	return conf, nil
}

// Gather values from .env file.
// if a value is empty panics.
func mustGetEnv(key string) string {
	var value string = os.Getenv(key)

	if value == "" {
		panic(fmt.Sprintf("The key %s is reqired in the .env file", key))
	}

	return value
}

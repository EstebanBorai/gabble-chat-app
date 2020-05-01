package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// EnvConfig satisfies Gabble server
// configuration from a .env file
type EnvConfig struct {
	port     string
	host     string
	logLevel int
}

// GetHost returns the "SERVER_HOST" value from the .env file
func (conf *EnvConfig) GetHost() string {
	return conf.host
}

// GetPort returns the "SERVER_PORT" value from the .env file
func (conf *EnvConfig) GetPort() string {
	return conf.port
}

// GetLogLevel returns the "LOG_LEVEL" value from the .env file
func (conf *EnvConfig) GetLogLevel() int {
	return conf.logLevel
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
	conf.logLevel = mustGetEnvInt("LOG_LEVEL")

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

func mustGetEnvInt(key string) int {
	var value string = mustGetEnv(key)

	i, err := strconv.Atoi(value)

	if err != nil {
		panic(fmt.Sprintf("The key %s has an invalid value %s. Expected a valid integer", key, value))
	}

	return i
}

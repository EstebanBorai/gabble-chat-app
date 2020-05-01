package config

import (
	"fmt"
	"os"
	"strconv"
)

// MustGetEnv gather environment variable and convert it to an
// int value. Panics if it fail
func MustGetEnv(key string) string {
	var value string = os.Getenv(key)

	if value == "" {
		panic(fmt.Sprintf("The key %s is reqired in the .env file", key))
	}

	return value
}

// MustGetEnvInt gather environment variable and convert it to an
// int value. Panics if it fail
func MustGetEnvInt(key string) int {
	var value string = MustGetEnv(key)

	i, err := strconv.Atoi(value)

	if err != nil {
		panic(fmt.Sprintf("The key %s has an invalid value %s. Expected a valid integer", key, value))
	}

	return i
}

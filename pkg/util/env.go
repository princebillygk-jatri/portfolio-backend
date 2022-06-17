package util

import (
	"errors"
	"os"
)

// MustHaveEnv gets a env that is required to be set. It it is not set it will
// panic.
func MustHaveEnv(key string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	panic(errors.New("Failed to get env \"" + key + "\"\n"))
}

// Env gets a env that may or may not be set. If it is not set then it will
// return a fallback string given as parameter.
func Env(key, fallback string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return fallback
}

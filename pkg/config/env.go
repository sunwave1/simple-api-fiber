package config

import (
	"errors"
	"os"
)

func GetEnv(key, defaultValue string) string {
	conf := os.Getenv(key)
	if conf != "" {
		return conf
	}
	return defaultValue
}

func CheckFileEnvExist() bool {
	_, err := os.Stat("./.env")
	return !errors.Is(err, os.ErrNotExist)
}

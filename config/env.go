package config

import (
	"os"
)

func LoadEnvConfig(name string) (string) {

	return os.Getenv(name)

}
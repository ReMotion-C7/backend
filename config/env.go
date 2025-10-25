package config

import (
	"ReMotion-C7/constant"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	envPaths := []string{
		filepath.Join(".", ".env"),
		filepath.Join("..", ".env"),
		filepath.Join("..", "..", ".env"),
	}

	for _, path := range envPaths {
		if _, err := os.Stat(path); err == nil {
			err = godotenv.Load(path)
			if err == nil {
				log.Println(constant.SuccessLoadEnvFile)
				return
			} 
		}
	}

	log.Fatalln(constant.ErrReadEnvFile)

}

func LoadEnvConfig(name string) string {

	return os.Getenv(name)

}

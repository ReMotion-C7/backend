package config

import (
	"ReMotion-C7/constant"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	err := godotenv.Load(".env")
	if err != nil {
		return
	}

	log.Println(constant.SuccessLoadEnvFile)

}

func LoadEnvConfig(name string) string {

	return os.Getenv(name)

}

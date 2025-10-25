package main

import (
	"ReMotion-C7/config"
	"ReMotion-C7/cmd/reset/clear"
)

func main() {

	config.LoadEnv()

	config.ConnectDatabase()

	database := config.GetDatabase()

	clear.ClearDatabase(database)

}

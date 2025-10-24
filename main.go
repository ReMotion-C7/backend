package main

import (
	"ReMotion-C7/config"
	"ReMotion-C7/constant"
	"ReMotion-C7/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)


func main() {

	app := fiber.New(fiber.Config{})

	err := godotenv.Load()
	if err != nil {
		log.Fatalf(constant.ErrReadEnvFile)
	}

	config.ConnectDatabase()
	config.CheckConnection()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	router.SetUp(app)

	log.Fatal(app.Listen(":8080"))

}
package main

import (
	"ReMotion-C7/config"
	"ReMotion-C7/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New(fiber.Config{})

	config.LoadEnv()

	config.ConnectDatabase()
	config.CheckConnection()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	router.SetUp(app)

	log.Fatal(app.Listen(":8080"))

}

package main

import (
	"ReMotion-C7/config"
	"ReMotion-C7/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New(fiber.Config{
		BodyLimit: 50 * 1024 * 1024,
	})

	config.LoadEnv()

	config.ConnectDatabase()
	config.CheckConnection()
	config.SetUpStorage()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	router.SetUp(app)

	log.Fatal(app.Listen(":8080"))

}

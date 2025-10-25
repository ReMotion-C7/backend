package router

import (
	"ReMotion-C7/app/controller"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	
	// AUTH ENPOINTS
	app.Post("/api/v1/login", controller.Login)
	app.Post("/api/v1/register", controller.Register)

}
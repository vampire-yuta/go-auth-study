package main

import (
	"main/database"
	"main/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello world")
	// })

	app.Listen(":80")
}

package routes

import (
	"main/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Home)
	app.Get("/other", controllers.Other)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Get("/api/logout", controllers.Logout)
	app.Post("/api/forgot", controllers.Forgot)
}

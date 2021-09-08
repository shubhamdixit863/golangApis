package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubhamdixit863/golangApis/handlers"
)

func RegisterRoutes(app *fiber.App) {

	app.Get("/signup", handlers.Login)

}

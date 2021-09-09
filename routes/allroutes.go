package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubhamdixit863/golangApis/handlers"
)

func RegisterRoutes(app *fiber.App) {
	/**
	Auth Routes
	*/

	app.Post("/signup", handlers.Signup)

	/*
		Listing Routes
	*/

	app.Post("/createListing", handlers.CreateListing)
	app.Get("/allListings", handlers.GetListing)

}

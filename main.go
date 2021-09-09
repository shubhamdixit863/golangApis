package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubhamdixit863/golangApis/database"
	"github.com/shubhamdixit863/golangApis/routes"
)

func main() {
	app := fiber.New()
	client := database.Connect("santo_vecino_db", "mongodb://localhost:27017")
	defer database.Disconnect(client) // Disconnecting once the main finished execution
	routes.RegisterRoutes(app)

	app.Listen(":3001")

}

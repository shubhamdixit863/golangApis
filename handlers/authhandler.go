package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubhamdixit863/golangApis/services"
)

func Login(c *fiber.Ctx) error {

	services.Register()

	return c.SendString("User Registered")

}

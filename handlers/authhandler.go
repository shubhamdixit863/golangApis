package handlers

import (
	"github.com/shubhamdixit863/golangApis/models"

	"github.com/gofiber/fiber/v2"
	"github.com/shubhamdixit863/golangApis/dao"
)

func Signup(c *fiber.Ctx) error {
	p := new(models.User)

	if err := c.BodyParser(p); err != nil {
		return err
	}
	_, err := dao.Register(p)
	if err != nil {

		return c.JSON(err)

	} else {
		response := models.Response{
			Message: "User Registered SuccessFully",
			Success: true,
		}
		return c.JSON(response)

	}

}

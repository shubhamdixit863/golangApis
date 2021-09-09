package handlers

import (
	"context"

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

func CreateListing(c *fiber.Ctx) error {
	p := new(models.Listings)

	if err := c.BodyParser(p); err != nil {
		return err
	}
	_, err := dao.CreateListing(p)
	if err != nil {

		return c.JSON(err)

	} else {
		response := models.Response{
			Message: "Listing Created SuccessFully",
			Success: true,
		}
		return c.JSON(response)

	}

}

func GetListing(c *fiber.Ctx) error {
	cursor, err := dao.GetListing()

	if err != nil {

		return c.JSON(err)

	} else {
		var results []models.Listings

		cursor.All(context.Background(), &results)
		response := models.Response{
			Message: "Query Success",
			Success: true,
			Output:  results,
		}
		return c.JSON(response)

	}
}

package handlers

import (
	"context"
	"fmt"

	"github.com/shubhamdixit863/golangApis/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
	"github.com/shubhamdixit863/golangApis/dao"
	"github.com/shubhamdixit863/golangApis/util"
)

func CreateListing(c *fiber.Ctx) error {

	file, err := c.FormFile("image")

	c.SaveFile(file, fmt.Sprintf("./uploads/%s", file.Filename))

	fileName := util.UploadS3(fmt.Sprintf("./uploads/%s", file.Filename))
	fmt.Println(fileName)

	p := new(models.Listings)

	if err := c.BodyParser(p); err != nil {
		return err
	}
	_, err = dao.CreateListing(p)
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

func GetListingById(c *fiber.Ctx) error {
	var response models.Response
	listing := new(models.Listings)
	objectId, err := primitive.ObjectIDFromHex(c.Params("id")) // converting string id int object Id
	if err != nil {
		response = models.Response{
			Message: "Error",
			Success: false,
			Output:  err.Error(),
		}

	}

	err = dao.GetListingById(objectId).Decode(listing)

	if err != nil {

		response = models.Response{
			Message: "Error",
			Success: false,
			Output:  err.Error(),
		}

	} else {

		response = models.Response{
			Message: "Query Success",
			Success: true,
			Output:  listing,
		}

	}
	return c.JSON(response)

}

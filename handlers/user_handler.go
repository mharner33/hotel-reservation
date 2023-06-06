package handlers

import (
	"github.com/mharner33/hotel-reservation/types"

	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {

	return c.JSON(c.App().Stack())
}

func HandleGetUser(c *fiber.Ctx) error{
	usr := types.User{
		FirstName: "Mike",
		LastName: "Jones",
	}
	return c.JSON(usr)
}

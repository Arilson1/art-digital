package main

import (
	"github.com/Arilson1/art-digital/data"
	"github.com/Arilson1/art-digital/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.Init()

	db.DB.AutoMigrate(&data.User{})

	db.DB.Create(&data.User{
		Name:  "Arilson",
		Email: "arilson@htomail.com",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")
}

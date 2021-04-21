package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/juliosinaysantos/go-simple-todo-app/database"
	"github.com/juliosinaysantos/go-simple-todo-app/utils"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	database.Connect()

	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Hello, world! 👋",
		})
	})

	app.Use(func(c *fiber.Ctx) error {
		url := c.Path()
		return fiber.NewError(fiber.StatusNotFound, "Not Found: "+url)
	})

	log.Fatal(app.Listen(":5000"))
}

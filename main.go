package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/juliosinaysantos/go-simple-todo-app/database"
)

func main() {
	app := fiber.New()

	database.Connect()

	app.Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Hello, world! 👋",
		})
	})

	log.Fatal(app.Listen(":5000"))
}

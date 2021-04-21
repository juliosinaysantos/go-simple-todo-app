package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetTodosHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"todos": nil,
	})
}

func CreateTodoHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "New ToDo",
	})
}

func GetTodoHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"id":   ctx.Params("id"),
		"todo": nil,
	})
}

func UpdateTodoHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"id":   ctx.Params("id"),
		"todo": nil,
	})
}

func DeleteTodoHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"id":      ctx.Params("id"),
		"message": "Delete ToDo",
	})
}

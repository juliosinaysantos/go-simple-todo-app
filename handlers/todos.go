package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juliosinaysantos/go-simple-todo-app/database"
)

func GetTodosHandler(ctx *fiber.Ctx) error {
	todos := database.GetAllTodos()
	return ctx.JSON(fiber.Map{
		"todos": todos,
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

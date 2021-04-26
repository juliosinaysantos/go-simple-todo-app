package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/juliosinaysantos/go-simple-todo-app/database"
	"github.com/juliosinaysantos/go-simple-todo-app/models"
)

func GetTodosHandler(ctx *fiber.Ctx) error {
	todos := database.GetAllTodos()
	return ctx.JSON(fiber.Map{
		"todos": todos,
	})
}

func CreateTodoHandler(ctx *fiber.Ctx) error {
	todo := new(models.Todo)

	if err := ctx.BodyParser(todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"stack":   err.Error(),
		})
	}

	if todo.Title == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Title is required",
		})
	}

	todo.ID = uint(database.CountTodos() + 1)
	todo.Completed = false
	todo.CreatedAt = time.Now().UTC()
	todo.UpdatedAt = time.Now().UTC()

	id := database.CreateTodo(todo)

	if id <= 0 {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return ctx.JSON(fiber.Map{
		"todo":    id,
		"message": "ToDo created successfully!",
	})
}

func GetTodoHandler(ctx *fiber.Ctx) error {
	id := ctx.Locals("id").(int)

	t := database.GetSingleTodo(id)

	if t == nil {
		return fiber.NewError(fiber.StatusNotFound, "ToDo Not Found")
	}
	return ctx.JSON(fiber.Map{
		"todo": t,
	})
}

func UpdateTodoHandler(ctx *fiber.Ctx) error {
	id := ctx.Locals("id").(int)

	t := database.GetSingleTodo(id)

	if t == nil {
		return fiber.NewError(fiber.StatusNotFound, "ToDo Not Found")
	}

	todo := new(models.Todo)

	if err := ctx.BodyParser(todo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
			"stack":   err.Error(),
		})
	}

	if todo.Title == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Title is required",
		})
	}

	t.Title = todo.Title
	t.Content = todo.Content
	t.Completed = todo.Completed
	t.UpdatedAt = time.Now().UTC()

	return ctx.JSON(fiber.Map{
		"id":      t.ID,
		"message": "ToDo updated successfully!",
	})
}

func DeleteTodoHandler(ctx *fiber.Ctx) error {
	id := ctx.Locals("id").(int)

	t := database.GetSingleTodo(id)

	if t == nil {
		return fiber.NewError(fiber.StatusNotFound, "ToDo Not Found")
	}

	now := time.Now().UTC()

	t.DeletedAt = &now

	return ctx.JSON(fiber.Map{
		"id":      t.ID,
		"message": "ToDo deleted successfully!",
	})
}

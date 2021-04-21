package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juliosinaysantos/go-simple-todo-app/handlers"
)

func TodosRouter(app fiber.Router) {
	router := app.Group("/todos")

	router.Get("/", handlers.GetTodosHandler)
	router.Post("/", handlers.CreateTodoHandler)
	router.Get("/:id", handlers.GetTodoHandler)
	router.Put("/:id", handlers.UpdateTodoHandler)
	router.Delete("/:id", handlers.DeleteTodoHandler)
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/juliosinaysantos/go-simple-todo-app/handlers"
	"github.com/juliosinaysantos/go-simple-todo-app/middlewares"
)

func TodosRouter(app fiber.Router) {
	router := app.Group("/todos")

	router.Get("/", handlers.GetTodosHandler)
	router.Post("/", handlers.CreateTodoHandler)
	router.Get("/:id", middlewares.ParseIDParam, handlers.GetTodoHandler)
	router.Put("/:id", middlewares.ParseIDParam, handlers.UpdateTodoHandler)
	router.Delete("/:id", middlewares.ParseIDParam, handlers.DeleteTodoHandler)
}

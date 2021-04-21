package database

import (
	"github.com/juliosinaysantos/go-simple-todo-app/models"
)

func GetAllTodos() []models.Todo {
	todos := make([]models.Todo, 0)

	for _, todo := range db.todos {
		if todo.DeletedAt == nil {
			todos = append(todos, todo)
		}
	}

	return todos
}

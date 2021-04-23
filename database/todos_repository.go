package database

import (
	"github.com/juliosinaysantos/go-simple-todo-app/models"
)

func GetAllTodos() []*models.Todo {
	todos := make([]*models.Todo, 0)

	for _, todo := range db.todos {
		if todo.DeletedAt == nil {
			todos = append(todos, todo)
		}
	}

	return todos
}

func CountTodos() int {
	return len(db.todos)
}

func CreateTodo(todo *models.Todo) uint {
	mu.Lock()
	db.todos = append(db.todos, todo)
	mu.Unlock()
	return todo.ID
}

func GetSingleTodo(id int) *models.Todo {
	for _, t := range db.todos {
		if t.ID == uint(id) {
			return t
		}
	}
	return nil
}

package database

import (
	"fmt"
	"sync"

	"github.com/juliosinaysantos/go-simple-todo-app/models"
)

type DB struct {
	todos []models.Todo
}

var (
	db *DB
	mu sync.Mutex
)

func Connect() {
	db = &DB{
		todos: make([]models.Todo, 0),
	}
	fmt.Println("Connected with Database")
}

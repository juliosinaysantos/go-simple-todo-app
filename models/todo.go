package models

import "time"

type Todo struct {
	ID        uint       `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content,omitempty"`
	Completed bool       `json:"completed"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

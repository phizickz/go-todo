package storage

import "go-todo/internal/entities/task"

type TaskRepository interface {
	Create(task *task.Task) error
	FindByID(id int) (*task.Task, error)
}

type Repository struct {
	Task TaskRepository
}

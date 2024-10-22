package storage

import "go-todo/internal/entities/task"

type TaskRepository interface {
	Create(task *task.Task) error
	Update(task *task.Task) error
}

type Repository struct {
	Task TaskRepository
}

package service

import "go-todo/internal/entities/task"

type TaskService interface {
	Create(task *task.Task) error
	// Update(ctx context.Context, task *task.Task) error
	// Delete()
	FindByID(id int) (*task.Task, error)
}

type Service struct {
	TaskService TaskService
}

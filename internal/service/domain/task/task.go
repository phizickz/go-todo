package task

import (
	taskEntity "go-todo/internal/entities/task"
	"go-todo/internal/storage"
)

type TaskService struct {
	taskRepository storage.TaskRepository
}

func (s TaskService) Create(task *taskEntity.Task) error {
	return nil
}

func (s TaskService) Update(task *taskEntity.Task) error {
	return nil
}

func NewTaskService(taskRepository storage.TaskRepository) *TaskService {
	return &TaskService{
		taskRepository: taskRepository,
	}
}

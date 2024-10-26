package task

import (
	taskEntity "go-todo/internal/entities/task"
)

type TaskService interface {
	Create(taskEntity.Task) error
	GetAll() []*taskEntity.Task
}

type TaskController struct {
	service TaskService
}

func NewTaskController(service TaskService) *TaskController {
	return &TaskController{
		service: service,
	}
}

func (c *TaskController) CreateTask(task taskEntity.Task) error {
	return c.service.Create(task)
}

func (c *TaskController) GetAllTasks() []*taskEntity.Task {
	return c.service.GetAll()
}

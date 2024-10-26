package task

import (
	taskEntity "go-todo/internal/entities/task"
)

type TaskService struct {
	tasks []*taskEntity.Task
}

func (s *TaskService) Create(task *taskEntity.Task) error {
	return nil
}

func (s *TaskService) GetAll() []taskEntity.Task {
	return nil
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make([]*taskEntity.Task, 0),
	}
}

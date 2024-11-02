package task

import (
	taskEntity "go-todo/internal/entities/task"
)

type TaskService struct {
	tasks []taskEntity.Task
}

func (s *TaskService) Create(task taskEntity.Task) {
	s.tasks[len(s.tasks)] = task
}

func (s *TaskService) GetAll() []taskEntity.Task {
	return s.tasks
}

func NewTaskService() *TaskService {
	tasks := []taskEntity.Task{
		{
			Title: "abc",
			Body:  "def",
		},
		{
			Title: "hello",
			Body:  "world",
		},
		{
			Title: "my third",
			Body:  "card",
		},
	}
	return &TaskService{
		tasks: tasks,
	}
}

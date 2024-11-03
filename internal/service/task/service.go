package task

import (
	taskEntity "go-todo/internal/entities/task"
)

type TaskService struct {
	tasks map[int]taskEntity.Task
}

func (s *TaskService) Create(task taskEntity.Task) {
	s.tasks[len(s.tasks)] = taskEntity.Task{
		Title: task.Title,
		Body:  task.Body,
	}
}

func (s *TaskService) GetAll() map[int]taskEntity.Task {
	return s.tasks
}

func NewTaskService() *TaskService {
	tasks := make(map[int]taskEntity.Task)
	tasks[len(tasks)] = taskEntity.Task{
		Title: "abc",
		Body:  "def",
	}
	tasks[len(tasks)] = taskEntity.Task{
		Title: "hello",
		Body:  "world",
	}
	tasks[len(tasks)] = taskEntity.Task{
		Title: "my third",
		Body:  "card",
	}

	return &TaskService{
		tasks: tasks,
	}
}

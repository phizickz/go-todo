package task

import (
	taskEntity "go-todo/internal/entities/task"
)

type TaskService struct {
	tasks map[int]taskEntity.Task
}

func (s *TaskService) Create(task taskEntity.Task) {
	s.tasks[len(s.tasks)] = task
}

func (s *TaskService) GetAll() map[int]taskEntity.Task {
	return s.tasks
}

func NewTaskService() *TaskService {
	tasks := make(map[int]taskEntity.Task)
	tasks[0] = taskEntity.Task{Title: "abc", Body: "def"}
	tasks[1] = taskEntity.Task{Title: "mytask", Body: "mybody"}
	return &TaskService{
		tasks: tasks,
	}
}

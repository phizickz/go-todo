package task

import (
	"fmt"
	taskEntity "go-todo/internal/entities/task"
)

type TaskRepository struct {
	tasks []taskEntity.Task
}

func (r *TaskRepository) Create(task *taskEntity.Task) {
	r.tasks = append(r.tasks, task)
}

func (r *TaskRepository) FindByID(id int) (*taskEntity.Task, error) {
	for _, task := range r.tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("Task with ID %d not found", id)
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make([]taskEntity.Task, 0),
	}
}

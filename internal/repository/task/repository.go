package task

import (
	"errors"
	taskEntity "go-todo/internal/entities/task"
)

type TaskRepository struct {
	tasks   []taskEntity.Task
	counter int
}

// func (tr *TaskRepository) Create(title string, body string) error {
func (tr *TaskRepository) Create(task taskEntity.Task) error {
	tr.tasks = append(tr.tasks, taskEntity.Task{
		ID:    tr.counter,
		Title: task.Title,
		Body:  task.Body,
	},
	)
	tr.counter++
	return nil
}

func (tr *TaskRepository) Retrieve(id int) (taskEntity.Task, error) {
	for _, task := range tr.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return taskEntity.Task{}, errors.New("task not found")
}

func (tr *TaskRepository) RetrieveAll() []taskEntity.Task {
	return tr.tasks
}

func (tr *TaskRepository) Delete(id int) error {
	for i, task := range tr.tasks {
		if task.ID == id {
			tr.tasks = append(tr.tasks[:i], tr.tasks[i+1:]...)
			break
		}
	}
	return nil
}

func NewTaskRepository() *TaskRepository {
	tasks := []taskEntity.Task{
		{
			ID:    0,
			Title: "abc",
			Body:  "def",
		},
		{
			ID:    1,
			Title: "hello",
			Body:  "world",
		},
		{
			ID:    2,
			Title: "my third",
			Body:  "card",
		},
		{
			ID:    3,
			Title: "def",
			Body:  "abc",
		},
	}

	return &TaskRepository{
		tasks:   tasks,
		counter: len(tasks),
	}
}

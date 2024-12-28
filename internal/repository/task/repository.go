package task

import (
	"database/sql"
	"errors"
	"fmt"
	taskEntity "go-todo/internal/entities/task"
)

type TaskRepository struct {
	tasks   []taskEntity.Task
	db      *sql.DB
	counter int
}

func (tr *TaskRepository) Update(task taskEntity.Task) {
	for i, t := range tr.tasks {
		if task.ID == t.ID {
			tr.tasks[i] = task
		}
	}
}

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
	query := "SELECT * FROM tasks"
	rows, err := tr.db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var tasks []taskEntity.Task
	for rows.Next() {
		var task taskEntity.Task
		_ = rows.Scan(&task.ID, &task.Title, &task.Body)
		// if err := rows.Scan(&task.ID, &task.Title, &task.Body); err != nil {
		// 	return nil, err
		// }
		tasks = append(tasks, task)
	}
	// if err := rows.Err(); err != nil {
	//   return nil, err
	// }
	return tasks
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

func NewTaskRepository(db *sql.DB) *TaskRepository {
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
		db:      db,
	}
}

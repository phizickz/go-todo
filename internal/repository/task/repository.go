package task

import (
	"database/sql"
	"fmt"
	taskEntity "go-todo/internal/entities/task"
	"log"
)

type TaskRepository struct {
	db *sql.DB
}

func (tr *TaskRepository) Update(task taskEntity.Task) {
	query := `
  UPDATE tasks
  SET title = $1, body = $2
  WHERE id = $3;
  `
	tx, err := tr.db.Begin()
	if err != nil {
		log.Println(err)
	}

	_, err = tx.Exec(query, task.Title, task.Body, task.ID)
	if err != nil {
		log.Println(err)
		tx.Rollback()
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}
}

func (tr *TaskRepository) Create(task taskEntity.Task) error {
	query := ` 
  INSERT INTO tasks (title, body)
  VALUES($1, $2);
  `
	_, err := tr.db.Exec(query, task.Title, task.Body)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (tr *TaskRepository) Retrieve(id int) taskEntity.Task {
	query := `
  SELECT id, title, body FROM tasks
  WHERE id=$1
  `
	var task taskEntity.Task
	err := tr.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Body)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No task found with the given ID.")
		} else {
			log.Println("Error querying task:", err)
		}
	}
	return task
}

func (tr *TaskRepository) RetrieveAll() ([]taskEntity.Task, error) {
	query := `
  SELECT id, title, body FROM tasks
  `
	rows, err := tr.db.Query(query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var tasks []taskEntity.Task
	for rows.Next() {
		var task taskEntity.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Body); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tr *TaskRepository) Delete(id int) error {
	query := `
  DELETE FROM tasks
  WHERE id = $1
  `
	tx, err := tr.db.Begin()
	if err != nil {
		log.Println(err)
	}

	_, err = tx.Exec(query, id)
	if err != nil {
		log.Println(err)
		tx.Rollback()
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
	}
	return nil
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

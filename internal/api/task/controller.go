package task

import (
	taskEntity "go-todo/internal/entities/task"
	taskService "go-todo/internal/service/task"
	"net/http"
)

type TaskController struct {
	service *taskService.TaskService
}

func NewTaskController(service *taskService.TaskService) *TaskController {
	return &TaskController{
		service: service,
	}
}

func (c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	task := taskEntity.Task{
		Title: r.Form("title"),
		Body:  r.Form("Body"),
	}
	c.service.Create(task)
}

func (c *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := c.service.GetAll()
}

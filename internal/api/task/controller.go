package task

import (
	taskEntity "go-todo/internal/entities/task"
	taskService "go-todo/internal/service/task"
	"go-todo/web/views"
	"net/http"

	"github.com/a-h/templ"
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
		Title: "abc",
		Body:  "def",
		// Title: r.Form("title"),
		// Body:  r.Form("Body"),
	}
	c.service.Create(task)
}

func (c *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	handler := templ.Handler(views.Tasks(c.service.GetAll()))
	handler.ServeHTTP(w, r)
}

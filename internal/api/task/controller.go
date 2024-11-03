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
	c.service.Create(taskEntity.Task{
		Title: r.Form["title"][0],
		Body:  r.Form["body"][0],
	})
	w.Header().Set("HX-Trigger", "task-created")
	w.WriteHeader(http.StatusOK)
}

/*
func (c *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {

}
*/

func (c *TaskController) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := c.service.GetAll()
	tasksSlice := make([]taskEntity.Task, 0, len(tasks))
	for i := range len(tasks) {
		tasksSlice = append(tasksSlice, tasks[i])
	}
	handler := templ.Handler(views.Tasks(tasksSlice))
	handler.ServeHTTP(w, r)
}

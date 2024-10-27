package task

import (
	"net/http"
)

func (c *TaskController) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/tasks", c.GetAllTasks)
}

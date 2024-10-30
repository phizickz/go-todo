package task

import (
	"net/http"
)

func (c *TaskController) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/tasks", http.HandlerFunc(c.GetAllTasks))
}

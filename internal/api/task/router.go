package task

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, service TaskService) {
	handler := NewTaskController(service)
}

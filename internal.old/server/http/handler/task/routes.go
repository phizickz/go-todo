package task

import (
	"go-todo/internal/service"
	"net/http"
	"strconv"
)

func RegisterRoutes(
	mux *http.ServeMux,
	taskService service.TaskService,
) {
	handler := newHttpHandler(taskService)

	mux.HandleFunc("/task/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			handler.FindByID(w, r, id)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}

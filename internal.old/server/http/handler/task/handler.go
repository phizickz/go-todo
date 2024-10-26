package task

import (
	"go-todo/internal/service"
	"net/http"
	"strconv"
)

type httpHandler struct {
	taskService service.TaskService
}

func (h *httpHandler) FindByID(w http.ResponseWriter, r *http.Request, id int) {
	task, _ := h.taskService.FindByID(id)
	w.Write([]byte(strconv.Itoa(task.ID)))
}

func newHttpHandler(taskService service.TaskService) *httpHandler {
	return &httpHandler{
		taskService: taskService,
	}
}

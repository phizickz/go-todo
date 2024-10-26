package http

import (
	"go-todo/internal/server/http/handler/task"
	"net/http"
)

func (s *Server) router() http.Handler {
	mux := http.NewServeMux()
	task.RegisterRoutes(mux, s.service.TaskService)
	return mux
}

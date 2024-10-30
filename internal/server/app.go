package server

import (
	indexController "go-todo/internal/api/index"
	taskController "go-todo/internal/api/task"
	taskService "go-todo/internal/service/task"
	"net/http"
)

type Controller interface {
	RegisterRoutes(mux *http.ServeMux)
}

type Server struct {
	Mux         *http.ServeMux
	Controllers []Controller
}

func NewServer() *Server {
	indexCon := indexController.NewIndexController()

	taskSer := taskService.NewTaskService()
	taskCon := taskController.NewTaskController(taskSer)

	controllers := []Controller{
		indexCon,
		taskCon,
	}

	return &Server{
		Mux:         http.NewServeMux(),
		Controllers: controllers,
	}
}

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	for _, controller := range s.Controllers {
		controller.RegisterRoutes(mux)
	}
}

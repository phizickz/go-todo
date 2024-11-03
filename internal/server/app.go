package server

import (
	indexController "go-todo/internal/api/index"
	taskController "go-todo/internal/api/task"
	taskService "go-todo/internal/service/task"
	"net/http"
)

type Controller interface{}

type Server struct {
	Mux *http.ServeMux
}

func NewServer(mux *http.ServeMux) *Server {
	indexCon := indexController.NewIndexController()

	taskSer := taskService.NewTaskService()
	taskCon := taskController.NewTaskController(taskSer)

	mux.HandleFunc("GET /", indexCon.ServeIndex)
	mux.HandleFunc("GET /task/all", taskCon.GetAllTasks)
	mux.HandleFunc("POST /task/create", taskCon.CreateTask)

	return &Server{
		Mux: http.NewServeMux(),
	}
}

package server

import (
	taskHandler "go-todo/internal/handler/task"
	taskRepo "go-todo/internal/repository/task"
	"go-todo/web/views"
	"net/http"

	"github.com/a-h/templ"
)

type Server struct {
	Mux *http.ServeMux
}

func NewServer(mux *http.ServeMux) *Server {
	taskRepo := taskRepo.NewTaskRepository()
	taskHan := taskHandler.NewTaskHandler(taskRepo)

	// Static files
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	// Index
	mux.Handle("/", templ.Handler(views.Index()))

	// Tasks
	mux.HandleFunc("/task/all", taskHan.GetAll)
	mux.HandleFunc("POST /task/create", taskHan.Create)
	mux.HandleFunc("GET /task/update/{id}", taskHan.Edit)
	mux.HandleFunc("POST /task/update/{id}", taskHan.Update)
	mux.HandleFunc("DELETE /task/update/{id}", taskHan.Delete)

	return &Server{
		Mux: http.NewServeMux(),
	}
}

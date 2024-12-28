package server

import (
	"database/sql"
	"fmt"
	taskHandler "go-todo/internal/handler/task"
	taskRepo "go-todo/internal/repository/task"
	"go-todo/web/views"
	"log"
	"net/http"

	"github.com/a-h/templ"
	_ "github.com/lib/pq"
)

type Server struct {
	Mux *http.ServeMux
}

func NewServer(mux *http.ServeMux) *Server {
	host := "localhost"
	port := 5432
	user := "go_user"
	password := "go_password"
	dbname := "go_app_db"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	taskRepo := taskRepo.NewTaskRepository(db)
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

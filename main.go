package main

import (
	"log"
	"net/http"
	"github.com/a-h/templ"
	"go-todo/web/views"
)

func main() {
	http.Handle("/", templ.Handler(views.Root()))
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))

	//http.Handle("/test", http.NotFoundHandler())
	//http.HandleFunc("/test/", handlers.ApiTesting)
	
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

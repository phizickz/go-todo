package main

import (
	"log"
	"net/http"
	"github.com/a-h/templ"
	"go-todo/web/views"
)

func main() {
	http.Handle("/", templ.Handler(views.Index()))
	
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

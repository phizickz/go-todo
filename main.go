package main

import (
	"go-todo/web/views"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	http.Handle("/", templ.Handler(views.Index()))
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
